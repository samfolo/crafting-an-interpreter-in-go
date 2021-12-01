package parser_test

import (
	"testing"

	"monkey/ast"
	"monkey/lexer"
	"monkey/parser"
)

func TestLetStatements(t *testing.T) {
	input := `
	let x 5;
	let = 10;
	let 838383;
	`

	lex := lexer.New(input)
	p := parser.New(lex)

	program := p.ParseProgram()
	checkParserErrors(t, p)
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, test := range tests {
		statement := program.Statements[i]
		if !testLetStatement(t, statement, test.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, statement ast.Statement, name string) bool {
	if statement.TokenLiteral() != "let" {
		t.Errorf("statement.TokenLiteral() not 'let', got=%q", statement.TokenLiteral())
		return false
	}

	letStatement, ok := statement.(*ast.LetStatement)
	if !ok {
		t.Errorf("statement not *ast.LetStatement. got=%q", statement.TokenLiteral())
		return false
	}

	if letStatement.Name.Value != name {
		t.Errorf("letStatement.Name.TokenLiteral() not '%s'. got=%q", name, letStatement.Name.Value)
		return false
	}

	return true
}

func checkParserErrors(t *testing.T, p *parser.Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser error: %d", len(errors))
	for _, message := range errors {
		t.Errorf("parser error: '%q'", message)
	}
	t.FailNow()
}

func TestReturnStatements(t *testing.T) {
	input := `
	return 5;
	return 10;
	return 993322;
	`

	lex := lexer.New(input)
	p := parser.New(lex)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does nit contain 3 statements. got='%q'", len(program.Statements))
	}

	for _, statement := range program.Statements {
		returnStatement, ok := statement.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("statements not *ast.ReturnStatement. got='%q'", statement)
			continue
		}

		if returnStatement.TokenLiteral() != "return" {
			t.Errorf("returnStatement.TokenLiteral() is not 'return'. got='%q'", returnStatement.TokenLiteral())
		}
	}
}
