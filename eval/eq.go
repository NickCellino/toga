package eval

import (
	"errors"
	"fmt"
)

type Eq struct {
	Args []Expression
}

func (eq Eq) Eval(context Context) (Value, error) {
	if len(eq.Args) == 0 {
		return BoolValue{true}, nil
	}
	first, err := eq.Args[0].Eval(context)
	if err != nil {
		return BoolValue{false}, errors.New("error evaluating arg 0 of Eq")
	}
	for i, exp := range eq.Args[1:] {
		val, err := exp.Eval(context)
		if err != nil {
			return BoolValue{false}, fmt.Errorf("error evaluating arg %v of Eq", i)
		}
		eq, err := first.Eq(val)
		if err != nil {
			return BoolValue{false}, err
		}
		if !eq {
			return BoolValue{false}, nil
		}
	}
	return BoolValue{true}, nil
}
