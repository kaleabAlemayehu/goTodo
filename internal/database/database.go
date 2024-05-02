package database

import (

	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/readpref"

	"context"
	"fmt"
	"os"
	"time"

	"github.com/lpernett/godotenv"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"

	// "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// type Database interface{
// 	Collection(ctx context.Context, dbName, colName string) (*mongo.Collection, error)

// 	Context context.Context

// }


func DBConnect() (*mongo.Collection , context.Context){


	err := godotenv.Load(".env")
	if err != nil {
		// TODO adding the current file name so that it can found the current error file 
		log.Fatal("Error loading .env file: %s form database.go", err)
	}

	//get url string
	mongo_url := os.Getenv("DBLOCAL");
	
	// connect to the db
	clientOption := options.Client().ApplyURI(mongo_url)
	client, err := mongo.Connect(context.Background(), clientOption)
	if err != nil {
	 log.Fatal(err)
	}


	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	
	defer client.Disconnect(ctx)

	//check connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
	 log.Fatal(err)
	}
	// create collection
	collection := client.Database("TodoDB").Collection("todo")
	if err != nil {
    	log.Fatal(err)
	}

	fmt.Println("Connected to db")

	return collection, ctx

}


