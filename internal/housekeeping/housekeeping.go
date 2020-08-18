package housekeeping

import (
	"time"

	"github.com/avinashmk/goTicketSystem/internal/model"

	"github.com/avinashmk/goTicketSystem/logger"
)

var (
	chartsOpenDate time.Time
	// stopChartsOpenDate chan bool
)

// Init inits
func Init() (result bool) {
	logger.Enter.Println("Init()")
	defer logger.Leave.Println("Init()")
	result = true

	/*
		- Prepare Charts db
		- Prepare Tickets db
		- go Monitor Charts db
		- go Monitor Tickets db
	*/
	// stopChartsOpenDate = make(chan bool)
	// setupChartsOpenDate()

	if result = initCharts(); !result {
		logger.Err.Println("Unable to init and setup Charts")
	}
	if result = initTickets(); !result {
		logger.Err.Println("Unable to init and setup Tickets")
	}

	exportStationList()
	return
}

// Finalize Finalizes
func Finalize() {
	logger.Enter.Println("Finalize()")
	defer logger.Leave.Println("Finalize()")

	// stopChartsOpenDate <- true
	// <-stopChartsOpenDate
	// logger.Debug.Println("Closed Charts Open Date")

	/*
		- Close Monitor Tickets db
		- Close Monitor Charts db
	*/
}

func exportStationList() {
	logger.Enter.Println("exportStationList()")
	defer logger.Leave.Println("exportStationList()")

	for k := range stationsList {
		s := model.Station{
			Name:  k,
			Value: k,
		}
		model.StationsList = append(model.StationsList, s)
	}
}

// func setupChartsOpenDate() {
// 	logger.Enter.Println("setupChartsOpenDate()")
// 	defer logger.Leave.Println("setupChartsOpenDate()")

// 	go func() {
// 		for {
// 			select {
// 			case <-stopChartsOpenDate:
// 				logger.Debug.Println("Closing Charts Open Date...")
// 				stopChartsOpenDate <- true
// 				return
// 			case <-time.After(1 * time.Hour):
// 				logger.Debug.Println("Updating Charts Open Date...")
// 			}
// 		}
// 	}()
// }
