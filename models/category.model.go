package models

import (
	"time"
)

type Category struct {
	Id          string    `json:"id" bson:"_id,omitempty"`
	Title       string    `json:"title"  validate:"required" bson:"title"`
	Description string    `json:"description" validate:"required" bson:"description"`
	CreatedAt   time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt" bson:"updatedAt"`
}
