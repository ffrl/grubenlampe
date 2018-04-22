package server

import (
	"context"
	"fmt"

	pb "github.com/ffrl/grubenlampe/api"
	"github.com/ffrl/grubenlampe/database"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	db *database.Connection
}

// New creates a new API server instance
func New(db *database.Connection) *grpc.Server {
	a := &Server{db}
	auth := auth{db}
	s := grpc.NewServer(grpc.StreamInterceptor(auth.streamInterceptor), grpc.UnaryInterceptor(auth.unaryInterceptor))
	pb.RegisterGrubenlampeServer(s, a)
	reflection.Register(s)
	return s
}

// AddUser creates a user
func (s *Server) AddUser(ctx context.Context, req *pb.AddUserRequest) (*pb.GenericResponse, error) {
	users := s.db.Users()
	exists, err := users.EmailExists(req.Email)
	if err != nil {
		return nil, fmt.Errorf("error while processing")
	}
	if exists {
		return &pb.GenericResponse{Message: "User already exists"}, nil
	}

	u := &database.User{
		Email:      req.Email,
		Password:   req.Password,
		RIPEHandle: req.RipeHandle,
	}
	err = users.Save(u)
	if err != nil {
		return nil, fmt.Errorf("could not store user")
	}

	return &pb.GenericResponse{Success: true}, nil
}

func (s *Server) AddOrg(ctx context.Context, req *pb.AddOrgRequest) (*pb.GenericResponse, error) {
	orgs := s.db.Orgs()
	exists, err := orgs.ShortNameExists(req.ShortName)
	if err != nil {
		return nil, fmt.Errorf("error while processing")
	}
	if exists {
		return &pb.GenericResponse{Message: "Org already exists"}, nil
	}

	o := &database.Org{
		Name:      req.Name,
		ShortName: req.ShortName,
	}
	err = orgs.Save(o)
	if err != nil {
		return nil, fmt.Errorf("could not store org")
	}

	return &pb.GenericResponse{Success: true}, nil
}

func (s *Server) AddASN(ctx context.Context, req *pb.AddASNRequest) (*pb.GenericResponse, error) {
	orgs := s.db.Orgs()
	o, err := orgs.GetByShortName(req.OrgShortName)
	if err != nil {
		return nil, fmt.Errorf("error while processing")
	}

	asns := s.db.ASNs()
	exists, err := asns.CheckedASNExists(req.Asn)
	if err != nil {
		return nil, fmt.Errorf("error while processing")
	}
	if exists {
		return &pb.GenericResponse{Message: "ASN already exists"}, nil
	}

	a := &database.ASN{
		ASN: req.Asn,
		Org: o,
	}
	err = asns.Save(a)
	if err != nil {
		return nil, fmt.Errorf("could not store ASN")
	}

	return &pb.GenericResponse{Success: true}, nil
}

func (s *Server) AddTunnel(context.Context, *pb.AddTunnelRequest) (*pb.GenericResponse, error) {
	return &pb.GenericResponse{Success: false, Message: "not implemented"}, nil
}

func (s *Server) DeleteTunnel(context.Context, *pb.DeleteTunnelRequest) (*pb.GenericResponse, error) {
	return &pb.GenericResponse{Success: false, Message: "not implemented"}, nil
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
