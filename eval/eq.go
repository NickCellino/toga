package eval

import (
	"errors"
	"fmt"
	"reflect"
)

type Eq struct {
	args []Resolvable
}

func (eq Eq) Eval() (bool, error) {
	if len(eq.args) == 0 {
		return true, nil
	}
	firstVal := eq.args[0].Resolve()
	firstValType := reflect.TypeOf(firstVal)
	for i, val := range eq.args[1:] {
		valResolved := val.Resolve()
		valType := reflect.TypeOf(valResolved)
		if valType != firstValType {
			return false, errors.New(fmt.Sprintf("The type of the value at index %v does not match the type of the value at index 0.", i))
		}
		if valResolved != firstVal {
			return false, nil
		}
	}
	return true, nil
}
