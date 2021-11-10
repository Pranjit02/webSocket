package main

import (
	"encoding/json"
	"fmt"
	"log"

	"golang.org/x/net/websocket"
)

func main() {
	ws, err := websocket.Dial("ws://localhost:8000/ws", "", "http://localhost:8000")

	if err != nil {
		fmt.Println(err)
		return
	}
	type Message struct {
		Email    string `json:"email"`
		Username string `json:"username"`
		Message  string `json:"message"`
	}

	for {
		var m string
		msg1 := &Message{Email: "rohit", Username: "raja", Message: "f"}
		k, _ := json.Marshal(msg1)
		log.Println(k)
		//websocket.Message.Send(ws, k)
		websocket.Message.Receive(ws, &m)

	}
}
