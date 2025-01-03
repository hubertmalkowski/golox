package ast

import (
	"fmt"
	"golox/runtime"
	"strconv"
	"strings"
)

type Scanner struct {
	source  string
	reader  *strings.Reader
	runtime *runtime.Runtime
	tokens  TokenList
	start   int
	current int
	line    int
}

func NewScanner(s string, r *runtime.Runtime) *Scanner {
	return &Scanner{
		source:  s,
		reader:  strings.NewReader(s),
		tokens:  []*Token{},
		start:   0, // Character were we started scanning a token
		current: 0, // Cursor which rune we are looking at
		line:    1,
	}
}

func (s *Scanner) ScanTokens() TokenList {
	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}

	s.tokens = append(s.tokens, NewToken(EOFTT, "", nil, s.line))

	return s.tokens
}

func (s *Scanner) scanToken() {
	c := s.advance()
	switch c {
	case '(':
		s.addToken(LeftParenTT)
	case ')':
		s.addToken(RightParenTT)
	case '{':
		s.addToken(LeftBraceTT)
	case '}':
		s.addToken(RightBraceTT)
	case ',':
		s.addToken(CommaTT)
	case '-':
		s.addToken(MinusTT)
	case '+':
		s.addToken(PlusTT)
	case ';':
		s.addToken(SemicolonTT)
	case '*':
		s.addToken(StarTT)
	case '!':
		s.addToken(ifToken(s.match('='), BangEqualTT, BangTT))
	case '=':
		s.addToken(ifToken(s.match('='), EqualEqualTT, EqualTT))
	case '<':
		s.addToken(ifToken(s.match('='), LessEqualTT, LessTT))
	case '>':
		s.addToken(ifToken(s.match('='), GreaterEqualTT, GreaterTT))
	case '/':
		if s.match('/') {
			for s.peek() != '\n' && !s.isAtEnd() {
				s.advance()
			}
			break
		}
		s.addToken(SlashTT)
	case ' ', '\r', '\t':
		break
	case '\n':
		s.line++
	case '"':
		s.string()

	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		s.number()
	default:
		if isAlpha(c) {
			s.identifier()
		} else {
			s.runtime.Error(s.line, fmt.Sprintf("Unexpected character '%x'", c))
		}

	}
}

func (s *Scanner) advance() rune {
	char, size, err := s.reader.ReadRune()
	if err != nil {
		panic(fmt.Sprintf("advance error: %v", err))
	}
	s.current += size
	return char
}

// It’s like a conditional advance(). We only consume the current character if it’s what we’re looking for.
func (s *Scanner) match(expected rune) bool {
	if s.isAtEnd() {
		return false
	}

	char, size, err := s.reader.ReadRune()
	if err != nil {
		panic(fmt.Sprintf("match error: %v", err))
	}

	if char != expected {
		if err := s.reader.UnreadRune(); err != nil {
			panic(fmt.Sprintf("match unread error: %v", err))
		}
		return false
	}

	s.current += size
	return true
}

func (s *Scanner) peek() rune {
	if s.isAtEnd() {
		return 0
	}
	char, _, err := s.reader.ReadRune()
	if err != nil {
		panic(fmt.Sprintf("peek error: %v", err))
	}
	if err := s.reader.UnreadRune(); err != nil {
		panic(fmt.Sprintf("peek unread error: %v", err))
	}
	return char
}

func (s *Scanner) peekNext() rune {
	var (
		i   int
		err error
		ch  rune
	)

	// scan forward
	for ; i < 2 && !s.isAtEnd() && err == nil; i++ {
		ch, _, err = s.reader.ReadRune()
	}

	// unwind
	if _, err = s.reader.Seek(int64(s.current), 0); err != nil {
		panic(fmt.Sprintf("peekNext read error: %v", err))
	}

	// failed to peek n
	if i < 2 {
		return 0
	}

	return ch
}

func (s *Scanner) addToken(tokType TokenType) {
	s.addTokenWithLiteral(tokType, nil)
}

func (s *Scanner) addTokenWithLiteral(tokType TokenType, literal interface{}) {
	lexeme := s.source[s.start:s.current]
	token := NewToken(tokType, lexeme, literal, s.line)
	s.tokens = append(s.tokens, token)
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) string() {
	isEscape := false // define isEscape to handle \"
	for (isEscape || s.peek() != '"') && !s.isAtEnd() {
		if s.peek() == '\n' {
			s.line++
		}

		if s.peek() == '\\' {
			isEscape = !isEscape
		} else {
			isEscape = false
		}

		s.advance()
	}

	if s.isAtEnd() {
		s.runtime.Error(s.line, "Unterminated string")
		return
	}

	s.advance()
	literal := s.source[s.start+1 : s.current-1]
	s.addTokenWithLiteral(StringTT, literal)
}

func (s *Scanner) number() {
	for isDigit(s.peek()) {
		s.advance()
	}

	if s.peek() == '.' && isDigit(s.peekNext()) {
		s.advance()

		for isDigit(s.peek()) {
			s.advance()
		}
	}

	v := s.source[s.start:s.current]
	f, err := strconv.ParseFloat(v, 64)
	if err != nil {
		s.runtime.Error(s.line, fmt.Sprintf("number error: %v", err))
		return
	}
	s.addTokenWithLiteral(NumberTT, f)
}

func (s *Scanner) identifier() {
	for isAlphaNumeric(s.peek()) {
		s.advance()
	}

	literal := s.source[s.start:s.current]
	if keyword, ok := keywords[literal]; ok {
		s.addToken(keyword)
		return
	}
	s.addTokenWithLiteral(IdentifierTT, literal)
}

func ifToken(c bool, a TokenType, b TokenType) TokenType {
	if c {
		return a
	}
	return b
}

func isAlpha(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '_'
}

func isDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}

func isAlphaNumeric(ch rune) bool {
	return isAlpha(ch) || isDigit(ch)
}
