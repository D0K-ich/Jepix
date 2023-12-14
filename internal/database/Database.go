package Database

//todo Сделать checkheals, disconnect, init и ping
//todo Перенести файл базы в другую папку

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var collection *mongo.Collection
var ctx = context.TODO()

type Task struct {
	ID        int `bson:"_id"`
	CreatedAt int `bson:"created_at"`
}

func createTask(task *Task) error {
	_, err := collection.InsertOne(ctx, task)
	return err
}

func Disconnect() {
	clientOptions := options.Client().ApplyURI("mongodb+srv://root:root@jepix.yshipgd.mongodb.net/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("jepix").Collection("jepix")

	task1 := &Task{
		ID:        12,
		CreatedAt: 13,
	}
	createTask(task1)

}
