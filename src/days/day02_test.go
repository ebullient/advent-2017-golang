package days

import (
 "bufio"
 "fmt"
 "os"
 "strings"
 "testing"
)

var test_2part1 = []testpair {
  {"5\t1\t9\t5\n7\t5\t3\n2\t4\t6\t8", 18},
}

var test_2part2 = []testpair {
	{"5\t9\t2\t8\n9\t4\t7\t3\n3\t8\t6\t5\n", 9},
}

func TestSampleData_2part1(t *testing.T) {
	for _, pair := range test_2part1 {
		reader := strings.NewReader(pair.input)
		v := Checksum(reader)
		verifyTestPair(pair, v, t)
	}
}

func TestInput_2part1(t *testing.T) {
  f, err := os.Open("day02_input.txt")
  check(err)
  defer f.Close()

  reader := bufio.NewReader(f)
  _, err = reader.Peek(2)
	check(err)

	v := Checksum(reader)
  fmt.Println("Day 2 / Part 1 Result", v);
}

func TestSampleData_2part2(t *testing.T) {
	for _, pair := range test_2part2 {
		reader := strings.NewReader(pair.input)
		v := DivisibleChecksum(reader)
		verifyTestPair(pair, v, t)
	}
}

func TestInput_2part2(t *testing.T) {
  f, err := os.Open("day02_input.txt")
  check(err)
  defer f.Close()

  reader := bufio.NewReader(f)
  _, err = reader.Peek(2)
	check(err)

	v := DivisibleChecksum(reader)
  fmt.Println("Day 2 / Part 2 Result", v);
}
