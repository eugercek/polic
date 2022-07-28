Simple CLI for expanding your IAM policies' actions.

Expand cryptic policies to achieve least privilege

# Usage

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

- Handle sum type problem in policy type
- Add sort option
