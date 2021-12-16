package websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

//					Inputs VVV 								Outputs VVV
func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	up := websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024, CheckOrigin: func(r *http.Request) bool { return true }}
	ws, err := up.Upgrade(w, r, nil) //Upgrade outputs 2 variables, ws for the websocket and err to check for errors
	if err != nil {
		log.Println(err)
		return ws, err
	}
	return ws, nil
}
