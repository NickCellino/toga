package eval

import "fmt"

type Eq struct {
	Args []Expression
}

func (eq Eq) Eval(context Context) (Value, error) {
	if len(eq.Args) == 0 {
		return BoolValue{true}, nil
	}
	first, err := eq.Args[0].Eval(context)
	if err != nil {
		return nil, err
	}
	for _, exp := range eq.Args[1:] {
		val, err := exp.Eval(context)
		if err != nil {
			return nil, err
		}
		eq, err := first.Eq(val)
		if err != nil {
			return nil, fmt.Errorf("error comparing %v with %v. message: %v", first, val, err)
		}
		if !eq {
			return BoolValue{false}, nil
		}
	}
	return BoolValue{true}, nil
}
