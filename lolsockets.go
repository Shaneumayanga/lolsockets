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
	if err = u.upgradeToWs(conn, challengeKey); err != nil {
		return nil, err
	}
	client := &Client{
		Conn:           conn,
		ReadBufferSize: u.ReadBufferSize,
	}
	return client, nil
}
