package connect

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Mongo(opDocker bool) error {
	if opDocker != false {
		fmt.Print("docker ativado!")

		os.Exit(0)
	}
	fmt.Print("local!")
	clientOptions := options.Client().ApplyURI("mongodb://" + os.Getenv("MONGODB_HOST") + ":27017/" + os.Getenv("MONGODB_DBNAME"))
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return nil
}
