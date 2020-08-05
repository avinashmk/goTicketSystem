package housekeeping

import (
	"github.com/avinashmk/goTicketSystem/logger"
)

// Init inits
func Init() (result bool) {
	logger.Info.Println("Init")
	result = true
	return
}

// Finalize Finalizes
func Finalize() {
	logger.Info.Println("Finalize")
}
