package iampolicy

import (
	"encoding/json"
	"testing"
)

func TestMarshall(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		t.Parallel()
		policyStr := `
{
    "Version": "2012-10-17",
    "Statement": {
            "Effect": "Allow",
            "Actions": "s3:*",
            "Resource": "*"
	}
}
`
		var policy Policy
		json.Unmarshal([]byte(policyStr), &policy)

		if policy.Version == oldVersion {
			t.Fail()
		}

		if policy.Statements[0].Actions[0] != "s3:*" {
			t.Fail()
		}

		if policy.Statements[0].Resources[0] != "*" {
			t.Fail()
		}
	})
	t.Run("Array", func(t *testing.T) {
		t.Parallel()
		policyStr := `
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Actions": ["s3:*", "ec2:*"],
            "Resource": ["*", "arn:aws:ec2:*"]
        }
    ]
}
`
		var policy Policy
		json.Unmarshal([]byte(policyStr), &policy)

		if policy.Version == oldVersion {
			t.Fail()
		}

		if policy.Statements[0].Actions[0] != "s3:*" &&
			policy.Statements[0].Actions[1] != "ec2:*" {
			t.Fail()
		}

		if policy.Statements[0].Resources[0] != "*" &&
			policy.Statements[9].Resources[1] != "arn:aws:ec2:*" {
			t.Fail()
		}
	})
}
