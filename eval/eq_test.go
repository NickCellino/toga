package eval

import (
	"testing"
)

func TestEqRuleTwoEqualConstants(t *testing.T) {
	expression := Eq{
		[]Expression{
			NumberValue{1},
			NumberValue{1},
		},
	}
	result, _ := expression.Eval(Context{})
	if r, _ := result.AsBool(); !r {
		t.Fatalf("eq should return true when comparing 2 equal constant values.")
	}
}

func TestEqRuleTwoUnequalConstants(t *testing.T) {
	expression := Eq{
		[]Expression{
			NumberValue{1},
			NumberValue{2},
		},
	}
	result, _ := expression.Eval(Context{})
	if r, _ := result.AsBool(); r {
		t.Fatalf("eval should return false when comparing 2 unequal constant values")
	}
}

func TestEqRuleMultipleConstants(t *testing.T) {
	expression := Eq{
		[]Expression{
			NumberValue{2},
			NumberValue{2},
			NumberValue{2},
			NumberValue{2},
		},
	}
	result, _ := expression.Eval(Context{})
	if r, _ := result.AsBool(); !r {
		t.Fatalf("eval should return true when comparing multiple equal constant values")
	}
}

func TestEqRuleTypeError(t *testing.T) {
	expression := Eq{
		[]Expression{
			NumberValue{1},
			BoolValue{true},
		},
	}
	_, err := expression.Eval(Context{})
	if err == nil {
		t.Fatal("expected to get type error when comparing a number to a bool")
	}
}
