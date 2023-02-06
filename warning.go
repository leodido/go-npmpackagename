package npmpackagename

// see https://www.digitalocean.com/community/tutorials/how-to-add-extra-information-to-errors-in-go

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
