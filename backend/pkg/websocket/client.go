package websocket

import (
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
	Type   int    `json:"type"`
	Body   string `json:"body"`
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		//messageType, p, err := c.Conn.ReadMessage()
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		if string(p) == "btnpress" {
			c.Pool.Update <- string(p)
		} else {
			//var i = c.D
			message := Message{Sender: c.ID, Type: messageType, Body: string(p)}
			c.Pool.Broadcast <- message
			fmt.Printf("Message Receive: %+v\n", message)
		}
	}
}
