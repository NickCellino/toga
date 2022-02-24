package parse

import (
	"fmt"
	"toga/eval"
)

type ExpressionParser func(interface{}) (eval.Expression, error)

func ConvertToAst(rawExpression interface{}) (eval.Expression, error) {
	var keywords = map[string]ExpressionParser{
		"eq": func(exp interface{}) (eval.Expression, error) {
			args, ok := exp.([]interface{})
			if !ok {
				return nil, fmt.Errorf("argument to eq operator should be an array, instead got: %v", exp)
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
		},
		"context": func(exp interface{}) (eval.Expression, error) {
			rawArg, ok := exp.(string)
			if !ok {
				return nil, fmt.Errorf("argument to context operator should be string, instead got: %v", exp)
			}
			arg, err := ConvertToAst(rawArg)
			if err != nil {
				return nil, err
			}
			return eval.ContextExpression{Key: arg}, nil
		},
		"and": func(exp interface{}) (eval.Expression, error) {
			args, ok := exp.([]interface{})
			if !ok {
				return nil, fmt.Errorf("argument to and operator should be an array, instead got: %v", exp)
			}
			parsedArgs := []eval.Expression{}
			for _, arg := range args {
				parsedArg, err := ConvertToAst(arg)
				if err != nil {
					return nil, err
				}
				parsedArgs = append(parsedArgs, parsedArg)
			}
			return eval.And{Args: parsedArgs}, nil
		},
		"or": func(exp interface{}) (eval.Expression, error) {
			args, ok := exp.([]interface{})
			if !ok {
				return nil, fmt.Errorf("argument to or operator should be an array, instead got: %v", exp)
			}
			parsedArgs := []eval.Expression{}
			for _, arg := range args {
				parsedArg, err := ConvertToAst(arg)
				if err != nil {
					return nil, err
				}
				parsedArgs = append(parsedArgs, parsedArg)
			}
			return eval.And{Args: parsedArgs}, nil
		},
	}

	switch expression := rawExpression.(type) {
	case map[string]interface{}:
		for keyword, parser := range keywords {
			value, in := expression[keyword]
			if in {
				exp, err := parser(value)
				if err != nil {
					return nil, err
				}
				return exp, nil
			}
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

func ParseContext(rawContext map[string]interface{}) (eval.Context, error) {
	parsedContext := eval.Context{}
	for key, val := range rawContext {
		expr, err := ConvertToAst(val)
		if err != nil {
			return nil, fmt.Errorf("error parsing context key %v. error: %v", key, err)
		}
		parsedContext[key] = expr
	}
	return parsedContext, nil
}
