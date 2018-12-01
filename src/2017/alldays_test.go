package days

import (
	"testing"
)

type testpair struct {
	input    string
	expected int
}

type testIntPair struct {
	input    int
	expected int
}

type testBoolPair struct {
	input    string
	expected bool
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func verifyTestPair(pair testpair, v int, t *testing.T) {
	if v != pair.expected {
		t.Error(
			"For", pair.input,
			"expected", pair.expected,
			"got", v,
		)
	}
}

func verifyTestIntPair(pair testIntPair, v int, t *testing.T) {
	if v != pair.expected {
		t.Error(
			"For", pair.input,
			"expected", pair.expected,
			"got", v,
		)
	}
}

func verifyTestBoolPair(pair testBoolPair, v bool, t *testing.T) {
	if v != pair.expected {
		t.Error(
			"For", pair.input,
			"expected", pair.expected,
			"got", v,
		)
	}
}
