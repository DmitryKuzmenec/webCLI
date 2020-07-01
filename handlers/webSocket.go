package handlers

import (
	"log"
	"webCli/cli"

	"golang.org/x/net/websocket"
)

func WebSoketHandler(ws *websocket.Conn) {
	c, err := cli.NewCli()
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Created new Cli")
	cmd := make(chan []byte, 1024)

	// get input from ws
	go func(c *cli.Cli, ws *websocket.Conn) {
		for {
			data := []byte{}
			websocket.Message.Receive(ws, &data)
			if len(data) > 0 {
				cmd <- data
			}
		}
	}(c, ws)

	// send output to ws
	go func(c *cli.Cli, ws *websocket.Conn) {
		for {
			select {
			case com := <-cmd:
				c.Do(com)
			case data := <-c.Out:
				websocket.Message.Send(ws, string(data))
			}
		}
	}(c, ws)
	for {
	}
}
