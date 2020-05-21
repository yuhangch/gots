package wkt

import (
	"fmt"
	"strings"
	"unicode"
)

type Token struct {
	position int
	value    string
	token    TokenType
}

func NewToken(postion int, value string, t TokenType) Token {
	return Token{
		position: postion,
		value:    value,
		token:    t,
	}
}

type TokenType int

const (
	TEXT TokenType = iota + 1
	LEFT_PAREN
	RIGHT_PAREN
	NUMBER
	COMMA
	EOF
)

type Lexer struct {
	wkt   string
	index int
}

func NewLexer(wkt string) *Lexer {
	return &Lexer{
		wkt:   wkt,
		index: -1,
	}
}

func isLetter(char rune) bool {
	return unicode.IsLetter(char)
}

func isWhiteSpace(char rune) bool {
	return unicode.IsSpace(char)
}

func isInCoordinate(char rune) bool {
	return unicode.IsDigit(char) ||
		char == '.' ||
		char == '-' ||
		char == '+'
}

func (l *Lexer) readCoordinate() string {
	start := l.index
	for {
		if !isInCoordinate(l.nextChar()) {
			break
		}
	}
	l.index--

	return l.wkt[start : l.index+1]

}

func (l *Lexer) readText() string {
	start := l.index
	for {
		if !isLetter(l.nextChar()) {
			break
		}
	}
	l.index--

	return strings.ToUpper(l.wkt[start : l.index+1])

}

func (l *Lexer) nextChar() rune {
	if l.index >= (len(l.wkt) - 1) {
		return 0
	}
	l.index++
	return rune(l.wkt[l.index])
}

func (l *Lexer) nextToken() Token {
	char := l.nextChar()
	position := l.index
	value := ""
	var t TokenType
	switch {
	case char == '(':
		t = LEFT_PAREN
		break
	case char == ',':
		t = COMMA
		break
	case char == ')':
		t = RIGHT_PAREN
		break
	case isLetter(char):
		t = TEXT
		value = l.readText()
		break
	case isInCoordinate(char):
		t = NUMBER
		value = l.readCoordinate()
	case isWhiteSpace(char):
		return l.nextToken()
	case char == 0:
		t = EOF
		break
	default:
		_ = fmt.Errorf("error at%d", char)

	}

	return NewToken(position, value, t)
}
