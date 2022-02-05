package eval

type Comparable interface {
	Compare()
}

type Evalable interface {
	Eval() (bool, error)
}

type Resolvable interface {
	Resolve() interface{}
}

func Eval(rules []Evalable) bool {
	for _, rule := range rules {
		if result, _ := rule.Eval(); !result {
			return false
		}
	}
	return true
}
