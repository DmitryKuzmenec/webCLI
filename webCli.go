package main

import (
	"log"
	"net/http"
	fs "webCli/commonFS"
	"webCli/dirBox"
	h "webCli/handlers"

	"golang.org/x/net/websocket"
)

func main() {
	box := dirBox.NewBox()
	fs := fs.CommonFS{box}
	http.HandleFunc("/", h.Index)
	http.Handle("/common/", http.StripPrefix("/common/", http.FileServer(fs)))
	http.Handle("/ws/", websocket.Handler(h.WebSoketHandler))

	err := http.ListenAndServe("localhost:9000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
