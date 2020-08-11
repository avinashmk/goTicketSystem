package housekeeping

import (
	"github.com/avinashmk/goTicketSystem/logger"
)

// Init inits
func Init() (result bool) {
	logger.Enter.Println("Init()")
	defer logger.Leave.Println("Init()")
	result = true

	/*
		- Prepare Charts db
		- Prepare Tickets db
		- go Monitor Charts db
		- go Monitor Tickets db
	*/

	return
}

// Finalize Finalizes
func Finalize() {
	logger.Enter.Println("Finalize()")
	defer logger.Leave.Println("Finalize()")

	/*
		- Close Monitor Tickets db
		- Close Monitor Charts db
	*/
}
