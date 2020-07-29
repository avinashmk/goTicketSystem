package data

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
	handlerStarted       chan bool = make(chan bool)
	stopHandler          chan bool = make(chan bool)
	stopHandlerCompleted chan bool = make(chan bool)

	// Collections support concurrency.
	usersCollection *mongo.Collection = nil
)

// Init Inits
func Init() {
	logger.InfoLog.Println("Init")
	go setupHandler()

	logger.InfoLog.Println("Waiting for Handler setup...")
	<-handlerStarted

	logger.InfoLog.Println("Handler up & running!")
}

func setupHandler() {
	// TODO: Have a pool of clients for each collection.
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoDbURL))
	if err != nil {
		logger.ErrLog.Println(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	if err = client.Connect(ctx); err != nil {
		logger.ErrLog.Println(err)
	}

	db := client.Database("trainTicket")
	usersCollection = db.Collection("users")
	logger.InfoLog.Println("setupHandler Complete")

	// Wait as long as the MongoDB connection is needed.
	// (i.e., span of Program)
	handlerStarted <- true
	<-stopHandler
	logger.InfoLog.Println("setupHandler Shut down")
	stopHandlerCompleted <- true
}

// Stop Stops
func Stop() {
	logger.InfoLog.Println("Stop")
	stopHandler <- true
	<-stopHandlerCompleted
}
