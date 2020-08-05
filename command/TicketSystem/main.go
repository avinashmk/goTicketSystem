package main

import (
	"github.com/avinashmk/goTicketSystem/internal/core"
	"github.com/avinashmk/goTicketSystem/logger"
)

func main() {
	_ = logger.Init()
	defer logger.Finalize()

	_ = core.Start()
	logger.Info.Println("Program shut down properly.")
}
