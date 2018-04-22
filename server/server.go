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
		return nil, fmt.Errorf("Error while processing: %v", err)
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
		return nil, fmt.Errorf("Could not store user: %v", err)
	}

	return &pb.GenericResponse{Success: true}, nil
}

// AddOrg creates an organization
func (s *Server) AddOrg(ctx context.Context, req *pb.AddOrgRequest) (*pb.GenericResponse, error) {
	orgs := s.db.Orgs()
	exists, err := orgs.ShortNameExists(req.ShortName)
	if err != nil {
		return nil, fmt.Errorf("Error while processing: %v", err)
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
		return nil, fmt.Errorf("Could not store org: %v", err)
	}

	return &pb.GenericResponse{Success: true}, nil
}

// AddASN creates an Autonomous System Number
func (s *Server) AddASN(ctx context.Context, req *pb.AddASNRequest) (*pb.GenericResponse, error) {
	orgs := s.db.Orgs()
	o, err := orgs.GetByShortName(req.OrgShortName)
	if err != nil {
		return nil, fmt.Errorf("Error while processing: %s", err)
	}

	asns := s.db.ASNs()
	exists, err := asns.CheckedASNExists(req.Asn)
	if err != nil {
		return nil, fmt.Errorf("Error while processing: %v", err)
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
		return nil, fmt.Errorf("Could not store ASN: %v", err)
	}

	return &pb.GenericResponse{Success: true}, nil
}

// AddTunnel crates a tunnel
func (s *Server) AddTunnel(ctx context.Context, req *pb.AddTunnelRequest) (*pb.GenericResponse, error) {
	res, err := s.validateAddTunnel(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("Validation failed: %v", err)
	}

	if res != nil {
		return res, nil
	}

	err = s.db.Tunnels().AddTunnel(req.Asn, req.Address)
	if err != nil {
		return nil, fmt.Errorf("Unable to add tunnel")
	}

	return &pb.GenericResponse{Success: true}, nil
}

func (s *Server) validateAddTunnel(ctx context.Context, req *pb.AddTunnelRequest) (*pb.GenericResponse, error) {
	tunnels, err := s.db.Tunnels().GetTunnelsByAddress(req.Address)
	if err != nil {
		return nil, fmt.Errorf("Unable to get tunnels by address: %v", err)
	}

	if len(tunnels) != 0 {
		return &pb.GenericResponse{Message: "Tunnel already exists"}, nil
	}

	asn, err := s.db.ASNs().GetCheckedASN(req.Asn)
	if err != nil {
		return nil, fmt.Errorf("Unable to get ASN: %v", err)
	}

	if asn == nil {
		return &pb.GenericResponse{Message: "ASN not checked or does not exist"}, nil
	}

	user, ok := ctx.Value(ctxUserKey{}).(*database.User)
	if !ok {
		return nil, fmt.Errorf("Auth failure")
	}

	if user.HasOrg(asn.OrgID) {
		return &pb.GenericResponse{Message: "Authorization error"}, nil
	}

	return nil, nil
}

// DeleteTunnel deletes a tunnel
func (s *Server) DeleteTunnel(context.Context, *pb.DeleteTunnelRequest) (*pb.GenericResponse, error) {
	return &pb.GenericResponse{Success: false, Message: "Not implemented"}, nil
}

// AddIPv4Address creates an IPv4 address
func (s *Server) AddIPv4Address(context.Context, *pb.AddIPv4AddressRequest) (*pb.AddIPv4AddressResponse, error) {
	return nil, nil
}

// ReleaseIPv4Address releases an IPv4 address
func (s *Server) ReleaseIPv4Address(context.Context, *pb.ReleaseIPv4AddressRequest) (*pb.GenericResponse, error) {
	return nil, nil
}

// AddIPv6Prefix creates an IPv6 /48 prefix
func (s *Server) AddIPv6Prefix(context.Context, *pb.AddIPv6PrefixRequest) (*pb.AddIPv6PrefixReply, error) {
	return nil, nil
}

// ReleaseIPv6Prefix releases an IPv6 prefix
func (s *Server) ReleaseIPv6Prefix(context.Context, *pb.ReleaseIPv6PrefixRequest) (*pb.GenericResponse, error) {
	return nil, nil
}

// GetBGPStatus gets status of BGP sessions
func (s *Server) GetBGPStatus(context.Context, *pb.GetBGPStatusRequest) (*pb.GetBGPStatusReply, error) {
	return nil, nil
}

// GetLogs gets Logs
func (s *Server) GetLogs(context.Context, *pb.GetLogsRequest) (*pb.GetLogsReply, error) {
	return nil, nil
}
