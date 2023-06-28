package main

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
)

type ServerInfo struct {
	Name   string
	Status string
	Msg    string
}

func healthzServer(conn *grpc.ClientConn) (r ServerInfo) {
	cli := grpc_health_v1.NewHealthClient(conn)
	res, err := cli.Check(context.TODO(), &grpc_health_v1.HealthCheckRequest{Service: ""})
	if err != nil {
		r.Msg = status.Convert(err).Message()
	}
	r.Status = res.GetStatus().String()
	return
}
