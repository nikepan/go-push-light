# go-push-light

[![Build Status](https://travis-ci.org/nikepan/go-push-light.svg?branch=master)](https://travis-ci.org/nikepan/go-push-light)
[![download binaries](https://img.shields.io/badge/binaries-download-blue.svg)](https://github.com/nikepan/go-push-light/releases)

This is a simple push server by Websocket in Go (uses "github.com/gorilla/websocket")

usage:

  sub - **`ws://127.0.0.1:8095/sub`** - send `{"op": "sub", "intent": "channelname.taskname.12345"}` // or "unsub"
  
  pub - **`http://127.0.0.1:8095/pub?intent=channelname.taskname.12345?obj={"id":123}`** // can send in post/get
  
  server will sent `{"op": "intent", "intent": "channelname.taskname.12345", "obj": "{\"id\":123}"}`

  You can use pusher.js on page:
```javascript
// subscribe
pushSub(intent, function(obj){}); // intent - string, obj - additional data from server

// unsubscribe
pushUnsub(intent);
```
  Also you can use pusher from python/django apps. See `pusher.py` module (use requests).

  Use:
```python
from pusher import push_intent

  push_intent(intent, obj) # obj - string or dict
```

To change listen port: add param `-addr=8080`