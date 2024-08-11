package mongosh

import (
	pb "budget/genproto/budgeting_service"
	"log/slog"

	"context"
	_ "database/sql"
	"fmt"
	"time"

	// "go.mongodb.org/mongo-driver/bson"
	// _ "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	// _ "go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/mongo"
	// _ "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
	// "google.golang.org/protobuf/types/known/emptypb"
	// "google.golang.org/protobuf/types/known/timestamppb"
)

type BudgetMongo struct {
	Coll *mongo.Collection
	log  slog.Logger
}

func NewBudgetMongo(db *mongo.Collection,log slog.Logger) *BudgetMongo {
	return &BudgetMongo{
		Coll: db,
		log: log,
	}
}

func (mongodb *BudgetMongo) CreateBudget(ctx context.Context, request *pb.BudgetRequest) (*pb.Budget, error) {
	var (
		err         error
		response    pb.Budget
		currentTime = timestamppb.Now()
	)

	budget := bson.D{
		{Key: "user_id", Value: request.GetUserId()},
		{Key: "category_id", Value: request.GetCategoryId()},
		{Key: "amount", Value: request.GetAmount()},
		{Key: "period", Value: request.GetPeriod()},
		{Key: "start_time", Value: request.GetStartTime()},
		{Key: "end_time", Value: request.GetEndTime()},
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
func (mongodb *BudgetMongo) UpdateBudget(ctx context.Context, request *pb.Budget) (*pb.Budget, error) {
	var (
		params      = bson.M{}
		response    pb.Budget
		currentTime = time.Now()
	)

	if request.GetUserId() != "" {
		params["user_id"] = request.GetUserId()
	}
	if request.GetCategoryId() != "" {
		params["category_id"] = request.GetCategoryId()
	}
	if request.GetAmount() > 0 {
		params["amount"] = request.GetAmount()
	}
	if request.GetPeriod() != "" {
		params["period"] = request.GetPeriod()
	}
	if request.GetStartTime() != "" {
		params["start_time"] = request.GetStartTime()
	}
	if request.GetEndTime() != "" {
		params["end_time"] = request.GetEndTime()
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

func (mongodb *BudgetMongo) GetBudget(ctx context.Context, request *pb.PrimaryKey) (*pb.Budget, error) {
	var (
		err      error
		response pb.Budget
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
func (mongodb *BudgetMongo) GetAllBudget(ctx context.Context, request *pb.GetListRequest) (*pb.Budgets, error) {
	var (
		budgets = []*pb.Budget{}
		offset  = (request.GetPage() - 1) * request.GetLimit()
		count   int64
		err     error
		filter  = bson.M{}
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
		var budget pb.Budget
		if err = cursor.Decode(&budget); err != nil {
			return nil, err
		}

		budgets = append(budgets, &budget)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &pb.Budgets{
		Budgets: budgets,
		Count:   int32(count),
	}, nil
}
func (mongodb *BudgetMongo) DeleteBudget(ctx context.Context, request *pb.PrimaryKey) (*emptypb.Empty, error) {
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

//func (mongodb *TransactionMongo) GetUserTotalBudget(ctx context.Context, request *pb.PrimaryKey) (*pb.GetUserMoney, error) {
//	var (
//		err      error
//		response pb.GetUserMoney
//	)
//	matchStage := bson.D{{"$match", bson.D{
//		{"user_id", request.GetId()},
//		{"type", "income"},
//	}}}
//	groupStage := bson.D{{"$group", bson.D{
//		{"_id", nil},
//		{"totalAmount", bson.D{{"$sum", "$amount"}}},
//	}}}
//
//	cursor, err := mongodb.Coll.Aggregate(ctx, mongo.Pipeline{matchStage, groupStage})
//	if err != nil {
//		return nil, err
//	}
//	defer cursor.Close(ctx)
//
//	if cursor.Next(ctx) {
//		err := cursor.Decode(&response)
//		if err != nil {
//			return nil, err
//		}
//	}
//
//	return &response, nil
//}
