package lolsockets

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
)

type Upgrader struct {
	ReadBufferSize int
	CheckOrigin    func(r *http.Request) bool
}

func (u *Upgrader) upgradeToWs(conn net.Conn, challengeKey string) error {
	w := bufio.NewWriter(conn)
	w.Write([]byte("HTTP/1.1 101 Switching Protocols"))
	w.Write([]byte("\r\n"))
	w.Write([]byte("Upgrade: websocket"))
	w.Write([]byte("\r\n"))
	w.Write([]byte("Connection: Upgrade"))
	w.Write([]byte("\r\n"))
	w.Write([]byte(fmt.Sprintf("Sec-WebSocket-Accept: %s", computeAcceptKey(challengeKey))))
	w.Write([]byte("\r\n"))
	w.Write([]byte("\r\n"))
	if err := w.Flush(); err != nil {
		return err
	}
	return nil
}
