package server

import (
	"context"
	"eventsie/events/models"
	pb "eventsie/pb/events"
	"net/http"

	"github.com/kamva/mgm/v3"
)

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
