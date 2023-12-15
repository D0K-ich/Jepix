package Database

import (
	cfg "Jepix/internal/Config"
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _ = godotenv.Load("../../internal/config/.env")
var serverAPI = options.ServerAPI(options.ServerAPIVersion1)
var uri = cfg.GetConfDB()
var opts = options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
var client *mongo.Client

func CheackHealth() string {
	client, _ = mongo.Connect(context.TODO(), opts)
	var result bson.M
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}
	return "Pinged your deployment. You successfully connected to MongoDB!"
}

func Disconnect() string {
	err := client.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
	return "Disconnected from MongoDB successfully!"
}

func AddUser(id int, created_at int) {

	collection := client.Database("jepix").Collection("jepix")
	_, err := collection.InsertOne(context.TODO(), bson.D{{"_id", id}, {"created_at", created_at}})
	if err != nil {
		panic(err)
	}
}
