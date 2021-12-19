package websocket

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID   int
	Conn *websocket.Conn
	Pool *Pool
}

type Message struct {
	Sender int    `json:"sender"`
	Type   string `json:"type"`
	Body   string `json:"body"`
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		//messageType, p, err := c.Conn.ReadMessage()
		_, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		var msgJson Message
		json.Unmarshal([]byte(string(p)), &msgJson)

		if msgJson.Type == "ButtonEvent" {
			c.Pool.Update <- msgJson.Body
		} else {
			message := Message{Sender: c.ID, Type: msgJson.Type, Body: msgJson.Body}
			c.Pool.Broadcast <- message
			fmt.Printf("Message Receive: %+v\n", message)
		}
	}
}
