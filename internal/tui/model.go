package tui

import (
    "strconv"
    "strings"

    tea "github.com/charmbracelet/bubbletea"
    "github.com/muesli/termenv"
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
        cursor := colorCursor(string(m.text[m.index]), p)
        incorrect := colorIncorrectText(m.incorrect, p)
        completed := colorCompletedText(m.text[:m.index], p)
        todo := colorTodoText(m.text[m.index + 1:], p)

        s = completed + incorrect + cursor + todo
    }
    s = wrapString(s, 20)

    return s + "\n\n" + strconv.Itoa(m.numCorrect()) + "\n"
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
