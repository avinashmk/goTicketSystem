package login

import (
	"github.com/avinashmk/goTicketSystem/console"
	"github.com/avinashmk/goTicketSystem/internal/data"
	"github.com/avinashmk/goTicketSystem/internal/data/types"
	"github.com/avinashmk/goTicketSystem/internal/login/pwd"
	"github.com/avinashmk/goTicketSystem/logger"
	"github.com/avinashmk/goTicketSystem/util"
)

// UserLogin logs/registers users.
func UserLogin() (success bool) {
	userID := console.GetString("Username")
	newUser, userDoc := data.VerifyUser(userID)
	if newUser {
		success = registerUser(userID)
	} else {
		success = verifyCredentials(userDoc)
		if success {
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
		success = util.Attempt3("Mismatched passwords!",
			func() bool {
				password := console.GetString("Password")
				if !pwd.Good(password) {
					return false
				}
				if password == console.GetString("Re-confirm Password") {
					encrypted := pwd.Encrypt(password)
					if data.AddUser(types.Users{
						Username: userID,
						Pwd:      encrypted,
						Role:     types.UserRole,
						Status:   types.ActiveStatus,
					}) {
						console.Prompt("User Registered: " + userID)
						return true
					}
				}
				return false
			})
		if !success {
			console.Prompt("User Registration: Failed")
		}
	}
	return
}

func verifyCredentials(userDoc types.Users) (verified bool) {
	verified = false
	if userDoc.Status == types.LockedStatus {
		console.Prompt("User ID Locked! Contact @avinashmk, Call +46 769502099")
		return
	}

	verified = util.Attempt3("Incorrect Password!",
		func() bool {
			cleartext := console.GetString("Password")
			if pwd.Validate(cleartext, userDoc.Pwd) {
				logger.InfoLog.Println("Logged in! ", userDoc.Username)
				return true
			}
			return false
		})

	if !verified {
		lockUser(userDoc)
	}
	return
}

func lockUser(userDoc types.Users) {
	userDoc.Status = types.LockedStatus
	if data.UpdateUser(userDoc) {
		console.Prompt("ACCOUNT LOCKED! Contact @avinashmk, Call +46 769502099")
	} else {
		logger.ErrLog.Println("Unable to Lock user account: ", userDoc.Username)
	}
}