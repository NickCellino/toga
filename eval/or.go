package eval

import "fmt"

type Or struct {
	Args []Expression
}

func (or Or) Eval(context Context) (Value, error) {
	if len(or.Args) == 0 {
		return BoolValue{true}, nil
	}
	for _, exp := range or.Args {
		val, err := exp.Eval(context)
		if err != nil {
			return nil, err
		}
		evaledVal, err := val.AsBool()
		if err != nil {
			return nil, fmt.Errorf("error evaluating %v, err: %v", val, err)
		}
		if evaledVal {
			return BoolValue{true}, nil
		}
	}
	return BoolValue{false}, nil
}
