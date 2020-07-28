package user

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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
func GetLoginInfo() (userID string, pwd string) {
	fmt.Println("Welcome!")
	fmt.Print("Username: ")
	reader := bufio.NewReader(os.Stdin)
	userID, _ = reader.ReadString('\n')
	userID = strings.Replace(userID, "\r\n", "", -1)
	fmt.Print("Password: ")
	pwd, _ = reader.ReadString('\n')
	pwd = strings.Replace(pwd, "\r\n", "", -1)
	return
}
