package core

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/avinashmk/goTicketSystem/internal/housekeeping"
	"github.com/avinashmk/goTicketSystem/internal/server"
	"github.com/avinashmk/goTicketSystem/internal/store"
	"github.com/avinashmk/goTicketSystem/logger"
)

var (
	stopCore chan bool
)

// Start Starts
func Start() (result bool) {
	result = false
	logger.Debug.Println("Starting...")
	stopCore = make(chan bool)
	handleSignal()
	sInit := store.Init()
	if sInit {
		hkInit := housekeeping.Init()
		if hkInit {
			hInit := server.Init()
			if hInit {
				result = server.Run()
			}
		}
	}
	<-stopCore
	return
}

// Stop Stops
func Stop() {
	logger.Debug.Println("Stopping...")
	server.Finalize()
	housekeeping.Finalize()
	store.Finalize()
	logger.Info.Println("Stopped")
	stopCore <- true
}

// handleSignal creates a goroutine that monitors for ctrl+C event
//              and shuts the program down gracefully.
func handleSignal() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		logger.Info.Println("Cleaning up")
		Stop()
	}()
}
