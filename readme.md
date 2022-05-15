
```
  go get github.com/Shaneumayanga/lolsockets
```


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

func (client *Client) Read() {
	go func() {
		//listen for incoming messages from the websocket connection
		for msg := range client.C.ReadMessages() {
			fmt.Printf("msg: %v\n", string(msg))
			//send the reply
			reply := fmt.Sprintf("Hello you said %s !", string(msg))
			client.C.WriteMessage([]byte(reply))
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