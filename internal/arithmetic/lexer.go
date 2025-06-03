package arithmetic

import (
	"strings"
	"unicode"
)

type TokenType int16

const (
	C_Start TokenType = iota
	C_End
	C_Num
	C_LeftBracket
	C_RightBracket
	C_Plus
	C_Minus
	C_Multiply
	C_Divide
)

type lexer struct {
	text   []byte
	cursor int
}

func NewLexer(text string) *lexer {
	text = strings.ReplaceAll(text, " ", "")
	return &lexer{
		text: []byte(text),
	}
}
func (r *lexer) nextToken() (TokenType, ExpType) {
	cur := r.cursor
	if len(r.text) == cur {
		return C_End, -1
	}
	b := r.text[cur]
	cur++
	if unicode.IsDigit(rune(b)) {
		x := ExpType(b - '0')
		for len(r.text) != cur && unicode.IsDigit(rune(r.text[cur])) {
			x = x*10 + ExpType(r.text[cur]-'0')
			cur++
		}
		return C_Num, x
	}
	switch b {
	case '(':
		return C_LeftBracket, 0
	case ')':
		return C_RightBracket, 0
	case '+':
		return C_Plus, 0
	case '-':
		return C_Minus, 0
	case '*':
		return C_Multiply, 0
	case '/':
		return C_Divide, 0
	}

	return C_End, -1
}
func (r *lexer) consumeToken() {
	cur := r.cursor
	if len(r.text) == cur {
		return
	}
	b := r.text[cur]
	cur++
	if unicode.IsDigit(rune(b)) {
		x := ExpType(b - '0')
		for len(r.text) != cur && unicode.IsDigit(rune(r.text[cur])) {
			x = x*10 + ExpType(r.text[cur]-'0')
			cur++
		}
	}
	r.cursor = cur
}

func (r *lexer) readToken() (TokenType, ExpType) {
	x, kind := r.nextToken()
	r.consumeToken()
	return x, kind
}
