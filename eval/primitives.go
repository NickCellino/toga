package eval

import "errors"

type NumberValue struct {
	val float64
}

func (n NumberValue) AsNumber() (float64, error) {
	return n.val, nil
}
func (n NumberValue) AsBool() (bool, error) {
	return false, errors.New("number value cannot be output as a bool")
}
func (n NumberValue) AsString() (string, error) {
	return "", errors.New("number value cannot be output as a string")
}
func (n NumberValue) Eval(_ Context) (Value, error) {
	return n, nil
}

type StringValue struct {
	val string
}

func (n StringValue) AsNumber() (float64, error) {
	return 0, errors.New("string value cannot be output as a number")
}
func (n StringValue) AsBool() (bool, error) {
	return false, errors.New("string value cannot be output as a bool")
}
func (n StringValue) AsString() (string, error) {
	return n.val, nil
}
func (s StringValue) Eval(_ Context) (Value, error) {
	return s, nil
}

type BoolValue struct {
	val bool
}

func (n BoolValue) AsNumber() (float64, error) {
	return 0, errors.New("bool value cannot be output as a number")
}
func (n BoolValue) AsBool() (bool, error) {
	return n.val, nil
}
func (n BoolValue) AsString() (string, error) {
	return "", errors.New("bool value cannot be output as a string")
}
func (b BoolValue) Eval(_ Context) (Value, error) {
	return b, nil
}
