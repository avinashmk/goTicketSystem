package handler

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/avinashmk/goTicketSystem/internal/consts"
	"github.com/avinashmk/goTicketSystem/internal/server/session"
	"github.com/avinashmk/goTicketSystem/logger"
)

// AddTrainSchema (Kind of) redirects to the form
func AddTrainSchema(w http.ResponseWriter, r *http.Request) {
	logger.Enter.Println("AddTrainSchema()")
	defer logger.Leave.Println("AddTrainSchema()")

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

	if t, err := template.ParseFiles(consts.AddTrainSchemaFormTemplate); err == nil {
		t.Execute(w, s.Gen)
	} else {
		logger.Err.Println("Unable to parse: ", err)
		return
	}
}

// AddTrainSchemaForm Handles POST request from the Add new train Form.
func AddTrainSchemaForm(w http.ResponseWriter, r *http.Request) {
	logger.Enter.Println("AddTrainSchemaForm()")
	defer logger.Leave.Println("AddTrainSchemaForm()")

	if r.Method != http.MethodPost {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	s, httpStatus := session.Get(r)
	if httpStatus != http.StatusOK {
		http.Error(w, "Unable to authenticate session!", httpStatus)
		return
	}

	fmt.Fprintf(w, "AddTrainSchemaForm got!\n")
	fmt.Fprintf(w, "TrainNumber: %s\n", r.PostFormValue("TrainNumber"))
	fmt.Fprintf(w, "TrainName: %s\n", r.PostFormValue("TrainName"))

	fmt.Fprintf(w, "Frequency: %s\n", r.PostFormValue("Mon"))
	fmt.Fprintf(w, "Frequency: %s\n", r.PostFormValue("Tue"))
	fmt.Fprintf(w, "Frequency: %s\n", r.PostFormValue("Wed"))
	fmt.Fprintf(w, "Frequency: %s\n", r.PostFormValue("Thu"))
	fmt.Fprintf(w, "Frequency: %s\n", r.PostFormValue("Fri"))
	fmt.Fprintf(w, "Frequency: %s\n", r.PostFormValue("Sat"))
	fmt.Fprintf(w, "Frequency: %s\n", r.PostFormValue("Sun"))

	fmt.Fprintf(w, "t1c: %s\n", r.PostFormValue("ticket[1][class]"))
	fmt.Fprintf(w, "t2c: %s\n", r.PostFormValue("ticket[2][class]"))

	s.Refresh(w)
}
