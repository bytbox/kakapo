package lisp

import (
	"strings"
	"testing"
)

type readTokenTest struct{
	str string
	tok token
}

var readTokenTests = []readTokenTest{
	{"1\n", "1"},
	{"32.5a\n", "32.5a"},
	{"32.5 a\n", "32.5"},
	{"\"b\"\n", "\"b\""},
	{" \t5;6", "5"},
	{"(", "("},
	{")(", ")"},
}

func TestReadToken(t *testing.T) {
	for _, test := range readTokenTests {
		r := strings.NewReader(test.str)
		tok, err := readToken(r)
		if err != nil {
			t.Errorf("ERROR: %s", err.Error())
		}
		if tok != test.tok {
			t.Errorf("%s != %s", tok, test.tok)
		}
	}
}

type parseTest struct{
	str string
	res sexpr
}

var parseTests = []parseTest{
	{"1\n", 1.0},
	{"5.5\n", 5.5},
	{"5e-9\n", 5e-9},
	{"x\n", sym("x")},
	{"5%x\n", sym("5%x")},
	{"\"a\"", "a"},

	{"()", Nil},
	{"(())", cons{nil, nil}},
	{"(1)", cons{1.0, nil}},
	{"(1 (2 3) ())",
		cons{1.0, cons{cons{2.0, cons{3.0, nil}}, cons{nil, nil}}}},
}

func eqS(a sexpr, b sexpr) bool {
	ac, ok := a.(cons)
	if ok {
		bc, ok := b.(cons)
		if !ok {
			return false
		}
		return eqS(ac.car, bc.car) && eqS(ac.cdr, bc.cdr)
	}
	return a == b
}

func TestParse(t *testing.T) {
	for _, test := range parseTests {
		r := strings.NewReader(test.str)
		res, err := parse(r)
		if err != nil {
			t.Errorf("ERROR: %s", err.Error())
		}
		if !eqS(res, test.res) {
			t.Errorf("%s != %s", res, test.res)
		}
	}
}
