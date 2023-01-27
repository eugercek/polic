package internal

import (
	"errors"
	"strings"
	"sync"
)

const (
	DownloadUrl  = "https://awspolicygen.s3.amazonaws.com/js/policies.js"
	RemovePrefix = "app.PolicyEditorConfig="
)

var GlobalDocument Document

type Document struct {
	Doc *PolicyDocument
	m   sync.Mutex
}

func (d *Document) Set(document *PolicyDocument) {
	d.m.Lock()
	defer d.m.Unlock()

	d.Doc = document
}
func (d *Document) Get() *PolicyDocument {
	d.m.Lock()
	defer d.m.Unlock()

	return d.Doc
}

type PolicyDocument struct {
	ServiceMap map[string]Service `json:"serviceMap"`
}

type Service struct {
	StringPrefix string   `json:"StringPrefix"`
	Actions      []string `json:"Actions"`
}

func ExpandAction(inp string) (ret []string, str string, err error) {
	doc := GlobalDocument.Get()
	args := strings.Split(inp, ":")

	if len(args) != 2 {
		return nil, "", errors.New("wrong type of input")
	}

	service := args[0]
	folded := args[1]

	if !strings.Contains(folded, "*") {
		return []string{folded}, "", nil
	}

	var actions []string

	for _, v := range doc.ServiceMap {
		if v.StringPrefix == service {
			actions = v.Actions
			break
		}
	}

	// strings.Contains("foo", "") -> true
	s := strings.Replace(folded, "*", "", 1)

	// TODO Optimize
	for _, a := range actions {
		if strings.Contains(a, s) {
			ret = append(ret, service+":"+a)
		}
	}

	return ret, folded, nil
}
