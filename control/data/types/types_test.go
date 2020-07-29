package types

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

func TestFill(t *testing.T) {
	var got Users
	want := Users{"arey", "ooo", "samba"}
	data := bson.M{
		"username": want.Username,
		"passkey":  want.Pwd,
		"role":     want.Role,
	}
	got.Fill(data)
	if got != want {
		t.Errorf("got: %v \nwant: %v", got, want)
	}
}
