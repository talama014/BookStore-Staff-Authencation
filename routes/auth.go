package auth

import (
	"context"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/julienschmidt/httprouter"
	db "github.com/talama014/bookstore-staff-authencation/database"
	"github.com/talama014/bookstore-staff-authencation/models"
	res "github.com/talama014/bookstore-staff-authencation/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	username := r.PostFormValue("username")
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	if govalidator.IsNull(username) || govalidator.IsNull(email) || govalidator.IsNull(password) {
		res.JSON(w, 400, "Data can not empty")
		return
	}

	if !govalidator.IsEmail(email) {
		res.JSON(w, 400, "Email is invalid")
		return
	}

	username = models.Santize(username)
	email = models.Santize(email)
	password = models.Santize(password)

	collection := db.ConnectUsers()
	var result bson.M
	errFindUsername := collection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&result)
	errFindEmail := collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&result)

	if errFindUsername == nil || errFindEmail == nil {
		res.JSON(w, 409, "User does exists")
		return
	}

	password, err := models.Hash(password)

	if err != nil {
		res.JSON(w, 500, "Register has failed")
		return
	}

	newUser := bson.M{"username": username, "email": email, "password": password}

	_, errs := collection.InsertOne(context.TODO(), newUser)

	if errs != nil {
		res.JSON(w, 500, "Register has failed")
		return
	}

	res.JSON(w, 201, "Register Succesfully")
}
