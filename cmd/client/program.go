package main

import (
	"flag"
	"log"

	"github.com/sebsprenger/chatterschool/client"
)

var (
	ip   = flag.String("ip", "localhost", "server ip")
	port = flag.String("port", "9001", "server port")
	name = flag.String("name", "nobody", "name used for chat")
)

func main() {
	flag.Parse()

	msgCreator := MyInputFormatter{
		sender: *name,
	}
	msgFormatter := MyOutputFormatter{}
	client := client.NewChatClient(msgCreator, msgFormatter)

	err := client.Connect(*ip, *port)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect()

	client.JoinChat()
}
