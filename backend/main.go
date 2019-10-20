package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var subscribers = map[string]*websocket.Conn{}

// We'll need to define an Upgrader
// this will require a Read and Write buffer size
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

// define a reader which will listen for
// new messages being sent to our WebSocket
// endpoint
func reader(conn *websocket.Conn) {
	for {
		// read in a message
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		// print out that message for clarity
		fmt.Println(string(p))

		if err := conn.WriteMessage(messageType, []byte("Hello client, I just received your message")); err != nil {
			log.Println(err)
			return
		}

		fmt.Print("done", messageType)

	}
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	// upgrade this connection to a WebSocket
	// connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if _, ok := subscribers[r.RemoteAddr]; !ok {
		subscribers[r.RemoteAddr] = ws
	}
	if err != nil {
		log.Println(err)
	}
	// helpful log statement to show connections
	log.Println("Client Connected")
	reader(ws)
}

func sendUpdateMessage(m string) {
	for addr, c := range subscribers {
		c.WriteMessage(1, []byte(fmt.Sprintf("Hello %v, %v", addr, m)))
	}
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndpoint)
}

func main() {
	fmt.Println("Hello World")
	setupRoutes()
	go func() {
		time.Sleep(10 * time.Second)
		sendUpdateMessage("this is the latest message from server")
	}()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
