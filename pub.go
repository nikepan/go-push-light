package main

import (
	"encoding/json"
	"log"
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
	Op   string `json:"op"`
	Task string `json:"task"`
}

func sendPub(task string) string {
	sub, ok := h.subs[task]
	if ok {
		for l := range sub {
			msg := PubMsg{"event", task}
			l.sendData(msg)
			log.Printf("pub: %+v", msg)
		}
	}
	return ""
}
