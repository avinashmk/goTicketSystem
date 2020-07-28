package main

import (
	"github.com/avinashmk/goTicketSystem/control"
	"github.com/avinashmk/goTicketSystem/logger"
)

func main() {
	defer logger.Final()
	defer control.Stop()

	logger.Init()
	control.Init()
	control.Start()
}
