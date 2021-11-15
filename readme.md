## 一 开发环境配置
- ubuntu
  Go 1.13 及以上（推荐）,使用 Go Module
   ```
   $ go env -w GO111MODULE=on
   $ go env -w GOPROXY=https://goproxy.cn,direct
   ```

- protobuf编译器和插件
   ```
   $ apt get install -y protobuf-compiler 
   $ protoc --version  # Ensure compiler version is 3+

   $ go get google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
   $ go get google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1

   $ export PATH="$PATH:$(go env GOPATH)/bin"
      //临时环境变量
   ```
   //注: 安装插件可能会失败，由于环境、版本的不同，查阅官网即可 [gRPC](https://grpc.io/docs/languages/go/quickstart/)

   - 生成golang源码
   ```
   $ protoc --go_out=. --go-grpc_out=. proto/*.proto
   //或
   $ make gen
   ```

## 二 架构及流程
 ### 2.1 点对点分布式模型
 * 本项目架构采用点对节模型，没有中心节点，每个节点都有维护一个节点池，可以理解为每个节点都是中心节点，节点之间通过心跳检测来维护节点池信息。新增一个节点可以在任意运行中的节点完成注册。
 * 缓存淘汰策略-使用与Mysql缓冲池同样的LRU算法实现
 * 采用一致性哈希算法处理节点的选择。
   > 步骤：一致性哈希算法，首尾相连的一个取值范围在2^32的圆环上，形成一个哈希圆环。
   假如：
   节点A地址为194.14.5.23:9090的哈希值为2000 
   节点B地址为173.10.22.6:8080的哈希值为5000
   那么映射到A节点的哈希值范围为2000 =< hash(key) < 5000
   映射到B节点的哈希值范围为hash(key) >= 5000
   如果要从节点获取key="xia"的值，假如hash(key)=3333,那么请求就会命中A节点来处理。

 ### 2.2  调用流程

   * 客户端调用很简单，最多一次rpc远程调用。
   ```

        正好命中节点直接返回|
                        否|      
   |--------key------->是否需rpc
            |           是|
            |             |rpc远程调用->返回缓存值
   一致性哈希选择节点            |未查找到
                               |
                        调用回调函数获取值并缓存-返回缓存值
   ```
## 三 如何使用

写本项目初衷 **`仅学习交流`** 练手用，感兴趣朋友可任意修改、完善。编码有些随意，内有详细中文注释。
集成到项目中或单独运行都可，写有两种示例，直接运行即可，见[examples](./examples)


