package main

import (
	"encoding/json"
	"fmt"
    "com.wildrain/aicd"
)

func main() {
	a := aicd.Aicd{ApplicationName: "Test1",
		Revision: "0.1",
		Events: []aicd.ParametrizedEndpoint{
			{Name: "TestEvent", Params: []aicd.Param{aicd.Param{ParamName: "Clicks", ParamType: "int", Required: true}}},
		},
		Commands: []aicd.ParametrizedEndpoint{
			{Name: "TestCommand", Params: []aicd.Param{aicd.Param{ParamName: "Repeat", ParamType: "int", Required: false}}},
		},
		Queries: []aicd.Query{
			{
				Name:   "TestQuery",
				Params: []aicd.Param{aicd.Param{ParamName: "FromEpoch", ParamType: "boolean", Required: true}},
				Return: []aicd.Param{aicd.Param{ParamName: "Repeat", ParamType: "int", Required: false}},
			},
		},
	}
	b, _ := json.Marshal(a)
	fmt.Println(string(b))
}

