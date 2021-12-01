package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
)

type Parser struct {
	lex *lexer.Lexer

	currentToken token.Token
	peekToken    token.Token
}

func New(lex *lexer.Lexer) *Parser {
	p := &Parser{lex: lex}
	// read two tokens, so currentToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.lex.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
