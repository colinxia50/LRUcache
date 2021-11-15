package test

import (
	"testing"

	"github.com/colinxia/lrucache/cache"
	"github.com/colinxia/lrucache/node"
	"github.com/colinxia/lrucache/pb"
)

func TestCacheGet(t *testing.T) {
	cache := cache.NewCacheStore(2<<10, nil)
	data := &pb.Cache{
		Key:   "key1",
		Value: []byte("123456"),
	}
	cache.Add("key1", data)

	if v, ok := cache.Get("key1"); ok != nil || string(v.Value) != "123456" {
		t.Fatalf("cache hit key1=1234 failed")
	}
	if _, ok := cache.Get("key2"); ok != nil {
		t.Fatalf("cache miss key2 failed")
	}
}

type Value int

func (v Value) String() string {
	return v.String()
}

func TestNodePool(t *testing.T) {
	cacheServer := node.RunNode("127.0.0.1:9090", nil)
	//注册节点信息 地址空表示单体应用
	cacheServer.Reg()

	key, err := cacheServer.Set("key1", Value(1234))
	if err != nil {
		t.Error(err)
	}
	t.Log(key)

}
