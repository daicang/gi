package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT"
	INT   = "INT"

	ASSIGN = "="
	PLUS   = "+"

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "FUNCTiON"
	LET      = "LET"
)

type TokenType string

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

type Token struct {
	Type    TokenType
	Literal string
}

func LookupIdent(ident string) TokenType {
	if identType, ok := keywords[ident]; ok {
		return identType
	}
	return IDENT
}
