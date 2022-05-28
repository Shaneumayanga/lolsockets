package lolsockets

import (
	"log"
	"net/http"
)

func Upgrade(rw http.ResponseWriter, r *http.Request) *Client {
	if r.Method != http.MethodGet {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Bad request"))
		return nil
	}
	challengeKey := r.Header.Get("Sec-Websocket-Key")
	if !isValidChallengeKey(challengeKey) {
		log.Fatal("Challenge key is not valid")
	}

	log.Println(challengeKey)

	h := rw.(http.Hijacker)

	conn, _, err := h.Hijack()
	if err != nil {
		log.Println(err.Error())
	}
	_, err = conn.Write([]byte("HTTP/1.1 101 Switching Protocols\r\n" +
		"Upgrade: websocket\r\n" +
		"Connection: Upgrade\r\n" +
		"Sec-WebSocket-Accept: " + computeAcceptKey(challengeKey) + "\r\n" + "\r\n"))
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Websocket handshake completed")
	client := NewClient(conn)
	return client
}
