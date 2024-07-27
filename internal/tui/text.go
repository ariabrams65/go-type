package tui

import (
    "strings"

    "github.com/muesli/reflow/wordwrap"
    "github.com/muesli/termenv"
)

func wrapString(s string, limit int) string {
    s = strings.ReplaceAll(s, " ", "|")

    f := wordwrap.NewWriter(limit)
    f.Breakpoints = []rune{'|'}
    f.Write([]byte(s))
    f.Close()

    return strings.ReplaceAll(f.String(), "|", " ")
}

func colorCursor(s string, p termenv.Profile) string {
        return termenv.String(s).
            Foreground(p.Color("#000000")).
            Background(p.Color("#FFFFFF")).
            String()
}

func colorIncorrectText(s string, p termenv.Profile) string {
    return termenv.String(s).
        Foreground(p.Color("#FF0000")).
        Underline().
        String()
}

func colorCompletedText(s string, p termenv.Profile) string {
    return termenv.String(s).
        Foreground(p.Color("#FFFFFF")).
        String()
}

func colorTodoText(s string, p termenv.Profile) string {
    return termenv.String(s).
        Foreground(p.Color("#808080")).
        String()
}
