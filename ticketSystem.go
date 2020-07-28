package main

import (
	"fmt"

	"github.com/avinashmk/goTicketSystem/control"
	"github.com/avinashmk/goTicketSystem/logger"
)

func main() {
	fmt.Println("Hello World")
	defer logger.Final()
	defer control.Stop()

	logger.Init()
	control.Init()
	control.Start()
}
