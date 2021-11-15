package test

import (
	"testing"

	"github.com/colinxia/lrucache/node"
)

func TestHashNode(t *testing.T) {
	consistent := node.NewMap(5, nil)
	nodes := []string{"127.0.0.1:9001", "127.0.0.1:9002", "127.0.0.1:9003"}
	consistent.Add(nodes...)
	k, n := consistent.GetNode()
	if k != len(nodes)*5 {
		t.Error("创建节点失败")
	}
	t.Log(n)
	key1, key2 := "key1key1", "key2key2"
	url1 := consistent.Get(key1)
	url2 := consistent.Get(key2)
	t.Log(key1, "命中节点:", url1, "\n")
	t.Log(key2, "命中节点:", url2, "\n")

	//移除节点
	consistent.Remove("127.0.0.1:9002")
	k1, n1 := consistent.GetNode()
	if k1 != 10 {
		t.Error("节点移除错误")
	}
	t.Log(n1)
}
