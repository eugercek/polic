package cmd

import (
	"fmt"
	"strings"

	"github.com/eugercek/aws-iam-policy-expander/cmd/fetch"
)

func Single(action string) int {
	data, err := fetch.GetData()

	if err != nil {
		fmt.Println(err)
		return 1
	}

	actions, base, err := fetch.ExpandAction(action, data)

	if err != nil {
		fmt.Println(err, action)
		return 1
	}

	for _, v := range actions {
		fmt.Println(color(v, base))
	}
	return 0
}

// paint c in s
func color(s string, c string) string {
	return strings.Replace(s, c, fmt.Sprintf("\x1b[32m%s\x1b[0m", c), 1)
}
