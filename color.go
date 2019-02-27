package color

import "fmt"

func AddColor(s string, color int) string {
	return fmt.Sprintf("\x1b[38;5;%dm%s\x1b[0m", color, s)
}