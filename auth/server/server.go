package server

import (
	"context"
	"net/http"

	"eventsie/auth/config"
	"eventsie/auth/models"
	pb "eventsie/pb/auth"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
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
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":        user.ID.Hex(),
		"firstName": user.FirstName,
		"lastName":  user.LastName,
		"email":     user.Email,
	})

	cfg := config.GetConfig()
	tokenString, err := token.SignedString(cfg.JWT_SECRET)

	if err != nil {
		return &pb.LoginResponse{Status: http.StatusInternalServerError, Error: true, Message: "Could not generate login token"}, nil
	}

	return &pb.LoginResponse{Status: http.StatusUnauthorized, Token: tokenString}, nil
}

func (s *Server) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	user := &models.User{
		FirstName:       in.GetFirstName(),
		LastName:        in.GetLastName(),
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
	mgm.Coll(user).SimpleFind(&existingUser, bson.M{"email": user.Email}, options.Find().SetLimit(1))

	if len(existingUser) > 0 {
		return &pb.RegisterResponse{Status: http.StatusBadRequest, Error: true, Message: "User already exists"}, nil
	}

	// Add user to db
	mgm.Coll(user).Create(user)

	return &pb.RegisterResponse{Status: http.StatusOK, Message: "User registered successfully"}, nil
}
