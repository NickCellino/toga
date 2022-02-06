package eval

import (
	"fmt"
)

type Eq struct {
	args []Expression
}

func (eq Eq) Eval(context Context) (Value, error) {
	if len(eq.args) == 0 {
		return BoolValue{true}, nil
	}
	_, err := eq.args[0].Eval(context)
	if err != nil {
		return BoolValue{false}, fmt.Errorf("error evaluating arg 0 of Eq")
	}
	// TODO
	return BoolValue{true}, nil
}
