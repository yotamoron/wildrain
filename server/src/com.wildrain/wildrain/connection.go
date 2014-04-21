package main

import (
	"com.wildrain/aicd"
	"github.com/gorilla/websocket"
)

type Connection struct {
	ws       *websocket.Conn
	send     chan []byte
	instance ApplicationInstance
	aicd     *aicd.Aicd
}

type IncomigMessage struct {
	msg  []byte
	conn *Connection
}

func (c *Connection) reader() {
	for {
		_, message, err := c.ws.ReadMessage()
		if err != nil {
			break
		}
		IncomingMessagesFromApps <- &IncomigMessage{msg: message, conn: c}
	}
	c.ws.Close()
}

func (c *Connection) writer() {
	for message := range c.send {
		err := c.ws.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			break
		}
	}
	c.ws.Close()
}
