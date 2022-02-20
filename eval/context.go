package eval

import "fmt"

type ContextExpression struct {
	Key Expression
}

func (contextExpr ContextExpression) Eval(context Context) (Value, error) {
	key, err := contextExpr.Key.Eval(context)
	if err != nil {
		return nil, err
	}
	keyStr, err := key.AsString()
	if err != nil {
		return nil, err
	}
	contextVal, in := context[keyStr]
	if !in {
		return nil, fmt.Errorf("value '%v' not specified in context", keyStr)
	}
	return contextVal.Eval(context)
}
