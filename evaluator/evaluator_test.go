package evaluator_test

import (
	"testing"

	"monkey/evaluator"
	"monkey/lexer"
	"monkey/object"
	"monkey/parser"
)

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
	}

	for _, test := range tests {
		evaluated := testEval(test.input)
		testIntegerObject(t, evaluated, test.expected)
	}
}

func testEval(input string) object.Object {
	lex := lexer.New(input)
	p := parser.New(lex)
	program := p.ParseProgram()

	return evaluator.Eval(program)
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Fatalf("object was not an *object.Integer. got '%T (%+v)' instead.", obj, obj)
		return false
	}

	if result.Value != expected {
		t.Errorf("obj.Value was not %d. got %d instead", result.Value, expected)
		return false
	}

	return true
}
