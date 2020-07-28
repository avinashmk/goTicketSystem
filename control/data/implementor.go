package data

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/avinashmk/goTicketSystem/control/data/types"
	"github.com/avinashmk/goTicketSystem/logger"
)

// VerifyUser verifies login credentials
func VerifyUser(userID string) (newUser bool, userDoc types.Users) {
	newUser = false
	var result bson.M
	filter := bson.D{{"username", userID}}

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
