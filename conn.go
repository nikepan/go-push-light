package main

import (
	"github.com/gorilla/websocket"
	"net/http"
)

type connection struct {
	// The websocket connection.
	ws *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte
}

func (c *connection) reader() {
	for {
		_, message, err := c.ws.ReadMessage()
		if err != nil {
			break
		}
		msg := InboundMsg{c, message}
		h.broadcast <- msg
	}
	c.ws.Close()
}

func (c *connection) writer() {
	for message := range c.send {
		err := c.ws.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			break
		}
	}
	c.ws.Close()
}

var upgrader = &websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024}

func subHandler(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		// allow all connections by default
		return true
	}
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	c := &connection{send: make(chan []byte, 256), ws: ws}
	h.register <- c
	defer func() { h.unregister <- c }()
	go c.writer()
	c.reader()
}
