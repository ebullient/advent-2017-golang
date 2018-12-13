package days

import (
	"fmt"
	"io/ioutil"
	//"reflect"
	//"regexp"
	//"sort"
	//"strconv"
	"strings"
	"testing"
)

var test_polymers = []testStringStringPair{
	{"aA", ""},
	{"abBA", ""},
	{"abAB", "abAB"},
	{"aabAAB", "aabAAB"},
	{"dabAcCaCBAcCcaDA", "dabCBAcaDA"},
}

func Reactive(x byte, y byte) bool {
	if x > y {
		return x-y == 32
	}
	return y-x == 32
}

// Reduce letters next to each other that are
// the same type but opposite polarity: a and A
// uppercase/lowercase bytes are separated by 32
func Reduce(polymer string) string {
	splices := []string{}
	last := 0

	for i, j := 0, 1; j < len(polymer); i, j = i+1, j+1 {
		if Reactive(polymer[i], polymer[j]) {
			splices = append(splices, polymer[last:i])
			i++
			j++
			last = j
		}
	}
	if last > 0 {
		if last < len(polymer) {
			splices = append(splices, polymer[last:])
		}
		return strings.Join(splices, "")
	}
	return polymer
}

func RepeatReduce(polymer string) string {
	prev := len(polymer)
	for {
		//fmt.Println("-> ", prev, polymer)
		polymer = Reduce(polymer)
		l := len(polymer)
		//fmt.Println("<- ", l, polymer)
		if l == prev || l == 0 {
			return polymer
		}
		prev = l
	}
}

func TestSampleData_5part1(t *testing.T) {
	for _, pair := range test_polymers {
		result := RepeatReduce(pair.input)
		if result != pair.expected {
			t.Error("For", pair.input, "expected", pair.expected, "got", result)
		}
	}
}

func TestInput_5part1(t *testing.T) {
	content, err := ioutil.ReadFile("day05_input.txt")
	check(err)

	defer elapsed("Day 5 / Part 1")() // time execution of the rest

	s := strings.TrimSpace(string(content))

	//fmt.Println(len(s))
	result := RepeatReduce(s)
	fmt.Println("Day 5 / Part 1 Result", len(result))
}

type PolymerImprovement struct {
	unit    rune
	polymer string
}

func RemoveUnit(attempt PolymerImprovement, done chan<- PolymerImprovement) {
	result := strings.Replace(attempt.polymer, string(attempt.unit), "", -1)
	result = strings.Replace(result, string(attempt.unit+32), "", -1)
	result = RepeatReduce(result)
	attempt.polymer = result
	done <- attempt
}

func TestSampleData_5part2(t *testing.T) {
	polymer := "dabAcCaCBAcCcaDA"
	done := make(chan PolymerImprovement)
	go RemoveUnit(PolymerImprovement{'A', polymer}, done)
	go RemoveUnit(PolymerImprovement{'B', polymer}, done)
	go RemoveUnit(PolymerImprovement{'C', polymer}, done)
	go RemoveUnit(PolymerImprovement{'D', polymer}, done)

	shortest := len(polymer)
	for i := 0; i < 4; i++ {
		result := <-done
		x := len(result.polymer)
		if x < shortest {
			shortest = x
		}
	}
	if shortest != 4 {
		t.Error("Expected", 4, "got", shortest)
	}
}

func TestInput_5part2(t *testing.T) {
	content, err := ioutil.ReadFile("day05_input.txt")
	s := strings.TrimSpace(string(content))
	check(err)

	defer elapsed("Day 5 / Part 2")() // time execution of the rest

	done := make(chan PolymerImprovement)
	for i := 'A'; i <= 'Z'; i++ {
		go RemoveUnit(PolymerImprovement{i, s}, done)
	}

	unit := ""
	shortest := len(s)
	for i := 'A'; i <= 'Z'; i++ {
		result := <-done
		x := len(result.polymer)
		if x < shortest {
			shortest = x
			unit = string(result.unit)
		}
	}

	fmt.Println("Day 5 / Part 2 Result", shortest, unit)
}
