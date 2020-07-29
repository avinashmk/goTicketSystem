package util

import "github.com/avinashmk/goTicketSystem/console"

// Attempt3 attempts given function 3 times.
func Attempt3(prompt string, attempt func() bool) (success bool) {
	success = false
	for _, v := range []string{"2", "1", "0"} {
		if success = attempt(); success {
			break
		} else {
			// TODO: Lock account after 3 attempts.
			console.Prompt(prompt + " Attempts left: " + v)
		}
	}
	return
}
