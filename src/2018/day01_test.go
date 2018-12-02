package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
	"unicode"
)

var test_1part1 = []testStringIntPair{
	{"+1 -2 +3 +1", 3},
	{"+1 +1 +1", 3},
	{"+1 +1 -2", 0},
	{"-1 -2 -3", -6},
}

//var test_1part2 = []testStringIntPair{}

func tune(freq int, delta string) int {
	var (
		m = 1
		r = 0
	)
	delta = strings.TrimLeftFunc(delta, func(r rune) bool {
		if r == '-' {
			m = -1
		}
		return !unicode.IsNumber(r)
	})
	i, _ := strconv.Atoi(delta)
	r = freq + (m * i)
	//fmt.Println(freq, " + (", m, "*", delta, ") = ", r)
	return r
}

func TestSampleData_1part1(t *testing.T) {
	for _, pair := range test_1part1 {

		var freq int = 0
		for _, delta := range strings.Split(pair.input, " ") {
			freq = tune(freq, delta)
		}

		if freq != pair.expected {
			t.Error("For", pair.input, "expected", pair.expected, "got", freq)
		}
	}
}

func TestInput_1part1(t *testing.T) {
	f, err := os.Open("day01_input.txt")
	check(err)
	defer f.Close()

	reader := bufio.NewReader(f)
	_, err = reader.Peek(2)
	check(err)

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	var freq int = 0
	for scanner.Scan() {
		freq = tune(freq, scanner.Text())
	}

	fmt.Println("Day 1 / Part 1 Result", freq)
}

func TestSampleData_1part2(t *testing.T) {
	//	for _, pair := range test_1part2 {
	//		b := []byte(pair.input)
	//		v := HalfCaptcha(b)
	//		verifyTestPair(pair, v, t)
	//	}
}

func TestInput_1part2(t *testing.T) {
	//	content, err := ioutil.ReadFile("day01_input.txt")
	//	check(err)

	//	content = bytes.TrimRightFunc(content, func(r rune) bool {
	//		return (r < '0' || '9' < r)
	//	})

	//	v := HalfCaptcha(content)
	//	fmt.Println("Day 1 / Part 2 Result", v)
}
