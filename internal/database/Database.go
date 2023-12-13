package Database

//todo Сделать checkheals, disconnect, init и ping
//todo Перенести файл базы в другую папку

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	//cfg "Jepix/internal/Config"
	//models "Jepix/internal/Models"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

func Disconnect() {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	uri := "mongodb+srv://root:root@jepix.yshipgd.mongodb.net/"
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	// Send a ping to confirm a successful connection
	var result bson.M
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
}
