package main

import "encoding/json"

type Policy struct {
	Version    string          `json:"Version"`
	Id         string          `json:"ID,omitempty"`
	Statements StatmentOrArray `json:"Statement"` // Single Statement or Array of Statements
}

// All interface{} are either: string or an array of strings
type Statement struct {
	Id            string                   `json:"Sid,omitempty"`
	Effect        string                   `json:"Effect"`
	Action        StringOrArray            `json:"Action"`
	NotAction     StringOrArray            `json:"NotAction,omitempty"`
	Principals    map[string]StringOrArray `json:"Principal,omitempty"`
	NotPrincipals map[string]StringOrArray `json:"NotPrincipal,omitempty"`
	Resources     StringOrArray            `json:"Resource,omitempty"`
	NotResources  StringOrArray            `json:"NotResource,omitempty"`
	Condition     StringOrArray            `json:"Condition,omitempty"`
}

type StringOrArray []string

func (s *StringOrArray) UnmarshalJSON(d []byte) error {
	if d[0] == '"' {
		var v string
		err := json.Unmarshal(d, &v)
		*s = StringOrArray{v}
		return err

	}

	var v []string

	err := json.Unmarshal(d, &v)

	*s = StringOrArray(v)

	return err
}

type StatmentOrArray []Statement

func (s *StatmentOrArray) UnmarshalJSON(d []byte) error {
	if d[0] == '{' {
		var v Statement
		err := json.Unmarshal(d, &v)
		*s = StatmentOrArray{v}
		return err

	}

	var v []Statement

	err := json.Unmarshal(d, &v)

	*s = StatmentOrArray(v)

	return err
}
