package mongo

import (
	"budgeting_service/configs"
	"budgeting_service/pkg/logger"
	"budgeting_service/storage"
	"budgeting_service/storage/redis"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GoalMongo struct {
	Coll *mongo.Collection
	log  logger.ILogger
}

type Store struct {
	client    *mongo.Client
	db        *mongo.Database
	cfg       configs.Config
	log       logger.ILogger
	redisCash *redis.RedisRepo
}

func (s *Store) Accounts() storage.IAccountStorage {
	return NewAccountMongo(s.db.Collection("account"), s.log, s.redisCash)
}

func (s *Store) Categories() storage.ICategoryStorage {
	return NewCategoryMongo(s.db.Collection("category"), s.log)
}

func (s *Store) Budgets() storage.IBudgetStorage {
	return NewBudgetMongo(s.db.Collection("budget"), s.log)
}

func NewGoalMongo(db *mongo.Collection, lg logger.ILogger) storage.IGoalStorage {
	return &GoalMongo{
		Coll: db,
		log:  lg,
	}
}

func (s *Store) Goals() storage.IGoalStorage {
	fmt.Println("husanbek")
	return NewGoalMongo(s.db.Collection("goal"), s.log)
}

func (s *Store) Transactions() storage.ITransactionStorage {
	return NewTransactionMongo(s.db.Collection("transaction"), s.log)
}

func NewStore(ctx context.Context, cfg configs.Config, log logger.ILogger, redisCash *redis.RedisRepo) (*Store, error) {
	uri := fmt.Sprintf(
		`mongodb://%s:%s@%s:%s/%s?authSource=admin&authMechanism=SCRAM-SHA-256`,
		cfg.MongoUser,
		cfg.MongoPassword,
		cfg.MongoHost,
		cfg.MongoPort,
		cfg.MongoDB,
	)

	clientOptions := options.Client().ApplyURI(uri).
		SetMaxPoolSize(100).
		SetConnectTimeout(10 * time.Second)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Error("Failed to connect to MongoDB", logger.Error(err))
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Error("Failed to ping MongoDB", logger.Error(err))
		return nil, err
	}

	db := client.Database(cfg.MongoDB)

	return &Store{
		client:    client,
		db:        db,
		cfg:       cfg,
		log:       log,
		redisCash: redisCash,
	}, nil
}

func (s *Store) Close() {
	if err := s.client.Disconnect(context.Background()); err != nil {
		s.log.Error("Error disconnecting from MongoDB", logger.Error(err))
	}
}
