package models

import (
	pb "eventsie/pb/events"

	"github.com/kamva/mgm/v3"
)

type EventPrice struct {
	From float64 `json:"from" bson:"from"`
	To   float64 `json:"to" bson:"to"`
}

type EventLocation struct {
	Address  string `json:"address" bson:"address"`
	City     string `json:"city" bson:"city"`
	Postcode string `json:"psotcode" bson:"psotcode"`
}

type Event struct {
	mgm.DefaultModel `bson:",inline"`
	Title            string   `json:"title" bson:"title"`
	Date             string   `json:"date" bson:"date"`
	Description      string   `json:"description" bson:"description"`
	Tags             []string `json:"tags" bson:"tags"`
	Category         string   `json:"category" bson:"category"`
	Location         *EventLocation
	Price            *EventPrice
}

func EventFromProto(e *pb.Event) *Event {
	return &Event{
		Title:       e.Title,
		Date:        e.Date,
		Description: e.Description,
		Tags:        e.Tags,
		Category:    e.Category,
		Location: &EventLocation{
			Address:  e.Location.Address,
			City:     e.Location.City,
			Postcode: e.Location.Postcode,
		},
		Price: &EventPrice{
			From: e.Price.From,
			To:   e.Price.To,
		},
	}
}

func EventToProto(e *Event) *pb.Event {
	return &pb.Event{
		Id:          e.ID.Hex(),
		Title:       e.Title,
		Date:        e.Date,
		Description: e.Description,
		Tags:        e.Tags,
		Category:    e.Category,
		Location: &pb.Location{
			Address:  e.Location.Address,
			City:     e.Location.City,
			Postcode: e.Location.Postcode,
		},
		Price: &pb.Price{
			From: e.Price.From,
			To:   e.Price.To,
		},
	}
}

func EventsToProto(e []*Event) []*pb.Event {
	result := []*pb.Event{}

	for _, event := range e {
		result = append(result, EventToProto(event))
	}

	return result
}
