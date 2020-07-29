package pwd

import (
	"github.com/avinashmk/goTicketSystem/console"
	"github.com/avinashmk/goTicketSystem/logger"

	"golang.org/x/crypto/bcrypt"
)

const hashcost = 13

// Encrypt encrypts string
func Encrypt(cleartext string) (password string) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(cleartext), hashcost)
	if err != nil {
		logger.ErrLog.Println("Encryption failed: ", err)
	}
	password = string(bytes)
	return
}

// Validate verifies if pwd matches
func Validate(cleartext, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(cleartext))
	return err == nil
}

// Good if password is good enough
func Good(cleartext string) bool {
	if (len(cleartext) >= 6) && (len(cleartext) <= 16) {
		return true
	}
	console.Prompt(" - Password should have length atleast 6.\n - Password should have length atmost 16.")
	return false
}
