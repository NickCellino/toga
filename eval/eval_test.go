package eval

import (
	"testing"
)

func TestEvalNoRules(t *testing.T) {
	rules := make([]Evalable, 0)
	result := Eval(rules)
	if !result {
		t.Fatalf("Eval with empty rules slice should return true")
	}
}

func TestEqRuleTwoEqualConstants(t *testing.T) {
	rules := []Evalable{
		Eq{
			[]Resolvable{
				IntVal{1},
				IntVal{1},
			},
		},
	}
	result := Eval(rules)
	if !result {
		t.Fatalf("Eval should return false when comparing 2 equal constant values.")
	}
}

func TestEqRuleTwoUnequalConstants(t *testing.T) {
	rules := []Evalable{
		Eq{
			[]Resolvable{
				IntVal{1},
				IntVal{2},
			},
		},
	}
	result := Eval(rules)
	if result {
		t.Fatalf("eval should return true when comparing 2 unequal constant values")
	}
}

func TestEqRuleTypeError(t *testing.T) {
	rules := []Evalable{
		Eq{
			[]Resolvable{
				IntVal{1},
				BoolVal{true},
			},
		},
	}
	result := Eval(rules)
	if result {
		t.Fatalf("eval should return false for values of 2 different types")
	}
}
