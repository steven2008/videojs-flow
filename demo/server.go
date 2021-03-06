package main

import (
	"golang.org/x/net/websocket"
	"net/http"
	"fmt"
	"flag"
	"time"
)

func main()  {
	var port int
	flag.IntVar(&port, "l", 8088, "the port to listen for websocket")
	flag.Parse()

	http.Handle("/", websocket.Handler(func(c *websocket.Conn){
		var msg string
		websocket.Message.Receive(c, &msg)
		fmt.Println("Got websocket connection", c.RemoteAddr().String(), msg)

		for {
			now := time.Now().Format("2006-01-02 15:04:05.000")
			if err := websocket.Message.Send(c, fmt.Sprintf("%v, %v", now, msg)); err != nil {
				fmt.Println("Send failed, err is", err)
				return
			}
			time.Sleep(1 * time.Second)
		}
	}))

	fmt.Println(fmt.Sprintf("Client can connect to ws://127.0.0.1:%v/echo", port))
	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil); err != nil {
		fmt.Println("Serve failed, err is", err)
	}
}
