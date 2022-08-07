package cli

import (
	"flag"
	"fmt"

	"github.com/eugercek/aws-iam-policy-expander/cmd"
)

func Run() int {
	single := flag.Bool("single", false, "convert single")
	file := flag.String("file", "", "expand inline in a file")
	repl := flag.Bool("repl", false, "open in repl mode")
	out := flag.String("out", "", "output file name (only for file flag)")
	inline := flag.Bool("inline", false, "change the (input) policy file (only for file flag")

	flag.Parse()

	if *single && *file == "" && !*repl {
		if flag.Args() == nil {
			fmt.Println("No action given")
			return 1
		}

		return cmd.Single(flag.Args()[0])
	} else if !*single && *file != "" && !*repl {
		var resultFile string

		if *out != "" && !*inline {
			resultFile = *out
		} else if *out == "" && *inline {
			resultFile = *file
		} else {
			fmt.Println("Choose either: out or inline")
			return 1
		}

		return cmd.File(*file, resultFile)
	} else if !*single && *file == "" && *repl {
		return cmd.Repl()
	} else {
		fmt.Println("Wrong flag. Given")
		return 1
	}

}
