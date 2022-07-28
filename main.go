package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const (
	DOWNLOAD_URL  = "https://awspolicygen.s3.amazonaws.com/js/policies.js"
	REMOVE_PREFIX = "app.PolicyEditorConfig="
)

type PolicyDocument struct {
	ServiceMap map[string]Service `json:"serviceMap"`
}

type Service struct {
	StringPrefix string   `json:"StringPrefix"`
	Actions      []string `json:"Actions"`
}

func main() {
	data, err := getData()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for {
		fmt.Print("Enter an AWS action:")
		var inp string
		fmt.Scanln(&inp)
		args := strings.Split(inp, ":")

		if len(args) != 2 {
			fmt.Println("Wrong type of input!")
			continue
		}

		actions := expandAction(args[0], args[1], data)

		for _, v := range actions {
			fmt.Println(v)
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

func expandAction(service string, folded string, data *PolicyDocument) (ret []string) {
	if !strings.Contains(folded, "*") {
		return []string{folded}
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
			ret = append(ret, v)
		}
	}

	return ret
}
