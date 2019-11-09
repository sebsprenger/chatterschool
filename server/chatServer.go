package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sebsprenger/chatterschool/shared"
	"golang.org/x/net/websocket"
)

type ChatServer struct {
	clients          map[string]*websocket.Conn
	addClientChan    chan *websocket.Conn
	removeClientChan chan *websocket.Conn
	broadcastChan    chan shared.Message
	formatter        PassThroughFormatter
}

func NewChatServer(formatter PassThroughFormatter) *ChatServer {
	return &ChatServer{
		clients:          make(map[string]*websocket.Conn),
		addClientChan:    make(chan *websocket.Conn),
		removeClientChan: make(chan *websocket.Conn),
		broadcastChan:    make(chan shared.Message),
		formatter:        formatter,
	}
}

func (server *ChatServer) Start(port string) error {
	go server.distribute()

	mux := http.NewServeMux()
	mux.Handle("/", websocket.Handler(server.maintainConnection))

	serverAddress := fmt.Sprintf(":%s", port)
	webserver := http.Server{
		Addr:    serverAddress,
		Handler: mux,
	}

	fmt.Printf("starting web server on: %s\n", serverAddress)
	return webserver.ListenAndServe()
}

func (server *ChatServer) connectClient(connection *websocket.Conn) {
	server.maintainConnection(connection)
}

func (server *ChatServer) maintainConnection(connection *websocket.Conn) {
	server.addClientChan <- connection
	defer server.removeClient(connection)

	for {
		var msg shared.Message
		err := websocket.JSON.Receive(connection, &msg)
		if err != nil {
			log.Printf("Error with connection %s - will not serve client any longer\n", connection.RemoteAddr())
			return
		}
		msg = server.formatter.Modify(msg)
		server.broadcastChan <- msg
	}
}

func (server *ChatServer) distribute() {
	fmt.Println("ready to distribute messages...")
	for {
		select {
		case conn := <-server.addClientChan:
			server.addClient(conn)
		case conn := <-server.removeClientChan:
			server.removeClient(conn)
		case m := <-server.broadcastChan:
			server.broadcastMessage(m)
		}
	}
}

func (server *ChatServer) addClient(conn *websocket.Conn) {
	server.clients[conn.RemoteAddr().String()] = conn
}

func (server *ChatServer) removeClient(conn *websocket.Conn) {
	delete(server.clients, conn.RemoteAddr().String())
}

func (server *ChatServer) broadcastMessage(msg shared.Message) {
	for _, conn := range server.clients {
		err := websocket.JSON.Send(conn, msg)
		if err != nil {
			fmt.Printf("Error broadcasting message: %s\n", err)
			fmt.Printf("Removing client from the server\n")
			server.removeClient(conn)
		}
	}
}
