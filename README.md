# go-push-light

This is a simple push server by Websocket in Go (uses "github.com/gorilla/websocket")

usage:

  sub - ws://127.0.0.1:8095/sub - send {"op": "sub", "intent": "channelname.taskname.12345"} // or "unsub"
  
  pub - http://127.0.0.1:8095/pub?intent=channelname.taskname.12345?obj={"id":123} // can send in post/get
  
  server will sent {"op": "intent", "intent": "channelname.taskname.12345", "obj": "{\"id\":123}"}
