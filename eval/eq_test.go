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

/*
func TestEqRuleTypeError(t *testing.T) {
	rules := []Evalable{
		Eq{
			[]Resolvable{
				IntVal{1},
				BoolVal{true},
			},
		},
	}
	result := Eval(rules, Context{})
	if result {
		t.Fatalf("eval should return false for values of 2 different types")
	}
}
*/
