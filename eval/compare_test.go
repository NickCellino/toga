package eval

import "testing"

func TestLtTwoIntsLt(t *testing.T) {
	ltRule := Lt{
		IntVal{8},
		IntVal{19},
	}
	result, _ := ltRule.Eval()
	if !result {
		t.Fatalf("expected less than comparison to be true for 8 and 19")
	}
}

func TestLtTwoIntsGt(t *testing.T) {
	ltRule := Lt{
		IntVal{208},
		IntVal{12},
	}
	result, _ := ltRule.Eval()
	if result {
		t.Fatalf("expected less than comparison to be false for 208 and 12")
	}
}

func TestLtIntVsFloat(t *testing.T) {
	ltRule := Lt{
		IntVal{100},
		FloatVal{100.18},
	}
	result, _ := ltRule.Eval()
	if !result {
		t.Fatalf("expected less than comparison to be true for 100 and 100.18")
	}
}

func TestLtTwoFloats(t *testing.T) {
	ltRule := Lt{
		FloatVal{99.1315},
		FloatVal{99.000001},
	}
	result, _ := ltRule.Eval()
	if result {
		t.Fatalf("expected less than comparison to be false for 99.1315 and 99.000001")
	}
}
