Simple CLI for expanding your IAM policies' actions.

Expand cryptic policies to achieve least privilege

# Usage

## Expand policy file and save as

```sh
$ ./polic --file=foo.json --out=out.json # Also can --sort
```

## Expand policy file and change the file

```sh
$ ./polic --file=foo.json --inline --sort
```

## Expand single action

```sh
$ ./polic -single waf:Get*
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
$ ./polic
Enter an AWS action: (enter-something)
.... # results
Enter an AWS action: (enter-something)
.... # results
Enter an AWS action: exit # Or <Ctrl-C>
```

# TODO

