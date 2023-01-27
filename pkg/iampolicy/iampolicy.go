package iampolicy

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

const (
	oldVersion = "2008-10-17"
	newVersion = "2012-10-17"
)

type Policy struct {
	Version    string           `json:"Version"`
	Id         string           `json:"ID,omitempty"`
	Statements StatementOrArray `json:"Statement"`
}

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

	*s = v

	return err
}

type StatementOrArray []Statement

func (s *StatementOrArray) UnmarshalJSON(d []byte) error {
	if d[0] == '{' {
		var v Statement
		err := json.Unmarshal(d, &v)
		*s = StatementOrArray{v}
		return err

	}

	var v []Statement

	err := json.Unmarshal(d, &v)

	*s = v

	return err
}

func New(r io.Reader) (*Policy, error) {
	bytes, err := io.ReadAll(r)

	if err != nil {
		return nil, err
	}

	var policy Policy

	err = json.Unmarshal(bytes, &policy)

	if err != nil {
		return nil, err
	}

	// Check if policy is correct
	// TODO: Check if all Sids are unique
	// TODO: Check is 2008-10-17 versioned policy is valid
	if policy.Version != oldVersion && policy.Version != newVersion {
		return nil, errors.New("wrong Version")
	} else if len(policy.Statements) == 0 {
		return nil, errors.New("there should be at least one Statement")
	}

	for _, st := range policy.Statements {
		if !isExclusive(st.Action, st.NotAction) {
			return nil, fmt.Errorf("Statement %v: Action/NotAction error", st)
		}

		if !isExclusive(st.Resources, st.NotResources) {
			return nil, fmt.Errorf("Statement %v: Resource/NotResource error", st)
		}

		if st.Principals != nil && st.NotPrincipals != nil {
			return nil, fmt.Errorf("Statement %v: cannot have both Principal and NotPrincipal", st)
		}
	}

	return &policy, nil
}

func isExclusive(a, b []string) bool {
	if len(a) == 0 && len(b) == 0 {
		return false
	} else if len(a) != 0 && len(b) != 0 {
		return false
	}

	return true
}
