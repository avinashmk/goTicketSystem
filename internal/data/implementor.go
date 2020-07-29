package data

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/avinashmk/goTicketSystem/internal/data/types"
	"github.com/avinashmk/goTicketSystem/logger"
)

// VerifyUser verifies login credentials
func VerifyUser(userID string) (newUser bool, userDoc types.Users) {
	newUser = false
	var result bson.M
	filter := bson.D{{
		Key:   "username",
		Value: userID,
	}}

	err := usersCollection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			newUser = true
		} else {
			logger.ErrLog.Fatalln(err)
		}
	} else {
		userDoc.Fill(result)
	}
	return
}

// AddUser adds doc to users collection
func AddUser(userDoc types.Users) (success bool) {
	var bsonDoc = bson.M{
		"username": userDoc.Username,
		"passkey":  userDoc.Pwd,
		"role":     userDoc.Role,
		"status":   userDoc.Status,
	}
	res, err := usersCollection.InsertOne(context.Background(), bsonDoc)
	if err != nil {
		success = false
		logger.ErrLog.Println("Unable to insert document!")
		logger.ErrLog.Println(err)
	} else {
		success = true
		id := fmt.Sprintf("%v", res.InsertedID)
		logger.InfoLog.Println("New document inserted: " + id)
	}
	return
}

// UpdateUser updates document in users collection
func UpdateUser(userDoc types.Users) (success bool) {
	success = true
	filter := bson.D{{
		Key:   "username",
		Value: userDoc.Username,
	}}
	update := bson.D{
		{Key: "$set",
			Value: bson.D{
				{
					Key:   "passkey",
					Value: userDoc.Pwd,
				},
				{
					Key:   "status",
					Value: userDoc.Status,
				},
			},
		},
	}
	err := usersCollection.FindOneAndUpdate(context.Background(), filter, update).Err()
	if err != nil {
		logger.ErrLog.Println("Unable to update: ", err)
		success = false
	}
	return
}
