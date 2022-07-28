Simple CLI for expanding your IAM policies' actions.

Expand cryptic policies to achieve least privilege

# Usage

## Expand single action
```sh
$ ./aws-iam-policy-expander -single waf:Get*
GetByteMatchSet
GetChangeToken
GetChangeTokenStatus
GetGeoMatchSet
GetIPSet
GetLoggingConfiguration
GetPermissionPolicy
GetRateBasedRule
GetRateBasedRuleManagedKeys
GetRegexMatchSet
GetRegexPatternSet
GetRule
GetRuleGroup
GetSampledRequests
GetSizeConstraintSet
GetSqlInjectionMatchSet
GetWebACL
GetXssMatchSet
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

Add whole policy expander