package days

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

type testDay2Boxes struct {
	input    string
	twos     int
	threes   int
	checksum int
}

var test_2part1 = []testDay2Boxes{
	{"abcdef bababc abbcde abcccd aabcdd abcdee ababab", 4, 3, 12},
}

var test_2part2 = []testDay2Boxes{}

// Analyze a word. Return two integers with binary value:
// If the word contains only 2 of a given letter, return 1, otherwise return 0
// If the word contains only 3 of a given letter, return 1, otherwise return 0
func CountLetters(word string) (int, int) {
	count := map[rune]int{}
	for _, r := range word {
		count[r] = count[r] + 1
	}

	twos := 0
	threes := 0
	for _, value := range count {
		if value == 2 {
			twos = 1
		} else if value == 3 {
			threes = 1
		} else if twos == 1 && threes == 1 {
			break
		}
	}

	return twos, threes
}

func CalculateChecksum(words []string) (int, int, int) {
	twos := 0
	threes := 0
	for _, word := range words {
		x, y := CountLetters(word)
		twos = twos + x
		threes = threes + y
	}
	checksum := twos * threes

	return twos, threes, checksum
}

func TestSampleData_2part1(t *testing.T) {
	for _, sample := range test_2part1 {

		twos, threes, checksum := CalculateChecksum(strings.Fields(sample.input))

		if twos != sample.twos {
			t.Error("For", sample.input, "expected", sample.twos, "words with two of the same letter, got", twos)
		}

		if threes != sample.threes {
			t.Error("For", sample.input, "expected", sample.threes, "words with three of the same letter, got", threes)
		}

		if checksum != sample.checksum {
			t.Error("For", sample.input, "expected checksum", sample.checksum, "got", checksum)
		}
	}
}

func TestInput_2part1(t *testing.T) {
	content, err := ioutil.ReadFile("day02_input.txt")
	check(err)

	twos, threes, checksum := CalculateChecksum(strings.Fields(string(content)))

	fmt.Println("Day 2 / Part 1 Result", twos, threes, checksum)
}

func TestSampleData_2part2(t *testing.T) {
	//	for _, sample.:= range test_2part2 {
	//	delta := strings.Fields(sample.input)

	//	freq := iterate(delta)
	//	if freq != sample.expected {
	//		t.Error("For", sample.input, "expected", sample.expected, "got", freq)
	//	}
	//	}
}

func TestInput_2part2(t *testing.T) {
	//	content, err := ioutil.ReadFile("day01_input.txt")
	//	check(err)

	//	delta := strings.Fields(string(content))
	//	freq := iterate(delta)

	//	fmt.Println("Day 1 / Part 2 Result", freq)
}
