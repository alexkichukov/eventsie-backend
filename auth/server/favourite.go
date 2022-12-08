package server

import (
	"context"
	"net/http"

	"eventsie/auth/models"
	"eventsie/auth/util"
	pb "eventsie/pb/auth"

	"github.com/kamva/mgm/v3"
)

func (s *Server) FavouriteEvent(ctx context.Context, in *pb.FavouriteEventRequest) (*pb.FavouriteEventResponse, error) {
	user := &models.User{}

	tokenData, err := util.ParseJWTToken(in.Token)
	if err != nil {
		return &pb.FavouriteEventResponse{Status: http.StatusUnauthorized, Error: true, Message: "Unauthorized request"}, nil
	}

	mgm.Coll(user).FindByID(tokenData.ID, user)

	if user.Email == "" {
		return &pb.FavouriteEventResponse{Status: http.StatusBadRequest, Error: true, Message: "Could not favourite event"}, nil
	}

	user.FavouriteEvents = append(user.FavouriteEvents, in.EventID)
	mgm.Coll(user).Update(user)

	return &pb.FavouriteEventResponse{Status: http.StatusOK, FavouriteEvents: user.FavouriteEvents}, nil
}

func (s *Server) UnfavouriteEvent(ctx context.Context, in *pb.FavouriteEventRequest) (*pb.FavouriteEventResponse, error) {
	user := &models.User{}

	tokenData, err := util.ParseJWTToken(in.Token)
	if err != nil {
		return &pb.FavouriteEventResponse{Status: http.StatusUnauthorized, Error: true, Message: "Unauthorized request"}, nil
	}

	mgm.Coll(user).FindByID(tokenData.ID, user)
	if user.Email == "" {
		return &pb.FavouriteEventResponse{Status: http.StatusBadRequest, Error: true, Message: "Could not favourite event"}, nil
	}

	// Remove event from slice
	for i, event := range user.FavouriteEvents {
		if event == in.EventID {
			user.FavouriteEvents = append(user.FavouriteEvents[:i], user.FavouriteEvents[i+1:]...)
			break
		}
	}

	// Save changes
	mgm.Coll(user).Update(user)

	return &pb.FavouriteEventResponse{Status: http.StatusOK, FavouriteEvents: user.FavouriteEvents}, nil
}
