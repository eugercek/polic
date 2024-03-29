package cmd

import (
	"fmt"
	"github.com/eugercek/polic/internal"
	"strings"
)

func Single(action string) int {
	actions, base, err := internal.ExpandAction(action)

	if err != nil {
		fmt.Println(err, action)
		return 1
	}

	for _, v := range actions {
		fmt.Println(color(v, base[:len(base)-1]))
	}
	return 0
}

// paint c in s
func color(s string, c string) string {
	return strings.Replace(s, c, fmt.Sprintf("\x1b[32m%s\x1b[0m", c), 1)
}
