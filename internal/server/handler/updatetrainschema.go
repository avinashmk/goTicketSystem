package handler

import (
	"fmt"
	"net/http"

	"github.com/avinashmk/goTicketSystem/internal/server/session"
	"github.com/avinashmk/goTicketSystem/logger"
)

// UpdateTrainSchema UpdateTrainSchema
func UpdateTrainSchema(w http.ResponseWriter, r *http.Request) {
	logger.Enter.Println("UpdateTrainSchema()")
	defer logger.Leave.Println("UpdateTrainSchema()")

	if r.Method != http.MethodPost {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	s, httpStatus := session.Get(r)
	if httpStatus == http.StatusOK {
		s.Refresh(w)
	} else {
		http.Error(w, "Unable to authenticate session!", httpStatus)
		return
	}

	fmt.Fprintf(w, "UpdateTrainSchema got!\n")
}
