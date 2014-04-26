package main

import (
	"fmt"
	"github.com/robertkrimen/otto"
)

type MessageTypeWrapper struct {
	Type string
	Body *interface{}
}

func getDispatcher(inbox chan *interface{}, vm *otto.Otto, wrapper func(body *interface{}) *interface{}) func(call otto.FunctionCall) otto.Value {
	return func(call otto.FunctionCall) otto.Value {
		instance := getInstance(call.Argument(0))
		body, _ := call.Argument(1).Export()
		outgoing := RequestFromFlow{body: wrapper(&body), instance: instance, inbox: inbox}
		IncomingRequestFromFlow <- &outgoing
		incoming := <-inbox
		value, _ := vm.ToValue(*incoming)
		return value
	}
}

func getInstance(value otto.Value) ApplicationInstance {
	exportedInstance, _ := value.Export()
	m := exportedInstance.(map[string]interface{})
	return ApplicationInstance{ApplicationName: m["ApplicationName"].(string), ApplicationVersion: m["ApplicationVersion"].(string)}
}

func getDispatcherWrapped(inbox chan *interface{}, vm *otto.Otto, t string) func(call otto.FunctionCall) otto.Value {
	f := func(body *interface{}) *interface{} {
		ret := interface{}(MessageTypeWrapper{Type: t, Body: body})
		return &ret
	}
	return getDispatcher(inbox, vm, f)
}

func getSendCommand(inbox chan *interface{}, vm *otto.Otto) func(call otto.FunctionCall) otto.Value {
	return getDispatcherWrapped(inbox, vm, "COMMAND")
}

func getSendQuery(inbox chan *interface{}, vm *otto.Otto) func(call otto.FunctionCall) otto.Value {
	return getDispatcherWrapped(inbox, vm, "QUERY")
}

func setBuiltIns(vm *otto.Otto) {
	inbox := make(chan *interface{})
	vm.Set("sendQuery", getSendQuery(inbox, vm))
	vm.Set("sendCommand", getSendCommand(inbox, vm))
}

func printOut(call otto.FunctionCall) otto.Value {
	fmt.Println(call.Argument(0).String())
	return otto.UndefinedValue()
}

func NewFlowVm(trigger *Trigger) *otto.Otto {
	vm := otto.New()
	setBuiltIns(vm)
	vm.Set("TRIGGER", trigger)
	vm.Set("printOut", printOut)
	return vm
}
