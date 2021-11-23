package test

import (
	"context"
	"testing"

	"github.com/colinxia50/LRUcache/cache"
	"github.com/colinxia50/LRUcache/node"
	"github.com/colinxia50/LRUcache/pb"
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

func TestNodePool(t *testing.T) {
	cacheServer, _ := node.RunNode("127.0.0.1:9090", context.Background())
	//注册节点信息 地址空表示单体应用
	cacheServer.Reg("")

	key, err := cacheServer.Set("key1", []byte("1234"))
	if err != nil {
		t.Error(err)
	}
	t.Log(key)

}
