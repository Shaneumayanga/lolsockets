package lolsockets

import (
	"log"
	"net"
)

type Client struct {
	Conn net.Conn
}

func NewClient(conn net.Conn) *Client {
	return &Client{
		Conn: conn,
	}
}

func (c *Client) WriteMessage(msg []byte) {
	_, err := c.Conn.Write(encodePayload(msg))
	if err != nil {
		log.Println(err.Error())
	}
}

func (c *Client) ReadMessages() chan []byte {
	msg := make(chan []byte)
	go func(ch chan []byte) {
		for {
			b := make([]byte, 2048)
			_, err := c.Conn.Read(b)
			if err != nil {
				panic(err)
			}
			decodedData := decodePayload(b)
			if decodedData == nil {
				log.Println("Failed to decode data")
			} else {
				msg <- decodedData
			}
		}
	}(msg)
	return msg
}
