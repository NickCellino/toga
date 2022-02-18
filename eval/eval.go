package eval

type Expression interface {
	Eval(context Context) (Value, error)
}

type Value interface {
	Expression
	AsNumber() (float64, error)
	AsBool() (bool, error)
	AsString() (string, error)
	Eq(Value) (bool, error)
}

type Context map[string]Expression
