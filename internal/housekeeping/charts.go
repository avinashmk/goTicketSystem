package housekeeping

import (
	"strconv"
	"time"

	"github.com/avinashmk/goTicketSystem/internal/consts"
	"github.com/avinashmk/goTicketSystem/internal/store"
	"github.com/avinashmk/goTicketSystem/logger"
)

var (
	// map of Weekday vs. list of train numbers on that day.
	daySchema = make(map[string][]int)
)

func initCharts() bool {
	logger.Enter.Println("initCharts()")
	defer logger.Leave.Println("initCharts()")
	/*
		- get all schema.
		- against each schema, verify if a chart has been created.
		- if chart not present already, create a new chart for it.
	*/

	schemaList, err := store.GetAllSchema()
	if err != nil {
		logger.Err.Println("Unable to fetch all trains' schema")
		return false
	}

	/*
		- re-organize schema data as map[weekday] vs. List of *SchemaDoc
		- create a list of Dates in range[today, today + 5days]
		- Get all TrainChartDocs
		- For each of Dates,
			- get List of *SchemaDoc for weekday in Date
			- For each of SchemaDoc
				- Verify if SchemaDoc+Date combo is present in TrainChartDocs
					- If exists, remove that combo from TrainChartDoc
					- If not, create new TrainChartDoc
	*/

	// - re-organize schema data as map[weekday] vs. List of *SchemaDoc
	populateDaySchema(schemaList)

	// - create a list of Dates in range[today, today + 5days]

	// - Get all TrainChartDocs
	chartList, err := store.GetAllCharts()
	if err != nil {
		logger.Err.Println("Unable to fetch charts!")
		return false
	}

	now := time.Now()
	year, month, date := now.Date()
	for dt := date; dt < date+5; dt++ {
		// - For each of Dates,
		timestamp := time.Date(year, month, dt, 0, 0, 0, 0, time.UTC)
		day := timestamp.Weekday().String()[:3]
		logger.Debug.Println("Iterating...", day, ": ", timestamp.String())

		// - get List of *SchemaDoc for weekday in Date
		for _, trainNum := range daySchema[day] {

			// - For each of SchemaDoc
			foundIndex, found := func() (index int, f bool) {
				f = false
				var chart store.ChartDoc
				for index, chart = range chartList {

					// logger.Debug.Println("trainNum: ", trainNum)
					// logger.Debug.Println("chart.TrainNumber: ", chart.TrainNumber)
					// logger.Debug.Println("chart.Date: ", chart.Date)
					// logger.Debug.Println("timestamp: ", timestamp)
					// logger.Debug.Println("timestamp.AddDate(0, 0, 1): ", timestamp.AddDate(0, 0, 1))

					// - Verify if SchemaDoc+Date combo is present in TrainChartDocs
					if trainNum == chart.TrainNumber &&
						(chart.Date.Equal(timestamp) || chart.Date.After(timestamp)) &&
						chart.Date.Before(timestamp.AddDate(0, 0, 1)) {
						f = true
						logger.Debug.Println("Chart already exists for: ", trainNum, " Date: ", timestamp)
						break
					} else {
						// logger.Debug.Println("Chart deemed mismatch")
					}
				}
				return
			}()

			if found {
				// - If exists, remove that combo from TrainChartDoc
				chartList[foundIndex] = chartList[len(chartList)-1]
				chartList[len(chartList)-1] = store.ChartDoc{} // really needed?
				chartList = chartList[:len(chartList)-1]
			} else {
				// 	- If not, create new TrainChartDoc
				if createChart(trainNum, timestamp) {
					logger.Debug.Println("Created charts for Train: ", trainNum, " Date: ", timestamp)
				} else {
					logger.Err.Println("Failed to create charts for Train: ", trainNum, " Date: ", timestamp)
				}
			}
		}
	}
	// logger.Debug.Println("Chart list: ", chartList)
	return true
}

func populateDaySchema(schemaList []store.SchemaDoc) {
	logger.Enter.Println("populateDaySchema()")
	defer logger.Leave.Println("populateDaySchema()")

	for _, schema := range schemaList {
		for _, day := range schema.Frequency {
			daySchema[day] = append(daySchema[day], schema.TrainNumber)
		}
	}
	logger.Debug.Println("daySchema: %v", daySchema)
}

func createChart(trainNum int, date time.Time) bool {
	logger.Enter.Println("createChart()")
	defer logger.Leave.Println("createChart()")

	schema, err := store.FindSchema(trainNum)
	if err != nil {
		logger.Err.Println("Unable to fetch SchemaDoc! ", err)
		return false
	}

	chartDoc := store.ChartDoc{
		TrainSchemaID: schema.ID,
		Date:          date,
		Availability:  populateTickets(schema.Availability),
		TicketIDs:     []string{},
		ExpireAt:      getChartExpiry(date, schema.Stops),
		TrainNumber:   schema.TrainNumber,
	}
	return chartDoc.AddChart()
}

func getChartExpiry(d time.Time, stops []store.StationSchema) (t time.Time) {
	logger.Enter.Println("getChartExpiry()")
	defer logger.Leave.Println("getChartExpiry()")

	for _, stop := range stops {
		if stop.Position == consts.DestinPos {
			hr, min := stop.GetArriveTime()
			yr, mon, dt := d.Date()
			t = time.Date(yr, mon, dt+stop.ArriveOffset+1, hr, min, 0, 0, time.UTC)
			break
		}
	}
	return
}

func populateTickets(avail []store.TicketSchema) (tickets []string) {
	logger.Enter.Println("populateTickets()")
	defer logger.Leave.Println("populateTickets()")

	for _, class := range avail {
		for i := 1; i <= class.SeatsTotal; i++ {
			ticket := class.Class + "_" + strconv.Itoa(i)
			tickets = append(tickets, ticket)
		}
	}

	return
}
