package lexer

import (
	"github.com/livghit/gopreter/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `=+(){}-`

	test := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSING, "="},
		{token.PLUS, "+"},
		{token.LPAR, "("},
		{token.RPAR, ")"},
		{token.LSQURLY, "{"},
		{token.RSQURLY, "}"},
		{token.MINUS, "-"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range test {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
