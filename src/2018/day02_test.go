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

type testDay2FabricBoxes struct {
	input  string
	common string
}

var test_2part1 = []testDay2Boxes{
	{"abcdef bababc abbcde abcccd aabcdd abcdee ababab", 4, 3, 12},
}

var test_2part2 = []testDay2FabricBoxes{
	{"abcde fghij klmno pqrst fguij axcye wvxyz", "fgij"},
}

// Analyze an id. Return two integers with binary value:
// If the id contains only 2 of a given letter, return 1, otherwise return 0
// If the id contains only 3 of a given letter, return 1, otherwise return 0
func CountLetters(id string) (int, int) {
	count := map[rune]int{}
	for _, r := range id {
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

func CalculateChecksum(ids []string) (int, int, int) {
	twos := 0
	threes := 0
	for _, id := range ids {
		x, y := CountLetters(id)
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

// --- PART 2 --------

func CompareID(first string, second string) (bool, string) {
	diffs := 0
	last := 0

	for i := 0; i < len(first); i++ {
		if first[i] != second[i] {
			diffs++
			if diffs >= 2 {
				//fmt.Println("Bailing: ", first, second)
				return false, ""
			}
			last = i
		}
	}

	r := first[0:last] + first[(last+1):]

	//	fmt.Println("FOUND", first, second, "-->", r)
	return true, r
}

func CompareBoxIds(ids []string) string {
	for _, first := range ids {
		for _, second := range ids {
			if first == second {
				continue
			}

			found, common := CompareID(first, second)
			if found {
				return common
			}
		}
	}
	fmt.Println("Did not find any common box ids")
	return ""
}

func TestSampleData_2part2(t *testing.T) {
	for _, sample := range test_2part2 {

		common := CompareBoxIds(strings.Fields(sample.input))

		// expensive! but this is only a test...
		if strings.Compare(sample.common, common) != 0 {
			t.Error("For [", sample.input, "] expected [", sample.common, "] got [", common, "]")
		}
	}
}

func TestInput_2part2(t *testing.T) {
	content, err := ioutil.ReadFile("day02_input.txt")
	check(err)

	common := CompareBoxIds(strings.Fields(string(content)))

	fmt.Println("Day 2 / Part 2 Result", common)
}
