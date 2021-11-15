package main

import (
	"context"
	"fmt"

	"github.com/colinxia/lrucache/node"
)

type Value int

func (v Value) String() string {
	return fmt.Sprintf("xia%d", v)
}
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cacheServer := node.RunNode("127.0.0.1:9090", ctx)
	//注册节点信息 地址空表示单体应用
	//当作新加节点就传入任意在线远程节点地址就可
	cacheServer.Reg()

	key, err := cacheServer.Set("keyhhhhh", Value(1234))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(key)
	<-ctx.Done()
	fmt.Println("服务停止")
}
