package days

import (
	"fmt"
	"testing"
)

var test_3part1 = []testIntPair{
	{1, 0},
	{12, 3},
	{23, 2},
	{27, 4},
	{1024, 31},
}

var test_3part2 = []testIntPair{
	{1, 1},
	{2, 1},
	{3, 2},
	{7, 11},
	{21, 362},
	{27, 1968},
}

func TestSampleData_3part1(t *testing.T) {
	for _, pair := range test_3part1 {
		v := Spiral(pair.input)
		verifyTestIntPair(pair, v, t)
	}
}

func TestInput_3part1(t *testing.T) {
	v := Spiral(265149)
	fmt.Println("Day 3 / Part 1 Result", v)
}

func TestSampleData_3part2(t *testing.T) {
	for _, pair := range test_3part2 {
		v := SpiralSum(pair.input, 0)
		verifyTestIntPair(pair, v, t)
	}
}

func TestInput_3part2(t *testing.T) {
	v := SpiralSum(0, 265149)
	fmt.Println("Day 3 / Part 2 Result", v)
}
