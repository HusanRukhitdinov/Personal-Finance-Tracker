package mongosh

import (
	pb "budget/genproto/budgeting_service"
	"context"
	_ "database/sql"
	"encoding/json"
	"fmt"
	"log/slog"
	"time"
// 	"github.com/segmentio/kafka-go"
// 	"go.mongodb.org/mongo-driver/bson"
// 	_ "go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	_ "go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	_ "go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// 	"google.golang.org/protobuf/types/known/emptypb"
)

type TransactionMongo struct {
	Coll *mongo.Collection
	log  slog.Logger
}

func NewTransactionMongo(db *mongo.Collection, log slog.Logger) *TransactionMongo {
	return &TransactionMongo{
		Coll: db,
		log:  log,
	}
}

func (mongodb *TransactionMongo) CreateTransaction(ctx context.Context, request *pb.TransactionRequest) (*pb.Transaction, error) {
	var (
		err         error
		response    pb.Transaction
		currentTime = time.Now()
	)

	budget := bson.D{
		{Key: "user_id", Value: request.GetUserId()},
		{Key: "account_id", Value: request.GetAccountId()},
		{Key: "category_id", Value: request.GetCategoryId()},
		{Key: "amount", Value: request.GetAmount()},
		{Key: "type", Value: request.GetType()},
		{Key: "description", Value: request.GetDescription()},
		{Key: "date", Value: request.GetDate()},
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
func CreateTransactionMessages() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"kafka:9092"},
		Topic:   "add-to-basket",
		GroupID: "basket-group",
	})
	defer reader.Close()

	for {
		msg, err := reader.FetchMessage(context.Background())
		if err != nil {
			fmt.Printf("error fetching message: %v\n", err)
			continue
		}

		var req pb.TransactionRequest
		err = json.Unmarshal(msg.Value, &req)
		if err != nil {
			fmt.Printf("error unmarshalling message: %v\n", err)
			continue
		}

		err = CreateTransaction(context.Background(), &req)
		if err != nil {
			fmt.Printf("error adding product to basket: %v\n", err)
			continue
		}

		if err := reader.CommitMessages(context.Background(), msg); err != nil {
			fmt.Printf("error committing message: %v\n", err)
		}
	}
}
func (mongodb *TransactionMongo) UpdateTransaction(ctx context.Context, request *pb.Transaction) (*pb.Transaction, error) {
	var (
		params      = bson.M{}
		response    pb.Transaction
		currentTime = time.Now()
	)

	if request.GetUserId() != "" {
		params["user_id"] = request.GetUserId()
	}
	if request.GetAccountId() != "" {
		params["account_id"] = request.GetAccountId()
	}
	if request.GetCategoryId() != "" {
		params["category_id"] = request.GetCategoryId()
	}
	if request.GetAmount() > 0 {
		params["amount"] = request.GetAmount()
	}
	if request.GetType() != "" {
		params["type"] = request.GetType()
	}
	if request.GetDescription() != "" {
		params["description"] = request.GetDescription()
	}
	if request.GetDate() != "" {
		params["date"] = request.GetDate()
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

func (mongodb *TransactionMongo) GetTransaction(ctx context.Context, request *pb.PrimaryKey) (*pb.Transaction, error) {
	var (
		err      error
		response pb.Transaction
	)
	id, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, err
	}
	err = mongodb.Coll.FindOne(ctx, bson.M{"id": id}).Decode(&response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
func (mongodb *TransactionMongo) GetAllTransaction(ctx context.Context, request *pb.GetListRequest) (*pb.Transactions, error) {
	var (
		transactions = []*pb.Transaction{}
		offset       = (request.GetPage() - 1) * request.GetLimit()
		count        int64
		err          error
		filter       = bson.M{}
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
		var transaction pb.Transaction
		if err = cursor.Decode(&transaction); err != nil {
			return nil, err
		}
		fmt.Println("transaction+++++", &transaction)

		transactions = append(transactions, &transaction)
	}
	fmt.Println("transactionssss+++++", transactions)

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &pb.Transactions{
		Transactions: transactions,
		Count:        int32(count),
	}, nil
}
func (mongodb *TransactionMongo) DeleteTransaction(ctx context.Context, request *pb.PrimaryKey) (*emptypb.Empty, error) {
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

func (mongodb *TransactionMongo) GetUserTotalSpend(ctx context.Context, request *pb.GetUserMoneyRequest) (*pb.GetUserMoneyResponse, error) {
	var (
		err      error
		response pb.GetUserMoneyResponse
	)

	matchStage := bson.D{
		{"$match", bson.D{
			{"user_id", request.GetUserId()},
			{"type", "expense"},
			{"date", bson.D{
				{"$gte", request.GetStartTime()},
				{"$lte", request.GetEndTime()},
			}},
		}},
	}

	groupStage := bson.D{
		{"$group", bson.D{
			{"_id", "$category_id"},
			{"total_income", bson.D{{"$sum", "$amount"}}},
			{"date", bson.D{{"$first", "$date"}}},
		}},
	}

	projectStage := bson.D{
		{"$project", bson.D{
			{"category_id", "$_id"},
			{"total_income", 1},
			{"date", 1},
		}},
	}

	cursor, err := mongodb.Coll.Aggregate(ctx, mongo.Pipeline{matchStage, groupStage, projectStage})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (mongodb *TransactionMongo) GetUserTotalIncome(ctx context.Context, request *pb.GetUserMoneyRequest) (*pb.GetUserMoneyResponse, error) {
	var (
		err      error
		response pb.GetUserMoneyResponse
	)

	matchStage := bson.D{
		{"$match", bson.D{
			{"user_id", request.GetUserId()},
			{"type", "income"},
			{"date", bson.D{
				{"$gte", request.GetStartTime()},
				{"$lte", request.GetEndTime()},
			}},
		}},
	}

	groupStage := bson.D{
		{"$group", bson.D{
			{"_id", "$category_id"},
			{"total_income", bson.D{{"$sum", "$amount"}}},
			{"date", bson.D{{"$first", "$date"}}},
		}},
	}

	projectStage := bson.D{
		{"$project", bson.D{
			{"category_id", "$_id"},
			{"total_income", 1},
			{"date", 1},
		}},
	}

	cursor, err := mongodb.Coll.Aggregate(ctx, mongo.Pipeline{matchStage, groupStage, projectStage})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
