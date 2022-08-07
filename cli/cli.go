package cli

import (
	"flag"
	"fmt"

	"github.com/eugercek/aws-iam-policy-expander/cmd"
)

func Run() int {
	single := flag.Bool("single", false, "convert single")
	file := flag.Bool("file", false, "expand inline in a file")
	repl := flag.Bool("repl", false, "open in repl mode")

	flag.Parse()

	if *single && !*file && !*repl {
		if flag.Args() == nil {
			fmt.Println("No action given")
			return 1
		}

		return cmd.Single(flag.Args()[0])
	} else if !*single && *file && !*repl {
		if flag.Args() == nil {
			fmt.Println("No file given")
			return 1
		}

		return cmd.File(flag.Args()[0])
	} else if !*single && !*file && *repl {
		return cmd.Repl()
	} else {
		fmt.Println("Wrong flag. Given")
		return 1
	}

}
