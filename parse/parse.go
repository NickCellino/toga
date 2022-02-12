package parse

import (
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
	switch expression := rawExpression.(type) {
	case map[string]interface{}:
		value, in := expression["eq"]
		if in {
			args, ok := value.([]interface{})
			if !ok {
				return nil, fmt.Errorf("argument to eq operator should be an array, instead got: %v", value)
			}
			parsedArgs := []eval.Expression{}
			for _, arg := range args {
				parsedArg, err := ConvertToAst(arg)
				if err != nil {
					return nil, err
				}
				parsedArgs = append(parsedArgs, parsedArg)
			}
			return eval.Eq{Args: parsedArgs}, nil
		}
	case string:
		return eval.StringValue{Val: expression}, nil
	case float64:
		return eval.NumberValue{Val: expression}, nil
	case bool:
		return eval.BoolValue{Val: expression}, nil
	}
	return nil, fmt.Errorf("we don't know how to parse '%v'", rawExpression)
}
