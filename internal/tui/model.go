package tui

import (
    "strconv"
    "strings"

    tea "github.com/charmbracelet/bubbletea"
    "github.com/muesli/termenv"
    "github.com/ariabrams65/go-type/internal/tui/utils"
)

type model struct {
    text string
    incorrect string
    index int
}

func (m model) Init() tea.Cmd {
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {

    case tea.KeyMsg:
        key := msg.String()
        if key == "ctrl+c" {
            return m, tea.Quit
        } else if key == "backspace" {
            if len(m.incorrect) > 0 {
                m.incorrect = m.incorrect[:len(m.incorrect)-1]
            } else if m.index > 0 {
                m.index--
            }            
        } else if len(key) == 1 {
            if key == string(m.text[m.index]) && len(m.incorrect) == 0 {
                m.index++
                if m.index == len(m.text) {
                    return m, tea.Quit
                }
            } else if key != " "{
                m.incorrect += key
            }
        }
    }
    return m, nil
}

func (m model) View() string {
    s := m.text
    if (m.index != len(m.text)) {
        p := termenv.ColorProfile()
        cursor := termenv.String(string(m.text[m.index])).Foreground(p.Color("#000000")).Background(p.Color("#FFFFFF")).String()
        incorrect := termenv.String(m.incorrect).Foreground(p.Color("#FF0000")).Underline().String()
        completed := termenv.String(m.text[:m.index]).Foreground(p.Color("#FFFFFF")).String()
        todo := termenv.String(m.text[m.index + 1:]).Foreground(p.Color("#808080")).String()
        s = completed + incorrect + cursor + todo
    }
    return utils.WrapString(s, 20) + "\n\n" + strconv.Itoa(m.numCorrect()) + "\n"
}

func (m model) numCorrect() int {
    return m.index - strings.Count(m.text[:m.index], " ")
}


func InitialModel() model {
    return model{
        text: "This is a test to see if this works lets see",
        incorrect: "",
        index: 0,
    }
}
