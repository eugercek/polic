package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/eugercek/aws-iam-policy-expander/cmd/expander"
	"github.com/eugercek/aws-iam-policy-expander/cmd/policy"
)

func File(filename, resultFile string) int {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
		return 1
	}

	policy, err := policy.New(file)

	if err != nil {
		log.Fatal(err)
	}

	sts := policy.Statements

	for i, st := range policy.Statements {
		var actions []string
		var setter func(int, []string)
		var elems []string

		if st.Action != nil {
			setter = func(i int, as []string) {
				sts[i].Action = as
			}
			elems = st.Action
		} else if st.NotAction != nil {
			setter = func(i int, nas []string) {
				sts[i].NotAction = nas
			}
			elems = st.NotAction
		} else {
			log.Fatal("Action or NotAction must be given.")
		}

		for _, str := range elems {
			if strings.Contains(str, "*") {
				exps, _, _ := expander.ExpandAction(str)
				actions = append(actions, exps...)
			} else {
				actions = append(actions, str)
			}
			setter(i, actions)
		}
	}

	file.Close()

	f, _ := json.MarshalIndent(policy, "", "\t")
	err = ioutil.WriteFile(resultFile, f, 0644)
	if err != nil {
		fmt.Println(err)
	}
	return 0
}
