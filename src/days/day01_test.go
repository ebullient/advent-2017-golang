package days

import (
 "bufio"
 "bytes"
 "fmt"
 "io/ioutil"
 "os"
 "strings"
 "testing"
)

var test_1part1 = []testpair {
  {"1122", 3},
  {"1111", 4},
  {"1234", 0},
  {"91212129", 9},
}

var test_1part2 = []testpair {
  {"1212", 6},
  {"1221", 0},
  {"123425", 4},
  {"123123", 12},
  {"12131415", 4},
}

func TestSampleData_1part1(t *testing.T) {
  for _, pair := range test_1part1 {
    reader := strings.NewReader(pair.input)
    v := Captcha(reader)
		verify(pair, v, t)
  }
}

func TestInput_1part1(t *testing.T) {
  f, err := os.Open("day01_input.txt")
  check(err)
  defer f.Close()

  reader := bufio.NewReader(f)
  _, err = reader.Peek(2)
  check(err)

  v := Captcha(reader)
  fmt.Println("Day 1 / Part 1 Result", v);
}

func TestSampleData_1part2(t *testing.T) {
  for _, pair := range test_1part2 {
    b := []byte(pair.input)
    v := HalfCaptcha(b)
		verify(pair, v, t)
  }
}


func TestInput_1part2(t *testing.T) {
  content, err := ioutil.ReadFile("day01_input.txt")
	check(err)

  content = bytes.TrimRightFunc(content, func(r rune) bool {
    return ( r < '0' || '9' < r )
  })

  v := HalfCaptcha(content)
  fmt.Println("Day 1 / Part 2 Result", v)
}
