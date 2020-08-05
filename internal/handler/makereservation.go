package handler

import (
	"fmt"
	"net/http"

	"github.com/avinashmk/goTicketSystem/logger"
)

func makereservationHandler(w http.ResponseWriter, r *http.Request) {
	logger.Enter.Println("makereservationHandler()")
	defer logger.Leave.Println("makereservationHandler()")

	if r.Method != http.MethodPost {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "makereservationHandler got!\n")
}
