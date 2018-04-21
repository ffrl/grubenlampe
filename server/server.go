package server

import (
	"context"

	pb "github.com/ffrl/grubenlampe/api"
	"github.com/ffrl/grubenlampe/database"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
}

// New creates a new API server instance
func New(db *database.Connection) *grpc.Server {
	a := &Server{}
	auth := auth{db}
	s := grpc.NewServer(grpc.StreamInterceptor(auth.streamInterceptor), grpc.UnaryInterceptor(auth.unaryInterceptor))
	pb.RegisterGrubenlampeServer(s, a)
	reflection.Register(s)
	return s
}

func (s *Server) AddUser(context.Context, *pb.AddUserRequest) (*pb.GenericResponse, error) {
	return nil, nil
}

func (s *Server) AddOrg(context.Context, *pb.AddOrgRequest) (*pb.GenericResponse, error) {
	return nil, nil
}

func (s *Server) AddASN(context.Context, *pb.AddASNRequest) (*pb.GenericResponse, error) {
	return nil, nil
}

func (s *Server) AddTunnel(context.Context, *pb.AddTunnelRequest) (*pb.GenericResponse, error) {
	return nil, nil
}

func (s *Server) DeleteTunnel(context.Context, *pb.DeleteTunnelRequest) (*pb.GenericResponse, error) {
	return nil, nil
}

func (s *Server) AddIPv4Address(context.Context, *pb.AddIPv4AddressRequest) (*pb.AddIPv4AddressResponse, error) {
	return nil, nil
}

func (s *Server) ReleaseIPv4Address(context.Context, *pb.ReleaseIPv4AddressRequest) (*pb.GenericResponse, error) {
	return nil, nil
}

func (s *Server) AddIPv6Prefix(context.Context, *pb.AddIPv6PrefixRequest) (*pb.AddIPv6PrefixReply, error) {
	return nil, nil
}

func (s *Server) ReleaseIPv6Prefix(context.Context, *pb.ReleaseIPv6PrefixRequest) (*pb.GenericResponse, error) {
	return nil, nil
}

func (s *Server) GetBGPStatus(context.Context, *pb.GetBGPStatusRequest) (*pb.GetBGPStatusReply, error) {
	return nil, nil
}

func (s *Server) GetLogs(context.Context, *pb.GetLogsRequest) (*pb.GetLogsReply, error) {
	return nil, nil
}
