package server

import (
	"context"
	"eventsie/events/models"
	pb "eventsie/pb/events"
	"net/http"

	"github.com/kamva/mgm/v3"
	"github.com/kamva/mgm/v3/operator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
