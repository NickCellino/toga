package eval

type ContextVal struct {
	name string
}

func (contextVal ContextVal) Resolve() interface{} {
	return contextVal.name
}
