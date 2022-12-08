package server

import (
	pb "eventsie/pb/auth"
)

type Server struct {
	pb.UnimplementedAuthServer
}
