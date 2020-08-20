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

type iLogger interface {
	Println(v ...interface{})
}

type ignoreLogger struct {
}

const (
	debugs = true
)

// Log utility for logging.
var (
	Err   iLogger
	Warn  iLogger
	Info  iLogger
	Debug iLogger
	Enter iLogger
	Leave iLogger

	fileName string
	file     *os.File
)

func (il *ignoreLogger) Println(v ...interface{}) {
	// ignore
}

// Init inits logger
func Init() bool {
	now := time.Now().UTC().Unix()
	nowString := strconv.FormatInt(now, 10)
	fileName = "logs/log_" + nowString + ".txt"
	var err error
	file, err = os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	logFlags := log.Ltime | log.Lshortfile

	Err = log.New(file, "[ER] --- ", logFlags)
	Warn = log.New(file, "[WA] --- ", logFlags)
	Info = log.New(file, "[IN] --- ", logFlags)
	if debugs {
		Debug = log.New(file, "[DE]     ", logFlags)
		Enter = log.New(file, "[DE] >>> ", logFlags)
		Leave = log.New(file, "[DE] <<< ", logFlags)
	} else {
		dummyLogger := new(ignoreLogger)
		Debug = dummyLogger
		Enter = dummyLogger
		Leave = dummyLogger
	}

	return true
}

// Finalize Dumps logs to console
func Finalize() {
	if err := file.Close(); err != nil {
		Err.Println(err)
	}

	fmt.Println("\n\nLogging info:")
	if dat, err := ioutil.ReadFile(fileName); err == nil {
		fmt.Print(string(dat))
	} else {
		Err.Println(err)
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
