package handler

import (
	"fmt"
	"net/http"

	"github.com/avinashmk/goTicketSystem/internal/server/session"
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

	s, httpStatus := session.Get(r)
	if httpStatus == http.StatusOK {
		s.Refresh(w)
	} else {
		http.Error(w, "Unable to authenticate session!", httpStatus)
		return
	}

	fmt.Fprintf(w, "SearchTrain got!\n User: %s", s.Gen.UserID)
}