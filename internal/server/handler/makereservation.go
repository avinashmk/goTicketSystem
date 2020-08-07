package handler

import (
	"fmt"
	"net/http"

	"github.com/avinashmk/goTicketSystem/logger"
)

// MakeReservation MakeReservation
func MakeReservation(w http.ResponseWriter, r *http.Request) {
	logger.Enter.Println("MakeReservation()")
	defer logger.Leave.Println("MakeReservation()")

	if r.Method != http.MethodPost {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "MakeReservation got!\n")
}
