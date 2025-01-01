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
	return fmt.Sprintf("%v %v %v", t.Type, t.Lexeme, t.Literal)
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
