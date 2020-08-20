package housekeeping

import (
	"strconv"
	"time"

	"github.com/avinashmk/goTicketSystem/internal/consts"
	"github.com/avinashmk/goTicketSystem/internal/store"
	"github.com/avinashmk/goTicketSystem/logger"
	"go.mongodb.org/mongo-driver/mongo"
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

	// Re-organize schema data as map[weekday] vs. List of *SchemaDoc
	populateDaySchema(schemaList)

	// Verify charts for tomorrow to tomorrow+5 days & create if not exists.
	if !setupCharts() {
		return false
	}

	return true
}

func setupCharts() bool {
	/*
		- For each of Dates,
			- get List of *SchemaDoc for weekday in Date
			- For each of SchemaDoc
				- Verify if SchemaDoc+Date combo is present in TrainChartDocs
					- If exists, remove that combo from TrainChartDoc
					- If not, create new TrainChartDoc
	*/

	now := time.Now().UTC()
	for counter := range []int{1, 2, 3, 4, 5} {

		// - For each of Dates,
		timestamp := now.AddDate(0, 0, counter) // Prepare charts only from tomorrow for consistency.
		day := timestamp.Weekday().String()[:3]
		logger.Debug.Println("Iterating...", day, ": ", timestamp.String())

		// - get List of *SchemaDoc for weekday in Date
		for _, trainNum := range daySchema[day] {

			if createChart(trainNum, timestamp.Format(consts.DateLayout)) {
				logger.Debug.Println("Created charts for Train: ", trainNum, " Date: ", timestamp)
			} else {
				logger.Warn.Println("Charts NOT created for Train: ", trainNum, " Date: ", timestamp)
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

		for _, stop := range schema.Stops {
			stationsList[stop.Name] = '0'
		}
	}
	logger.Debug.Println("daySchema: %v", daySchema)
}

func createFutureCharts(today time.Time) {
	logger.Enter.Println("createFutureCharts()")
	defer logger.Leave.Println("createFutureCharts()")

	futureDate := today.AddDate(0, 0, 5)
	futureDay := futureDate.Weekday().String()[:3]
	for _, trainNum := range daySchema[futureDay] {
		if createChart(trainNum, futureDate.Format(consts.DateLayout)) {
			logger.Debug.Println("Created charts for Train: ", trainNum, " Date: ", futureDate)
		} else {
			logger.Warn.Println("Charts NOT created for Train: ", trainNum, " Date: ", futureDate)
		}
	}
}

func handleNewSchema(schema store.SchemaDoc) {
	logger.Enter.Println("handleNewSchema()")
	defer logger.Leave.Println("handleNewSchema()")

	for _, day := range schema.Frequency {
		daySchema[day] = append(daySchema[day], schema.TrainNumber)
	}

	now := time.Now().UTC()
	for counter := range []int{1, 2, 3, 4, 5} {
		timestamp := now.AddDate(0, 0, counter) // Prepare charts only from tomorrow for consistency.
		for _, day := range schema.Frequency {
			if day == timestamp.Weekday().String()[:3] {
				if createChart(schema.TrainNumber, timestamp.Format(consts.DateLayout)) {
					logger.Debug.Println("Created charts for Train: ", schema.TrainNumber, " Date: ", timestamp)
				} else {
					logger.Err.Println("Charts NOT created for Train: ", schema.TrainNumber, " Date: ", timestamp)
				}
				break
			}
		}
	}

	slMux.Lock()
	for _, stop := range schema.Stops {
		stationsList[stop.Name] = '0'
	}
	slMux.Unlock()
}

func createChart(trainNum int, date string) bool {
	logger.Enter.Println("createChart()")
	defer logger.Leave.Println("createChart()")

	_, err := store.FindChart(trainNum, date)
	if err != mongo.ErrNoDocuments {
		if err == nil {
			logger.Debug.Println("Charts already exist for ", trainNum, " date:", date)
		} else {
			logger.Err.Println("Unable to retrieve chart for ", trainNum, " date:", date)
			logger.Err.Println("Error: ", err)
		}
		return false
	}

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

func getChartExpiry(date string, stops []store.StationSchema) (t time.Time) {
	logger.Enter.Println("getChartExpiry()")
	defer logger.Leave.Println("getChartExpiry()")

	for _, stop := range stops {
		if stop.Position == consts.DestinPos {
			arrival := stop.GetArriveTime(date)
			t = arrival.AddDate(0, 0, 1)
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
