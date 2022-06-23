package lolsockets

import (
	"errors"
	"net/http"
)

func (u *Upgrader) Upgrade(rw http.ResponseWriter, r *http.Request) (*Client, error) {
	if r.Method != http.MethodGet {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Bad request"))
		return nil, errors.New("bad request")
	}
	challengeKey := r.Header.Get("Sec-Websocket-Key")

	if isEmpty(challengeKey) {
		return nil, errors.New("challenge key is empty")
	}

	if !isValidChallengeKey(challengeKey) {
		return nil, errors.New("invalid challenge key")
	}

	if !u.CheckOrigin(r) {
		return nil, errors.New("origin now allowed")
	}

	h := rw.(http.Hijacker)

	conn, _, err := h.Hijack()
	if err != nil {
		return nil, err
	}
	_, err = conn.Write([]byte("HTTP/1.1 101 Switching Protocols\r\n" +
		"Upgrade: websocket\r\n" +
		"Connection: Upgrade\r\n" +
		"Sec-WebSocket-Accept: " + computeAcceptKey(challengeKey) + "\r\n" + "\r\n"))
	if err != nil {
		return nil, err
	}
	client := &Client{
		Conn:           conn,
		ReadBufferSize: u.ReadBufferSize,
	}
	return client, nil
}
