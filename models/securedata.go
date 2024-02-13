package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)
type SavedPassword struct{
    ID primitive.ObjectID `bson:"_id"`
	Gmail *string `json:"gmail" validate:"required" bson:"gmail"`
	Secure_Key *string `json:"secure_password" validate:"required,min=8" bson:"secure_password"`
	Site_id *string `json:"site_id" validate:"required" bson:"site_id"`
}