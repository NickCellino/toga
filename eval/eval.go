package eval

type Expression interface {
	Eval(context Context) (Value, error)
}

type Value interface {
	Expression
	AsNumber() (float64, error)
	AsBool() (bool, error)
	AsString() (string, error)
}

type Context map[string]interface{}
