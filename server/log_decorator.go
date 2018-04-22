package server

import (
	"context"
	"encoding/json"

	pb "github.com/ffrl/grubenlampe/api"
	"github.com/ffrl/grubenlampe/database"
)

func WithLogging(db *database.Connection, service pb.GrubenlampeServer) pb.GrubenlampeServer {
	return &serverLogDecorator{db, service}
}

type serverLogDecorator struct {
	db      *database.Connection
	service pb.GrubenlampeServer
}

// AddUser creates a user
func (s *serverLogDecorator) AddUser(ctx context.Context, req *pb.AddUserRequest) (*pb.GenericResponse, error) {
	res, err := s.service.AddUser(ctx, req)
	s.logRequest(ctx, req, res, err)

	return res, err
}

// AddOrg creates an organization
func (s *serverLogDecorator) AddOrg(ctx context.Context, req *pb.AddOrgRequest) (*pb.GenericResponse, error) {
	res, err := s.service.AddOrg(ctx, req)
	s.logRequest(ctx, req, res, err)

	return res, err
}

// AddASN creates an Autonomous System Number
func (s *serverLogDecorator) AddASN(ctx context.Context, req *pb.AddASNRequest) (*pb.GenericResponse, error) {
	res, err := s.service.AddASN(ctx, req)
	s.logRequest(ctx, req, res, err)

	return res, err
}

// AddTunnel crates a tunnel
func (s *serverLogDecorator) AddTunnel(ctx context.Context, req *pb.AddTunnelRequest) (*pb.GenericResponse, error) {
	res, err := s.service.AddTunnel(ctx, req)
	s.logRequest(ctx, req, res, err)

	return res, err
}

// DeleteTunnel deletes a tunnel
func (s *serverLogDecorator) DeleteTunnel(ctx context.Context, req *pb.DeleteTunnelRequest) (*pb.GenericResponse, error) {
	res, err := s.service.DeleteTunnel(ctx, req)
	s.logRequest(ctx, req, res, err)

	return res, err
}

// AddIPv4Address creates an IPv4 address
func (s *serverLogDecorator) AddIPv4Address(ctx context.Context, req *pb.AddIPv4AddressRequest) (*pb.AddIPv4AddressResponse, error) {
	res, err := s.service.AddIPv4Address(ctx, req)
	s.logRequest(ctx, req, res, err)

	return res, err
}

// ReleaseIPv4Address releases an IPv4 address
func (s *serverLogDecorator) ReleaseIPv4Address(ctx context.Context, req *pb.ReleaseIPv4AddressRequest) (*pb.GenericResponse, error) {
	res, err := s.service.ReleaseIPv4Address(ctx, req)
	s.logRequest(ctx, req, res, err)

	return res, err
}

// AddIPv6Prefix creates an IPv6 /48 prefix
func (s *serverLogDecorator) AddIPv6Prefix(ctx context.Context, req *pb.AddIPv6PrefixRequest) (*pb.AddIPv6PrefixResponse, error) {
	res, err := s.service.AddIPv6Prefix(ctx, req)
	s.logRequest(ctx, req, res, err)

	return res, err
}

// ReleaseIPv6Prefix releases an IPv6 prefix
func (s *serverLogDecorator) ReleaseIPv6Prefix(ctx context.Context, req *pb.ReleaseIPv6PrefixRequest) (*pb.GenericResponse, error) {
	res, err := s.service.ReleaseIPv6Prefix(ctx, req)
	s.logRequest(ctx, req, res, err)

	return res, err
}

// GetBGPStatus gets status of BGP sessions
func (s *serverLogDecorator) GetBGPStatus(ctx context.Context, req *pb.GetBGPStatusRequest) (*pb.GetBGPStatusResponse, error) {
	res, err := s.service.GetBGPStatus(ctx, req)
	s.logRequest(ctx, req, res, err)

	return res, err
}

// GetLogs gets Logs
func (s *serverLogDecorator) GetLogs(ctx context.Context, req *pb.GetLogsRequest) (*pb.GetLogsResponse, error) {
	res, err := s.service.GetLogs(ctx, req)
	s.logRequest(ctx, req, res, err)

	return res, err
}

// GetTunnels retrieves tunnels for an org
func (s *serverLogDecorator) GetTunnels(ctx context.Context, req *pb.GetTunnelsRequest) (*pb.GetTunnelsResponse, error) {
	res, err := s.service.GetTunnels(ctx, req)
	s.logRequest(ctx, req, res, err)

	return res, err
}

func (s *serverLogDecorator) logRequest(ctx context.Context, req, res interface{}, e error) {
	errorMessage := ""
	if e != nil {
		errorMessage = e.Error()
	}

	l := &database.Log{
		User:            s.userFromContext(ctx),
		Error:           errorMessage,
		RequestMessage:  s.marshalToJSON(req),
		ResponseMessage: s.marshalToJSON(res),
	}
	err := s.db.Logs().Insert(l)
	if err != nil {
		// TODO: handle log failiure
	}
}

func (*serverLogDecorator) userFromContext(ctx context.Context) *database.User {
	u, ok := ctx.Value(ctxUserKey{}).(*database.User)
	if !ok {
		return nil
	}

	return u
}

func (*serverLogDecorator) marshalToJSON(x interface{}) string {
	if x == nil {
		return ""
	}

	b, err := json.Marshal(x)
	if err != nil {
		return ""
	}

	return string(b)
}
