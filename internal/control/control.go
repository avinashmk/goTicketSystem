package control

import (
	"github.com/avinashmk/goTicketSystem/console"
	"github.com/avinashmk/goTicketSystem/internal/data"
	"github.com/avinashmk/goTicketSystem/internal/login"
	"github.com/avinashmk/goTicketSystem/logger"
)

// Init Inits
func Init() {
	logger.InfoLog.Println("Init using logger")
	data.Init()
}

// Start Starts
func Start() {
	logger.InfoLog.Println("Start")
	console.Prompt("Welcome!")
	defer console.Prompt("Exiting...")
	if !login.UserLogin() {
		return
	}
}

// Stop Stops
func Stop() {
	logger.InfoLog.Println("Stopping...")
	data.Stop()
	defer logger.InfoLog.Println("Stopped")
}
