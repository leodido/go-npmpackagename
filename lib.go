package npmpackagename

import (
	"bytes"
	"regexp"

	"github.com/leodido/go-encodeuricomponent"
)

var withScopeRe = regexp.MustCompile(`^(?:@(?P<scope>[^/]+?)[/])?(?P<name>[^/]+?)$`)

// Validate validates the input npm package name.
//
// Notice that npm versions follow stricter naming rules.
// For an input package name to be valid for newer npm rules,
// this function must give no error AND no warnings (nil).
//
// When this function returns an error, the input package name
// is invalid even for previous npm naming rules.
//
// In case this function returns warnings only, then the input
// package name is only valid for the previous npm naming rules.
func Validate(input []byte) (bool, []Warning, error) {
	warnings := []Warning{}

	if len(input) == 0 {
		return false, warnings, ErroEmpty
	}

	if bytes.HasPrefix(input, []byte(".")) {
		return false, warnings, ErroDot
	}

	if bytes.HasPrefix(input, []byte("_")) {
		return false, warnings, Erro_
	}

	if !bytes.Equal(bytes.Trim(input, " "), input) {
		return false, warnings, ErroTrim
	}

	if _, found := blacklist[string(input)]; found {
		return false, warnings, ErroBlacklist
	}

	packName := string(input)
	matches := withScopeRe.FindStringSubmatch(string(input))
	if len(matches) > 1 {
		// Check the scope name (if any) for non URL-friendly characters
		scopeIndex := withScopeRe.SubexpIndex("scope")
		if scopeIndex < len(matches) {
			if scope := encodeuricomponent.EncodeURIComponent(matches[scopeIndex]); scope != matches[scopeIndex] {
				return false, warnings, ErroNonURLFriendly
			}
		}

		nameIndex := withScopeRe.SubexpIndex("name")
		if nameIndex < len(matches) {
			packName = matches[nameIndex]
		}
	}
	// Always check package name for non URL-friendly characters
	name := encodeuricomponent.EncodeURIComponent(packName)
	if name != packName {
		return false, warnings, ErroNonURLFriendly
	}

	if len(input) > 214 {
		warnings = append(warnings, WarnTooLong)
	}

	if _, found := builtins[string(input)]; found {
		warnings = append(warnings, WarnBuiltin)
	}

	if !bytes.Equal(bytes.ToLower(input), input) {
		warnings = append(warnings, WarnMixCase)
	}

	// TODO > Check whether the special chars are allowed in the scope part or not (otherwise use packName here)
	if bytes.ContainsAny(input, "~'!()*") {
		warnings = append(warnings, WarnSpecialChars)
	}

	return true, warnings, nil
}
