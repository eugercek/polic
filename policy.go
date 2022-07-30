package main

type Policy struct {
	Version    string      `json:"Version"`
	Id         string      `json:"ID,omitempty"`
	Statements []Statement `json:"Statement"`
}

// TODO Do it correctly
type Statement struct {
	Id            string              `json:"Sid,omitempty"`
	Effect        string              `json:"Effect"`
	Action        interface{}         `json:"Action"`              // string or array
	NotAction     interface{}         `json:"NotAction,omitempty"` // string or array
	Principals    map[string][]string `json:"Principal,omitempty"`
	NotPrincipals map[string][]string `json:"NotPrincipal,omitempty"`
	Resources     []string            `json:"Resource,omitempty"`
	NotResources  []string            `json:"NotResource,omitempty"`
	Condition     []string            `json:"Condition,omitempty"`
}
