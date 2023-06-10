package ast

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expresion interface {
	Node
	expresionNode()
}

type Programm struct {
	Statements []Statement
}

func (p *Programm) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}
