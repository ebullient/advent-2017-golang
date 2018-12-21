package days

import (
	"fmt"
	"strconv"
	"time"
)

type Point struct {
	x, y int
}

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

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func ToInt(input string) int {
	i, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return i
}

func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("\t\t-> %v\n", time.Since(start))
	}
}
