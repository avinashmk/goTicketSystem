package model

import (
	"strconv"
	"time"
)

// Search search
type Search struct {
	Gen              General
	Stations         []Station
	MinDate          string // in the format "yyyy-mm-dd"
	Results          []Result
	AllowReservation string
	From             string
	To               string
}

// Station station
type Station struct {
	Name  string
	Value string
}

// Result result
type Result struct {
	TrainNumber  int
	TrainName    string
	FromTime     string
	ToTime       string
	JourneyTime  string
	Availability int
}

var (
	// StationsList stations
	StationsList []Station
)

// MakeSearch make search form page
func MakeSearch(gen General) (s Search) {
	yr, mon, dt := time.Now().Date()
	minDate := strconv.Itoa(yr) + "-" + strconv.Itoa(int(mon)) + "-" + strconv.Itoa(dt+1)
	s = Search{
		Gen:              gen,
		Stations:         StationsList,
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
