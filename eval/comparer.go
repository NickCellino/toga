package eval

import "fmt"

type ComparisonOperator struct {
	First    Expression
	Second   Expression
	Comparer func(float64, float64) bool
}

func (c ComparisonOperator) Eval(context Context) (Value, error) {
	firstVal, err := c.First.Eval(context)
	if err != nil {
		return nil, err
	}
	first, err := firstVal.AsNumber()
	if err != nil {
		return nil, fmt.Errorf("error evaluating %v: %w", c, err)
	}

	secondVal, err := c.Second.Eval(context)
	if err != nil {
		return nil, err
	}
	second, err := secondVal.AsNumber()
	if err != nil {
		return nil, fmt.Errorf("error evaluating %v, %w", c, err)
	}

	return BoolValue{c.Comparer(first, second)}, nil
}
