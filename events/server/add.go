package server

import (
	"context"
	"eventsie/events/models"
	pb "eventsie/pb/events"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/kamva/mgm/v3"
)

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
