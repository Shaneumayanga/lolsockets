package lolsockets

import (
	"log"
	"net"
)

type Client struct {
	Conn           net.Conn
	ReadBufferSize int
}

func NewClient(conn net.Conn) *Client {
	return &Client{
		Conn:           conn,
		ReadBufferSize: 1024,
	}
}

func (c *Client) WriteMessage(msg []byte) error {
	_, err := c.Conn.Write(encodePayload(msg))
	return err
}

func (c *Client) ReadMessages() chan []byte {
	msg := make(chan []byte)
	go func(ch chan []byte) {
		b := make([]byte, c.ReadBufferSize)
		for {
			_, err := c.Conn.Read(b)
			if err != nil {
				log.Println(err.Error())
				return
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
