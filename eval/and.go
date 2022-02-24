package eval

import "fmt"

type And struct {
	Args []Expression
}

func (and And) Eval(context Context) (Value, error) {
	if len(and.Args) == 0 {
		return BoolValue{true}, nil
	}
	for _, exp := range and.Args {
		val, err := exp.Eval(context)
		if err != nil {
			return nil, err
		}
		evaledVal, err := val.AsBool()
		if err != nil {
			return nil, fmt.Errorf("error evaluating %v, err: %v", val, err)
		}
		if !evaledVal {
			return BoolValue{false}, nil
		}
	}
	return BoolValue{true}, nil
}
