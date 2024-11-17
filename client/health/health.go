package health

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health/grpc_health_v1"
)

type HealthChecker struct {
	client grpc_health_v1.HealthClient
	conn   *grpc.ClientConn
}

func NewHealthChecker(addr string) (*HealthChecker, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &HealthChecker{
		client: grpc_health_v1.NewHealthClient(conn),
		conn:   conn,
	}, nil
}

func (h *HealthChecker) Check(dur time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), dur)
	defer cancel()

	resp, err := h.client.Check(ctx, &grpc_health_v1.HealthCheckRequest{
		Service: "system",
	})
	if err != nil {
		return err
	}

	if resp.Status != grpc_health_v1.HealthCheckResponse_SERVING {
		return fmt.Errorf("service unhealthy: %v", resp.Status)
	}

	return nil
}

func (h *HealthChecker) Close() error {
	return h.conn.Close()
}