package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Account struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" bson:"_id"`
	UserId    string             `bson:"user_id" bson:"user_id"`
	Name      string             `bson:"name" bson:"name"`
	Type      string             `bson:"type" bson:"type"`
	Balance   float32            `bson:"balance" bson:"balance"`
	Currency  string             `bson:"currency" bson:"currency"`
	CreatedAt string             `bson:"created_at" bson:"created_at"`
	UpdatedAt string             `bson:"updated_at" bson:"updated_at"`
}

type CreateAccount struct {
	UserId   string  `bson:"user_id"`
	Name     string  `bson:"name" `
	Type     string  `bson:"type" `
	Balance  float32 `bson:"balance" `
	Currency string  `bson:"currency" `
}

type UpdateAccount struct {
	ID       string  `bson:"-"`
	UserId   string  `bson:"user_id"`
	Name     string  `bson:"name"`
	Type     string  `bson:"type" `
	Balance  float32 `bson:"balance" `
	Currency string  `bson:"currency" `
}

type AccountsResponse struct {
	Accounts []Account `bson:"accounts" `
	Count    int32     `bson:"count" `
}
