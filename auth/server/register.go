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
	token, err := util.GenerateJWTToken(user.ID.Hex(), user.FirstName, user.LastName, user.Email, user.Role)

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
