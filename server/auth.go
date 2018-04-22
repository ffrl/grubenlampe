package server

import (
	"context"

	"github.com/ffrl/grubenlampe/database"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

type auth struct {
	db *database.Connection
}

func (a *auth) streamInterceptor(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	if err := a.authorize(stream.Context()); err != nil {
		return err
	}

	return handler(srv, stream)
}

func (a *auth) unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if err := a.authorize(ctx); err != nil {
		return nil, err
	}

	return handler(ctx, req)
}

func (a *auth) authorize(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return grpc.Errorf(codes.Unauthenticated, "missing context metadata")
	}

	if len(md["username"]) == 0 || len(md["password"]) == 0 {
		return grpc.Errorf(codes.Unauthenticated, "username and password required")
	}

	res, err := a.db.Users().GetByCredentials(md["username"][0], md["password"][0])
	if err != nil {
		return grpc.Errorf(codes.Internal, "error while authenticating")
	}

	if res == nil {
		return grpc.Errorf(codes.Unauthenticated, "access denied")
	}

	return nil
}
