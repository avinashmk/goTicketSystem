package handler

import (
	"fmt"
	"net/http"

	"github.com/avinashmk/goTicketSystem/logger"
)

func searchtrainHandler(w http.ResponseWriter, r *http.Request) {
	logger.Enter.Println("searchtrainHandler()")
	defer logger.Leave.Println("searchtrainHandler()")

	if r.Method != http.MethodPost {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "searchtrainHandler got!\n")
}
