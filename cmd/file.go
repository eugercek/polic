package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/eugercek/polic/pkg/iampolicy"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/eugercek/polic/internal/expander"
)

func File(filename, resultFile string, sortFlag bool) int {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
		return 1
	}

	policy, err := iampolicy.New(file)

	if err != nil {
		log.Fatal(err)
	}

	sts := policy.Statements

	for i, st := range policy.Statements {
		var actions []string
		var setter func(int, []string)
		var elems []string

		if st.Actions != nil {
			setter = func(i int, as []string) {
				sts[i].Actions = as
			}
			elems = st.Actions
		} else if st.NotActions != nil {
			setter = func(i int, nas []string) {
				sts[i].NotActions = nas
			}
			elems = st.NotActions
		} else {
			log.Println("Actions or NotActions must be given.")
			return 1
		}

		for _, str := range elems {
			if strings.Contains(str, "*") {
				exps, _, _ := expander.ExpandAction(str)
				actions = append(actions, exps...)
			} else {
				actions = append(actions, str)
			}

			if sortFlag {
				sort.Strings(actions)
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
