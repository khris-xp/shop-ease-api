package models

import "time"

type User struct {
	Id        string    `json:"id" bson:"_id,omitempty"`
	FirstName string    `json:"firstName"  validate:"required" bson:"firstName"`
	LastName  string    `json:"lastName" validate:"required" bson:"lastName"`
	Email     string    `json:"email" validate:"required" bson:"email"`
	Password  string    `json:"password" validate:"required" bson:"password"`
	Profile   string    `json:"profile" validate:"required" bson:"profile"`
	Cart      []Cart    `json:"cart" validate:"required" bson:"cart"`
	Role      string    `json:"role" bson:"role"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}
