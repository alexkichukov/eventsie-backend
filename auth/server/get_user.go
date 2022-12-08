package server

import (
	"context"
	"net/http"

	"eventsie/auth/models"
	pb "eventsie/pb/auth"

	"github.com/kamva/mgm/v3"
)

func (s *Server) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user := &models.User{}

	if err := mgm.Coll(user).FindByID(in.Id, user); err != nil {
		return &pb.GetUserResponse{Status: http.StatusInternalServerError, Error: true, Message: "Unexpected error"}, nil
	}

	// Could not find such user
	if user.Email == "" {
		return &pb.GetUserResponse{Status: http.StatusNotFound, Error: true, Message: "No such user exists"}, nil
	}

	return &pb.GetUserResponse{
		Status: http.StatusOK,
		User: &pb.User{
			Id:              user.ID.Hex(),
			FirstName:       user.FirstName,
			LastName:        user.LastName,
			Email:           user.Email,
			Role:            user.Role,
			FavouriteEvents: user.FavouriteEvents,
			AttendingEvents: user.AttendingEvents,
		},
	}, nil
}
