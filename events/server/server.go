package server

import (
	pb "eventsie/pb/events"
)

type Server struct {
	pb.UnimplementedEventsServer
}
