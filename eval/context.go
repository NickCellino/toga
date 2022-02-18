package eval

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
	return context[keyStr].Eval(context)
}
