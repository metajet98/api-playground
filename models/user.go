package models

import (
	"api-playground/database"
	"log"

	"go.mongodb.org/mongo-driver/mongo/options"

	"context"

	"go.mongodb.org/mongo-driver/bson"
)

// User : Type for user
type User struct {
	Name  string `json:"name" form:"name" query:"name"`
	Email string `json:"email" form:"email" query:"email"`
}

// InsertNewUser : insert new user to DB
func InsertNewUser(u *User) error {
	_, err := database.GetUsersDB().Collection("Users").InsertOne(context.TODO(), u)
	if err != nil {
		return err
	}
	return nil
}

func GetUsers() []*User {
	var result []*User

	findOptions := options.Find()

	cur, err := database.GetUsersDB().Collection("Users").Find(context.Background(), bson.D{{}}, findOptions)

	defer cur.Close(context.TODO())

	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var element *User
		err := cur.Decode(&element)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, element)
	}

	return result
}
