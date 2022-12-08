package server

import (
	"context"
	"eventsie/events/models"
	pb "eventsie/pb/events"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Update an existing event
func (s *Server) Update(ctx context.Context, in *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	event := models.EventFromProto(in.Event)

	// Validate and set the ID
	eventID, err := primitive.ObjectIDFromHex(in.EventID)
	if err != nil {
		return &pb.UpdateResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid event data",
			Error:   true,
		}, nil
	}
	event.ID = eventID

	// Find the old event and make sure it was either created by the user or the user is an admin
	oldEvent := &models.Event{}
	if err := mgm.Coll(oldEvent).FindByID(in.EventID, oldEvent); err != nil || (oldEvent.CreatedBy != in.UserID && in.UserRole != "admin") {
		return &pb.UpdateResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid event data",
			Error:   true,
		}, nil
	}

	// Validate event
	validate := validator.New()
	if err := validate.Struct(event); err != nil {
		fmt.Println(err)
		return &pb.UpdateResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid event data",
			Error:   true,
		}, nil
	}

	// Update the event
	if err := mgm.Coll(event).Update(event); err != nil {
		return &pb.UpdateResponse{
			Status:  http.StatusInternalServerError,
			Message: "Unexpected error while trying to create event",
			Error:   true,
		}, nil
	}

	return &pb.UpdateResponse{Status: http.StatusOK, Message: "Event updated successfully"}, nil
}
