package store

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/avinashmk/goTicketSystem/internal/consts"
	"github.com/avinashmk/goTicketSystem/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// TicketSchema TicketSchema
type TicketSchema struct {
	Class      string `bson:"class"`
	SeatsTotal int    `bson:"seatstotal"`
	Fare       int    `bson:"fare"`
}

// StationSchema StationSchema
type StationSchema struct {
	Position int       `bson:"position"`
	Name     string    `bson:"name"`
	Arrive   time.Time `bson:"arive"`
	Depart   time.Time `bson:"depart"`
}

// SchemaDoc SchemaDoc
type SchemaDoc struct {
	TrainName    string          `bson:"TrainName"`
	TrainNumber  int             `bson:"TrainNumber"`
	Frequency    []string        `bson:"Frequency"`
	Availability []TicketSchema  `bson:"Availability"`
	Stops        []StationSchema `bson:"Stops"`
}

// AddSchema adds train schema
func (sd *SchemaDoc) AddSchema() (success bool) {
	logger.Enter.Println("AddSchema()")
	defer logger.Leave.Println("AddSchema()")

	bsonDoc := bson.M{
		consts.TrainName:   sd.TrainName,
		consts.TrainNumber: sd.TrainNumber,
		consts.Frequency:   sd.Frequency,
		consts.Tickets:     sd.Availability,
		consts.Stops:       sd.Stops,
	}
	res, err := schemaCollection.InsertOne(context.Background(), bsonDoc)
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

// FindSchema fetches trainSchema from db
func FindSchema(trainNumber int) (s SchemaDoc, err error) {
	logger.Enter.Println("FindSchema()")
	defer logger.Leave.Println("FindSchema()")

	var result bson.M
	filter := bson.D{{
		Key:   consts.TrainNumber,
		Value: trainNumber,
	}}
	err = schemaCollection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			logger.Debug.Println(err)
		} else {
			logger.Err.Println(err)
		}
	} else {
		s, _ = getSchemaDoc(result)
	}
	// logger.Debug.Println("TrainName   : ", s.TrainName)
	// logger.Debug.Println("TrainNumber : ", s.TrainNumber)
	// logger.Debug.Println("Frequency   : ", s.Frequency)
	// logger.Debug.Println("Availability: ", s.Availability)
	// logger.Debug.Println("Stops       : ", s.Stops)
	return
}

// GetAllSchema fetches trainSchema from db
func GetAllSchema(trainNumber int) (sArr []SchemaDoc, err error) {
	logger.Enter.Println("GetAllSchema()")
	defer logger.Leave.Println("GetAllSchema()")

	cursor, err := schemaCollection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	for _, result := range results {
		if s, valid := getSchemaDoc(result); valid {
			sArr = append(sArr, s)
		}
	}
	return
}

func getSchemaDoc(result bson.M) (s SchemaDoc, valid bool) {
	s.TrainName = fmt.Sprintf("%v", result[consts.TrainName])
	s.TrainNumber = int(result[consts.TrainNumber].(int32))
	for _, val := range []interface{}(result[consts.Frequency].(primitive.A)) {
		s.Frequency = append(s.Frequency, val.(string))
	}
	for _, val := range []interface{}(result[consts.Tickets].(primitive.A)) {
		var avail TicketSchema
		bsonBytes, _ := bson.Marshal(val)
		bson.Unmarshal(bsonBytes, &avail)
		s.Availability = append(s.Availability, avail)
	}
	for _, val := range []interface{}(result[consts.Stops].(primitive.A)) {
		var stop StationSchema
		bsonBytes, _ := bson.Marshal(val)
		bson.Unmarshal(bsonBytes, &stop)
		s.Stops = append(s.Stops, stop)
	}
	return
}
