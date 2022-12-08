package server

import (
	"context"
	"eventsie/events/models"
	pb "eventsie/pb/events"
	"net/http"

	"github.com/kamva/mgm/v3"
)

// Find one event by ID
func (s *Server) Delete(ctx context.Context, in *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	event := &models.Event{}

	if err := mgm.Coll(event).FindByID(in.EventID, event); err != nil {
		return &pb.DeleteResponse{
			Status:  http.StatusInternalServerError,
			Message: "Unexpected error",
			Error:   true,
		}, nil
	}

	if event.CreatedBy != in.UserID && in.UserRole != "admin" {
		return &pb.DeleteResponse{
			Status:  http.StatusUnauthorized,
			Error:   true,
			Message: "Not authorized to delete event",
		}, nil
	}

	if err := mgm.Coll(event).Delete(event); err != nil {
		return &pb.DeleteResponse{
			Status:  http.StatusInternalServerError,
			Message: "Unexpected error",
			Error:   true,
		}, nil
	}

	return &pb.DeleteResponse{
		Status:  200,
		Message: "Successfully deleted event",
	}, nil
}
