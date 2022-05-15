package lolsockets

import (
	"fmt"
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

func (c *Client) WriteMessage(msg string) {

}

func (c *Client) ReadMessages() chan string {
	msg := make(chan string)
	go func(ch chan string) {
		log.Println("lol")
		for {
			b := make([]byte, 2048)
			_, err := c.Conn.Read(b)
			if err != nil {
				panic(err)
			}

			length := b[1] - 0x80
			if length <= 125 {
				key := make([]byte, 4)
				key[0] = b[2]
				key[1] = b[3]
				key[2] = b[4]
				key[3] = b[5]
				decoded := make([]byte, length)
				for i := 0; i < int(length); i++ {
					realPos := 6 + i
					decoded[i] = b[realPos] ^ key[i%4]
				}
				ch <- string(decoded)
				fmt.Printf("decoded: %v\n", string(decoded))
			} else if length > 125 {
				len := uint16(b[2]) + uint16(b[3])
				key := []byte{b[4], b[5], b[6], b[7]}
				decoded := make([]byte, len)
				for i := 0; i < int(length); i++ {
					realPos := 8 + i
					decoded[i] = b[realPos] ^ key[i%4]
				}
				ch <- string(decoded)
				fmt.Printf("decoded: %v\n", string(decoded))
			}
		}
	}(msg)
	return msg
}
