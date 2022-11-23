package server

import (
	"context"
	"events/events/models"
	pb "events/pb/events"
	"net/http"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

type Server struct {
	pb.UnimplementedEventsServer
}

func (s *Server) FindOne(ctx context.Context, in *pb.FindOneRequest) (*pb.FindOneResponse, error) {
	event := &models.Event{}
	if err := mgm.Coll(event).FindByID(in.Id, event); err != nil {
		return &pb.FindOneResponse{
			Status: http.StatusInternalServerError,
			Error:  "Could not query the database.",
		}, nil
	}

	return &pb.FindOneResponse{
		Status: 200,
		Event:  models.EventToProto(event),
	}, nil
}

func (s *Server) FindMany(ctx context.Context, in *pb.FindManyRequest) (*pb.FindManyResponse, error) {
	events := []*models.Event{}

	filter := bson.M{}

	if in.Category != nil {
		filter["category"] = in.Category
	}

	// TODO: Figure out why greater than and less than operators dont work here
	// if in.Price != nil && (in.Price.From > 0 || in.Price.To > 0) {
	// 	filter["price"] = bson.M{}
	// }
	// if in.Price != nil {
	// 	if in.Price.From > 0 {
	// 		filter["price"].(bson.M)["from"] = bson.M{operator.Gte: in.Price.From}
	// 	}
	// 	if in.Price.To > 0 {
	// 		filter["price"].(bson.M)["to"] = bson.M{operator.Lte: in.Price.To}
	// 	}
	// }

	if err := mgm.Coll(&models.Event{}).SimpleFind(&events, filter); err != nil {
		return &pb.FindManyResponse{
			Status: http.StatusInternalServerError,
			Error:  "Could not query the database.",
		}, nil
	}

	return &pb.FindManyResponse{
		Status: 200,
		Events: models.EventsToProto(events),
	}, nil
}

func (s *Server) Add(ctx context.Context, in *pb.AddRequest) (*pb.AddResponse, error) {
	event := models.EventFromProto(in.Event)
	if err := mgm.Coll(event).Create(event); err != nil {
		return &pb.AddResponse{
			Status: http.StatusInternalServerError,
			Error:  "Unexpected error while trying to create an event.",
		}, err
	}

	return &pb.AddResponse{
		Status: http.StatusOK,
	}, nil
}