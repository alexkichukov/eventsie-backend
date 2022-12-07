package server

import (
	"context"
	pb "eventsie/pb/events"
)

// Find one event by ID
func (s *Server) GetCategories(ctx context.Context, in *pb.GetCategoriesRequest) (*pb.GetCategoriesResponse, error) {
	return &pb.GetCategoriesResponse{
		Categories: []*pb.Category{
			{Id: "music", Name: "Music"},
			{Id: "health", Name: "Health"},
			{Id: "sports", Name: "Sports"},
			{Id: "hobbies", Name: "Hobbies"},
			{Id: "tech", Name: "Technology"},
			{Id: "food-and-drink", Name: "Food and Drink"},
		},
	}, nil
}
