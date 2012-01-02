package lisp

import (
	"fmt"
)

type cons struct {
	car sexpr
	cdr sexpr
}

type sexpr interface{}
type atom interface{}
type sym string
type function func(*scope, []sexpr) sexpr

type native interface{}

func (v cons) String() string {
	return fmt.Sprintf("(%s . %s)", asString(v.car), asString(v.cdr))
}

func asString(v sexpr) string {
	switch v := v.(type) {
	case cons:
		return v.String()
	case sym:
		return fmt.Sprintf("<sym : %s>", string(v))
	case float64:
		return fmt.Sprintf("%G", v)
	case string:
		return fmt.Sprintf("\"%s\"", v)
	case function:
		return "<func>"
	case primitive:
		return "<primitive>"
	case native:
		return "<native>"
	case nil:
		return "nil"
	}
	return "<unknown>"
}

func isFunction(s sexpr) bool {
	_, ok := s.(function)
	return ok
}

func isPrimitive(s sexpr) bool {
	_, ok := s.(primitive)
	return ok
}
