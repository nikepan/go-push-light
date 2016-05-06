package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func (c *connection) sendJson(msg_json *[]byte) {
	select {
	case c.send <- []byte(*msg_json):
	default:
		dropConn(c)
	}
}

func (c *connection) sendData(msg interface{}) {
	msg_json, err := json.Marshal(msg)
	if err != nil {
		log.Printf(err.Error())
	}
	c.sendJson(&msg_json)
}

type PubMsg struct {
	Op     string `json:"op"`
	Intent string `json:"intent"`
	Obj    string `json:"obj,omitempty"`
}

func sendPub(intent string, obj string) string {
	log.Printf("intent: %+v %+v", intent, obj)
	sub, ok := h.subs[intent]
	if ok {
		for l := range sub {
			msg := PubMsg{"intent", intent, obj}
			l.sendData(msg)
			log.Printf("pub: %+v", msg)
		}
	}
	return "{\"status\": \"ok\"}"
}

func pubHandler(c http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	intent, ok := req.URL.Query()["intent"]
	if !ok {
		intent, ok = req.Form["intent"]
		if !ok {
			return
		}
	}
	obj, ok := req.URL.Query()["obj"]
	if !ok {
		obj, ok = req.Form["obj"]
		if !ok {
			obj = []string{""}
		}
	}
	msg_json := []byte(sendPub(intent[0], obj[0]))
	c.Header().Set("Content-Type", "application/json")
	c.Write(msg_json)
}
