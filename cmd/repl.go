package cmd

import (
	"fmt"
	"github.com/eugercek/polic/internal"
)

func Repl() int {
	for {
		fmt.Print("Enter an AWS action:")
		var inp string
		_, err := fmt.Scanln(&inp)
		if err != nil {
			return 1
		}

		if inp == "exit" {
			break
		}

		actions, base, err := internal.ExpandAction(inp)

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
