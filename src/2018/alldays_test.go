package days

import (
	"fmt"
	"time"
)

type testStringStringPair struct {
	input    string
	expected string
}

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
		fmt.Printf("\t\t-> %v\n", time.Since(start))
	}
}
