package handler

import (
	"fmt"
	"net/http"

	"github.com/avinashmk/goTicketSystem/logger"
)

// RemoveTrainSchema RemoveTrainSchema
func RemoveTrainSchema(w http.ResponseWriter, r *http.Request) {
	logger.Enter.Println("RemoveTrainSchema()")
	defer logger.Leave.Println("RemoveTrainSchema()")

	if r.Method != http.MethodPost {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "RemoveTrainSchema got!\n")
}
