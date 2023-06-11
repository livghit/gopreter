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
	ASSING = "="
	PLUS   = "+"
	MINUS  = "-"
	BANG   = "!"
	ASTER  = "*"
	SLASH  = "/"
	LT     = "<"
	GT     = ">"
	EQ     = "=="
	NOT_EQ = "!="

	//Delimiters ( , ), { , } , ;
	COMMA     = ","
	SEMICOLON = ";"
	LPAR      = "("
	RPAR      = ")"
	LSQURLY   = "{"
	RSQURLY   = "}"

	// keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	RETURN   = "RETURN"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"true":   TRUE,
	"false":  FALSE,
	"return": RETURN,
}

func LookupIdentifier(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
