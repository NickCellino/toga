package eval

type If struct {
	Condition Expression
	Then      Expression
	Else      Expression
}

func (_if If) Eval(context Context) (Value, error) {
	conditionVal, err := _if.Condition.Eval(context)
	if err != nil {
		return nil, err
	}

	conditionValBool, err := conditionVal.AsBool()
	if err != nil {
		return nil, err
	}

	if conditionValBool {
		return _if.Then.Eval(context)
	}

	return _if.Else.Eval(context)
}
