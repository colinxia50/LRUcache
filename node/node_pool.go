package node

import (
	"errors"
	"log"
	"net"
	"sync"
	"time"

	"github.com/colinxia/lrucache/cache"
	"github.com/colinxia/lrucache/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

const (
	defaultBasePath = "/cache/"
	defaultReplicas = 50
)

/*
节点维护--因为是去中心化架构，所以这里需要维护节点信息一致性
1 启动本节点时需要和远程节点同步
2 先查询所有在运行的远程节点信息同步到本节点
3 将本节点信息传递给各远程节点
*/

// self 当前节点地址
// basePath
// node_hash 一致性哈希 计算哈希值
// nodes 节点池
type NODEPool struct {
	self      string
	basePath  string
	mu        sync.Mutex
	node_hash *Map
	cache     *cache.Cache
	nodes     map[string]*RegBaseUrl
}

// 启动节点
//lca 本地节点地址
func RunNode(lca string, ctx context.Context) *NODEPool {
	nodePool := NODEPool{
		self:     lca,
		basePath: defaultBasePath,
		cache:    cache.NewCacheStore(2<<10, nil),
		nodes:    make(map[string]*RegBaseUrl),
	}
	rpcServer := NewRpcCacheServer(&nodePool)
	grpcServer := grpc.NewServer()
	pb.RegisterCacheServiceServer(grpcServer, rpcServer)
	healthserver := health.NewServer()
	healthserver.SetServingStatus("grpc.health", healthpb.HealthCheckResponse_SERVING)
	healthpb.RegisterHealthServer(grpcServer, healthserver)
	listener, err := net.Listen("tcp", lca)
	if err != nil {
		log.Fatal("监听节点服务失败", err)
	}
	go func() {
		log.Println("服务启动")
		log.Println(grpcServer.Serve(listener))
		ctx.Done()
	}()
	return &nodePool
}

// 向远程节点注册当前节点信息 返回全部节点信息
// 将返回得到的全部节点信息注册到当前节点 保证节点信息同步
// fromNode 远程节点地址
func (n *NODEPool) Reg(fromNode ...string) {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.node_hash = NewMap(defaultReplicas, nil)
	if len(fromNode) > 0 {
		RpcClient := NewClien(fromNode[0])
		nodeAddress, err := RpcClient.findAllCacheNode()
		if err != nil {
			log.Fatal("注册节点-查询节点信息失败 ", err)
		}
		c := make(chan string, len(nodeAddress.NodesAdders))
		for _, url := range nodeAddress.NodesAdders {
			//因为有健康检查
			//不考虑注册通知时有节点崩溃这种巧合
			go func(url string) {
				RpcClient := NewClien(url)
				RpcClient.regCacheNode(n.self)
				//n.nodes[url] = &cache.RegBaseUrl{baseURL: url}
				n.nodes[url] = &RegBaseUrl{BaseURL: url}
				n.node_hash.Add(url)
				c <- url
			}(url)
		}
		for k, _ := range nodeAddress.NodesAdders {
			//只要远程节点存活 默认注册节点成功。。。 很粗暴 只管通知不管结果
			//严谨点这里其实可以根据返回状态处理的 这里忽略小概率情况
			//这里只打印去通知的服务节点地址
			log.Printf("%d-向节点:%s 通知！\n", k, <-c)
		}
	}
	n.nodes[n.self] = &RegBaseUrl{BaseURL: n.self}
	n.node_hash.Add(n.self)
}

func (n *NODEPool) Health(s time.Duration) *time.Timer {
	t := time.NewTimer(s)
	go func() {
		for {
			<-t.C
			for _, v := range n.nodes {
				go func(v *RegBaseUrl) {
					if err := v.HealthCheek(); err != nil {
						n.node_hash.Remove(v.BaseURL)
						delete(n.nodes, v.BaseURL)
					}
				}(v)
			}
			t.Reset(s)
		}
	}()

	return t
}

/*
 缓存设置->根据key选择对应节点->rpc请求到节点设置
*/

//定义类型key->value类型
type Value interface {
	String() string
}

type GetFunc func(key string) Value

func (n *NODEPool) Set(key string, value Value) (*pb.Key, error) {
	n.mu.Lock()
	defer n.mu.Unlock()
	data := &pb.Cache{
		Key:   key,
		Value: []byte(value.String()),
	}
	if node := n.node_hash.Get(key); node != "" {
		return n.nodes[node].SetRpcValue(key, data)
	} else if node == n.self {
		n.cache.Add(key, data)
	}
	return nil, errors.New("设置缓存失败")
}

func (n *NODEPool) Get(key string, f GetFunc) (*pb.Cache, error) {
	//n.mu.Lock()
	//defer n.mu.Unlock()
	if node := n.node_hash.Get(key); node != "" {
		if node == n.self {
			value, err := n.cache.Get(key)
			return value, err
		}
		value, err := n.nodes[node].GetRpcValue(key)
		if err != nil {
			if f != nil {
				//节点池改变(节点崩溃或新增节点)导致未获取到数据
				if node := n.node_hash.Get(key); node != "" {
					value := f(key)
					data := &pb.Cache{
						Key:   key,
						Value: []byte(value.String()),
					}
					n.nodes[node].SetRpcValue(key, data)
					return data, nil
				}
			}

		}
		return value, err
	}
	return nil, errors.New("无此key数据")
}
