package housekeeping

import (
	"github.com/avinashmk/goTicketSystem/logger"
)

func initTickets() bool {
	logger.Enter.Println("initTickets()")
	defer logger.Leave.Println("initTickets()")

	return true
}
