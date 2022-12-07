package server

import (
	"context"

	"eventsie/auth/util"
	pb "eventsie/pb/auth"
)

// Check if JWT token is valid or not
func (s *Server) ValidateToken(ctx context.Context, in *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	parsed, err := util.ParseJWTToken(in.Token)

	if err != nil {
		return &pb.ValidateResponse{Valid: false}, nil
	}

	return &pb.ValidateResponse{
		Valid:     true,
		FirstName: parsed.FirstName,
		LastName:  parsed.LastName,
		Email:     parsed.Email,
		Id:        parsed.ID,
	}, nil
}
