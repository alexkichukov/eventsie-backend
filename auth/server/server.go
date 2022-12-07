package server

import (
	"context"
	"net/http"

	"eventsie/auth/models"
	"eventsie/auth/util"
	pb "eventsie/pb/auth"

	"github.com/go-playground/validator/v10"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type Server struct {
	pb.UnimplementedAuthServer
}

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
	token, err := util.GenerateJWTToken(user.ID.Hex(), user.FirstName, user.LastName, user.Email)

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

func (s *Server) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	user := &models.User{
		FirstName:       in.GetFirstName(),
		LastName:        in.GetLastName(),
		Email:           in.GetEmail(),
		Password:        in.GetPassword(),
		Role:            models.UserRole,
		FavouriteEvents: []string{},
		AttendingEvents: []string{},
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
	mgm.Coll(user).SimpleFind(&existingUser, bson.M{"email": user.Email}, options.Find().SetLimit(1))

	if len(existingUser) > 0 {
		return &pb.RegisterResponse{Status: http.StatusBadRequest, Error: true, Message: "User already exists"}, nil
	}

	// Add user to db
	mgm.Coll(user).Create(user)

	// Create a jwt token
	token, err := util.GenerateJWTToken(user.ID.Hex(), user.FirstName, user.LastName, user.Email)

	if err != nil {
		return &pb.RegisterResponse{Status: http.StatusInternalServerError, Error: true, Message: "Could not generate login token"}, nil
	}

	return &pb.RegisterResponse{Status: http.StatusOK, Token: token, User: &pb.User{
		Id:              user.ID.Hex(),
		FirstName:       user.FirstName,
		LastName:        user.LastName,
		Email:           user.Email,
		Role:            user.Role,
		FavouriteEvents: user.FavouriteEvents,
		AttendingEvents: user.AttendingEvents,
	}}, nil
}

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

	user.FavouriteEvents = append(user.FavouriteEvents, in.EventID)
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
	for i, event := range user.FavouriteEvents {
		if event == in.EventID {
			user.FavouriteEvents = append(user.FavouriteEvents[:i], user.FavouriteEvents[i+1:]...)
			break
		}
	}

	// Save changes
	mgm.Coll(user).Update(user)

	return &pb.AttendEventResponse{Status: http.StatusOK, AttendingEvents: user.AttendingEvents}, nil
}
