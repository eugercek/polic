Simple CLI for expanding your IAM policies' actions.

Expand cryptic policies to achieve least privilege

# Usage

## Expand policy file and save as

```sh
$ ./aws-iam-policy-expander -file=foo.json -out=out.json
```

## Expand policy file and change the file

```sh
$ ./aws-iam-policy-expander -file=foo.json -inline
```

## Expand single action

```sh
$ ./aws-iam-policy-expander -single waf:Get*
waf:GetByteMatchSet
waf:GetChangeToken
waf:GetChangeTokenStatus
waf:GetGeoMatchSet
waf:GetIPSet
waf:GetLoggingConfiguration
waf:GetPermissionPolicy
waf:GetRateBasedRule
waf:GetRateBasedRuleManagedKeys
waf:GetRegexMatchSet
waf:GetRegexPatternSet
waf:GetRule
waf:GetRuleGroup
waf:GetSampledRequests
waf:GetSizeConstraintSet
waf:GetSqlInjectionMatchSet
waf:GetWebACL
waf:GetXssMatchSet
```

## Expand in a repl

```sh
$ ./aws-iam-policy-expander
Enter an AWS action: (enter-something)
.... # results
Enter an AWS action: (enter-something)
.... # results
Enter an AWS action: exit # Or <Ctrl-C>
```

# TODO

- Add sort option
