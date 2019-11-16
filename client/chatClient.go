package client

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/sebsprenger/chatterschool/shared"
	"golang.org/x/net/websocket"
)

func NewChatClient() ChatClient {
	return ChatClient{
		ws: nil,
	}
}

type ChatClient struct {
	ws *websocket.Conn
}

func (client *ChatClient) Connect(ip, port string) error {
	theirIP := fmt.Sprintf("ws://%s:%s", ip, port)
	myIP := getMyIP()
	protocol := ""

	ws, err := websocket.Dial(theirIP, protocol, myIP)
	if err != nil {
		return err
	}

	client.ws = ws
	return nil
}

func getMyIP() string {
	// TODO lookup actual IP
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("http://%d.%d.%d.%d", rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(256))
}

func (client *ChatClient) Disconnect() error {
	if client.ws == nil {
		return fmt.Errorf("Disconnected before connection was established")
	}
	return client.ws.Close()
}

func (client *ChatClient) ReceiveChatMessagenOn(outputFormatter OutputFormatter) {
	go client.receive(outputFormatter)
}

func (client *ChatClient) receive(outputFormatter OutputFormatter) {
	for {
		var msg shared.Message
		err := websocket.JSON.Receive(client.ws, &msg)
		if err != nil {
			fmt.Println("Error receiving message: ", err.Error())
			break
		}

		outputFormatter.FormatMessage(msg)
	}
}

func (client *ChatClient) Send(msg shared.Message) error {
	err := websocket.JSON.Send(client.ws, msg)
	if err != nil {
		fmt.Printf("Error sending message: %s\n", err.Error())
		return err
	}

	return nil
}
