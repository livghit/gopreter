package lexer

import (
	"github.com/livghit/gopreter/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5 ;
      let ten = 10 ;
          let add = fn( x ,y ) {
        x + y ;
      } ;
        let result = add(five , ten) ;
        !-/*5
        
        if ( 5 < 10 ) {
      return true
      } else {
      return false
      }
      
      10 == 10
      10 != 9 
      `
	test := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSING, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSING, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSING, "="},
		{token.FUNCTION, "fn"},
		{token.LPAR, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAR, ")"},
		{token.LSQURLY, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RSQURLY, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSING, "="},
		{token.IDENT, "add"},
		{token.LPAR, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAR, ")"},
		{token.SEMICOLON, ";"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTER, "*"},
		{token.INT, "5"},
		{token.IF, "if"},
		{token.LPAR, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAR, ")"},
		{token.LSQURLY, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.RSQURLY, "}"},
		{token.ELSE, "else"},
		{token.LSQURLY, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.RSQURLY, "}"},
		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.INT, "10"},
		{token.NOT_EQ, "!="},
		{token.INT, "9"},
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
