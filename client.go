package lolsockets

import "net"

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

func (c *Client) ReadMessages() {

}
