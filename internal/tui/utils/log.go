package utils

import (
    "fmt"

	tea "github.com/charmbracelet/bubbletea"

)

func Log(s any) {
	f, err := tea.LogToFile("debug.log", "debug")
    if err != nil {
        panic(err)
    }
	defer f.Close()

    _, err = f.WriteString(fmt.Sprint(s) + "\n")
    if err != nil {
        panic(err)
    }
}
