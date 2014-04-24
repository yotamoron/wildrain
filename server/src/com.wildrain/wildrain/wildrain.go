package main

import (
	"encoding/json"
	"github.com/GeertJohan/go.rice"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"net/http"
)

type Reply struct {
	msg string
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

func uploadAicd(w http.ResponseWriter, r *http.Request) {
	ws := upgrade(w, r)
	if ws == nil {
		return
	}
	defer func() { ws.Close() }()
	SaveAicd(ws)
}

func _getApplications(ws *websocket.Conn) {
	aicds := GetAicds()
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

func connect(w http.ResponseWriter, r *http.Request) {
	ws := upgrade(w, r)
	if ws == nil {
		return
	}
	defer func() { ws.Close() }()
	NewConnection(ws)
}

func getFlows(w http.ResponseWriter, r *http.Request) {
	ws := upgrade(w, r)
	if ws == nil {
		return
	}
	defer func() { ws.Close() }()
	flows := GetFlows()
	arr, _ := json.Marshal(flows)
	ws.WriteMessage(websocket.TextMessage, arr)
}

func saveFlow(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	SaveFlow(body)
}

func main() {
	LoadStatic()
	StoreStaticFlow()
	go StartEngine()
	http.Handle("/", http.FileServer(rice.MustFindBox("static").HTTPBox()))
	http.HandleFunc("/uploadAicd", uploadAicd)
	http.HandleFunc("/getApplications", getApplications)
	http.HandleFunc("/connect", connect)
	http.HandleFunc("/saveFlow", saveFlow)
	http.HandleFunc("/getFlows", getFlows)
	http.ListenAndServe(":8080", nil)
}
