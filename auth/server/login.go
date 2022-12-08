package server

import (
	"context"
	"net/http"

	"eventsie/auth/models"
	"eventsie/auth/util"
	pb "eventsie/pb/auth"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func (s *Server) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	user := &models.User{}

	// Try to find user by email
	mgm.Coll(user).First(bson.M{"email": in.Email}, user)

	if user.Email == "" {
		return &pb.LoginResponse{Status: http.StatusBadRequest, Error: true, Message: "Could not find user with this email"}, nil
	}

	// Compare provided password and user's hashed password from the database
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(in.Password)); err != nil {
		return &pb.LoginResponse{Status: http.StatusUnauthorized, Error: true, Message: "Invalid credentials"}, nil
	}

	// Create a jwt token
	token, err := util.GenerateJWTToken(user.ID.Hex(), user.FirstName, user.LastName, user.Email, user.Role)

	if err != nil {
		return &pb.LoginResponse{Status: http.StatusInternalServerError, Error: true, Message: "Could not generate login token"}, nil
	}

	return &pb.LoginResponse{Status: http.StatusOK, Token: token, User: &pb.User{
		Id:              user.ID.Hex(),
		FirstName:       user.FirstName,
		LastName:        user.LastName,
		Email:           user.Email,
		Role:            user.Role,
		FavouriteEvents: user.FavouriteEvents,
		AttendingEvents: user.AttendingEvents,
	}}, nil
}
