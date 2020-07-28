package console

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/avinashmk/goTicketSystem/logger"
)

// Prompt displays prompt string to console.
func Prompt(prompt string) {
	fmt.Println(prompt)
}

// GetString reads a single string(without spaces) from console
func GetString(prompt string) (line string) {
	fmt.Print(prompt + ":")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		logger.ErrLog.Println(err)
	} else {
		line = strings.Replace(input, "\r\n", "", -1)
		line = strings.Fields(input)[0]
	}
	return
}

// // GetLine reads a single line(upto '\n') from console
// func GetLine(prompt string) (line string) {
// 	fmt.Print(prompt + ":")
// 	reader := bufio.NewReader(os.Stdin)
// 	input, err := reader.ReadString('\n')
// 	if err != nil {
// 		logger.ErrLog.Println(err)
// 	} else {
// 		line = strings.Replace(input, "\r\n", "", -1)
// 	}
// 	return
// }
