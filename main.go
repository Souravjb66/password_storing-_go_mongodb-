package main

import (
	"context"
	"fmt"
	"log"
	database "mysec/database"
	calling "mysec/routes"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)
func main(){
	fmt.Println("welcome")
	database.ConnectDb()
	myDb:=database.MyDb.Db
	defer myDb.Disconnect(context.TODO())
	router:=mux.NewRouter()
	Toretrive(router)
	Forlogin(router)
	
	headers:=handlers.AllowedHeaders([]string{"Content-Type","Authorization"})
    methods:=handlers.AllowedMethods([]string{"GET","HEAD","POST","PUT","OPTIONS"})
    origins:=handlers.AllowedOrigins([]string{"*"})


	log.Fatal(http.ListenAndServe(":8080",handlers.CORS(headers,methods,origins)(router)))


}
func Forlogin(router *mux.Router){
	router.HandleFunc("/signup",calling.SignupUser).Methods("POST")
	router.HandleFunc("/login",calling.LoginUser).Methods("POST")

}
func Toretrive(router *mux.Router){
	router.HandleFunc("/to-save",calling.KeyStoreform).Methods("GET")
	router.HandleFunc("/saved-password",calling.SaveKey).Methods("POST")
	router.HandleFunc("/get-stored-data",calling.GetSecureKey).Methods("POST")
	router.HandleFunc("/get-saveform",calling.Getsave).Methods("GET")
	router.HandleFunc("/index",calling.Index).Methods("GET")
	router.HandleFunc("/login-form",calling.GetLoginForm).Methods("GET")

	


}