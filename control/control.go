package control

import (
	"github.com/avinashmk/goTicketSystem/console"
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
	console.Prompt("Welcome!")
	defer console.Prompt("Exiting...")
	if !userLogin() {
		return
	}
}

// Stop Stops
func Stop() {
	logger.InfoLog.Println("Stopping...")
	defer user.Stop()
	defer data.Stop()
	defer logger.InfoLog.Println("Stopped")
}
