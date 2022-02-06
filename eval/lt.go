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
	first, err := firstVal.AsNumber()
	if err != nil {
		fmt.Errorf("Lt expected first argument to evaluate to a number")
		return BoolValue{false}, err
	}

	secondVal, err := lt.second.Eval(context)
	second, err := secondVal.AsNumber()
	if err != nil {
		fmt.Errorf("Lt expected second argument to evaluate to a number")
		return BoolValue{false}, err
	}

	return BoolValue{first < second}, nil
}
