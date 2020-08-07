package handler

import (
	"fmt"
	"net/http"

	"github.com/avinashmk/goTicketSystem/logger"
)

// SearchTrain SearchTrain
func SearchTrain(w http.ResponseWriter, r *http.Request) {
	logger.Enter.Println("SearchTrain()")
	defer logger.Leave.Println("SearchTrain()")

	if r.Method != http.MethodPost {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "SearchTrain got!\n")
}
