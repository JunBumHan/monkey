package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT" // 식별자 (add, foobar, x, y)
	INT   = "INT"   // 리터럴 (12313)

	// 연산자
	ASSIGN = "="
	PLUS   = "+"

	// 구분자
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// 예약어
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

/*

ILLEGAL : 불법적인 ( 잘못된 토큰을 나타내는 상수입니다)
EOF : end of line (소스코드의 킅을 나타내는 토큰입니다.)


IDENT[Identifier] : 식별자
ASSIGN : 할당하다, 배당하다

COMMA : 쉼표
SEMICOLON : 세미콜론

LPAREN[left parenthesis]
RPAREN[right parenthesis]
LBRACE[left brace]
RBRACE[right brace]

FUNCTION : FUNCTION
LET : LET

*/
