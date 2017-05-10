// connection
package skaro

import (
	"log"
	"net/url"

	"os"
	"os/signal"

	"github.com/gorilla/websocket"
)

type Connection struct {
	ws *websocket.Conn
	Account
}

type Account struct {
	login    string
	password string
}

/* Methods */

func (this *Connection) start(server string) error {
	var err error
	interrupt := make(chan os.Signal, 1)

	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: server, Path: "/skaro_client/"}
	log.Printf("connecting to %s", u.String())

	this.ws, _, err = websocket.DefaultDialer.Dial(u.String(), nil)
	defer log.Println("Closing connection")
	defer this.ws.Close()
	if err != nil {
		log.Fatal("dial:", err)
	}

	done := make(chan struct{})

	go func() {
		defer log.Println("Closing connection")
		defer this.ws.Close()
		defer close(done)
		for {
			_, message, err := this.ws.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()
	return err
}

func (this *Connection) Send(message []byte) error {
	return nil
}

func (this *Connection) Recieve() ([]byte, error) {
	return nil, nil
}

func (this *Connection) Authorize() error {
	return nil
}
