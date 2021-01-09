package service

import (
	"go.mongodb.org/mongo-driver/bson"
	"golang-boilerplate/models"
	"github.com/kamva/mgm/v3"
)

//Get document's collection
var user = &models.User{}
var coll = mgm.Coll(user)

func CreateUser(userBody) {
	book:=NewBook("Pride and Prejudice", 345)

	// Make sure pass the model by reference.
	err := mgm.Coll(book).Create(book)
}

func GetUserById(id string) {
	coll.FindByID(id, user)
}

func GetUserByEmail(email string) {
	coll.First(bson.M{"email": email}, user)
}

func UpdateUserByID(userID) {
	// Find your book
	book:= GetUserById(userID)

	// and update it
	book.Name="Moulin Rouge!"
	err:=mgm.Coll(book).Update(book)
}

func DeleteUserByID() {
	// Just find and delete your document
	err := mgm.Coll(book).Delete(book)
}
