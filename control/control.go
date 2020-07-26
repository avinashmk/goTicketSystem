package control

import (
	"fmt"

	"github.com/avinashmk/goTicketSystem/control/data"
	"github.com/avinashmk/goTicketSystem/control/user"
)

// Init Inits
func Init() {
	fmt.Println("Control::Init")
	user.Init()
	data.Init()
}

// Start Starts
func Start() {
	fmt.Println("Control::Start")
	userID, pwd := user.GetLoginInfo()
	// TODO: encrypt password
	_, _ = data.VerifyCredentials(userID, pwd)
}

// Stop Stops
func Stop() {
	fmt.Println("Control::Stop")
	user.Stop()
	data.Stop()
}
