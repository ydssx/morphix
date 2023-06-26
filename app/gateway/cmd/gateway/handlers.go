package main

import (
	"context"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func healthzServer(conn *grpc.ClientConn) http.HandlerFunc {
	cli := grpc_health_v1.NewHealthClient(conn)
	res, _ := cli.Check(context.TODO(), &grpc_health_v1.HealthCheckRequest{})
	res.GetStatus().String()
	return nil
}
