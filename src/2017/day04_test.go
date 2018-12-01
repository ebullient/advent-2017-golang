package days

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

var test_4part1 = []testBoolPair{
	{"aa bb cc dd ee", true},
	{"aa bb cc dd aa", false},
	{"aa bb cc dd aaa", true},
}

var test_4part2 = []testBoolPair{
	{"abcde fghij", true},
	{"abcde xyz ecdab", false},
	{"a ab abc abd abf abj", true},
	{"iiii oiii ooii oooi oooo", true},
	{"oiii ioii iioi iiio", false},
}

func TestSampleData_4part1(t *testing.T) {
	for _, pair := range test_4part1 {
		v := Passphrase(pair.input, false)
		verifyTestBoolPair(pair, v, t)
	}
}

func TestInput_4part1(t *testing.T) {
	f, err := os.Open("day04_input.txt")
	check(err)
	defer f.Close()

	reader := bufio.NewReader(f)
	_, err = reader.Peek(2)
	check(err)

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	var count int
	for scanner.Scan() {
		ok := Passphrase(scanner.Text(), false)
		if ok {
			count++
		}
	}
	fmt.Println("Day 4 / Part 1 Result", count)
}

func TestSampleData_4part2(t *testing.T) {
	for _, pair := range test_4part2 {
		v := Passphrase(pair.input, true)
		verifyTestBoolPair(pair, v, t)
	}
}

func TestInput_4part2(t *testing.T) {
	f, err := os.Open("day04_input.txt")
	check(err)
	defer f.Close()

	reader := bufio.NewReader(f)
	_, err = reader.Peek(2)
	check(err)

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	var count int
	for scanner.Scan() {
		ok := Passphrase(scanner.Text(), true)
		if ok {
			count++
		}
	}
	fmt.Println("Day 4 / Part 2 Result", count)
}
