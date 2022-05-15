```go

package main

import (
	"fmt"
	"log"
	"net/http"

	lolsockets "github.com/Shaneumayanga/lolsockets"
)

type Client struct {
	C lolsockets.Client
}

func (c *Client) Read() {
	go func() {
		for msg := range c.C.ReadMessages() {
			fmt.Printf("msg: %v\n", msg)
		}
	}()
}

func main() {
	http.HandleFunc("/ws", func(rw http.ResponseWriter, r *http.Request) {
		client := lolsockets.Upgrade(rw, r)
		c := Client{
			C: *client,
		}
		c.Read()
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}


```