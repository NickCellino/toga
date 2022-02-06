package eval

import "testing"

func TestLtTwoIntsLt(t *testing.T) {
	ltRule := Lt{
		NumberValue{8},
		NumberValue{19},
	}
	result, _ := ltRule.Eval(Context{})
	resultBool, err := result.AsBool()
	if err != nil {
		t.Fatalf("expected result to be a bool")
	}
	if !resultBool {
		t.Fatalf("expected 8 to be Lt 9")
	}
}

func TestLtTwoIntsGt(t *testing.T) {
	ltRule := Lt{
		NumberValue{208},
		NumberValue{12},
	}
	result, _ := ltRule.Eval(Context{})
	resultBool, err := result.AsBool()
	if err != nil {
		t.Fatalf("expected result to be a bool")
	}
	if resultBool {
		t.Fatalf("expected 208 not to be Lt 12")
	}
}

func TestLtIntVsFloat(t *testing.T) {
	ltRule := Lt{
		NumberValue{100},
		NumberValue{100.18},
	}
	result, _ := ltRule.Eval(Context{})
	resultBool, err := result.AsBool()
	if err != nil {
		t.Fatalf("expected result to be a bool")
	}
	if !resultBool {
		t.Fatalf("expected 100 not to be Lt 100.18")
	}
}

func TestLtTwoFloats(t *testing.T) {
	ltRule := Lt{
		NumberValue{99.1315},
		NumberValue{99.000001},
	}
	result, _ := ltRule.Eval(Context{})
	resultBool, err := result.AsBool()
	if err != nil {
		t.Fatalf("expected result to be a bool")
	}
	if resultBool {
		t.Fatalf("expected 99.1315 not to be Lt 99.00001")
	}
}
