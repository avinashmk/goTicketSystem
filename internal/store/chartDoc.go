package store

import (
	"context"
	"fmt"
	"time"

	"github.com/avinashmk/goTicketSystem/internal/consts"
	"github.com/avinashmk/goTicketSystem/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// ChartDoc ChartDoc
type ChartDoc struct {
	TrainSchemaID string    `bson:"trainschema_id"`
	Date          time.Time `bson:"Date"`
	Availability  []string  `bson:"Availability"`
	TicketIDs     []string  `bson:"traintickets_id"`
	ExpireAt      time.Time `bson:"expireAt"`
	TrainNumber   int       `bson:"TrainNumber"`
}

// AddChart adds train Chart
func (cd *ChartDoc) AddChart() (success bool) {
	logger.Enter.Println("AddChart()")
	defer logger.Leave.Println("AddChart()")

	bsonDoc := bson.M{
		consts.TrainSchemaID: cd.TrainSchemaID,
		consts.Date:          cd.Date,
		consts.Availability:  cd.Availability,
		consts.TicketIDs:     cd.TicketIDs,
		consts.ExpireAt:      cd.ExpireAt,
		consts.TrainNumber:   cd.TrainNumber,
	}
	res, err := chartsCollection.InsertOne(context.Background(), bsonDoc)
	if err != nil {
		success = false
		logger.Err.Println(err)
	} else {
		success = true
		id := fmt.Sprintf("%v", res.InsertedID)
		logger.Debug.Println("New document inserted: " + id)
	}
	return
}

// GetAllCharts fetches trainSchema from db
func GetAllCharts() (cArr []ChartDoc, err error) {
	logger.Enter.Println("GetAllCharts()")
	defer logger.Leave.Println("GetAllCharts()")

	var cursor *mongo.Cursor
	cursor, err = chartsCollection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		logger.Err.Println(err)
		return
	}

	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		logger.Err.Println(err)
		return
	}

	logger.Debug.Println("got chart docs: ", len(results))
	for _, result := range results {
		if chart, valid := getChartDoc(result); valid {
			cArr = append(cArr, chart)
		} else {
			logger.Err.Println("Unable to convert bson.M to ChartDoc")
		}
	}
	return
}

func getChartDoc(result bson.M) (c ChartDoc, valid bool) {
	logger.Enter.Println("getChartDoc()")
	defer logger.Leave.Println("getChartDoc()")
	valid = true

	c.TrainSchemaID = result[consts.TrainSchemaID].(primitive.ObjectID).String()
	c.Date = result[consts.Date].(primitive.DateTime).Time()
	c.ExpireAt = result[consts.ExpireAt].(primitive.DateTime).Time()
	c.TrainNumber = int(result[consts.TrainNumber].(float64))

	if avail := result[consts.Availability]; avail != nil {
		for _, val := range []interface{}(avail.(primitive.A)) {
			c.Availability = append(c.Availability, val.(string))
		}
	}

	if tickt := result[consts.TicketIDs]; tickt != nil {
		for _, val := range []interface{}(tickt.(primitive.A)) {
			c.TicketIDs = append(c.TicketIDs, val.(string))
		}
	}

	return
}
