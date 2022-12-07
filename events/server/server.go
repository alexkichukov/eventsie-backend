package server

import (
	"context"
	"eventsie/events/models"
	pb "eventsie/pb/events"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/kamva/mgm/v3"
	"github.com/kamva/mgm/v3/operator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Server struct {
	pb.UnimplementedEventsServer
}

// Find one event by ID
func (s *Server) FindOne(ctx context.Context, in *pb.FindOneRequest) (*pb.FindOneResponse, error) {
	event := &models.Event{}

	if err := mgm.Coll(event).FindByID(in.Id, event); err != nil {
		return &pb.FindOneResponse{
			Status:  http.StatusInternalServerError,
			Message: "Unexpected error while trying to find event",
			Error:   true,
		}, nil
	}

	return &pb.FindOneResponse{
		Status: 200,
		Event:  models.EventToProto(event),
	}, nil
}

// Find many events
func (s *Server) FindMany(ctx context.Context, in *pb.FindManyRequest) (*pb.FindManyResponse, error) {
	events := []*models.Event{}

	filter := bson.M{}

	// Build the filter
	if in.Categories != nil {
		filter["category"] = bson.M{operator.In: in.Categories}
	}
	if in.Id != nil {
		objectIds := make([]primitive.ObjectID, len(in.Id))
		for i, id := range in.Id {
			oId, err := primitive.ObjectIDFromHex(id)
			if err != nil {
				return &pb.FindManyResponse{
					Status:  http.StatusBadRequest,
					Message: "Bad request data",
					Error:   true,
				}, nil
			}
			objectIds[i] = oId
		}

		filter["_id"] = bson.M{operator.In: objectIds}
	}
	if in.Cities != nil {
		filter["location.city"] = bson.M{operator.In: in.Cities}
	}
	if in.CreatedBy != nil {
		filter["createdBy"] = bson.M{operator.In: in.CreatedBy}
	}
	if in.Tags != nil {
		filter["tags"] = bson.M{operator.In: in.Tags}
	}
	if in.PriceFrom != nil || in.PriceTo != nil {
		filter["price.from"] = bson.M{}
	}
	if in.PriceFrom != nil {
		filter["price.from"].(bson.M)[operator.Gte] = *in.PriceFrom
	}
	if in.PriceTo != nil {
		filter["price.from"].(bson.M)[operator.Lte] = *in.PriceTo
	}

	if err := mgm.Coll(&models.Event{}).SimpleFind(&events, filter); err != nil {
		return &pb.FindManyResponse{
			Status:  http.StatusInternalServerError,
			Message: "Unexpected error while trying to find events",
			Error:   true,
		}, nil
	}

	return &pb.FindManyResponse{
		Status: 200,
		Events: models.EventsToProto(events),
	}, nil
}

// Create a new event
func (s *Server) Add(ctx context.Context, in *pb.AddRequest) (*pb.AddResponse, error) {
	event := models.EventFromProto(in.Event)

	// Validate event
	validate := validator.New()
	if err := validate.Struct(event); err != nil {
		return &pb.AddResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid event data",
			Error:   true,
		}, nil
	}

	if err := mgm.Coll(event).Create(event); err != nil {
		return &pb.AddResponse{
			Status:  http.StatusInternalServerError,
			Message: "Unexpected error while trying to create event",
			Error:   true,
		}, nil
	}

	return &pb.AddResponse{Status: http.StatusOK, Message: "Event created successfully"}, nil
}
