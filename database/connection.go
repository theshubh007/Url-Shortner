package database

import (
	"context"
	"fmt"
	"os"
	"time"

	// "time"
	"log"
	"urlshortner/types"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
)

type manager struct {
	connection *mongo.Client
	ctx        context.Context
	cancel     context.CancelFunc
}

var Mgr Manager

type Manager interface {
	Insert(interface{}, string) (interface{}, error)
	GetUrlFromCode(string, string) (types.UrlDb, error)
}

func ConnectDb() {
	uri := os.Getenv("MONGO_DB_KEY")
	if uri == "" {
		panic("DB_URI environment variable not set")
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println("Failed to connect to MongoDB")
		fmt.Println(err)
	}

	if err == nil {
		fmt.Println("Successfully connected to MongoDB")
	}

	// Assign the context and cancel function to the manager
	Mgr = &manager{connection: client, ctx: ctx, cancel: cancel}

	// Mgr = &manager{connection: client, ctx: ctx, cancel: cancel}
}
