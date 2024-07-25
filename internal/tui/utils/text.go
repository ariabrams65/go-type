package utils

import (
    "strings"

    "github.com/muesli/reflow/wordwrap"
    "github.com/muesli/termenv"
)

func WrapString(s string, limit int) string {
    s = strings.ReplaceAll(s, " ", "|")

    f := wordwrap.NewWriter(limit)
    f.Breakpoints = []rune{'|'}
    f.Write([]byte(s))
    f.Close()

    return strings.ReplaceAll(f.String(), "|", " ")
}

func ColorCursor(s string, p termenv.Profile) string {
        return termenv.String(s).
            Foreground(p.Color("#000000")).
            Background(p.Color("#FFFFFF")).
            String()
}

func ColorIncorrectText(s string, p termenv.Profile) string {
    return termenv.String(s).
        Foreground(p.Color("#FF0000")).
        Underline().
        String()
}

func ColorCompletedText(s string, p termenv.Profile) string {
    return termenv.String(s).
        Foreground(p.Color("#FFFFFF")).
        String()
}

func ColorTodoText(s string, p termenv.Profile) string {
    return termenv.String(s).
        Foreground(p.Color("#808080")).
        String()
}
