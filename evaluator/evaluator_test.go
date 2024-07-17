package evaluator

import (
	"monkey/lexer"
	"monkey/object"
	"monkey/parser"
	"testing"
)

func TestEvalIntegetExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	return Eval(program)
}

func testIntegerObject(t *testing.T, obj object.Object, expexted int64) bool {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("Object is not Integer got=%T (%+v)", obj, obj)
	}
	if result.Value != expexted {
		t.Errorf("Object is wrong value. got %d wanted=%d", result.Value, expexted)
		return false
	}
	return true
}
