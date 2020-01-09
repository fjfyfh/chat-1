// Package chat 处理socket请求
package chat

import (
	"context"
	"encoding/json"
	"log"
)

// 默认适配器 全局存储
var socketHandler EventInterface

// EventInterface 事件接口
type EventInterface interface {
	RegisterEvent(cli *Client) error
	DestroyEvent(cli *Client) error
	HeartBeatEvent(msg Message, cli *Client) error
	BroadcastEvent(msg Message, cli *Client) error
	DefaultMessageEvent(MessageType int, msg Message, cli *Client) error
	Context() context.Context
	Init()
	Close()
	Status() map[string]interface{}
}

func init() {
	SetHandler(&DefaultHandler{})
}

// SetHandler 设置业务处理器
func SetHandler(adp EventInterface) {
	// 目前只能有一个去处理业务
	if socketHandler != nil {
		socketHandler.Close()
	}
	socketHandler = adp
	socketHandler.Init()
}

// GetHandler 获取业务处理器
func GetHandler() EventInterface {
	return socketHandler
}

// HandleRequest 处理
func HandleRequest(cli *Client, msg []byte) (err error) {
	log.Printf("获取信息: %s \n", msg)

	var msgBody Message
	err = json.Unmarshal(msg, &msgBody)

	if err != nil {
		return
	}

	// log.Printf("MessageBody: %#+v \n", msgBody)

	switch msgBody.Type {
	case BroadcastMessage, SystemMessage:
		socketHandler.BroadcastEvent(msgBody, cli)
	case HeartBeatMessage:
		socketHandler.HeartBeatEvent(msgBody, cli)
	default:
		// 自定义的type就层层套娃一下
		socketHandler.DefaultMessageEvent(msgBody.Type, msgBody, cli)
	}
	return
}
