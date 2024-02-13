package middlewares

import (
	"fmt"
	"log"
    db "mysec/database"
	"golang.org/x/crypto/bcrypt"
	"go.mongodb.org/mongo-driver/bson"
	"context"
)
func Makehash(password string)[]byte{
	bytePassword:=[]byte(password)
	hashPass,err:=bcrypt.GenerateFromPassword(bytePassword,bcrypt.DefaultCost)
	if err!=nil{
		fmt.Println(err)
		log.Panic("hash is not generate")
	}
	return hashPass
}
func Checkpassword(password *string,gmail *string)(bool,bool){
	type login struct{
		
		Gmail string `json:"gmail" bson:"gmail"`
		Password string `json:"password" bson:"password"`
	}
	var getlogin login
	filter:=bson.M{"gmail":*gmail}

	collection:=db.MyDb.Db.Database(db.MyDb.Dbname).Collection(db.MyDb.Dbcollection2)
	err:=collection.FindOne(context.TODO(),filter).Decode(&getlogin)
	if err!=nil{
		log.Println("error in query :",err)
		
	}
	
	
	if getlogin.Gmail==*gmail{
		bytePass:=[]byte(*password)
		
	    err:=bcrypt.CompareHashAndPassword([]byte(getlogin.Password),bytePass)
		if err!=nil{
			log.Println("in if of bcrypt",err)
			
			
			return false,true
		}
		
		return true,true
		
	   
	}else{
		return false,false

	}
	

}

