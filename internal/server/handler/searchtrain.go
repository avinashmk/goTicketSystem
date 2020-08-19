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
	// logger.Debug.Println("Search Result data: ", t)
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
	/*
	 - get all train schema matching against from, to, date
	 - get all train charts matching against train schema, date
	 - put each of the charts into the search model template.
	*/
	t.From = r.PostFormValue(consts.From)
	t.To = r.PostFormValue(consts.To)
	date := r.PostFormValue(consts.Date)

	results := getResults(t.From, t.To, date)
	for _, r := range results {
		t.AddResult(r)
	}
}

func getResults(from string, to string, date string) (rArr []model.Result) {
	logger.Enter.Println("getResults()")
	defer logger.Leave.Println("getResults()")

	tmpDate, _ := time.Parse(consts.DateLayout, date)
	day := tmpDate.Weekday().String()[:3]

	tempSchema, _ := store.FindMatchSchema(from, to, day)
	for _, schema := range tempSchema {
		var fromIndex int
		var toIndex int
		for stopIndex, stop := range schema.Stops {
			if stop.Name == from {
				fromIndex = stopIndex
			} else if stop.Name == to {
				toIndex = stopIndex
			}
		}

		if schema.Stops[fromIndex].Position > schema.Stops[toIndex].Position {
			continue
		}

		chart, err := store.FindChart(schema.TrainNumber, date)
		if err != nil {
			logger.Warn.Println("Unable to find charts for: ", schema.TrainNumber, " Date: ", date)
			continue
		}

		fromTime := schema.Stops[fromIndex].GetDepartTime(date)
		toTime := schema.Stops[toIndex].GetArriveTime(date)

		var r model.Result
		r.TrainNumber = schema.TrainNumber
		r.TrainName = schema.TrainName
		r.Availability = len(chart.Availability)
		r.FromTime = fromTime.Format(consts.TimeLayout)
		r.ToTime = toTime.Format(consts.TimeLayout)
		r.JourneyTime = toTime.Sub(fromTime).String()
		rArr = append(rArr, r)
	}
	return
}
