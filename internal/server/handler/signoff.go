package handler

import (
	"fmt"
	"net/http"

	"github.com/avinashmk/goTicketSystem/internal/server/session"
	"github.com/avinashmk/goTicketSystem/logger"
)

// Signoff Signoff
func Signoff(w http.ResponseWriter, r *http.Request) {
	logger.Enter.Println("Signoff()")
	defer logger.Leave.Println("Signoff()")

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
		s.Close(w)
		httpStatus = http.StatusSeeOther
	} else {
		http.Error(w, "Unable to authenticate session!", httpStatus)
	}
	http.Redirect(w, r, "/", httpStatus)
}
