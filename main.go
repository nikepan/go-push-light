package main

import (
	"flag"
	"go/build"
	"log"
	"net/http"
)

var (
	addr   = flag.String("addr", ":8095", "http service address")
	assets = flag.String("assets", defaultAssetPath(), "path to assets")
)

func defaultAssetPath() string {
	p, err := build.Default.Import("gary.burd.info/go-websocket-chat", "", build.FindOnly)
	if err != nil {
		return "."
	}
	return p.Dir
}

func pubHandler(c http.ResponseWriter, req *http.Request) {
	task := req.URL.Query().Get("task")
	msg_json := []byte(sendPub(task))
	c.Header().Set("Content-Type", "application/json")
	c.Write(msg_json)
}

func main() {

	flag.Parse()

	go h.run()
	http.HandleFunc("/pub", pubHandler)
	http.HandleFunc("/sub", subHandler)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}
