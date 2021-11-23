package node

import (
	"bytes"
	"fmt"
	"hash/crc32"
	"math/rand"
	"sort"
	"strconv"
	"sync"
	"time"
)

type Hash func(data []byte) uint32

// hash 自定义哈希函数 默认crc32.ChecksumIEEE
// replicas 虚拟节点个数
// keys 所有节点的哈希值
// hashMap 节点哈希值对应的真实节点地址
type Map struct {
	hash     Hash
	replicas int
	rw       sync.RWMutex
	keys     []int // Sorted
	hashMap  map[int]string
}

func NewMap(replicas int, fn Hash) *Map {
	m := &Map{
		replicas: replicas,
		hash:     fn,
		hashMap:  make(map[int]string),
	}
	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE
	}
	return m
}

/*
m.keys
[随机字符串+ip1,随机字符串+ip1, ...,随机字符串+ip2 ...]
将keys排序
----------------------
m.Map
Map[随机字符串+ip1]=ip1
Map[随机字符串+ip1]=ip1
	.
	.
Map[随机字符串+ip2]=ip2
Map[随机字符串+ip2]=ip2
	.
	.

*/
func (m *Map) Add(keys ...string) {
	m.rw.Lock()
	defer m.rw.Unlock()
	for _, key := range keys {
		for i := 0; i < m.replicas; i++ {
			//拼接随机字符串再哈希 使虚拟节点分布均匀
			//...真的蠢 这里不能用随机 又不是轮轴式架构
			//hash := int(m.hash([]byte(RandString(10) + key)))
			hash := int(m.hash([]byte(strconv.Itoa(i+10) + "colinxia50" + key)))
			m.keys = append(m.keys, hash)
			m.hashMap[hash] = key
		}
	}
	sort.Ints(m.keys)
}

/*
1 假如要查找key="9999"的缓存 key哈希后值为100
2 假如keys里各节点哈希值是[88,97,105,110,....]
3 那么100就对应105的节点哈希值 得到m.Map[105]的真实节点地址
*/
func (m *Map) Get(key string) string {
	if len(m.keys) == 0 {
		return ""
	}

	hash := int(m.hash([]byte(key)))

	idx := sort.Search(len(m.keys), func(i int) bool {
		return m.keys[i] >= hash
	})

	return m.hashMap[m.keys[idx%len(m.keys)]]
}

/*
一致性哈希算法虽然降低节点数量变化的影响范围
但依然有极大的缓存雪崩的风险
*/
// 	移除节点...
func (m *Map) Remove(key string) {
	for k, v := range m.hashMap {
		if v == key {
			idx := sort.Search(len(m.keys), func(i int) bool {
				return m.keys[i] >= k
			})
			m.keys = append(m.keys[:idx], m.keys[idx+1:]...)
			delete(m.hashMap, k)
		}
	}
}

//用于测试。。。
// keys 节点长度
// hashMap 节点信息
func (m *Map) GetNode() (keys int, hashMap []string) {
	keys = len(m.keys)
	for k, v := range m.hashMap {
		hashMap = append(hashMap, fmt.Sprintf("节点哈希值:%d,节点:%s\n", k, v))
	}

	return keys, hashMap
}

func RandString(len int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var st bytes.Buffer
	for i := 0; i < len; i++ {
		st.WriteByte(byte((33 + r.Intn(126-33))))
	}
	return st.String()
}

// 利用goroutine抢占 生成随机字符串
// 一次会产生len*26*2个孤儿协程 很豪横 废.....
func RandStringByte(len int) string {
	var st bytes.Buffer
	bytes_chan := make(chan byte, len)
	for i := 0; i < len; i++ {
		go func() {
			byte_chan := make(chan int, 1)
			for i := 65; i < 91; i++ {
				go func(i int) {
					byte_chan <- i
				}(i)
				go func(i int) {
					byte_chan <- i + 32
				}(i)
			}
			bytes_chan <- byte(<-byte_chan)
		}()
	}
	for i := 0; i < len; i++ {
		st.WriteByte(<-bytes_chan)
	}
	return st.String()
}
