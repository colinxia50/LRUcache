package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/colinxia/lrucache/node"
)

type Value struct {
	value []string
}

func (v Value) String() string {
	return fmt.Sprintf("%v", v.value)
}

type Srt string

func (s Srt) String() string {
	return string(s)
}

func main() {

	var (
		url      string
		from_url string
		choose   string
		health   bool
	)

	fmt.Println("--输入启动节点地址(如127.0.0.1:9001):")
	fmt.Scan(&url)

	fmt.Println("--输入已启动节点地址:(用于服务注册,初始节点键入false):")
	fmt.Scan(&from_url)

	fmt.Printf("----当前节点 %s-----注册节点 %s\n", url, from_url)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cacheServer := node.RunNode(url, ctx)
	if ctx.Err() != nil {
		log.Fatal("节点启动失败", ctx.Err())
	}
	//注册节点信息 地址空表示单体应用
	if from_url == "false" {
		cacheServer.Reg()
	} else {
		cacheServer.Reg(from_url)
	}

	//不能停止健康检查 这里只是为了查看效果
	h := cacheServer.Health(time.Second * 2)
	h.Stop()

	for {
		fmt.Println("----查找或设置缓存-----当前节点", url)
		fmt.Println("	输入 1 查找key")
		fmt.Println("	输入 2 设置key")
		fmt.Println("	输入 3 开启心跳健康检测")
		fmt.Println("	输入 4 退出当前节点")
		fmt.Scan(&choose)
		key, value := "", ""
		switch choose {
		case "1":
			fmt.Println("-----输入要查询的key------")
			fmt.Scan(&key)
			v, err := cacheServer.Get(key, func(key string) node.Value {
				//当取值失败，取其它数据源数据
				s := fmt.Sprintf("%s回调的值", key)
				return Srt(s)
			})
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("查询成功 key:%s--value:%s \n", v.GetKey(), v.GetValue())
		case "2":
			fmt.Println("----设置缓存----")
			fmt.Println("输入key:")
			fmt.Scan(&key)
			fmt.Println("输入value:")
			fmt.Scan(&value)
			//k, err := cacheServer.Set(key, Value{value: value})
			k, err := cacheServer.Set(key, Srt(value))
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("设置成功 key:%s\n", k.GetKey())
		case "3":
			fmt.Println("输true=>开启 false=>关闭")
			fmt.Scan(&health)
			if health {
				h.Reset(time.Second * 2)
			} else {
				h.Stop()
			}

			fmt.Printf("开启成功")
		case "4":
			fmt.Println("----已退出----")
			//os.Exit(0)
			return
		}
	}
}
