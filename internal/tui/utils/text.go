package utils

import (
    "strings"

    "github.com/muesli/reflow/wordwrap"
)

func WrapString(s string, limit int) string {
    s = strings.ReplaceAll(s, " ", "|")

    f := wordwrap.NewWriter(limit)
    f.Breakpoints = []rune{'|'}
    f.Write([]byte(s))
    f.Close()

    return strings.ReplaceAll(f.String(), "|", " ")
}
