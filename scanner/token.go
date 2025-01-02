package scanner

import "fmt"

type TokenType int

const (
	// Single-character tokens
	LeftParenTT TokenType = iota
	RightParenTT
	LeftBraceTT
	RightBraceTT
	CommaTT
	DotTT
	MinusTT
	PlusTT
	SemicolonTT
	SlashTT
	StarTT

	// One or two character tokens
	BangTT
	BangEqualTT
	EqualTT
	EqualEqualTT
	GreaterTT
	GreaterEqualTT
	LessTT
	LessEqualTT

	// Literals
	IdentifierTT
	StringTT
	NumberTT

	// Keywords
	AndTT
	ClassTT
	ElseTT
	FalseTT
	FunTT
	ForTT
	IfTT
	NilTT
	OrTT
	PrintTT
	ReturnTT
	SuperTT
	ThisTT
	TrueTT
	VarTT
	WhileTT

	EOFTT
)

var tokenTypeStrings = map[TokenType]string{
	LeftParenTT:    "LeftParenTT",
	RightParenTT:   "RightParenTT",
	LeftBraceTT:    "LeftBraceTT",
	RightBraceTT:   "RightBraceTT",
	CommaTT:        "CommaTT",
	DotTT:          "DotTT",
	MinusTT:        "MinusTT",
	PlusTT:         "PlusTT",
	SemicolonTT:    "SemicolonTT",
	SlashTT:        "SlashTT",
	StarTT:         "StarTT",
	BangTT:         "BangTT",
	BangEqualTT:    "BangEqualTT",
	EqualTT:        "EqualTT",
	EqualEqualTT:   "EqualEqualTT",
	GreaterTT:      "GreaterTT",
	GreaterEqualTT: "GreaterEqualTT",
	LessTT:         "LessTT",
	LessEqualTT:    "LessEqualTT",
	IdentifierTT:   "IdentifierTT",
	StringTT:       "StringTT",
	NumberTT:       "NumberTT",
	AndTT:          "AndTT",
	ClassTT:        "ClassTT",
	ElseTT:         "ElseTT",
	FalseTT:        "FalseTT",
	FunTT:          "FunTT",
	ForTT:          "ForTT",

	IfTT:     "IfTT",
	NilTT:    "NilTT",
	OrTT:     "OrTT",
	PrintTT:  "PrintTT",
	ReturnTT: "ReturnTT",

	SuperTT: "SuperTT",
	ThisTT:  "ThisTT",
	TrueTT:  "TrueTT",
	VarTT:   "VarTT",
	WhileTT: "WhileTT",
	EOFTT:   "EOFTT",
}

type Token struct {
	Type    TokenType
	Lexeme  string
	Literal interface{}
	Line    int
}

func NewToken(t TokenType, lexeme string, literal interface{}, line int) *Token {
	return &Token{Type: t, Lexeme: lexeme, Literal: literal, Line: line}
}

func (t *Token) String() string {
	return fmt.Sprintf("%v %v %v", t.Lexeme, tokenTypeStrings[t.Type], t.Line)
}

var keywords = map[string]TokenType{
	"and":    AndTT,
	"class":  ClassTT,
	"else":   ElseTT,
	"false":  FalseTT,
	"for":    ForTT,
	"fun":    FunTT,
	"if":     IfTT,
	"nil":    NilTT,
	"or":     OrTT,
	"print":  PrintTT,
	"return": ReturnTT,
	"super":  SuperTT,
	"this":   ThisTT,
	"true":   TrueTT,
	"var":    VarTT,
	"while":  WhileTT,
}

type TokenList []*Token

func (tl TokenList) String() string {
	var s string
	for _, t := range tl {
		s += t.String() + "\n"
	}
	return s
}
