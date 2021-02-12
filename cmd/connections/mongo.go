package connections

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)



func GetMongoDatabase(connectionUrl, dataBase string) (db *mongo.Database, err error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionUrl))
	if err != nil {
		log.Fatal(err)
	}

	// Create connect
	err = client.Connect(context.TODO())
	if err != nil {
		return
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return
	}

	db = client.Database(dataBase)
	return

}
