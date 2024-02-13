package database

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"context"
	"log"
)

type DBinstance struct {
	Db *mongo.Client
	Dbname string
	Dbcollection1 string
	Dbcollection2 string
	Dbcollection3 string
	
}

var MyDb DBinstance

func ConnectDb(){
	client,err:=mongo.Connect(context.TODO(),options.Client().ApplyURI("mongodb://localhost:27017"))
	if err!=nil{
		log.Fatal("problem in database connection")
	}
	MyDb.Db=client
	MyDb.Dbname="passwordsaver"
	MyDb.Dbcollection1="signupdetails"
	MyDb.Dbcollection2="logindetails"
	MyDb.Dbcollection3="savedInfo"

}