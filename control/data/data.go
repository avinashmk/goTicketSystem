package data

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Init Inits
func Init() {
	fmt.Println("Data::Init")
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("trainTicket")
	usersCollection := db.Collection("users")

	var result bson.M
	filter := bson.D{{"username", "surya"}}
	//filter := bson.D{{"hello", "world"}}
	res := usersCollection.FindOne(context.Background(), filter)
	err1 := res.Decode(&result)
	if err1 != nil {
		log.Fatal(err1)
	}
	// do something with result...
	fmt.Printf("Found document %v\n", result)
	fmt.Println("hello ", result["passkey"], "!")
	fmt.Println("Raw bytes: ", res)

	
}

// VerifyCredentials verifies creds
func VerifyCredentials(userID string, pwd string) (validUser bool, userRole string) {
	var result bson.M
	filter := bson.D{{"username", userID}}
	res := usersCollection.FindOne(context.Background(), filter)
	err1 := res.Decode(&result)
	fmt.Printf("Found result %v\n", result)
	if err1 != nil {
		log.Fatal(err1)
	}
	passkey := fmt.Sprintf("%v", result["passkey"])
	if pwd == passkey {
		validUser = true
		role := fmt.Sprintf("%v", result["role"])
		userRole = role
	}
	return
}

// Stop Stops
func Stop() {
	fmt.Println("Data::Stop")
}
