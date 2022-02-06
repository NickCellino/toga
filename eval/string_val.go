package eval

type StringVal struct {
	Val string
}

func (stringVal StringVal) Resolve() interface{} {
	return stringVal.Val
}
