// tocken/token.go
package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	// Special keywords for state
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	//Identifiers and literals
	IDENT = "IDENT" // add , foo , x ,y -> Variables in general ?
	INT   = "INT"   // integer like 1 ,2 ,3 ,4

	// Operators
	ASSING = "ASSING"
	PLUS   = "PLUS"
	MINUS  = "MINUS"

	//Delimiters ( , ), { , } , ;
	COMMA   = "COMMA"
	LPAR    = "LPAR"
	RPAR    = "RPAR"
	LSQURLY = "LSQURLY"
	RSQURLY = "RSQURLY"

	// keywords

	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupIdentifier(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
