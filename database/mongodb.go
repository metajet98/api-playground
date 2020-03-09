package database

import (
	"api-playground/helper"

	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Client

// ConnectDatabase : connect DB when init
func ConnectDatabase() {
	opt := options.Client().ApplyURI(helper.GetUserDBConnectionString())
	client, err := mongo.Connect(context.TODO(), opt)

	if err != nil {
		panic("Error when connect to MongoDB" + err.Error())
	}

	db = client
}

func GetUsersDB() *mongo.Database {
	return db.Database(helper.DBUsers)
}
