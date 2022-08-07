package cmd

import (
	"fmt"

	"github.com/eugercek/aws-iam-policy-expander/cmd/fetch"
)

func Repl() int {
	data, err := fetch.GetData()
	if err != nil {
		fmt.Println(err)
		return 1
	}

	for {
		fmt.Print("Enter an AWS action:")
		var inp string
		fmt.Scanln(&inp)

		if inp == "exit" {
			break
		}

		actions, base, err := fetch.ExpandAction(inp, data)

		if err != nil {
			fmt.Println(err)
			continue
		}

		for _, v := range actions {
			fmt.Println(color(v, base))
		}
	}
	return 0
}
