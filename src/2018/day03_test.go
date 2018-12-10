package days

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

type testFabricClaim struct {
	claims    string
	conflicts int
	distinct  string
}

type Square struct {
	left int
	top  int
}

var test_3part1 = []testFabricClaim{
	{"#1 @ 1,3: 1x3\n#2 @ 1,3: 3x1\n#3 @ 1,3: 2x2", 3, ""},
	{"#1 @ 1,3: 4x4\n#2 @ 3,1: 4x4\n#3 @ 5,5: 2x2", 4, "3"},
}

//var test_3part2 = []testDay2FabricBoxes{
//}

func ToInt(input string) int {
	i, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return i
}

func DefineClaim(fabric map[Square][]string, id string, start Square, width int, height int) {
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			x := start.left + j
			y := start.top + i
			fabric[Square{x, y}] = append(fabric[Square{x, y}], id)
		}
	}
}

func PlaceClaims(scanner *bufio.Scanner) (int, string) {
	// the compiler panics if this fails
	r := regexp.MustCompile(`#(\d+) @ (\d+),(\d+): (\d+)x(\d+)`)

	distinct := map[string]bool{}
	fabric := map[Square][]string{}
	for scanner.Scan() {
		results := r.FindStringSubmatch(scanner.Text())
		distinct[results[1]] = true
		DefineClaim(fabric, results[1],
			Square{ToInt(results[2]), ToInt(results[3])},
			ToInt(results[4]),
			ToInt(results[5]))
	}

	// fmt.Println(fabric)
	conflicts := 0
	for _, value := range fabric {
		length := len(value)
		if length > 1 {
			conflicts++
			for _, id := range value {
				delete(distinct, id)
			}
		}
	}

	id := ""
	for key := range distinct {
		id = key
	}

	return conflicts, id
}

func TestSampleData_3part1(t *testing.T) {
	for _, sample := range test_3part1 {
		scanner := bufio.NewScanner(strings.NewReader(sample.claims))
		conflicts, distinct := PlaceClaims(scanner)
		if conflicts != sample.conflicts {
			t.Error("For", sample.claims, "expected", sample.conflicts, "got", conflicts)
		}
		if distinct != sample.distinct {
			t.Error("For", sample.claims, "expected distinct id", sample.distinct, "got", distinct)
		}
	}
}

func TestInput_3part1(t *testing.T) {
	file, err := os.Open("day03_input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	conflicts, distinct := PlaceClaims(scanner)

	fmt.Println("Day 3 / Part 1 Result", conflicts)
	fmt.Println("Day 3 / Part 2 Result", distinct)
}
