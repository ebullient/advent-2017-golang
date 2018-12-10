package days

import (
	"fmt"
	"time"
)

type testStringIntPair struct {
	input    string
	expected int
}

type testIntPair struct {
	input    int
	expected int
}

type testStringBoolPair struct {
	input    string
	expected bool
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}
