package main

import (
	"com.wildrain/aicd"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
)

var aicds = make(map[string]map[string]aicd.Aicd)

var staticWaze []byte = []byte(`{"applicationName": "Waze", "version": "1.0", "events": [{ "name": "TIME_TO_HOME", "params": [{ "paramName": "minutes", "paramType": "int", "required": true }]}]}`)

func LoadStatic() {
	var waze aicd.Aicd
	json.Unmarshal(staticWaze, &waze)
	setAicd(&waze)
}

func GetAicds() *map[string]map[string]aicd.Aicd {
	return &aicds
}

func setApp(appName string) map[string]aicd.Aicd {
	if app, found := aicds[appName]; !found {
		newAicdVersions := make(map[string]aicd.Aicd)
		aicds[appName] = newAicdVersions
		return newAicdVersions
	} else {
		return app
	}
}

func setAppVersion(versionsMap map[string]aicd.Aicd, rev string, a *aicd.Aicd) bool {
	_, found := versionsMap[rev]
	if !found {
		versionsMap[rev] = *a
	}
	return found
}

func setAicd(a *aicd.Aicd) {
	app := a.ApplicationName
	rev := a.Version
	appVersion := setApp(app)
	setAppVersion(appVersion, rev, a)
}

func SaveAicd(ws *websocket.Conn) {
	_, msg, e := ws.ReadMessage()
	if e != nil {
		fmt.Println("Failed reading from socket", e)
	} else {
		var a aicd.Aicd
		unmarshalErr := json.Unmarshal(msg, &a)
		if unmarshalErr != nil {
			fmt.Println("Failed unmarshaling aicd", unmarshalErr)
		} else {
			setAicd(&a)
			ws.WriteMessage(websocket.TextMessage, []byte(`{"msg": "ok"}`))
		}
	}
}

func GetAicd(applicationName string, applicationVersion string) *aicd.Aicd {
	if app, foundApp := aicds[applicationName]; foundApp {
		if ver, foundVer := app[applicationVersion]; foundVer {
			return &ver
		}
	}
	return nil
}