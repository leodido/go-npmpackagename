package npmpackagename

type Warning int

const (
	WarnBuiltin = iota
	WarnTooLong
	WarnMixCase
	WarnSpecialChars
)

func (w Warning) String() string {
	switch w {
	case WarnBuiltin:
		return "using node builtin names is not allowed anymore"
	case WarnTooLong:
		return "names can no longer contain more than 214 characters"
	case WarnMixCase:
		return "names can no longer contain capital letters"
	case WarnSpecialChars:
		return "names can no longer contain special characters [~'!()*]"
	}

	return ""
}
