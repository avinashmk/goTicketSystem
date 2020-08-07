package handler

import (
	"fmt"
	"net/http"

	"github.com/avinashmk/goTicketSystem/logger"
)

// CancelReservation CancelReservation
func CancelReservation(w http.ResponseWriter, r *http.Request) {
	logger.Enter.Println("CancelReservation()")
	defer logger.Leave.Println("CancelReservation()")

	if r.Method != http.MethodPost {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "CancelReservation got!\n")
}
