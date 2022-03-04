package p2p

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/texasroh/junecoin/utils"
)

var upgrader = websocket.Upgrader{}

func Upgrade(rw http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := upgrader.Upgrade(rw, r, nil)
	utils.HandleErr(err)
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			// conn.Close()
			break
		}
		utils.HandleErr(err)
		fmt.Printf("Received: %s\n\n", p)
		message := fmt.Sprintf("Sent from server: %s\n\n", p)
		utils.HandleErr(conn.WriteMessage(websocket.TextMessage, []byte(message)))
	}
}
