package data

import (
	"context"
	"fmt"
	"log"

	"github.com/avinashmk/goTicketSystem/logger"
	"go.mongodb.org/mongo-driver/bson"
)

// VerifyCredentials verifies login credentials
func VerifyCredentials(userID string, pwd string) (validUser bool, userRole string) {
	var result bson.M
	filter := bson.D{{"username", userID}}

	err := usersCollection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	if pwd == result["passkey"] {
		logger.InfoLog.Println("Logged in! ", userID)
		validUser = true
		userRole = fmt.Sprintf("%v", result["role"])
	}
	return
}
