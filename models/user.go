package models
import(
	"go.mongodb.org/mongo-driver/bson/primitive"

)
type UserSignup struct{
	ID primitive.ObjectID `bson:"_id"`
	First_name *string `json:"first_name" validate:"required" bson:"first_name"` 
	Last_name *string `json:"last_name" validate:"required" bson:"last_name"`
	Gmail *string `json:"gmail" validate:"required" bson:"gmail"`
	Password *string `json:"password" validate:"required,min=6" bson:"password"`
	

}