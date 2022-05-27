
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
			if err := client.C.WriteMessage([]byte(reply)); err != nil {
				log.Println(err.Error())
				return
			}
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

	log.Println("Server running on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}



```