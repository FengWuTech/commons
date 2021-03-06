package wss

import (
	"encoding/json"
	"github.com/FengWuTech/commons/gredis"
	"github.com/FengWuTech/commons/logger"
	"github.com/FengWuTech/commons/wss/impl"
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"time"
)

var Upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Message struct {
	CompanyID int         `json:"company_id"`
	Type      int         `json:"type"`
	GroupID   int         `json:"group_id"`
	Content   interface{} `json:"content"`
	NoticeID  int         `json:"notice_id"`
	BizID     int         `json:"biz_id"`
}

type ClientConnect struct {
	CompanyID int
	StaffID   int
	WsConn    *websocket.Conn
	Conn      *impl.Connection
}

type Channel struct {
	Upgrader       websocket.Upgrader
	RedisSubPubKey string
	ClientConnect  map[int]map[string]ClientConnect
	SendCheckFunc  func(staffID int, groupID int, message Message) bool
}

func Init(channelName string) *Channel {
	return &Channel{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		RedisSubPubKey: channelName,
		ClientConnect:  make(map[int]map[string]ClientConnect),
	}
}

func (channel *Channel) addConnect(companyID int, staffID int, wsConn *websocket.Conn, conn *impl.Connection) {
	wsConn.SetReadDeadline(time.Now().AddDate(1, 0, 0))
	wsConn.SetWriteDeadline(time.Now().AddDate(1, 0, 0))
	connUUID := uuid.NewV1().String()
	if channel.ClientConnect[companyID] == nil {
		channel.ClientConnect[companyID] = make(map[string]ClientConnect)
	}
	channel.ClientConnect[companyID][connUUID] = ClientConnect{
		CompanyID: companyID,
		StaffID:   staffID,
		WsConn:    wsConn,
		Conn:      conn,
	}

	go func() {
		for {
			message, err := conn.ReadMessage()
			if err != nil {
				delete(channel.ClientConnect[companyID], connUUID)
				logger.Warnf("读取消息时发现消息通道关闭: %v", err)
				break
			}
			logger.Warnf("读取到客户端消息: %v", string(message))
		}
	}()
}

func (channel *Channel) sendMessage(wsMsg Message) {
	msgBytes, _ := json.Marshal(wsMsg)
	connMap, ok := channel.ClientConnect[wsMsg.CompanyID]
	if !ok {
		logger.Warnf("客户端没有连接，无法进行数据推送 companyID[%v] message[%v]", wsMsg.CompanyID, string(msgBytes))
		return
	}

	for connUUID, conn := range connMap {
		if channel.SendCheckFunc != nil && wsMsg.GroupID > 0 && conn.StaffID > 0 && channel.SendCheckFunc(conn.StaffID, wsMsg.GroupID, wsMsg) {
			err := conn.Conn.WriteMessage(msgBytes)
			if err != nil {
				delete(channel.ClientConnect[wsMsg.CompanyID], connUUID)
				logger.Warnf("推送消息时发现消息通道关闭: %v", err)
				continue
			}
		} else {
			logger.Warnf("不符合发送条件限制 %v", wsMsg)
		}
	}
}

func (channel *Channel) DispatchMessage(message Message) {
	str, _ := json.Marshal(message)
	gredis.Publish(channel.RedisSubPubKey, string(str))
}

func (channel *Channel) Serve(companyID int, staffID int, w http.ResponseWriter, r *http.Request, resHeader http.Header) error {
	var wsConn *websocket.Conn
	var conn *impl.Connection
	var err error

	if wsConn, err = channel.Upgrader.Upgrade(w, r, resHeader); err != nil {
		logger.Warnf("wss upgrader error: %v", err)
		return err
	}

	if conn, err = impl.InitConnection(wsConn); err != nil {
		logger.Warnf("wss init connection error: %v", err)
		return err
	}

	channel.addConnect(companyID, staffID, wsConn, conn)
	return nil
}

func (channel *Channel) Run() {
	go func() {
		for {
			err := gredis.Subscribe(channel.RedisSubPubKey, func(rdsMessage string) {
				var wsMsg Message
				_ = json.Unmarshal([]byte(rdsMessage), &wsMsg)
				channel.sendMessage(wsMsg)
			})
			logger.Warnf("Subscribe失败，将于1分钟后重新订阅，错误消息为: %v", err)
			time.Sleep(time.Minute)
		}
	}()
}
