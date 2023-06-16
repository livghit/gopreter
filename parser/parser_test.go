package parser

import (
	"github.com/livghit/gopreter/ast"
	"github.com/livghit/gopreter/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `
  let x = 5;
  let y = 10; 
  let foobar = 83838384;
  `

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgramm()

	if program == nil {
		t.Fatalf("ParseProgramm() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statement doese not contain 3 statements. got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatements(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatements(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got=%T", s)
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", letStmt.Name.Value)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value is not '%s' . got=%s", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("letStmt.Name.TokenLiteral() not '%s' got=%s", name, letStmt.Name.TokenLiteral())
		return false
	}

	return true
}
