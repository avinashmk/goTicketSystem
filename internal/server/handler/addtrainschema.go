package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"github.com/avinashmk/goTicketSystem/internal/consts"
	"github.com/avinashmk/goTicketSystem/internal/model"
	"github.com/avinashmk/goTicketSystem/internal/server/session"
	"github.com/avinashmk/goTicketSystem/internal/store"
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
	if httpStatus == http.StatusOK {
		s.Refresh(w)
	} else {
		http.Error(w, "Unable to authenticate session!", httpStatus)
		return
	}

	trainNumber, err := strconv.Atoi(strings.TrimSpace(r.PostFormValue(consts.TrainNumber)))
	if err != nil {
		logger.Debug.Println("Invalid number for train: ", err)
		return
	}
	trainName := strings.TrimSpace(r.PostFormValue(consts.TrainName))
	freq := getFrequency(r)
	avail := getAvail(r)
	stops := getStops(r)

	sd := store.SchemaDoc{
		TrainName:    trainName,
		TrainNumber:  trainNumber,
		Frequency:    freq,
		Availability: avail,
		Stops:        stops,
	}
	if sd.AddSchema() {
		s.Gen.Message = "Train Schema added! Train ID: " + trainName + "_" + string(strconv.Itoa(trainNumber))
		t, err := template.ParseFiles("./web/templates/menu.html")
		if err != nil {
			logger.Err.Println("Unable to parse template ", err)
			http.Error(w, "Unable to login: Internal Error Occurred", http.StatusInternalServerError)
		} else {
			t.Execute(w, model.MakeMainMenu(s.Gen))
		}
	} else {
		logger.Err.Println("Unable to add schema to DB")
		s.Gen.Message = "Unable to update! Try again!"
		if t, err := template.ParseFiles(consts.AddTrainSchemaFormTemplate); err == nil {
			t.Execute(w, s.Gen)
		} else {
			logger.Err.Println("Unable to parse: ", err)
			return
		}
	}
}

func getFrequency(r *http.Request) []string {
	logger.Enter.Println("getFrequency()")
	defer logger.Leave.Println("getFrequency()")

	freq := []string{}
	for _, day := range []string{
		consts.Monday,
		consts.Tuesday,
		consts.Wednesday,
		consts.Thursday,
		consts.Friday,
		consts.Saturday,
		consts.Sunday,
	} {
		if consts.CheckboxOn == r.PostFormValue(day) {
			freq = append(freq, day)
		}
	}
	return freq
}

func getAvail(r *http.Request) []store.TicketSchema {
	logger.Enter.Println("getAvail()")
	defer logger.Leave.Println("getAvail()")

	avail := []store.TicketSchema{}
	for i := 1; ; i++ {
		if availRow, valid := getAvailRow(r, i); valid {
			avail = append(avail, availRow)
		} else {
			break
		}
	}
	return avail
}

func getAvailRow(r *http.Request, index int) (availRow store.TicketSchema, valid bool) {
	logger.Enter.Println("getAvailRow()")
	defer logger.Leave.Println("getAvailRow()")

	class := strings.TrimSpace(r.PostFormValue(consts.AvailPrefix + strconv.Itoa(index) + consts.AvailClass))
	count := strings.TrimSpace(r.PostFormValue(consts.AvailPrefix + strconv.Itoa(index) + consts.AvailCount))
	fare := strings.TrimSpace(r.PostFormValue(consts.AvailPrefix + strconv.Itoa(index) + consts.AvailFare))

	for i, s := range []string{class, count, fare} {
		if len(s) == 0 {
			valid = false
			logger.Debug.Println("Empty string at: ", i, " row: ", index)
			return
		}
	}

	availRow.Class = class

	var err error
	availRow.SeatsTotal, err = strconv.Atoi(count)
	if err != nil {
		valid = false
		return
	}
	availRow.Fare, err = strconv.Atoi(fare)
	if err != nil {
		valid = false
		return
	}
	valid = true
	return
}

func getStops(r *http.Request) []store.StationSchema {
	logger.Enter.Println("getStops()")
	defer logger.Leave.Println("getStops()")

	stops := []store.StationSchema{}
	for i := 1; ; i++ {
		if stopsRow, valid := getStopsRow(r, i); valid {
			stops = append(stops, stopsRow)
		} else {
			break
		}
	}
	return stops
}

func getStopsRow(r *http.Request, index int) (stopsRow store.StationSchema, valid bool) {
	logger.Enter.Println("getStopsRow()")
	defer logger.Leave.Println("getStopsRow()")

	pos := strings.TrimSpace(r.PostFormValue(consts.StopPrefix + strconv.Itoa(index) + consts.StopPosition))
	name := strings.TrimSpace(r.PostFormValue(consts.StopPrefix + strconv.Itoa(index) + consts.StopStation))
	arrive := strings.TrimSpace(r.PostFormValue(consts.StopPrefix + strconv.Itoa(index) + consts.StopArrival))
	depart := strings.TrimSpace(r.PostFormValue(consts.StopPrefix + strconv.Itoa(index) + consts.StopDepart))
	arriveOffset := strings.TrimSpace(r.PostFormValue(consts.StopPrefix + strconv.Itoa(index) + consts.StopArriveOffset))
	departOffset := strings.TrimSpace(r.PostFormValue(consts.StopPrefix + strconv.Itoa(index) + consts.StopDepartOffset))

	for i, s := range []string{pos, name, arrive, depart, arriveOffset, departOffset} {
		if len(s) == 0 {
			valid = false
			logger.Warn.Println("Empty string at: ", i, " row: ", index)
			return
		}
	}

	var err error
	stopsRow.Position, err = strconv.Atoi(pos)
	if err != nil {
		valid = false
		logger.Err.Println("Unable to convert stopsRow.Position: ", err)
		return
	}

	stopsRow.Name = name
	stopsRow.Arrive = arrive
	stopsRow.ArriveOffset, _ = strconv.Atoi(arriveOffset)
	stopsRow.Depart = depart
	stopsRow.DepartOffset, _ = strconv.Atoi(departOffset)
	logger.Debug.Println("New Stop: ", stopsRow)
	valid = true
	return
}
