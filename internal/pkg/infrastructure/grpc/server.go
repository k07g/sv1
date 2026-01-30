package grpc

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

func NewServer(ctx context.Context) *grpc.Server {
	server := grpc.NewServer(
		grpc.KeepaliveEnforcementPolicy(
			keepalive.EnforcementPolicy{
				MinTime:             10 * time.Second,
				PermitWithoutStream: true,
			},
		),
		grpc.KeepaliveParams(
			keepalive.ServerParameters{
				MaxConnectionIdle:     0,
				MaxConnectionAge:      10 * time.Minute,
				MaxConnectionAgeGrace: 0,
				Time:                  20 * time.Second,
				Timeout:               10 * time.Second,
			},
		),
	)
	return server
}
