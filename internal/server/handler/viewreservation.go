package handler

import (
	"fmt"
	"net/http"

	"github.com/avinashmk/goTicketSystem/logger"
)

// ViewReservation ViewReservation
func ViewReservation(w http.ResponseWriter, r *http.Request) {
	logger.Enter.Println("ViewReservation()")
	defer logger.Leave.Println("ViewReservation()")

	if r.Method != http.MethodPost {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "ViewReservation got!\n")
}
