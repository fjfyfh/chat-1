### 一个简单的聊天室


#### 依赖文件

```
go get github.com/gorilla/websocket
```

运行: 

main.go
```go
package main

import (
	"chat"
)

func main() {
	chat.ServerStar()
}
```

```bash
$ go run main.go  默认地址 localhost:8080
//或
$ go run main.go -addr 127.0.0.1:8080
```
监听 `ws://localhost:8080/ws`

#### 一个具体例子

```go
// 启动服务端
go run example/server.go

// 启动客户端
go run example/client.go
```
客户端之间消息互相广播


##### TODO

- [x] 增加事件钩子, 业务解耦
- [ ] 结构化消息格式 
- [ ] 日志
- [ ] 消息持久化
- [ ] conn 和 user_id 的关联绑定
- [ ] 指定范围的广播(群组)
- [ ] socket转http
- [ ] 历史消息/限流
- [ ] 闲时心跳检测


##### 日志

* 2020/01/09

将消息分发逻辑拆分出来, 目前是通过chan来进行消息互通
优点:
将函数调用变成事件触发, 以此去解耦操作, 后续如果需要改变消息分发方式, 改动的代码会较少
缺点:
改动的地方不少, 估计后续开发会因为这个变动多一些代码量~~但是值~~


