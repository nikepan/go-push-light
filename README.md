# go-push-light

This is a simple push server by Websocket in Go (uses "github.com/gorilla/websocket")

usage:

  sub - ws://127.0.0.1:8095/sub - send {"op": "sub", "task": "channelname.taskname.12345"} // or "unsub"
  
  pub - http://127.0.0.1:8095/pub?task=channelname.taskname.12345

  
  server will sent {"op": "event", "task": "channelname.taskname.12345"}
