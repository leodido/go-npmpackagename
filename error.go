package npmpackagename

import (
	"errors"
)

// Define sentinel errors
var (
	ErroEmpty          = errors.New("name lenght must be greater than zero")
	ErroDot            = errors.New("name cannot start with a dot")
	Erro_              = errors.New("name cannot start with an underscore")
	ErroNonURLFriendly = errors.New("name can only contain URL friendly characters")
	ErroBlacklist      = errors.New("name is in blacklist")
	ErroTrim           = errors.New("name cannot containg leading or trailing spaces")
)
