package models

import "time"

type Product struct {
	Id          string    `json:"id" bson:"_id,omitempty"`
	Name        string    `json:"name"  validate:"required" bson:"name"`
	Description string    `json:"description" validate:"required" bson:"description"`
	Category    string    `json:"category" validate:"required" bson:"category"`
	Content     string    `json:"content" validate:"required" bson:"content"`
	Sold        int       `json:"sold" bson:"sold"`
	Price       int       `json:"price" validate:"required" bson:"price"`
	Discount    int       `json:"discount" bson:"discount"`
	Amount      int       `json:"amount" bson:"amount"`
	Thumbnail   string    `json:"thumbnail" bson:"thumbnail"`
	CreatedAt   time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt" bson:"updatedAt"`
}
