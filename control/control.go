package control

import (
	"fmt"

	"github.com/avinashmk/goTicketSystem/control/data"
	"github.com/avinashmk/goTicketSystem/control/user"
	"github.com/avinashmk/goTicketSystem/logger"
)

// Init Inits
func Init() {
	logger.InfoLog.Println("Init using logger")
	user.Init()
	data.Init()
}

// Start Starts
func Start() {
	logger.InfoLog.Println("Start")
	userID, pwd := user.GetLoginInfo()
	// TODO: encrypt password
	validUser, role := data.VerifyCredentials(userID, pwd)
	if validUser {
		fmt.Println("Successfully logged in as:", userID)
		logger.InfoLog.Println("Logged in as:", userID, " [Role:", role, "]")
	}
}

// Stop Stops
func Stop() {
	logger.InfoLog.Println("Stopping...")
	defer user.Stop()
	defer data.Stop()
	defer logger.InfoLog.Println("Stopped")
}
