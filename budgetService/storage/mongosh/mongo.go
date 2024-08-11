package mongosh

import (
	"budget/config"
	"budget/storage"
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Store struct {
	client *mongo.Client
	db     *mongo.Database
	cfg    config.Config
	log 	slog.Logger   }

func (s Store) Accounts() storage.IAccountStorage {
	return NewAccountMongo(s.db.Collection("account"), s.log)
}
func (s Store) Categories() storage.ICategoryStorage {
	return NewCategoryMongo(s.db.Collection("category"), s.log)
}
func (s Store) Budgets() storage.IBudgetStorage {
	return NewBudgetMongo(s.db.Collection("budget"), s.log)
}
func (s Store) Goals() storage.IGoalStorage {
	return NewGoalMongo(s.db.Collection("goal"), s.log)
}
func (s Store) Transactions() storage.ITransactionStorage {
	return NewTransactionMongo(s.db.Collection("transaction"), s.log)
}

func NewStore(ctx context.Context, cfg config.Config) (storage.IStorage, error) {
	uri := fmt.Sprintf(
		`mongodb://%s:%s@%s:%s/%s?authSource=admin&authMechanism=SCRAM-SHA-256`,
		cfg.MongoUser,
		cfg.MongoPassword,
		cfg.MongoHost,
		cfg.MongoPort,
		cfg.MongoDB,
	)

	clientOptions := options.Client().ApplyURI(uri)
	clientOptions.SetMaxPoolSize(100)
	clientOptions.SetConnectTimeout(10 * time.Second)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Error("Failed to connect to MongoDB", err)
		return Store{}, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Error("Failed to ping MongoDB", err)
		return Store{}, err
	}

	db := client.Database(cfg.MongoDB)

	return Store{
		client: client,
		db:     db,
		cfg:    cfg,
	}, nil
}

func (s Store) Close() {
	if err := s.client.Disconnect(context.Background()); err != nil {
		s.log.Error("Error disconnecting from MongoDB", err)
	}
}
