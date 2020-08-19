package store

import (
	"context"
	"errors"
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
	Date          string    `bson:"Date"`
	Availability  []string  `bson:"Availability"`
	TicketIDs     []string  `bson:"traintickets_id"`
	ExpireAt      time.Time `bson:"expireAt"`
	TrainNumber   int       `bson:"TrainNumber"`
}

// AddChart adds train Chart
func (cd *ChartDoc) AddChart() (success bool) {
	logger.Enter.Println("AddChart()")
	defer logger.Leave.Println("AddChart()")

	hex, _ := primitive.ObjectIDFromHex(cd.TrainSchemaID)
	bsonDoc := bson.M{
		consts.TrainSchemaID: hex,
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

// FindChart fetches chart from db
func FindChart(trainNum int, date string) (c ChartDoc, err error) {
	logger.Enter.Println("FindCharts()")
	defer logger.Leave.Println("FindCharts()")

	var result bson.M
	filter := bson.D{
		{
			Key:   consts.TrainNumber,
			Value: trainNum,
		},
		{
			Key:   consts.Date,
			Value: date,
		},
	}
	err = chartsCollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			logger.Debug.Println(err)
		} else {
			logger.Err.Println(err)
		}
	} else {
		var valid bool
		c, valid = getChartDoc(result)
		if !valid {
			err = errors.New("Unable to convert bson.M to ChartDoc")
		}
	}
	return
}

func getChartDoc(result bson.M) (c ChartDoc, valid bool) {
	logger.Enter.Println("getChartDoc()")
	defer logger.Leave.Println("getChartDoc()")
	valid = true

	c.TrainSchemaID = result[consts.TrainSchemaID].(primitive.ObjectID).Hex()
	c.Date = result[consts.Date].(string)
	c.ExpireAt = result[consts.ExpireAt].(primitive.DateTime).Time().In(time.UTC)
	c.TrainNumber = int(result[consts.TrainNumber].(int32))

	if avail := result[consts.Availability]; avail != nil {
		for _, val := range []interface{}(avail.(primitive.A)) {
			c.Availability = append(c.Availability, val.(string))
		}
	} else {
		valid = false
	}

	if tickt := result[consts.TicketIDs]; tickt != nil {
		for _, val := range []interface{}(tickt.(primitive.A)) {
			c.TicketIDs = append(c.TicketIDs, val.(string))
		}
	} else {
		valid = false
	}

	return
}
