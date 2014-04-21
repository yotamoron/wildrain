package main

import (
	"com.wildrain/aicd"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"strings"
)

type Message struct {
	ReqId int
	Body  interface{}
}

var IncomingMessagesFromApps = make(chan *IncomigMessage, 256)
var IncomingRequestFromFlow = make(chan *RequestFromFlow)

type ApplicationInstance struct {
	ApplicationName    string
	ApplicationVersion string
}

func authenticate(ws *websocket.Conn) (*ApplicationInstance, *aicd.Aicd, error) {
	_, msg, e := ws.ReadMessage()
	if e != nil {
		fmt.Println("Failed reading from socket", e)
		return nil, nil, e
	} else {
		var instance ApplicationInstance
		unmarshalErr := json.Unmarshal(msg, &instance)
		if unmarshalErr != nil {
			fmt.Println("Failed unmarshaling aicd", unmarshalErr)
			return nil, nil, unmarshalErr
		} else {
			aicd := GetAicd(instance.ApplicationName, instance.ApplicationVersion)
			if aicd == nil {
				errMsg := strings.Join([]string{"No such application/version aicd", instance.ApplicationName, instance.ApplicationVersion}, " ")
				return nil, nil, errors.New(errMsg)
			} else {
				return &instance, aicd, nil
			}
		}
	}
}

func listenForIncomigMessages() {
	currentReqId := 1
	pending := make(map[int]*RequestFromFlow)
	select {
	case fromApp := <-IncomingMessagesFromApps:
		var message Message
		json.Unmarshal(fromApp.msg, &message)
		msgReqId := message.ReqId
		// 0 is a special case
		if msgReqId == 0 {
			go NewFlow(fromApp.conn, &message)
		} else {
			fromFlow := pending[msgReqId]
			delete(pending, msgReqId)
			fromFlow.inbox <- &message.Body
		}
	case fromFlow := <-IncomingRequestFromFlow:
		instance := fromFlow.instance
		body := fromFlow.body
		conn := GetConnection(&instance)
		req := Message{ReqId: currentReqId, Body: body}
		jsonRequest, _ := json.Marshal(req)
		conn.send <- jsonRequest
		pending[currentReqId] = fromFlow
		currentReqId += 1
	}
}

func StartEngine() {
	go ConnectionStore()
	go listenForIncomigMessages()
}

func NewConnection(ws *websocket.Conn) {
	instance, aicd, _ := authenticate(ws)
	c := &Connection{send: make(chan []byte, 256), ws: ws, instance: *instance, aicd: aicd}
	StoreConnection <- c
	go c.writer()
	c.reader()
}
