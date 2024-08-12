package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Category struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" bson:"_id"`
	UserId    string             `bson:"user_id"`
	Name      string             `bson:"name"`
	Type      string             `bson:"type"`
	CreatedAt string             `bson:"created_at"`
	UpdatedAt string             `bson:"updated_at"`
}

type CreateCategory struct {
	UserId string `bson:"user_id"`
	Name   string `bson:"name"`
	Type   string `bson:"type"`
}

type UpdateCategory struct {
	ID     string `bson:"-"`
	UserId string `bson:"user_id"`
	Name   string `bson:"name"`
	Type   string `bson:"type"`
}

type CategoriesResponse struct {
	Categories []Category `bson:"categories"`
	Count      int32      `bson:"count"`
}
