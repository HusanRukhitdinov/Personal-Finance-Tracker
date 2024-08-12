package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	_ "go.mongodb.org/mongo-driver/bson/primitive"
)

type Budget struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" `
	UserId     string             `bson:"user_id"`
	CategoryId string             `bson:"category_id"`
	Amount     float32            `bson:"amount"`
	Period     string             `bson:"period"`
	StartTime  string             `bson:"start_time"`
	EndTime    string             `bson:"end_time"`
	CreatedAt  string             `bson:"created_at"`
	UpdatedAt  string             `bson:"updated_at"`
}

type CreateBudget struct {
	UserId     string  `bson:"user_id"`
	CategoryId string  `bson:"category_id"`
	Amount     float32 `bson:"amount"`
	Period     string  `bson:"period"`
	StartTime  string  `bson:"start_time"`
	EndTime    string  `bson:"end_time"`
}

type UpdateBudget struct {
	ID         string  `bson:"-"`
	UserId     string  `bson:"user_id"`
	CategoryId string  `bson:"category_id"`
	Amount     float32 `bson:"amount"`
	Period     string  `bson:"period"`
	StartTime  string  `bson:"start_time"`
	EndTime    string  `bson:"end_time"`
}

type BudgetsResponse struct {
	Budgets []Budget `bson:"budgets"`
	Count   int32    `bson:"count"`
}
