package user

import (
	"github.com/avinashmk/goTicketSystem/logger"
)

// Init Inits
func Init() {
	logger.InfoLog.Println("Init")
}

// Stop Stops
func Stop() {
	logger.InfoLog.Println("Stop")
}

// GetLoginInfo Gets the login info
func GetLoginInfo() {
	logger.InfoLog.Println("Welcome!")
}
