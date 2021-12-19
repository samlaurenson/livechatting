package websocket

import (
	"fmt"
	"strconv"
)

type Pool struct {
	Register    chan *Client
	Unregister  chan *Client
	Clients     map[*Client]bool
	Broadcast   chan Message
	Update      chan string
	ButtonCount int
}

type UserList struct {
	IDList []int  `json:"IDlist"`
	Type   string `json:"type"`
}

func NewPool() *Pool {
	return &Pool{
		Register:    make(chan *Client),
		Unregister:  make(chan *Client),
		Clients:     make(map[*Client]bool),
		Broadcast:   make(chan Message),
		Update:      make(chan string),
		ButtonCount: 0,
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			var clientList []int
			joinID := client.ID
			pool.Clients[client] = true
			if err := client.Conn.WriteJSON(Message{Type: "ButtonEvent", Body: strconv.Itoa(pool.ButtonCount)}); err != nil {
				fmt.Println(err)
			}
			fmt.Println("Size of connection pool: ", len(pool.Clients))
			for client, _ := range pool.Clients {
				clientList = append(clientList, client.ID)
				fmt.Println(client)
				client.Conn.WriteJSON(Message{Sender: joinID, Type: "UserJoin", Body: "New user joined... "})
			}

			for client, _ := range pool.Clients {
				client.Conn.WriteJSON(UserList{IDList: clientList, Type: "UserListUpdate"})
			}
			break
		case client := <-pool.Unregister:
			leavingID := client.ID
			delete(pool.Clients, client)
			fmt.Println("Size of connection pool: ", len(pool.Clients))
			for client, _ := range pool.Clients {
				client.Conn.WriteJSON(Message{Sender: leavingID, Type: "UserLeave", Body: "User disconnected"})
			}
			break
		case message := <-pool.Broadcast: //The <- operator represents the idea of passing a value from a channel to a reference
			fmt.Println("Sending message to all clients in Pool")
			for client, _ := range pool.Clients {
				if err := client.Conn.WriteJSON(message); err != nil {
					fmt.Println(err)
					return
				}
			}
		case <-pool.Update:
			fmt.Println("Button Updating")
			pool.ButtonCount++
			for client, _ := range pool.Clients {
				if err := client.Conn.WriteJSON(Message{Type: "ButtonEvent", Body: strconv.Itoa(pool.ButtonCount)}); err != nil {
					fmt.Println(err)
					return
				}
			}
			break
		}
	}
}
