package main

import (
	"encoding/json"
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
    // buffer := make([]byte, 1024)
    // for {
    //     mLen, err := conn.Read(buffer)
    //     if err != nil {
    //         panic(err)
    //     }
    //     fmt.Print(string(buffer[:mLen]))
    // }

    decoder := json.NewDecoder(conn)
    for {
        msg, err := messages.DecodeMessage(decoder)
        if err != nil {
            panic(err)
        }
        switch msg := msg.(type) {
        case messages.PositionMessage:
            fmt.Print(msg.Index) 
        }
    }
}
