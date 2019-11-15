package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sebsprenger/chatterschool/shared"
	"golang.org/x/net/websocket"
)

type ChatServer struct {
	clients   map[string]*websocket.Conn
	chatroom  chan shared.Message
	formatter PassThroughFormatter
}

func NewChatServer(formatter PassThroughFormatter) *ChatServer {
	return &ChatServer{
		clients:   make(map[string]*websocket.Conn),
		chatroom:  make(chan shared.Message),
		formatter: formatter,
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

func (server *ChatServer) maintainConnection(connection *websocket.Conn) {
	server.addClient(connection)
	defer server.removeClient(connection)

	for {
		var msg shared.Message
		err := websocket.JSON.Receive(connection, &msg)
		if err != nil {
			log.Printf("Error with connection %s - will not serve client any longer\n", connection.RemoteAddr())
			return
		}
		msg = server.formatter.Modify(msg)
		server.chatroom <- msg
	}
}

func (server *ChatServer) distribute() {
	fmt.Println("ready to distribute messages...")
	for {
		msg := <-server.chatroom
		server.sendToAllClients(msg)
	}
}

func (server *ChatServer) addClient(conn *websocket.Conn) {
	server.clients[conn.RemoteAddr().String()] = conn
}

func (server *ChatServer) removeClient(conn *websocket.Conn) {
	delete(server.clients, conn.RemoteAddr().String())
}

func (server *ChatServer) sendToAllClients(msg shared.Message) {
	for client, conn := range server.clients {
		err := websocket.JSON.Send(conn, msg)
		if err != nil {
			fmt.Printf("Error broadcasting message to client %s: %s\n", client, err)
			fmt.Printf("Removing client from the server\n")
			server.removeClient(conn)
		}
	}
}
