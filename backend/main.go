package main

import (
	"fmt"
	"net/http"

	"github.com/samlaurenson/livechatting/pkg/websocket"
)

var userIDIncrement = 0

func serve(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	ws, err := websocket.Upgrade(w, r)

	if err != nil {
		fmt.Fprintf(w, "%+V\n", err)
	}

	client := &websocket.Client{
		ID:   userIDIncrement,
		Conn: ws,
		Pool: pool,
	}
	userIDIncrement++

	pool.Register <- client
	client.Read()
}

func setupRoutes() {
	pool := websocket.NewPool()
	go pool.Start()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serve(pool, w, r)
	})
}

func main() {
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}