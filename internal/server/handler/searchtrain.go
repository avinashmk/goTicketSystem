package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/avinashmk/goTicketSystem/internal/consts"
	"github.com/avinashmk/goTicketSystem/internal/model"
	"github.com/avinashmk/goTicketSystem/internal/server/session"
	"github.com/avinashmk/goTicketSystem/internal/store"
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

	if t, err := template.ParseFiles(consts.SearchTrainFormTemplate); err == nil {
		t.Execute(w, model.MakeSearch(s.Gen))
	} else {
		logger.Err.Println("Unable to parse: ", err)
		return
	}
}

// SearchTrainForm SearchTrainForm
func SearchTrainForm(w http.ResponseWriter, r *http.Request) {
	logger.Enter.Println("SearchTrainForm()")
	defer logger.Leave.Println("SearchTrainForm()")

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

	t, err := template.ParseFiles(consts.SearchTrainFormTemplate)
	if err != nil {
		logger.Err.Println("Unable to parse: ", err)
		return
	}

	templateData := model.MakeSearch(s.Gen)
	msg, valid := validateSearchForm(r)
	if !valid {
		templateData.Gen.Message = msg
	} else {
		populateResults(r, &templateData)
	}
	t.Execute(w, templateData)
}

func validateSearchForm(r *http.Request) (msg string, valid bool) {
	logger.Enter.Println("validateSearchForm()")
	defer logger.Leave.Println("validateSearchForm()")

	valid = true
	msg = "Wait re..."
	return
}

func populateResults(r *http.Request, t *model.Search) {
	logger.Enter.Println("populateResults()")
	defer logger.Leave.Println("populateResults()")
	t.Gen.Message = "I know..."
	/*
	 - get all train schema matching against from, to, date
	 - get all train charts matching against train schema, date
	 - put each of the charts into the search model template.
	*/
	from := r.PostFormValue(consts.From)
	to := r.PostFormValue(consts.To)
	date := r.PostFormValue(consts.Date)

	_ = matchSchema(from, to, date)
}

func matchSchema(from string, to string, date string) (schema []store.SchemaDoc) {
	logger.Enter.Println("matchSchema()")
	defer logger.Leave.Println("matchSchema()")

	d, _ := time.Parse(consts.DateLayout, date)
	day := d.Weekday().String()[:3]

	schema, _ = store.FindMatchSchema(from, to, day)
	return
}
