package parse

import (
	"errors"
	"fmt"
	"toga/eval"
)

type ExpressionFactory func() (eval.Expression, error)

func Operations() map[string]ExpressionFactory {
	return map[string]ExpressionFactory{
		"eq": func() (eval.Expression, error) {
			return eval.Eq{}, nil
		},
	}
}

func ConvertToAst(rawExpression interface{}) (eval.Expression, error) {
	fmt.Println(rawExpression)
	expressionMap, _ := rawExpression.(map[string]interface{})
	_, in := expressionMap["op"]
	if !in {
		return nil, errors.New("expression should contain top level 'op' key")
	}
	fmt.Println("expr map: ", expressionMap)
	return nil, nil
}
