package node

import (
	"context"
	"log"

	"github.com/colinxia/lrucache/pb"
	"google.golang.org/grpc"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

type RegBaseUrl struct {
	BaseURL string
}

func (rb *RegBaseUrl) GetRpcValue(key string) (*pb.Cache, error) {
	RpcClient := NewClien(rb.BaseURL)
	return RpcClient.GetCacheValue(key)

}

func (rb *RegBaseUrl) SetRpcValue(key string, value *pb.Cache) (*pb.Key, error) {
	RpcClient := NewClien(rb.BaseURL)
	return RpcClient.SetCacheValue(key, value)
}

func (rb *RegBaseUrl) HealthCheek() error {
	ctx := context.Context(context.Background())
	conn, _ := grpc.Dial(rb.BaseURL, grpc.WithInsecure())
	client := healthpb.NewHealthClient(conn)
	req := &healthpb.HealthCheckRequest{
		Service: "grpc.health",
	}
	res, err := client.Check(ctx, req)
	log.Println(rb.BaseURL, res.GetStatus())
	return err
}
