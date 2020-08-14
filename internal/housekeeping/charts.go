package housekeeping

import (
	"time"

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

					// - Verify if SchemaDoc+Date combo is present in TrainChartDocs
					if trainNum == chart.TrainNumber &&
						chart.Date.After(timestamp) &&
						chart.Date.Before(timestamp.AddDate(0, 0, 1)) {
						f = true
						break
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
	logger.Debug.Println("Chart list: ", chartList)
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
	schema, err := store.FindSchema(trainNum)
	if err != nil {
		logger.Err.Println("Unable to fetch SchemaDoc! ", err)
		return false
	}

	chartDoc := store.ChartDoc{
		TrainSchemaID: "",                 // TODO:<-
		Date:          date,               // TODO:<-
		Availability:  []string{},         // TODO:<-
		TicketIDs:     []string{},         //
		ExpireAt:      time.Now(),         // TODO:<-
		TrainNumber:   schema.TrainNumber, //
	}

	return chartDoc.AddChart()
}
