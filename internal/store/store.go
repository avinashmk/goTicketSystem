package store

import (
	"context"
	"time"

	"github.com/avinashmk/goTicketSystem/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	mongoDbURL = "mongodb://127.0.0.1:27017"
)

var (
	clientStarter chan bool = make(chan bool)
	clientStopper chan bool = make(chan bool)

	// Collections support concurrency.
	usersCollection  *mongo.Collection = nil
	schemaCollection *mongo.Collection = nil
	chartsCollection *mongo.Collection = nil
)

// Init inits
func Init() (result bool) {
	logger.Enter.Println("Init")
	defer logger.Leave.Println("Init")

	go connectDb()
	logger.Debug.Println("Setting up MongoDB clients...")
	result = <-clientStarter
	logger.Info.Println("MongoDB clients setup & running!")
	return
}

// Finalize Finalizes
func Finalize() {
	clientStopper <- true
	logger.Debug.Println("Waiting for MongoDB clients to close...")
	<-clientStopper
	logger.Info.Println("Finalize")
}

func connectDb() {
	logger.Enter.Println("connectDb()")
	defer logger.Leave.Println("connectDb()")

	// TODO: Have a pool of clients for each collection.
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoDbURL))
	if err != nil {
		logger.Err.Println(err)
		clientStarter <- false
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	if err = client.Connect(ctx); err != nil {
		logger.Err.Println(err)
		clientStarter <- false
		return
	}

	db := client.Database("trainTicket")
	usersCollection = db.Collection("users")
	schemaCollection = db.Collection("trainschema")
	chartsCollection = db.Collection("traincharts")

	clientStarter <- true

	// Wait as long as the MongoDB connection is needed.
	// (i.e., span of Program)
	<-clientStopper
	clientStopper <- true
}
