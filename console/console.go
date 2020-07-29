package console

import (
	"bufio"
	"fmt"
	"io"
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
	fmt.Print(prompt + ": ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		if err != io.EOF { // TODO: Needed?
			logger.ErrLog.Println(err)
		}
	} else {
		line = strings.Replace(input, "\r\n", "", -1)
		line = strings.Fields(input)[0]
	}
	return
}
