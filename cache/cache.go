package cache

import (
	"container/list"
	"sync"

	"github.com/colinxia50/LRUcache/pb"
)

type Cache struct {
	mutex     sync.RWMutex
	maxBytes  int64
	usedBytes int64
	ll        *list.List
	cache     map[string]*list.Element
	callback  func(key string, value *pb.Cache)
}

func NewCacheStore(maxBytes int64, callback func(string, *pb.Cache)) *Cache {
	return &Cache{
		maxBytes: maxBytes,
		ll:       list.New(),
		cache:    make(map[string]*list.Element),
		callback: callback,
	}
}

func (c *Cache) Get(key string) (value *pb.Cache, ok error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*pb.Cache)
		return kv, nil
	}
	return
}

func (c *Cache) RemoveOldest() {
	ele := c.ll.Back()
	if ele != nil {
		c.ll.Remove(ele)
		kv := ele.Value.(*pb.Cache)
		delete(c.cache, kv.Key)
		c.usedBytes -= int64(len(kv.Key)) + int64(len(kv.Value))
		if c.callback != nil {
			c.callback(kv.Key, kv)
		}
	}
}

func (c *Cache) Add(key string, value *pb.Cache) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*pb.Cache)
		c.usedBytes += int64(len(value.Value)) - int64(len(kv.Value))
		kv.Value = value.Value
	} else {
		ele := c.ll.PushFront(value)
		c.cache[key] = ele
		c.usedBytes += int64(len(key)) + int64(len(value.Value))
	}
	for c.maxBytes != 0 && c.maxBytes < c.usedBytes {
		c.RemoveOldest()
	}
}
