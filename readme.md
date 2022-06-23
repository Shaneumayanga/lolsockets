
```
  go get github.com/Shaneumayanga/lolsockets
```


## Usage

```go

package main

import (
	"fmt"
	"log"
	"net/http"

	lolsockets "github.com/Shaneumayanga/lolsockets"
)

type Client struct {
	C *lolsockets.Client
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

var upgrader = &lolsockets.Upgrader{
	ReadBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	http.HandleFunc("/ws", func(rw http.ResponseWriter, r *http.Request) {
		client, err := upgrader.Upgrade(rw, r)
		if err != nil {
			panic(err)
		}
		c := Client{
			C: client,
		}
		c.Read()
	})

	log.Println("Server running on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}


```

## An example client

```go
package main

import (
	"fmt"
	"net/url"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var addr = "localhost:8080"

func main() {

	wg := &sync.WaitGroup{}
	u := url.URL{
		Scheme: "ws",
		Host:   addr,
		Path:   "/ws",
	}

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)

	if err != nil {
		panic(err)
	}
	defer c.Close()

	ticker := time.NewTicker(time.Second)

	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			<-ticker.C
			err := c.WriteMessage(websocket.TextMessage, []byte("Hemlooo"))
			if err != nil {
				return
			}
		}
	}()

	go func() {
		defer wg.Done()
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				return
			}
			fmt.Println(string(message))
		}
	}()
	wg.Wait()
}

```

### Any contributions are welcome!