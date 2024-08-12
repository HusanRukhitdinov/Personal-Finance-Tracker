package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Transaction struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	UserId      string             `bson:"user_id"`
	AccountId   string             `bson:"account_id"`
	CategoryId  string             `bson:"category_id"`
	Amount      float32            `bson:"amount"`
	Typed       string             `bson:"typed"`
	Description string             `bson:"description"`
	Date        string             `bson:"date"`
	CreatedAt   string             `bson:"created_at"`
	UpdatedAt   string             `bson:"updated_at"`
}

type CreateTransaction struct {
	UserId      string  `bson:"user_id"`
	AccountId   string  `bson:"account_id"`
	CategoryId  string  `bson:"category_id"`
	Amount      float32 `bson:"amount"`
	Typed       string  `bson:"typed"`
	Description string  `bson:"description"`
	Date        string  `bson:"date"`
}

type UpdateTransaction struct {
	ID          string  `bson:"-"`
	UserId      string  `bson:"user_id"`
	AccountId   string  `bson:"account_id"`
	CategoryId  string  `bson:"category_id"`
	Amount      float32 `bson:"amount"`
	Typed       string  `bson:"typed"`
	Description string  `bson:"description"`
	Date        string  `bson:"date"`
}

type TransactionsResponse struct {
	Transactions []Transaction `bson:"transactions"`
	Count        int32         `bson:"count"`
}
type GetUserMoneyRequest struct {
	UserId    string `bson:"user_id"`
	StartTime string `bson:"start_time"`
	EndTime   string `bson:"end_time"`
}
type GetUserMoneyResponse struct {
	CategoryId  string `bson:"category_id"`
	TotalAmount string `bson:"total_amount"`
	Time        string `bson:"time"`
}
