package node

import (
	"context"
	"log"
	"time"

	"github.com/colinxia/lrucache/pb"
	"google.golang.org/grpc"
)

type rpcClient struct {
	serverAddres string
	client       pb.CacheServiceClient
}

func NewClien(serverAddres string) *rpcClient {
	conn, err := grpc.Dial(serverAddres, grpc.WithInsecure())
	client := pb.NewCacheServiceClient(conn)
	if err != nil {
		log.Fatal("客户端请求服务端失败 ", err)
	}
	return &rpcClient{serverAddres, client}
}

func (r *rpcClient) findAllCacheNode() (*pb.NodesAdders, error) {
	req := &pb.RegCacheRequest{
		BaseURL: r.serverAddres,
	}
	cxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := r.client.FindAllCacheNode(cxt, req)
	if err != nil {
		log.Fatal("客户端发送搜索请求失败", err)
	}
	return res, err
}

func (r *rpcClient) regCacheNode(self string) error {
	req := &pb.RegCacheRequest{
		BaseURL: self,
	}

	cxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := r.client.RegCacheNode(cxt, req)
	if err != nil {
		log.Fatal("客户端注册通知失败", err)
	}
	return err
}

func (r *rpcClient) GetCacheValue(key string) (*pb.Cache, error) {
	req := &pb.Key{
		Key: key,
	}
	//超时测试
	cxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := r.client.GetCache(cxt, req)
	if err != nil {
		log.Fatal("获取远程缓存数据错误", err)
	}
	return res, err
}

func (r *rpcClient) SetCacheValue(key string, value *pb.Cache) (*pb.Key, error) {
	//超时测试
	cxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := r.client.SetCache(cxt, value)
	if err != nil {
		log.Fatal("保存缓存数据错误", err)
	}
	return res, err
}
