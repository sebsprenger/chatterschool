package main

import (
	"flag"
	"log"

	"github.com/sebsprenger/chatterschool/server"
)

var (
	port = flag.String("port", "9001", "server port")
)

func main() {
	flag.Parse()

	formatter := MyPassThroughFormatter{}

	chatServer := server.NewChatServer(formatter)
	log.Fatal(chatServer.Start(*port))
}
