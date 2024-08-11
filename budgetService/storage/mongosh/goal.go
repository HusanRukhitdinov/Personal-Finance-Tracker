package mongosh

import "C"
import (
	pb "budget/genproto/budgeting_service"
	"context"
	_ "database/sql"
	"fmt"
	"log/slog"
	"time"

	// "go.mongodb.org/mongo-driver/bson"
	// _ "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	// _ "go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/mongo"
	// _ "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
	// "google.golang.org/protobuf/types/known/emptypb"
)

type GoalMongo struct {
	Coll *mongo.Collection
	log  slog.Logger
}

func NewGoalMongo(db *mongo.Collection, log slog.Logger) *GoalMongo {
	return &GoalMongo{
		Coll: db,
		log: log,
	}
}

func (mongodb *GoalMongo) CreateGoal(ctx context.Context, request *pb.GoalRequest) (*pb.Goal, error) {
	var (
		err         error
		response    pb.Goal
		currentTime = time.Now()
	)

	budget := bson.D{
		{Key: "user_id", Value: request.GetUserId()},
		{Key: "name", Value: request.GetName()},
		{Key: "target_amount", Value: request.GetTargetAmount()},
		{Key: "status", Value: request.GetStatus()},
		{Key: "current_amount", Value: request.GetCurrentAmount()},
		{Key: "deadline", Value: request.GetDeadline()},
		{Key: "created_at", Value: currentTime},
	}
	result, err := mongodb.Coll.InsertOne(ctx, &budget)
	if err != nil {
		mongodb.log.Error("this error is insert one ", err)
		return nil, err
	}
	err = mongodb.Coll.FindOne(ctx, bson.M{"_id": result.InsertedID}).Decode(&response)
	if err != nil {
		return nil, err
	}
	return &response, nil

}
func (mongodb *GoalMongo) UpdateGoal(ctx context.Context, request *pb.Goal) (*pb.Goal, error) {
	var (
		params      = bson.M{}
		response    pb.Goal
		currentTime = time.Now()
	)

	if request.GetUserId() != "" {
		params["user_id"] = request.GetUserId()
	}
	if request.GetName() != "" {
		params["name"] = request.GetName()
	}
	if request.Deadline != "" {
		params["deadline"] = request.GetDeadline()
	}
	if request.CurrentAmount > 0 {
		params["current_amount"] = request.GetCurrentAmount()
	}
	if request.TargetAmount > 0 {
		params["target_amount"] = request.GetTargetAmount()
	}
	if request.Status != "" {
		params["status"] = request.GetStatus()
	}
	params["updated_at"] = currentTime

	id, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": params}

	result, err := mongodb.Coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	if result.MatchedCount == 0 {
		return nil, fmt.Errorf("this budget is not found than one")
	}

	err = mongodb.Coll.FindOne(ctx, filter).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (mongodb *GoalMongo) GetGoal(ctx context.Context, request *pb.PrimaryKey) (*pb.Goal, error) {
	var (
		err      error
		response pb.Goal
	)
	id, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, err
	}
	err = mongodb.Coll.FindOne(ctx, bson.M{"_id": id}).Decode(&response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
func (mongodb *GoalMongo) GetAllGoal(ctx context.Context, request *pb.GetListRequest) (*pb.Goals, error) {
	var (
		goals  = []*pb.Goal{}
		offset = (request.GetPage() - 1) * request.GetLimit()
		count  int64
		err    error
		filter = bson.M{}
	)

	if request.GetSearch() != "" {
		filter["name"] = bson.M{"$regex": request.GetSearch(), "$options": "i"}
	}

	count, err = mongodb.Coll.CountDocuments(ctx, filter)
	if err != nil {
		return nil, err
	}

	cursor, err := mongodb.Coll.Find(ctx, filter, options.Find().
		SetSkip(int64(offset)).
		SetLimit(int64(request.GetLimit())).
		SetSort(bson.M{"created_at": -1}))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var goal pb.Goal
		if err = cursor.Decode(&goal); err != nil {
			return nil, err
		}

		goals = append(goals, &goal)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &pb.Goals{
		Goals: goals,
		Count: int32(count),
	}, nil
}
func (mongodb *GoalMongo) DeleteGoal(ctx context.Context, request *pb.PrimaryKey) (*emptypb.Empty, error) {
	var (
		err    error
		filter bson.M
	)
	id, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, err
	}
	filter = bson.M{"_id": id}
	_, err = mongodb.Coll.DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
func (mongodb *GoalMongo) GenerateGoalProgressReport(ctx context.Context, request *pb.PrimaryKey) (*pb.GoalProgressesReport, error) {

	matchStage := bson.D{
		{"$match", bson.D{{"user_id", request.GetId()}}},
	}

	addFieldsStage := bson.D{
		{"$addFields", bson.D{
			{"progress", bson.D{
				{"$multiply", bson.A{
					bson.D{{"$divide", bson.A{"$current_amount", "$target_amount"}}},
					100,
				}},
			}},
			{"status", bson.D{
				{"$cond", bson.D{
					{"if", bson.D{{"$gte", bson.A{"$current_amount", "$target_amount"}}}},
					{"then", "achieved"},
					{"else", bson.D{
						{"$cond", bson.D{
							{"if", bson.D{{"$lt", bson.A{"$deadline", time.Now()}}}},
							{"then", "failed"},
							{"else", "in_progress"},
						}},
					}},
				}},
			}},
		}},
	}

	sortStage := bson.D{
		{"$sort", bson.D{{"created_at", -1}}},
	}

	cursor, err := mongodb.Coll.Aggregate(ctx, mongo.Pipeline{matchStage, addFieldsStage, sortStage})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var goals []*pb.GoalProgressReport
	if err = cursor.All(ctx, &goals); err != nil {
		return nil, err
	}

	return &pb.GoalProgressesReport{GoalProgressesReport: goals}, nil
}
