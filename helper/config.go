package helper

const MongoDbConnectionString string = "mongodb://Yasuo:yasuo@127.0.0.1:27017/"
const DBUsers string = "Users"

func GetUserDBConnectionString() string {
	return MongoDbConnectionString + DBUsers
}
