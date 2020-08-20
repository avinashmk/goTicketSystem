package model

import (
	"strconv"
	"time"

	"github.com/avinashmk/goTicketSystem/internal/housekeeping"
)

// Search search
type Search struct {
	Gen              General
	Stations         []string
	MinDate          string // in the format "yyyy-mm-dd"
	Results          []Result
	AllowReservation string
	From             string
	To               string
}

// Result result
type Result struct {
	ChartID      string
	TrainNumber  int
	TrainName    string
	FromTime     string
	ToTime       string
	JourneyTime  string
	Availability int
}

// MakeSearch make search form page
func MakeSearch(gen General) (s Search) {
	yr, mon, dt := time.Now().UTC().Date()
	minDate := strconv.Itoa(yr) + "-" + strconv.Itoa(int(mon)) + "-" + strconv.Itoa(dt+1)
	s = Search{
		Gen:              gen,
		Stations:         housekeeping.StationsList(),
		MinDate:          minDate,
		AllowReservation: "disabled",
	}
	return
}

// AddResult adds result row
func (s *Search) AddResult(r Result) {
	s.Results = append(s.Results, r)
	if r.Availability > 0 {
		s.AllowReservation = ""
	}
	return
}
