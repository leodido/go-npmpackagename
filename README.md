# go-npmpackagename

> Check whether a string is a valid npm package name or not

This Go module exports a single function (`Validate(input []byte)`) that takes a byte slice (ie., a string) as input and tells you whether it's a valid `npm` package name or not.

Notice that it also returns warnings to notify you about inputs that were okay by the previous `npm` naming rules but are no more.

For an input package name to be valid for the newer npm naming rules, this function must give no error and no warnings (empty slice).

## Rules

Below the **current** list of rules that valid `npm` package names must conform to:

- length greater than 0
- length less than 214 characters (allowed previously)
- no uppercase characters (allowed previously)
- must not contain any non URL safe characters
- must not start with a dot or an underscore
- must not contain any of characters in `[~'!()*]` (allowed previously)
- must not be the same as a Node.JS core module (allowed previously)
- must not be the same as a string in the blacklist
