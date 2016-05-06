package main

import (
	"encoding/json"
	"log"
)

type SubscribeMsg struct {
	Op     string `json:"op"`
	Intent string `json:"intent"`
}

type InboundMsg struct {
	Conn *connection
	Msg  []byte
}

type hub struct {
	// Registered connections.
	connections map[*connection]bool
	// Inbound messages from the connections.
	broadcast chan InboundMsg //[]byte

	// Register requests from the connections.
	register chan *connection

	// Unregister requests from connections.
	unregister chan *connection
	// Subscribers task > connection
	subs map[string]map[*connection]bool
}

var h = hub{
	broadcast:   make(chan InboundMsg),
	register:    make(chan *connection),
	unregister:  make(chan *connection),
	connections: make(map[*connection]bool),
	subs:        make(map[string]map[*connection]bool),
}

// ********************************************************* //

func dropConn(c *connection) {
	for sub := range h.subs {
		_, ok := h.subs[sub][c]
		if ok {
			delete(h.subs[sub], c)
		}
	}
	delete(h.connections, c)
}

func (h *hub) run() {

	log.Println("Pushserver started...")

	for {
		select {
		case c := <-h.register:
			h.connections[c] = true
			log.Printf("connect %+v conns %+v", c, len(h.connections))
		case c := <-h.unregister:
			log.Printf("disconnect %+v conns %+v", c, len(h.connections))
			dropConn(c)
		case m := <-h.broadcast:
			_, ok := h.connections[m.Conn]
			if ok {
				var msg SubscribeMsg
				err := json.Unmarshal(m.Msg, &msg)
				if err != nil {
					log.Printf(err.Error())
				}
				if msg.Op == "sub" {
					sub, ok := h.subs[msg.Intent]
					if !ok {
						sub = make(map[*connection]bool)
						h.subs[msg.Intent] = sub
					}
					sub[m.Conn] = true
					log.Printf("sub: %+v %+v", m.Conn, msg.Intent)
				} else if msg.Op == "unsub" {
					sub, ok := h.subs[msg.Intent]
					if ok {
						_, ok = sub[m.Conn]
						if ok {
							delete(sub, m.Conn)
							log.Printf("unsub: %+v %+v", m.Conn, msg.Intent)
						}
					}
				}
			}
		}
	}
}
