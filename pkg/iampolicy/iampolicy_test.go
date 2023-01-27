package iampolicy

import (
	"encoding/json"
	"testing"
)

func TestSimplePolicy(t *testing.T) {
	policyStr := `
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": "s3:*",
            "Resource": "*"
        }
    ]
}
`
	var policy Policy
	json.Unmarshal([]byte(policyStr), &policy)

	if policy.Version != oldVersion {
		t.Fail()
	}

	if policy.Statements[0].Actions != []string{"s3:*"} {
		t.Fail()
	}

	if policy.Statements[0].Resources != []string{"*"} {
		t.Fail()
	}
}
