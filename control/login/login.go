package login

import (
	"github.com/avinashmk/goTicketSystem/console"
	"github.com/avinashmk/goTicketSystem/control/data"
	"github.com/avinashmk/goTicketSystem/control/data/types"
	"github.com/avinashmk/goTicketSystem/logger"
)

// UserLogin logs/registers users.
func UserLogin() (success bool) {
	userID := console.GetString("Username")
	newUser, userDoc := data.VerifyUser(userID)
	if newUser {
		success = registerUser(userID)
	} else {
		validUser := verifyCredentials(userDoc)
		if validUser {
			console.Prompt("Successfully logged in as:" + userID)
			logger.InfoLog.Println("Logged in as:", userID, " [Role:", userDoc.Role, "]")
		} else {
			console.Prompt("Login Failed!")
		}
	}
	return
}

// registerUser registers new user
func registerUser(userID string) (success bool) {
	success = false
	register := console.GetString("Username not found! Register instead?[y/n]")
	if register == "y" || register == "Y" {
		for _, v := range []string{"2", "1", "0"} {
			pwd := console.GetString("Password")
			pwd2 := console.GetString("Re-confirm Password")
			if pwd == pwd2 {
				if success = data.AddUser(types.Users{userID, pwd, "user"}); success {
					console.Prompt("User Registered: " + userID)
				}
				break
			} else {
				console.Prompt("Mismatched passwords! Attempts left: " + v)
			}
		}
		if !success {
			console.Prompt("User Registration: Failed")
		}
	}
	return
}

// TODO: encrypt password
func verifyCredentials(userDoc types.Users) (validUser bool) {
	validUser = false
	for _, v := range []string{"2", "1", "0"} {
		pwd := console.GetString("Password")
		if pwd == userDoc.Pwd {
			logger.InfoLog.Println("Logged in! ", userDoc.Username)
			validUser = true
			break
		} else {
			// TODO: Lock account after 3 attempts.
			console.Prompt("Incorrect password! Attempts left: " + v)
		}
	}
	return
}
