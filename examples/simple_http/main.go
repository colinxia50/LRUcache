package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/colinxia50/LRUcache/node"
)

var (
	ip            string
	port, regport int
)

func init() {
	flag.StringVar(&ip, "ip", "127.0.0.1", "输入ip")
	flag.IntVar(&port, "port", 8080, "输入端口")
	flag.IntVar(&regport, "regport", 0, "输入已运行节点端口")
}

func main() {

	flag.Parse()
	//输入的端口都是http服务的
	//rpc服务端口用的是http服务端口减一百。。。。
	rpcAddr := fmt.Sprintf("%s:%d", ip, port-100)
	httpAddr := fmt.Sprintf("%s:%d", ip, port)
	regAddr := ""
	if regport > 0 {
		//如果是在docker上运行或是在不同机器上运行 那就得输入完整注册ip 就不能只是端口号了
		regAddr = fmt.Sprintf("%s:%d", ip, regport-100)
	}

	cacheServer, ctx := node.RunNode(rpcAddr, context.Background())
	//注册节点信息 地址空表示单体应用
	//当作新加节点就传入任意在线远程节点地址就可
	cacheServer.Reg(regAddr)
	_ = cacheServer.Health(time.Second * 10)

	ctxx, cancel := context.WithCancel(ctx)
	go func() {
		log.Println(http.ListenAndServe(httpAddr, cacheServer))
		cancel()
	}()

	<-ctxx.Done()
}
