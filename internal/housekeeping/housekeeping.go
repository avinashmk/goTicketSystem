package housekeeping

import (
	"sync"
	"time"

	"github.com/avinashmk/goTicketSystem/internal/store"
	"github.com/avinashmk/goTicketSystem/logger"
)

var (
	stopHousekeeping chan bool

	// NewTrainSchema Channel through which new schema added would be notified to housekeeping.
	NewTrainSchema chan store.SchemaDoc

	// TODO: these vars should need mutex if written after Program Init phase
	// map of Weekday vs. list of train numbers on that day.
	daySchema = make(map[string][]int)
	// Stations List of stations, to keep unique list hence map
	stationsList = make(map[string]byte)
	slMux        sync.Mutex
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

	if result = initCharts(); !result {
		logger.Err.Println("Unable to init and setup Charts")
	}

	if result = initTickets(); !result {
		logger.Err.Println("Unable to init and setup Tickets")
	}

	stopHousekeeping = make(chan bool)
	NewTrainSchema = make(chan store.SchemaDoc)
	startHousekeeping()

	return
}

// Finalize Finalizes
func Finalize() {
	logger.Enter.Println("Finalize()")
	defer logger.Leave.Println("Finalize()")

	stopHousekeeping <- true
	<-stopHousekeeping
	logger.Debug.Println("Housekeeping stopped")
}

func startHousekeeping() {
	logger.Enter.Println("startHousekeeping()")
	defer logger.Leave.Println("startHousekeeping()")

	go func() {
		yr, mon, dt := time.Now().UTC().Date()

		// hr, min, _ := time.Now().UTC().Clock()                                                      // TEST
		// tTest := time.Date(yr, mon, dt, hr, min+1, 0, 0, time.UTC)                                  // TEST
		// dur := time.Until(tTest)                                                                    // TEST
		// logger.Debug.Println("Timer set to: ", tTest.String(), " now: ", time.Now().UTC().String()) // TEST
		dur := time.Until(time.Date(yr, mon, dt+1, 0, 1, 0, 0, time.UTC)) // ORIG

		timer := time.NewTimer(dur)
		for {
			select {
			case <-stopHousekeeping:
				logger.Debug.Println("Stopping housekeeping...")
				if !timer.Stop() {
					<-timer.C
				}
				stopHousekeeping <- true
				return

			case <-timer.C:
				logger.Debug.Println("Housekeeping Charts...")
				now := time.Now().UTC()
				yr, mon, dt := now.Date()

				// hr, min, _ := time.Now().UTC().Clock()                               // TEST
				// dur := time.Until(time.Date(yr, mon, dt, hr, min+1, 0, 0, time.UTC)) // TEST
				dur := time.Until(time.Date(yr, mon, dt+1, 0, 1, 0, 0, time.UTC)) // ORIG

				timer.Reset(dur)
				createFutureCharts(now)

			case schema := <-NewTrainSchema:
				handleNewSchema(schema)
			}
		}
	}()
}

// StationsList gives all stations available
func StationsList() (sl []string) {
	logger.Enter.Println("StationsList()")
	defer logger.Leave.Println("StationsList()")

	slMux.Lock()
	for k := range stationsList {
		sl = append(sl, k)
	}
	slMux.Unlock()
	return
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
