package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const (
	DOWNLOAD_URL  = "https://awspolicygen.s3.amazonaws.com/js/policies.js"
	REMOVE_PREFIX = "app.PolicyEditorConfig="
	NO_COLOR      = false
)

type PolicyDocument struct {
	ServiceMap map[string]Service `json:"serviceMap"`
}

type Service struct {
	StringPrefix string   `json:"StringPrefix"`
	Actions      []string `json:"Actions"`
}

func main() {
	single := flag.Bool("single", false, "convert single")
	flag.Parse()

	fmt.Println("Downloading policies...")
	data, err := getData()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if *single {
		if flag.Args() == nil {
			fmt.Println("No action given")
		}

		actions, err := expandAction(flag.Args()[0], data)

		if err != nil {
			fmt.Println(err, flag.Args()[0])
		} else {
			for _, v := range actions {
				fmt.Println(v)
			}
		}
	} else {
		for {
			fmt.Print("Enter an AWS action:")
			var inp string
			fmt.Scanln(&inp)

			if inp == "exit" {
				break
			}

			actions, err := expandAction(inp, data)

			if err != nil {
				fmt.Println(err)
				continue
			}

			for _, v := range actions {
				fmt.Println(v)
			}
		}
	}
}

func getData() (data *PolicyDocument, err error) {
	resp, err := http.Get(DOWNLOAD_URL)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	body = body[len(REMOVE_PREFIX):] // It's used for editor config

	err = json.Unmarshal(body, &data)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func expandAction(inp string, data *PolicyDocument) (ret []string, err error) {
	args := strings.Split(inp, ":")

	if len(args) != 2 {
		return nil, errors.New("wrong type of input")
	}

	service := args[0]
	folded := args[1]

	if !strings.Contains(folded, "*") {
		return []string{folded}, nil
	}

	var actions []string

	for _, v := range data.ServiceMap {
		if v.StringPrefix == service {
			actions = v.Actions
			break
		}
	}

	s := strings.Replace(folded, "*", "", 1)

	// TODO Optimize
	for _, v := range actions {
		if strings.Contains(v, s) {
			if NO_COLOR {
				ret = append(ret, v)
			} else {
				ret = append(ret, colored(v, s))
			}
		}
	}

	return ret, nil
}

// paint c in s
func colored(s string, c string) string {
	return strings.Replace(s, c, fmt.Sprintf("\x1b[32m%s\x1b[0m", c), 1)
}
