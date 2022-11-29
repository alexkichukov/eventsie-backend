package server

import (
	"context"
	"net/http"

	"eventsie/auth/models"
	pb "eventsie/pb/auth"

	"github.com/go-playground/validator/v10"
	"github.com/kamva/mgm/v3"
	"github.com/kamva/mgm/v3/operator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type Server struct {
	pb.UnimplementedAuthServer
}

func (s *Server) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	user := &models.User{
		FirstName:       in.GetFirstName(),
		LastName:        in.GetLastName(),
		Username:        in.GetUsername(),
		Email:           in.GetEmail(),
		Password:        in.GetPassword(),
		Role:            models.UserRole,
		FavouriteEvents: []string{},
	}

	// Validate data
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		return &pb.RegisterResponse{Status: http.StatusBadRequest, Error: true, Message: "Invalid user register data"}, nil
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return &pb.RegisterResponse{Status: http.StatusInternalServerError, Error: true, Message: "Could not register user"}, nil
	}
	user.Password = string(hashedPassword)

	// Check if a user with email or username already exists
	existingUser := []*models.User{}
	mgm.Coll(user).SimpleFind(
		&existingUser,
		bson.M{operator.Or: []bson.M{{"username": user.Username}, {"email": user.Email}}},
		options.Find().SetLimit(1),
	)

	if len(existingUser) > 0 {
		return &pb.RegisterResponse{Status: http.StatusBadRequest, Error: true, Message: "User already exists"}, nil
	}

	// Add user to db
	mgm.Coll(user).Create(user)

	return &pb.RegisterResponse{Status: http.StatusOK, Message: "User registered successfully"}, nil
}
