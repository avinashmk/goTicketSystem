package logger

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

// Log utility for logging.
var (
	InfoLog *log.Logger
	WarnLog *log.Logger
	ErrLog  *log.Logger

	fileName string
	file     *os.File
)

// Init inits logger
func Init() {
	now := time.Now().Unix()
	nowString := strconv.FormatInt(now, 10)
	fileName = "logs_" + nowString + ".txt"
	var err error
	file, err = os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	InfoLog = log.New(file, "[IN] ", log.Ldate|log.Ltime|log.Lshortfile)
	WarnLog = log.New(file, "[WA] ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrLog = log.New(file, "[ER] ", log.Ldate|log.Ltime|log.Lshortfile)
}

// Final Dumps logs to console
func Final() {
	if err := file.Close(); err != nil {
		ErrLog.Println(err)
	}

	fmt.Println("\n\nLogging info:")
	if dat, err := ioutil.ReadFile(fileName); err == nil {
		fmt.Print(string(dat))
	} else {
		ErrLog.Println(err)
	}

	fmt.Print("\nKeep log file?: ")
	reader := bufio.NewReader(os.Stdin)
	response, _ := reader.ReadString('\n')
	response = strings.Replace(response, "\r\n", "", -1)
	if response != "y" {
		fmt.Println("Removing... " + fileName)
		if err := os.Remove(fileName); err != nil {
			log.Fatal(err)
		}
	}
}
