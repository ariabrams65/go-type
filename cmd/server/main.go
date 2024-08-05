package main

import (
	"fmt"
	"net"

	"github.com/ariabrams65/go-type/internal/messages"
)

const (
    SERVER_ADDRESS = "localhost:9998"
    SERVER_TYPE = "tcp"
)

func main() {
    server, err := net.Listen(SERVER_TYPE, SERVER_ADDRESS)
    if err != nil {
        panic(err)
    }
    defer server.Close()
    conn, err := server.Accept()
    if err != nil {
        panic(err)
    }
    for {
        msg, err := messages.DecodeMessage(conn)
        if err != nil {
            panic(err)
        }
        switch msg := msg.(type) {
        case messages.PositionMessage:
            fmt.Print(msg.Index) 
            messages.EncodeMessage(messages.TextMessage{
                Text: fmt.Sprintf("This is the text being send back from the server. You are currently on index %d", msg.Index),
            }, conn)

        case messages.JoinMessage:
            fmt.Println()
            fmt.Println(msg.Username)
            fmt.Println(msg.Roomname)
        }
    }
}
