package types

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

// Users specifies users collection schema relative to MongoDB
type Users struct {
	Username string
	Pwd      string
	Role     string
}

// Fill fills a type Users struct from a type bson.M data
func (u *Users) Fill(data bson.M) {
	u.Username = fmt.Sprintf("%v", data["username"])
	u.Pwd = fmt.Sprintf("%v", data["passkey"])
	u.Role = fmt.Sprintf("%v", data["role"])
}
