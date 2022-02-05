package eval

type IntVal struct {
	Val int
}

func (intVal IntVal) Resolve() interface{} {
	return intVal.Val
}
