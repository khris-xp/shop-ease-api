package models

import (
	"time"
)

type Cart struct {
	ID        string    `json:"id" bson:"_id"`
	UserId    string    `json:"userId" bson:"userId"`
	ProductId string    `json:"productId" bson:"productId"`
	Quantity  int       `json:"quantity" bson:"quantity"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}
