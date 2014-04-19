package main

import (
	"com.wildrain/aicd"
	"encoding/json"
	"fmt"
	"github.com/GeertJohan/go.rice"
	"github.com/gorilla/websocket"
	"net/http"
)

type Reply struct {
	msg string
}

var aicds = make(map[string]map[string]aicd.Aicd)

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

func upgrade(w http.ResponseWriter, r *http.Request) *websocket.Conn {
	ws, err := websocket.Upgrade(w, r, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(w, "Not a websocket handshake", 400)
		ws.Close()
		return nil
	} else if err != nil {
		return nil
	}
	return ws
}

func saveAicd(ws *websocket.Conn) {
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

func uploadAicd(w http.ResponseWriter, r *http.Request) {
	ws := upgrade(w, r)
	if ws == nil {
		return
	}
	defer func() { ws.Close() }()
	saveAicd(ws)
}

func _getApplications(ws *websocket.Conn) {
	arr, _ := json.Marshal(aicds)
	ws.WriteMessage(websocket.TextMessage, arr)
}

func getApplications(w http.ResponseWriter, r *http.Request) {
	ws := upgrade(w, r)
	if ws == nil {
		return
	}
	defer func() { ws.Close() }()
	_getApplications(ws)
}

func main() {
	http.Handle("/", http.FileServer(rice.MustFindBox("static").HTTPBox()))
	http.HandleFunc("/uploadAicd", uploadAicd)
	http.HandleFunc("/getApplications", getApplications)
	http.ListenAndServe(":8080", nil)
}
