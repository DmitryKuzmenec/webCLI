package main

import (
	"log"
	"net/http"
	fs "webCli/commonFS"
	h "webCli/handlers"

	"golang.org/x/net/websocket"
)

func main() {
	fs := fs.CommonFS{http.Dir("./common")}
	http.HandleFunc("/", h.Index)
	http.Handle("/common/", http.StripPrefix("/common/", http.FileServer(fs)))
	http.Handle("/ws/", websocket.Handler(h.WebSoketHandler))

	err := http.ListenAndServe("localhost:9000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
