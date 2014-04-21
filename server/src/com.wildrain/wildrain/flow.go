package main

import "fmt"

type RequestFromFlow struct {
	// Send this message
	body *interface{}
	// To this guy
	instance ApplicationInstance
	// Put the reply here
	inbox chan *interface{}
}

func NewFlow(conn *Connection, msg *Message) {
	fmt.Println(conn, msg)
}
