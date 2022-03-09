package p2p

import (
	"fmt"

	"github.com/gorilla/websocket"
)

var Peers map[string]*peer = make(map[string]*peer)

type peer struct {
	conn  *websocket.Conn
	inbox chan []byte
}

func (p *peer) read() {
	// delete peer in case of error
	for {
		_, m, err := p.conn.ReadMessage()
		if err != nil {
			break
		}
		fmt.Printf("%s", m)
	}
}

func (p *peer) write() {
	for {
		m := <-p.inbox
		err := p.conn.WriteMessage(websocket.TextMessage, m)
		if err != nil {
			break
		}
	}
}

func initPeer(conn *websocket.Conn, address, port string) *peer {
	p := &peer{
		conn,
		make(chan []byte),
	}
	key := fmt.Sprintf("%s:%s", address, port)
	Peers[key] = p
	go p.read()
	go p.write()
	return p
}
