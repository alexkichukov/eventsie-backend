package models

import (
	pb "eventsie/pb/events"

	"github.com/kamva/mgm/v3"
)

type EventPrice struct {
	From float64 `json:"from" bson:"from" validate:"omitempty,gte=0"`
	To   float64 `json:"to" bson:"to" validate:"omitempty,gtfield=From"`
}

type EventLocation struct {
	Address  string `json:"address" bson:"address" validate:"required"`
	City     string `json:"city" bson:"city" validate:"required"`
	Postcode string `json:"postcode" bson:"postcode" validate:"required"`
}

type Event struct {
	mgm.DefaultModel `bson:",inline"`
	Title            string         `json:"title" bson:"title" validate:"required"`
	Date             string         `json:"date" bson:"date" validate:"required"`
	Description      string         `json:"description" bson:"description" validate:"required"`
	Tags             []string       `json:"tags" bson:"tags" validate:"required,min=1,dive,required,min=2"`
	Category         string         `json:"category" bson:"category" validate:"required,oneof=music health sports hobbies tech food-and-drink"`
	CreatedBy        string         `json:"createdBy" bson:"createdBy" validate:"required"`
	Location         *EventLocation `validate:"required"`
	Price            *EventPrice    `validate:"required"`
}

func EventFromProto(e *pb.Event) *Event {
	return &Event{
		Title:       e.GetTitle(),
		Date:        e.GetDate(),
		Description: e.GetDescription(),
		Tags:        e.GetTags(),
		Category:    e.GetCategory(),
		CreatedBy:   e.GetCreatedBy(),
		Location: &EventLocation{
			Address:  e.GetLocation().GetAddress(),
			City:     e.GetLocation().GetCity(),
			Postcode: e.GetLocation().GetPostcode(),
		},
		Price: &EventPrice{
			From: e.GetPrice().GetFrom(),
			To:   e.GetPrice().GetTo(),
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
		CreatedBy:   e.CreatedBy,
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
