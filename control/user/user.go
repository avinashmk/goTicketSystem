package user

import (
	"fmt"

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
	fmt.Println("Welcome!")
}
