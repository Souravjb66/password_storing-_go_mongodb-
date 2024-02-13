package routes
import(
	"fmt"
	"log"
	db "mysec/database"
	"net/http"
	"encoding/json"
	"context"
	midleware "mysec/middlewares"
	"go.mongodb.org/mongo-driver/bson"
	"html/template"
	
)


type Sign struct{
	First_name *string `json:"first_name" validate:"required" bson:"first_name"` 
	Last_name *string `json:"last_name" validate:"required" bson:"last_name"`
	Gmail *string `json:"gmail" validate:"required" bson:"gmail"`
	Password *string `json:"password" validate:"required,min=6" bson:"password"`
	
}

func SignupUser(w http.ResponseWriter,r *http.Request){
	var mysign Sign
	r.ParseForm()
	Pfirst_name:=r.FormValue("first_name")
	Plast_name:=r.FormValue("last_name")
	Pgmail:=r.FormValue("gmail")
	Ppassword:=r.FormValue("password")
	err:=json.NewDecoder(r.Body).Decode(&mysign)
	if err!=nil{
		log.Println(err)
	}
    var arr []byte=midleware.Makehash(Ppassword)
	Hpassword:=string(arr)
	mysign=Sign{
		First_name:&Pfirst_name,
		Last_name: &Plast_name,
		Gmail: &Pgmail,
		Password: &Hpassword,
	}
	
	

	firstDb:= db.MyDb
	signCollect:=db.MyDb.Db.Database(firstDb.Dbname).Collection(firstDb.Dbcollection1)
	key1,err:=signCollect.InsertOne(context.TODO(),mysign)
	if err!=nil{
		log.Println(err)
		
	}
	fmt.Println("geting from insert :",key1)
	type Login struct{
		Gmail *string
		Password *string
	}
	mylogin:=Login{
		Gmail:mysign.Gmail,
		Password:mysign.Password,
	}
	loginCollect:=db.MyDb.Db.Database(firstDb.Dbname).Collection(firstDb.Dbcollection2)
	loginCollect.InsertOne(context.TODO(),mylogin)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mysign)
	// http.Redirect(w,r,"/login-form",http.StatusSeeOther)
}
var gmailInstance *string

func LoginUser(w http.ResponseWriter,r *http.Request){
	type Login struct{
		Gmail *string `json:"gmail" bson:"gmail"`
		Password *string `json:"password" bson:"password"`
	}
	var login Login
	r.ParseForm()
	Pgmail:=r.FormValue("gmail")
	Ppassword:=r.FormValue("password")
	login=Login{
		Gmail:&Pgmail,
		Password:&Ppassword,
	}

	// json.NewDecoder(r.Body).Decode(&login)
	hash,correctGmail:=midleware.Checkpassword(login.Password,login.Gmail)
	if !correctGmail{
		w.WriteHeader(http.StatusNotFound)
		log.Println("gmail account not signup yet")
		log.Println("hash :",hash)
		log.Println("coorectgmail :",correctGmail)
		return
	}
	if !hash{
		log.Println("wrong password")
		log.Println("hash :",hash)
		log.Println("coorectgmail :",correctGmail)
		return
	}
	gmailInstance=login.Gmail
	fmt.Println("success login :",*gmailInstance)
	http.Redirect(w, r, "/to-save", http.StatusSeeOther)

}

type secretKey struct{
	Secret_key *string `json:"secret_key" `
	Site_id *string `json:"site_id"`
	Email *string `json:"email"`
}
func SaveKey(w http.ResponseWriter,r *http.Request){
	var Key secretKey
	r.ParseForm()
	Psecret:=r.FormValue("storepassword")
	Psite:=r.FormValue("sitename")
	
	// err:=json.NewDecoder(r.Body).Decode(&Key)
	Key=secretKey{
		Secret_key: &Psecret,
		Site_id: &Psite,
		Email:gmailInstance,
	}
	
	collection:=db.MyDb.Db.Database(db.MyDb.Dbname).Collection(db.MyDb.Dbcollection3)
	collection.InsertOne(context.TODO(),Key)
    w.WriteHeader(http.StatusOK)
	log.Println("saved succesfuly")
}

func GetSecureKey(w http.ResponseWriter,r *http.Request){
	type Login struct{
		Gmail *string `json:"gmail"`
		Password *string `json:"password"`
	}
	type Info struct{
		Secret *string `json:"secret"`
		Site *string `json:"site"`
	}
	var getKey secretKey
	var login Login
	var info Info
	r.ParseForm()
	Pgmail:=r.FormValue("gmail")
	Ppassword:=r.FormValue("password")
	login=Login{
		Gmail: &Pgmail,
		Password: &Ppassword,
	}
	
	// json.NewDecoder(r.Body).Decode(&login)
	hashCode,gmail:=midleware.Checkpassword(login.Password,login.Gmail)
	if !hashCode && !gmail{
		log.Println("use correct gmail,password")
		return
	}
	if !hashCode{
		log.Println("wrong password")
		return
	}
    if !gmail{
		log.Println("wrong gmail")
		return
	}

	filter1:=bson.M{"gmail":getKey.Email}
	collection:=db.MyDb.Db.Database(db.MyDb.Dbname).Collection(db.MyDb.Dbcollection3)
	collection.FindOne(context.TODO(),filter1).Decode(&info)
	json.NewEncoder(w).Encode(info)

}
func KeyStoreform(w http.ResponseWriter,r *http.Request){
	temp,err:=template.ParseFiles("templates/keystore.html")
	if err!=nil{
		fmt.Println(err)
		http.Error(w,"problem in sending html",http.StatusInternalServerError)
	}
	temp.Execute(w,nil)

}
func Getsave(w http.ResponseWriter,r *http.Request){
	temp,err:=template.ParseFiles("templates/getstore.html")
	if err!=nil{
		fmt.Println(err)
	}
	temp.Execute(w,nil)
}
func Index(w http.ResponseWriter,r *http.Request){
	temp,err:=template.ParseFiles("templates/index.html")
	if err!=nil{
		fmt.Println(err)

	}
	temp.Execute(w,nil)

}
func GetLoginForm(w http.ResponseWriter,r *http.Request){
	temp,err:=template.ParseFiles("templates/login.html")
	if err!=nil{
		http.Error(w,"htmp parse faile",http.StatusNotFound)
		return
	}
	temp.Execute(w,nil)

}