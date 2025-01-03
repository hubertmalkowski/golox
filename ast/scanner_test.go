package ast

import (
	"golox/runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScanner(t *testing.T) {
	runtime := runtime.NewRuntime()

	var tests = []struct {
		name     string
		expected TokenList
		code     string
	}{
		{
			name: "assign val",
			expected: TokenList{
				NewToken(IdentifierTT, "x", "x", 1),
				NewToken(EqualTT, "=", nil, 1),
				NewToken(NumberTT, "1", 1.0, 1),
				NewToken(EOFTT, "", nil, 1),
			},
			code: "x = 1",
		},
		{
			name: "assign string",
			code: "x = \"hello\"",
			expected: TokenList{
				NewToken(IdentifierTT, "x", "x", 1),
				NewToken(EqualTT, "=", nil, 1),
				NewToken(StringTT, "\"hello\"", "hello", 1),
				NewToken(EOFTT, "", nil, 1),
			},
		},
		{
			name: "if block",
			expected: TokenList{
				NewToken(IfTT, "if", nil, 1),
				NewToken(IdentifierTT, "hoge", "hoge", 1),
				NewToken(LeftBraceTT, "{", nil, 1),
				NewToken(IdentifierTT, "x", "x", 1),
				NewToken(SemicolonTT, ";", nil, 1),
				NewToken(RightBraceTT, "}", nil, 1),
				NewToken(ElseTT, "else", nil, 1),
				NewToken(LeftBraceTT, "{", nil, 1),
				NewToken(IfTT, "if", nil, 1),
				NewToken(IdentifierTT, "piyo", "piyo", 1),
				NewToken(LeftBraceTT, "{", nil, 1),
				NewToken(IdentifierTT, "y", "y", 1),
				NewToken(SemicolonTT, ";", nil, 1),
				NewToken(RightBraceTT, "}", nil, 1),
				NewToken(ElseTT, "else", nil, 1),
				NewToken(LeftBraceTT, "{", nil, 1),
				NewToken(IdentifierTT, "z", "z", 1),
				NewToken(SemicolonTT, ";", nil, 1),
				NewToken(RightBraceTT, "}", nil, 1),
				NewToken(RightBraceTT, "}", nil, 1),
				NewToken(EOFTT, "", nil, 1),
			},
			code: "if hoge { x; } else { if piyo { y; } else { z; } }",
			// if hoge {
			//   x
			// } else {
			//   if piyo {
			//     y;
			//   } else {
			//     z;
			//   }
			// }
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			scanner := NewScanner(tt.code, runtime)
			actual := scanner.ScanTokens()
			if len(actual) != len(tt.expected) {
				t.Errorf("got: \n%v\nwant:\n %v", actual, tt.expected)
			}
			for i := range actual {
				// if  {
				// 	t.Errorf("got: \n%v\nwant:\n%v", actual, tt.expected)
				// }
				assert.Equal(t, tt.expected[i], actual[i])

			}
		})
	}
}
