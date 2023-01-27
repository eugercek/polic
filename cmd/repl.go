package cmd

import (
	"fmt"

	"github.com/eugercek/polic/internal/expander"
)

func Repl() int {
	for {
		fmt.Print("Enter an AWS action:")
		var inp string
		fmt.Scanln(&inp)

		if inp == "exit" {
			break
		}

		actions, base, err := expander.ExpandAction(inp)

		if err != nil {
			fmt.Println(err)
			continue
		}

		for _, v := range actions {
			fmt.Println(color(v, base[:len(base)-1]))
		}
	}
	return 0
}
