package eval

import (
	"encoding/json"
	"errors"
	"fmt"
)

type NumberValue struct {
	Val float64
}

func (n NumberValue) AsNumber() (float64, error) {
	return n.Val, nil
}
func (n NumberValue) AsBool() (bool, error) {
	return false, fmt.Errorf("%v cannot be output as a bool", n)
}
func (n NumberValue) AsString() (string, error) {
	return "", fmt.Errorf("%v cannot be output as string", n)
}
func (n NumberValue) Eval(_ Context) (Value, error) {
	return n, nil
}
func (n NumberValue) Eq(other Value) (bool, error) {
	otherAsNumber, err := other.AsNumber()
	if err != nil {
		return false, errors.New("incompatible types")
	}
	return n.Val == otherAsNumber, nil
}
func (n NumberValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.Val)
}

type StringValue struct {
	Val string
}

func (n StringValue) AsNumber() (float64, error) {
	return 0, fmt.Errorf("%v cannot be output as number", n)
}
func (n StringValue) AsBool() (bool, error) {
	return false, fmt.Errorf("%v cannot be output as a bool", n)
}
func (n StringValue) AsString() (string, error) {
	return n.Val, nil
}
func (s StringValue) Eval(_ Context) (Value, error) {
	return s, nil
}
func (s StringValue) Eq(other Value) (bool, error) {
	otherAsString, err := other.AsString()
	if err != nil {
		return false, errors.New("incompatible types")
	}
	return s.Val == otherAsString, nil
}
func (s StringValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Val)
}

type BoolValue struct {
	Val bool
}

func (n BoolValue) AsNumber() (float64, error) {
	return 0, fmt.Errorf("%v cannot be output as a number", n)
}
func (n BoolValue) AsBool() (bool, error) {
	return n.Val, nil
}
func (n BoolValue) AsString() (string, error) {
	return "", fmt.Errorf("%v cannot be output as a string", n)
}
func (b BoolValue) Eval(_ Context) (Value, error) {
	return b, nil
}
func (b BoolValue) Eq(other Value) (bool, error) {
	otherAsBool, err := other.AsBool()
	if err != nil {
		return false, fmt.Errorf("%v and %v have incompatible types", b, other)
	}
	return b.Val == otherAsBool, nil
}
func (b BoolValue) String() string {
	if b.Val {
		return "true"
	}
	return "false"
}
func (b BoolValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.Val)
}
