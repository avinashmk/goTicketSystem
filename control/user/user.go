package user

import (
	"bufio"
	"fmt"
	"os"
)

// Init Inits
func Init() {
	fmt.Println("User::Init")
}

// Stop Stops
func Stop() {
	fmt.Println("User::Stop")
}

// GetLoginInfo Gets the login info
func GetLoginInfo() (userID string, pwd string) {
	fmt.Println("Welcome!")
	fmt.Print("Username: ")
	reader := bufio.NewReader(os.Stdin)
	userID, _ = reader.ReadString('\n')
	fmt.Print("Password: ")
	pwd, _ = reader.ReadString('\n')
	return
}
