package tui

import (
    "net"
    "strconv"
    "strings"

    "github.com/ariabrams65/go-type/internal/messages"
    tea "github.com/charmbracelet/bubbletea"
    "github.com/muesli/termenv"
)


type model struct {
    text string
    incorrect string
    index int
    username string
    roomname string
    players map[string]int
    conn net.Conn
    err error
}

func (m model) Init() tea.Cmd {
    return tea.Batch(
        sendMessage(messages.JoinMessage{
            Username: m.username, 
            Roomname: m.roomname,
        }, m.conn),
        receiveMessage(m.conn),
    )
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.Type {
        case tea.KeyCtrlC:
            return m, tea.Quit

        case tea.KeyBackspace:
            if len(m.incorrect) > 0 {
                m.incorrect = m.incorrect[:len(m.incorrect)-1]
            } else if m.index > 0 {
                m.index--
                return m, sendMessage(messages.PositionMessage{
                    Username: m.username,
                    Index: m.index,
                }, m.conn)
            }            

        case tea.KeyRunes, tea.KeySpace:
            key := msg.String()
            if key == string(m.text[m.index]) && len(m.incorrect) == 0 {
                m.index++
                if m.index == len(m.text) {
                    return m, tea.Quit
                }
                return m, sendMessage(messages.PositionMessage{
                    Username: m.username,
                    Index: m.index,
                }, m.conn)
            } else if key != " "{
                m.incorrect += key
            }
        }

    case messages.TextMessage:
        m.text = msg.Text
        return m, receiveMessage(m.conn)

    case messages.PositionMessage:
        m.players[msg.Username] = msg.Index
        return m, receiveMessage(m.conn)
        
    case errMsg:
        m.err = msg.err
        return m, tea.Quit
        
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
    s = wrapString(s, 20) + "\n\n" + strconv.Itoa(m.numCorrect()) + "\n"
    if m.err != nil {
        s += m.err.Error() + "\n"
    }
    return s
}

func (m model) numCorrect() int {
    return m.index - strings.Count(m.text[:m.index], " ")
}

type errMsg struct {err error}

func sendMessage(m messages.Message, conn net.Conn) tea.Cmd {
    return func() tea.Msg {
        err := messages.EncodeMessage(m, conn)
        if err != nil {
            return errMsg{err}
        }
        return nil
    }
}

func receiveMessage(conn net.Conn) tea.Cmd {
    return func() tea.Msg {
        msg, err := messages.DecodeMessage(conn)
        if err != nil {
            return errMsg{err}
        }
        return msg
    }
}

func InitialModel(roomname string, username string, conn net.Conn) model {
    return model{
        text: "This is a test to see if this works lets see",
        incorrect: "",
        index: 0,
        username: username,
        roomname: roomname,
        players: make(map[string]int),
        conn: conn,
    }
}

