package main

import (
	"fmt"
)

type RequestFromFlow struct {
	// Send this message
	body *interface{}
	// To this guy
	instance ApplicationInstance
	// Put the reply here
	inbox chan *interface{}
}

type trigger struct {
	Event  string
	Params map[string]interface{}
}

func NewFlow(instance ApplicationInstance, msg *Message) {
	incoming := msg.Body.(map[string]interface{})
	newTrigger := trigger{Event: incoming["Event"].(string), Params: incoming["Params"].(map[string]interface{})}
	flow := GetFlow(instance.ApplicationName, instance.ApplicationVersion, newTrigger.Event)
	fmt.Println(flow)
}
