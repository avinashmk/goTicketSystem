package handler

import (
	"fmt"
	"net/http"

	"github.com/avinashmk/goTicketSystem/logger"
)

func cancelreservationHandler(w http.ResponseWriter, r *http.Request) {
	logger.Enter.Println("cancelreservationHandler()")
	defer logger.Leave.Println("cancelreservationHandler()")

	if r.Method != http.MethodPost {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "cancelreservationHandler got!\n")
}
