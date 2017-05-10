package skaro

import (
	"log"
)

func StartConnection(host string) {
	log.Println("Opening WebSocket connect...")
	var connection Connection
	connection.start(host)
}
