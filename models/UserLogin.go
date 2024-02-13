package models
import "go.mongodb.org/mongo-driver/bson/primitive"

type UserLogin struct{
	ID primitive.ObjectID `bson:"_id"`
	Gmail *string `json:"gmail" validate:"required" bson:"gmail"`
	Password *string  `json:"password" validate:"required" bson:"password"`

}