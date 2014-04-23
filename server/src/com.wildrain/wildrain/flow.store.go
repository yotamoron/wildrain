package main

import (
	"encoding/json"
)

// flows[ApplicationName][ApplicationVersion][EventName][FlowName] --> the flow itself
var flows = make(map[string]map[string]map[string]map[string]string)

type savedFlow struct {
	Flow               string
	FlowName           string
	ApplicationName    string
	ApplicationVersion string
	EventName          string
}

func GetFlows() map[string]map[string]map[string]map[string]string {
	return flows
}

func GetFlow(applicationName string, applicationVersion string, eventName string) map[string]string {
	if app, foundApp := flows[applicationName]; foundApp {
		if ver, foundVer := app[applicationVersion]; foundVer {
			if event, foundEvent := ver[eventName]; foundEvent {
				return event
			}
		}
	}
	return nil
}
func SaveFlow(body []byte) {
	var f savedFlow
	json.Unmarshal(body, &f)
	applicationName := f.ApplicationName
	app, foundApp := flows[applicationName]
	if !foundApp {
		app = make(map[string]map[string]map[string]string)
		flows[applicationName] = app
	}
	applicationVersion := f.ApplicationVersion
	ver, foundVer := app[applicationVersion]
	if !foundVer {
		ver = make(map[string]map[string]string)
		app[applicationVersion] = ver
	}
	eventName := f.EventName
	event, foundEvent := ver[eventName]
	if !foundEvent {
		event = make(map[string]string)
		ver[eventName] = event
	}
	event[f.FlowName] = f.Flow
}
