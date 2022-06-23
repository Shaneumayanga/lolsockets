package lolsockets

import "net/http"

type Upgrader struct {
	ReadBufferSize int
	CheckOrigin    func(r *http.Request) bool
}
