package server

import (
	"context"
	"net/http"

	"eventsie/auth/models"
	"eventsie/auth/util"
	pb "eventsie/pb/auth"

	"github.com/kamva/mgm/v3"
)

func (s *Server) AttendEvent(ctx context.Context, in *pb.AttendEventRequest) (*pb.AttendEventResponse, error) {
	user := &models.User{}

	tokenData, err := util.ParseJWTToken(in.Token)
	if err != nil {
		return &pb.AttendEventResponse{Status: http.StatusUnauthorized, Error: true, Message: "Unauthorized request"}, nil
	}

	mgm.Coll(user).FindByID(tokenData.ID, user)
	if user.Email == "" {
		return &pb.AttendEventResponse{Status: http.StatusBadRequest, Error: true, Message: "Could not favourite event"}, nil
	}

	user.AttendingEvents = append(user.AttendingEvents, in.EventID)
	mgm.Coll(user).Update(user)

	return &pb.AttendEventResponse{Status: http.StatusOK, AttendingEvents: user.AttendingEvents}, nil
}

func (s *Server) UnattendEvent(ctx context.Context, in *pb.AttendEventRequest) (*pb.AttendEventResponse, error) {
	user := &models.User{}

	tokenData, err := util.ParseJWTToken(in.Token)
	if err != nil {
		return &pb.AttendEventResponse{Status: http.StatusUnauthorized, Error: true, Message: "Unauthorized request"}, nil
	}

	mgm.Coll(user).FindByID(tokenData.ID, user)
	if user.Email == "" {
		return &pb.AttendEventResponse{Status: http.StatusBadRequest, Error: true, Message: "Could not favourite event"}, nil
	}

	// Remove event from slice
	for i, event := range user.AttendingEvents {
		if event == in.EventID {
			user.AttendingEvents = append(user.AttendingEvents[:i], user.AttendingEvents[i+1:]...)
			break
		}
	}

	// Save changes
	mgm.Coll(user).Update(user)

	return &pb.AttendEventResponse{Status: http.StatusOK, AttendingEvents: user.AttendingEvents}, nil
}
