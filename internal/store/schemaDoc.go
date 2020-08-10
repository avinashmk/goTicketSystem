package store

import (
	"time"
)

// TicketSchema TicketSchema
type TicketSchema struct {
	Class      string
	SeatsTotal int
	Fare       int
}

// StationSchema StationSchema
type StationSchema struct {
	Pos    int
	Name   string
	Arrive time.Time
	Depart time.Time
}

// SchemaDoc SchemaDoc
type SchemaDoc struct {
	TrainName    string
	TrainNumber  string
	Frequency    []string
	Availability []TicketSchema
	Stops        []StationSchema
}
