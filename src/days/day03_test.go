package days

import (
 "fmt"
 "testing"
)

var test_3part1 = []testIntPair {
	{1, 0},
	{12, 3},
	{23, 2},
	{27, 4},
	{1024, 31},
}

func TestSampleData_3part1(t *testing.T) {
	for _, pair := range test_3part1 {
		v := Spiral(pair.input)
		verifyTestIntPair(pair, v, t)
	}
}

func TestInput_3part1(t *testing.T) {
  v:= Spiral(265149)
  fmt.Println("Day 3 / Part 1 Result", v);
}


