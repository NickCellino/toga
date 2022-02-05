package eval

type IntVal struct {
	Val int
}

func (intVal IntVal) Resolve() interface{} {
	return intVal.Val
}

type BoolVal struct {
	Val bool
}

func (boolVal BoolVal) Resolve() interface{} {
	return boolVal.Val
}

type FloatVal struct {
	Val float64
}

func (floatVal FloatVal) Resolve() interface{} {
	return floatVal.Val
}
