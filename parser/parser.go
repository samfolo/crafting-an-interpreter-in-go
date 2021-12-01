package parser

import (
	"fmt"

	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
)

type Parser struct {
	lex    *lexer.Lexer
	errors []string

	currentToken token.Token
	peekToken    token.Token
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(ttype token.TokenType) {
	message := fmt.Sprintf("expcted next token to be %s, got %s instead", ttype, p.peekToken.Type)
	p.errors = append(p.errors, message)
}

func New(lex *lexer.Lexer) *Parser {
	p := &Parser{lex: lex, errors: []string{}}
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
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for !p.currentTokenIs(token.EOF) {
		statement := p.parseStatement()
		if statement != nil {
			program.Statements = append(program.Statements, statement)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.currentToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() ast.Statement {
	statement := &ast.LetStatement{Token: p.currentToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	statement.Name = &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// skip the expression for now, will handle later
	for !p.currentTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return statement
}

func (p *Parser) currentTokenIs(ttype token.TokenType) bool {
	return p.currentToken.Type == ttype
}

func (p *Parser) peekTokenIs(ttype token.TokenType) bool {
	return p.peekToken.Type == ttype
}

func (p *Parser) expectPeek(ttype token.TokenType) bool {
	if p.peekTokenIs(ttype) {
		p.nextToken()
		return true
	} else {
		p.peekError(ttype)
		return false
	}
}
