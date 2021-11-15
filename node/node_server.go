package node

import (
	"context"
	"log"

	"github.com/colinxia/lrucache/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RpcCacheServer struct {
	nodePool *NODEPool
	pb.UnimplementedCacheServiceServer
}

func NewRpcCacheServer(nodePool *NODEPool) *RpcCacheServer {
	return &RpcCacheServer{nodePool, pb.UnimplementedCacheServiceServer{}}
}

func (server *RpcCacheServer) FindAllCacheNode(
	ctx context.Context,
	req *pb.RegCacheRequest,
) (*pb.NodesAdders, error) {
	if err := contextError(ctx); err != nil {
		return nil, err
	}
	nodes := server.nodePool.nodes
	var Address []string
	for node, _ := range nodes {
		Address = append(Address, node)
	}

	res := &pb.NodesAdders{
		NodesAdders: Address,
	}
	return res, nil
}

func (server *RpcCacheServer) RegCacheNode(
	ctx context.Context,
	req *pb.RegCacheRequest,
) (*pb.RegCacheResponse, error) {
	address := req.GetBaseURL()
	if err := contextError(ctx); err != nil {
		return nil, err
	}
	server.nodePool.nodes[address] = &RegBaseUrl{BaseURL: address}
	server.nodePool.node_hash.Add(address)
	log.Printf("接收注册通知成功 地址为:%s", address)
	res := &pb.RegCacheResponse{
		Status: 1,
	}
	return res, nil

}

func (server *RpcCacheServer) GetCache(
	ctx context.Context,
	req *pb.Key,
) (*pb.Cache, error) {
	key := req.GetKey()
	val, err := server.nodePool.cache.Get(key)
	if err != nil {
		return nil, logError(status.Error(codes.NotFound, "该节点未设置此key"))
	}
	return val, nil
}

func (server *RpcCacheServer) SetCache(
	ctx context.Context,
	req *pb.Cache,
) (*pb.Key, error) {
	// 偷下懒
	server.nodePool.cache.Add(req.Key, req)
	res := &pb.Key{
		Key: req.Key,
	}
	return res, nil
}

func contextError(ctx context.Context) error {
	switch ctx.Err() {
	case context.Canceled:
		return logError(status.Error(codes.Canceled, "请求断开"))
	case context.DeadlineExceeded:
		return logError(status.Error(codes.DeadlineExceeded, "请求超时"))
	default:
		return nil
	}
}

func logError(err error) error {
	if err != nil {
		log.Print(err)
	}
	return err
}
