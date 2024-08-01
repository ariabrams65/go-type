package main

import (
    "encoding/json"
    "net"
    "os"

    "github.com/ariabrams65/go-type/internal/tui"
    tea "github.com/charmbracelet/bubbletea"
    "github.com/muesli/termenv"
)

const (
    SERVER_ADDRESS = "localhost:9998"
    SERVER_TYPE = "tcp"
)

func main() {
    restoreConsole, err := termenv.EnableVirtualTerminalProcessing(termenv.DefaultOutput())
    if err != nil {
        panic(err)
    }
    defer restoreConsole()

    conn, err := net.Dial(SERVER_TYPE, SERVER_ADDRESS)
    if err != nil {
        panic(err)
    }

    encoder := json.NewEncoder(conn)
    decoder := json.NewDecoder(conn)

    p := tea.NewProgram(tui.InitialModel(encoder, decoder))
    if _, err := p.Run(); err != nil {
        os.Exit(1)
    }
}
