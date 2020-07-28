package main

import (
	"fmt"

	"github.com/avinashmk/goTicketSystem/control"
	"github.com/avinashmk/goTicketSystem/logger"
)

func main() {
	fmt.Println("Hello World")
	defer control.Stop()
	defer logger.Final()

	logger.Init()
	control.Init()
	control.Start()
}
