package lexer

import (
	"monkey/token"
)

/*
Monkey Lexer는 ASCLL를 기본으로 한다.
*/

// Lexer
type Lexer struct {
	input        string
	position     int  // 입력에서 현재 위치 (현재 문자를 가라킨다. [index])
	readPosition int  // 입력에서 현재 읽는 위치 (현재 문자의 다음을 가리킨다.[position + 1])
	ch           byte // 현재 조사하고 있는 문자
}

func new(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// 렉서가 현재 보고 있는 위치를 다음으로 이동하는 메서드 입니다.
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // NULL (ASCLL)
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

// token
// 현재 검사하고 있는 문자(l.ch)를 보고 이에 대응하는 토큰을 반환한다.
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}

	// 반환하기 전에, 입력 문자열을 가리키는 position과 readPosition을 증가시킴으로써, NextToken() 함수를 실행시킬 때 이미 l.ch 필드는 업데이트가 된 상태가 된다.
	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
