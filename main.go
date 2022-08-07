package main

import (
	"os"

	"github.com/eugercek/aws-iam-policy-expander/cli"
)

func main() {
	os.Exit(cli.Run())
}
