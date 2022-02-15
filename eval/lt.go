package eval

import (
	"fmt"
)

type Lt struct {
	first  Expression
	second Expression
}

func (lt Lt) Eval(context Context) (Value, error) {
	firstVal, err := lt.first.Eval(context)
	if err != nil {
		return nil, err
	}
	first, err := firstVal.AsNumber()
	if err != nil {
		return nil, fmt.Errorf("error evaluating %v, first arg %v is not a number", lt, firstVal)
	}

	secondVal, err := lt.second.Eval(context)
	if err != nil {
		return nil, err
	}
	second, err := secondVal.AsNumber()
	if err != nil {
		return nil, fmt.Errorf("error evaluating %v, second arg %v is not a number", lt, secondVal)
	}

	return BoolValue{first < second}, nil
}
