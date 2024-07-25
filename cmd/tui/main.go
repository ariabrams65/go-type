package main

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/termenv"
    "github.com/ariabrams65/go-type/internal/tui"
)
func main() {
    restoreConsole, err := termenv.EnableVirtualTerminalProcessing(termenv.DefaultOutput())
    if err != nil {
        panic(err)
    }
    defer restoreConsole()

    p := tea.NewProgram(tui.InitialModel())
    if _, err := p.Run(); err != nil {
        os.Exit(1)
    }
}
