package eval

import (
	"fmt"
)

type Lt struct {
	first  Resolvable
	second Resolvable
}

func (lt Lt) Eval() (bool, error) {
	firstVal := lt.first.Resolve()
	secondVal := lt.second.Resolve()
	var firstValFloat, secondValFloat float64

	switch firstValTyped := firstVal.(type) {
	case int:
		firstValFloat = float64(firstValTyped)
	case float32:
		firstValFloat = float64(firstValTyped)
	case float64:
	default:
		return false, fmt.Errorf("first arg is not an int, float32, or float64")
	}

	switch secondValTyped := secondVal.(type) {
	case int:
		secondValFloat = float64(secondValTyped)
	case float64:
		secondValFloat = float64(secondValTyped)
	default:
		return false, fmt.Errorf("second arg is not an int, float32, or float64")
	}

	return firstValFloat < secondValFloat, nil
}
