package days

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"testing"
)

var test_1part1 = []testStringIntPair{
	{"+1 -2 +3 +1", 3},
	{"+1 +1 +1", 3},
	{"+1 +1 -2", 0},
	{"-1 -2 -3", -6},
}

var test_1part2 = []testStringIntPair{
	{"+1 -2 +3 +1 +1 -2", 2},
	{"+3 +3 +4 -2 -4", 10},
	{"-6 +3 +8 +5 -6", 5},
	{"+7 +7 -2 -7 -4", 14},
}

func tune(freq int, delta string) int {
	var r int = 0
	i, _ := strconv.Atoi(delta)
	r = freq + i
	//fmt.Println(freq, " + (", m, "*", delta, ") = ", r)
	return r
}

func TestSampleData_1part1(t *testing.T) {
	for _, pair := range test_1part1 {
		var freq int = 0
		for _, delta := range strings.Fields(pair.input) {
			freq = tune(freq, delta)
		}

		if freq != pair.expected {
			t.Error("For", pair.input, "expected", pair.expected, "got", freq)
		}
	}
}

func TestInput_1part1(t *testing.T) {
	content, err := ioutil.ReadFile("day01_input.txt")
	check(err)

	var freq int = 0
	for _, delta := range strings.Fields(string(content)) {
		freq = tune(freq, delta)
	}

	fmt.Println("Day 1 / Part 1 Result", freq)
}

func iterate(delta []string) int {
	var freq int = 0
	seen := make(map[int]bool)

	for i := 0; true; i = i + 1 {
		if i >= len(delta) {
			i = 0
		}
		freq = tune(freq, delta[i])

		_, present := seen[freq]
		//fmt.Println(i, "] ", delta[i], freq, present)
		if present {
			return freq
		}
		seen[freq] = true
	}

	return 0
}

func TestSampleData_1part2(t *testing.T) {
	for _, pair := range test_1part2 {
		delta := strings.Fields(pair.input)

		freq := iterate(delta)
		if freq != pair.expected {
			t.Error("For", pair.input, "expected", pair.expected, "got", freq)
		}
	}
}

func TestInput_1part2(t *testing.T) {
	content, err := ioutil.ReadFile("day01_input.txt")
	check(err)

	delta := strings.Fields(string(content))
	freq := iterate(delta)

	fmt.Println("Day 1 / Part 2 Result", freq)
}
