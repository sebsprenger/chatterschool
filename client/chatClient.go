package client

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/sebsprenger/chatterschool/shared"
	"golang.org/x/net/websocket"
)

func NewChatClient(inputFormatter InputFormatter, outputFormatter OutputFormatter) ChatClient {
	return ChatClient{
		ws:           nil,
		msgCreator:   inputFormatter,
		msgFormatter: outputFormatter,
	}
}

type ChatClient struct {
	ws           *websocket.Conn
	msgCreator   InputFormatter
	msgFormatter OutputFormatter
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

func (client *ChatClient) Disconnect() error {
	if client.ws == nil {
		return fmt.Errorf("Disconnected before connection was established")
	}
	return client.ws.Close()
}

func (client *ChatClient) JoinChat() {
	go client.receive()
	client.send()
}

func getMyIP() string {
	// TODO lookup actual IP
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("http://%d.%d.%d.%d", rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(256))
}

func (client *ChatClient) receive() {
	for {
		var msg shared.Message
		err := websocket.JSON.Receive(client.ws, &msg)
		if err != nil {
			fmt.Println("Error receiving message: ", err.Error())
			break
		}

		fmt.Printf("%s\n", client.msgFormatter.FormatMessage(msg))
	}
}

func (client *ChatClient) send() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		text := scanner.Text()

		switch text {
		case "":
			continue
		case "/quit":
			break
		}

		msg := client.msgCreator.CreateMessage(text)

		err := websocket.JSON.Send(client.ws, msg)
		if err != nil {
			fmt.Printf("Error sending message: %s\n", err.Error())
			fmt.Printf("I am going to stop sending...\n")
			break
		}
	}
}
