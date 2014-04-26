package main

type RequestFromFlow struct {
	// Send this message
	body *interface{}
	// To this guy
	instance ApplicationInstance
	// Put the reply here
	inbox chan *interface{}
}

type Trigger struct {
	Event  string
	Params map[string]interface{}
}

func runFlow(flowName string, flow string, instance ApplicationInstance, t *Trigger) {
	vm := NewFlowVm(t)
	vm.Run(flow)
}

func NewFlow(instance ApplicationInstance, msg *Message) {
	incoming := msg.Body.(map[string]interface{})
	newTrigger := Trigger{Event: incoming["Event"].(string), Params: incoming["Params"].(map[string]interface{})}
	flows := GetFlow(instance.ApplicationName, instance.ApplicationVersion, newTrigger.Event)
	if nil != flows {
		for flowName, flow := range flows {
			go runFlow(flowName, flow, instance, &newTrigger)
		}
	}
}
