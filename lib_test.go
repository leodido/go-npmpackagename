package npmpackagename

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	title string
	input []byte
	valid bool
	warns []Warning
	_erro error
}

var cases = []testCase{
	// INVALID
	{
		"empty",
		[]byte(""),
		false,
		[]Warning{},
		ErroEmpty,
	},
	{
		"starting-with-dot",
		[]byte(".startwithdot"),
		false,
		[]Warning{},
		ErroDot,
	},
	{
		"starting-with-underscore",
		[]byte("_underscore"),
		false,
		[]Warning{},
		Erro_,
	},
	{
		"with-colon",
		[]byte("with:colon"),
		false,
		[]Warning{},
		ErroNonURLFriendly,
	},
	{
		"with-slashes",
		[]byte("s/l/a/s/h/e/s"),
		false,
		[]Warning{},
		ErroNonURLFriendly,
	},
	{
		"leading-space",
		[]byte(" space"),
		false,
		[]Warning{},
		ErroTrim,
	},
	{
		"trailing-space",
		[]byte("space "),
		false,
		[]Warning{},
		ErroTrim,
	},
	{
		"blacklist-1",
		[]byte("node_modules"),
		false,
		[]Warning{},
		ErroBlacklist,
	},
	{
		"blacklist-2",
		[]byte("favicon.ico"),
		false,
		[]Warning{},
		ErroBlacklist,
	},
	{
		"invalid-scope",
		[]byte("@a space/pack"),
		false,
		[]Warning{},
		ErroNonURLFriendly,
	},
	{
		"invalid-scoped-package",
		[]byte("@scope/💩"),
		false,
		[]Warning{},
		ErroNonURLFriendly,
	},
	// VALID WITHOUT WARNINGS
	// NEWER NPM NAMING RULES
	{
		"with-dash",
		[]byte("some-package"),
		true,
		[]Warning{},
		nil,
	},
	{
		"with-dot",
		[]byte("example.com"),
		true,
		[]Warning{},
		nil,
	},
	{
		"with-underscore",
		[]byte("under_score"),
		true,
		[]Warning{},
		nil,
	},
	{
		"starting-with-numbers",
		[]byte("123numeric"),
		true,
		[]Warning{},
		nil,
	},
	{
		"max",
		[]byte("ifyouwanttogetthesumoftwonumberswherethosetwonumbersarechosenbyfindingthelargestoftwooutofthreenumbersandsquaringthemwhichismultiplyingthembyitselfthenyoushouldinputthreenumbersintothisfunctionanditwilldothatforyou"),
		true,
		[]Warning{},
		nil,
	},
	{
		"with-scope",
		[]byte("@org/pack"),
		true,
		[]Warning{},
		nil,
	},
	// VALID WITH WARNINGS
	// PREVIOUS NPM NAMING RULES (ONLY)
	{
		"with-special-chars!",
		[]byte("special!"),
		true,
		[]Warning{
			WarnSpecialChars,
		},
		nil,
	},
	{
		"with-special-chars'",
		[]byte("special'"),
		true,
		[]Warning{
			WarnSpecialChars,
		},
		nil,
	},
	{
		"with-special-chars*",
		[]byte("special*"),
		true,
		[]Warning{
			WarnSpecialChars,
		},
		nil,
	},
	{
		"with-special-chars(",
		[]byte("special("),
		true,
		[]Warning{
			WarnSpecialChars,
		},
		nil,
	},
	{
		"with-special-chars)",
		[]byte("special)"),
		true,
		[]Warning{
			WarnSpecialChars,
		},
		nil,
	},
	{
		"with-special-chars~",
		[]byte("special~"),
		true,
		[]Warning{
			WarnSpecialChars,
		},
		nil,
	},
	{
		"builtin-1",
		[]byte("http"),
		true,
		[]Warning{
			WarnBuiltin,
		},
		nil,
	},
	{
		"builtin-2",
		[]byte("process"),
		true,
		[]Warning{
			WarnBuiltin,
		},
		nil,
	},
	{
		"too-long",
		[]byte("ifyouwanttogetthesumoftwonumberswherethosetwonumbersarechosenbyfindingthelargestoftwooutofthreenumbersandsquaringthemwhichismultiplyingthembyitselfthenyoushouldinputthreenumbersintothisfunctionanditwilldothatforyou-"),
		true,
		[]Warning{
			WarnTooLong,
		},
		nil,
	},
	{
		"mixed-case",
		[]byte("CAPITAL-LETTERS"),
		true,
		[]Warning{
			WarnMixCase,
		},
		nil,
	},
}

func TestIsValid(t *testing.T) {
	t.Helper()

	for _, tc := range cases {
		tc := tc

		t.Run(tc.title, func(t *testing.T) {
			ok, w, e := Validate(tc.input)

			if !tc.valid {
				assert.False(t, ok)
				assert.Error(t, e)
				assert.Equal(t, tc._erro, e)
			} else {
				assert.True(t, ok)
				assert.Nil(t, e)
			}

			assert.Equal(t, tc.warns, w)
		})
	}
}
