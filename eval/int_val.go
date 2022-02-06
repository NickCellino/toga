package eval

type IntVal struct {
	Val int
}

func (intVal IntVal) Eval(_ Context) interface{} {
	return intVal.Val
}
