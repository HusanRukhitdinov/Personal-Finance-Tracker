package mongosh
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

type AccountMongo struct {
	Coll *mongo.Collection
	log  slog.Logger
}

func NewAccountMongo(db *mongo.Collection, log slog.Logger) *AccountMongo {
	return &AccountMongo{
		Coll: db,
		log: log,
	}
}

func (mongodb *AccountMongo) CreateAccount(ctx context.Context, request *pb.AccountRequest) (*pb.Account, error) {
	var (
		err         error
		response    pb.Account
		currentTime = time.Now()
	)

	// Prepare the account document for insertion
	account := bson.D{
		{Key: "name", Value: request.Name},
		{Key: "type", Value: request.Type},
		{Key: "balance", Value: request.Balance},
		{Key: "currency", Value: request.Currency},
		{Key: "created_at", Value: currentTime},
	}

	// Log the account document to be inserted
	mongodb.log.Info(fmt.Sprintf("Inserting document: %+v", account))

	// Insert the document
	result, err := mongodb.Coll.InsertOne(ctx, account)
	if err != nil {
		mongodb.log.Error(fmt.Sprintf("InsertOne failed: %v", err))
		return nil, err
	}

	// Log the result of the insertion
	mongodb.log.Info(fmt.Sprintf("InsertOne result: %+v", result))

	// Retrieve the inserted document
	err = mongodb.Coll.FindOne(ctx, bson.M{"_id": result.InsertedID}).Decode(&response)
	if err != nil {
		mongodb.log.Error(fmt.Sprintf("FindOne failed: %v", err))
		return nil, err
	}

	// Log the retrieved document
	mongodb.log.Info(fmt.Sprintf("Document retrieved: %+v", &response))

	// Print to console for debugging
	fmt.Printf("Retrieved Account: %+v\n", &response)

	return &response, nil
}

func (mongodb *AccountMongo) UpdateAccount(ctx context.Context, request *pb.Account) (*pb.Account, error) {
	var (
		params      = bson.M{}
		response    pb.Account
		currentTime = time.Now()
	)

	if request.GetUserId() != "" {
		params["user_id"] = request.GetUserId()
	}
	if request.GetName() != "" {
		params["name"] = request.GetName()
	}
	if request.GetType() != "" {
		params["type"] = request.GetType()
	}
	if request.GetCurrency() != "" {
		params["currency"] = request.GetCurrency()
	}
	if request.GetBalance() > -1 {
		params["balance"] = request.GetBalance()
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
		return nil, fmt.Errorf("no document found with the given id")
	}

	err = mongodb.Coll.FindOne(ctx, filter).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (mongodb *AccountMongo) GetAccount(ctx context.Context, request *pb.PrimaryKey) (*pb.Account, error) {
	var (
		err      error
		response pb.Account
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
func (mongodb *AccountMongo) GetAllAccount(ctx context.Context, request *pb.GetListRequest) (*pb.Accounts, error) {
	var (
		accounts = []*pb.Account{}
		offset   = (request.GetPage() - 1) * request.GetLimit()
		count    int64
		err      error
		filter   = bson.M{}
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
		var account pb.Account
		if err = cursor.Decode(&account); err != nil {
			return nil, err
		}

		accounts = append(accounts, &account)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &pb.Accounts{
		Accounts: accounts,
		Count:    int32(count),
	}, nil
}
func (mongodb *AccountMongo) DeleteAccount(ctx context.Context, request *pb.PrimaryKey) (*emptypb.Empty, error) {
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
