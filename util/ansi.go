package util

import "fmt"

func RGB(r, g, b byte) func(text string) string {
	return func(text string) string {
		return fmt.Sprintf(
			"\u001b[38;2;%d;%d;%dm%s\u001b[0m",
			r, g, b, text,
		)
	}
}
