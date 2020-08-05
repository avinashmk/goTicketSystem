package store

import (
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/avinashmk/goTicketSystem/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserDoc placeholder for User credentials.
type UserDoc struct {
	Username string
	Passkey  string
	Role     string
	Status   string
}

var (
	// Collections support concurrency.
	usersCollection *mongo.Collection = nil
)

// NewUser creates a new UserDoc obj
func NewUser(username, password string) (user *UserDoc, err error) {
	user = new(UserDoc)
	user.Username = username
	var bytes []byte
	bytes, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logger.Err.Println("Encryption failed: ", err)
		return
	}
	user.Passkey = string(bytes)
	user.Role = "user"
	user.Status = "active"
	return
}

// FindUser fetches user from db
func FindUser(username string) (user *UserDoc, err error) {
	logger.Enter.Println("FindUser()")
	defer logger.Leave.Println("FindUser()")

	var result bson.M
	filter := bson.D{{
		Key:   "username",
		Value: username,
	}}
	err = usersCollection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			logger.Debug.Println(err)
		} else {
			logger.Err.Println(err)
		}
	} else {
		user = new(UserDoc)
		user.Username = fmt.Sprintf("%v", result["username"])
		user.Passkey = fmt.Sprintf("%v", result["passkey"])
		user.Role = fmt.Sprintf("%v", result["role"])
		user.Status = fmt.Sprintf("%v", result["status"])
	}
	return
}

// RegisterUser adds doc to users collection
func (user *UserDoc) RegisterUser() (success bool) {
	var bsonDoc = bson.M{
		"username": user.Username,
		"passkey":  user.Passkey,
		"role":     user.Role,
		"status":   user.Status,
	}
	res, err := usersCollection.InsertOne(context.Background(), bsonDoc)
	if err != nil {
		success = false
		logger.Err.Println(err)
	} else {
		success = true
		id := fmt.Sprintf("%v", res.InsertedID)
		logger.Debug.Println("New document inserted: " + id)
	}
	return
}

// UpdateUserPassword updates password for existing user
func (user *UserDoc) UpdateUserPassword(newPwd string) (updated bool) {
	if updated = user.updateUser("passkey", newPwd); updated {
		user.Passkey = newPwd
	}
	return
}

// UpdateUserStatus updates password for existing user
func (user *UserDoc) UpdateUserStatus(newStatus string) (updated bool) {
	if updated = user.updateUser("status", newStatus); updated {
		user.Status = newStatus
	}
	return
}

func (user *UserDoc) updateUser(field string, newValue string) bool {
	filter := bson.D{{
		Key:   "username",
		Value: user.Username,
	}}
	update := bson.D{{
		Key: "$set",
		Value: bson.D{
			{
				Key:   field,
				Value: newValue,
			},
		},
	},
	}
	err := usersCollection.FindOneAndUpdate(context.Background(), filter, update).Err()
	if err != nil {
		logger.Err.Println("Unable to update: ", err)
		return false
	}
	return true
}

// Authenticate authenticates user against input password
func (user *UserDoc) Authenticate(pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Passkey), []byte(pwd))
	return err == nil
}
