package days

import (
 "bufio"
 "fmt"
 "os"
 "strings"
 "testing"
)

type testpair struct {
  input string
  expected int
}

var tests = []testpair {
  {"1122", 3},
  {"1111", 4},
  {"1234", 0},
  {"91212129", 9},
}

func TestSampleData(t *testing.T) {
  for _, pair := range tests {  
    reader := strings.NewReader(pair.input)
    v := Captcha(reader)    
    if v != pair.expected {
      t.Error(
        "For", pair.input,
        "expected", pair.expected,
        "got", v,
      )
    }
  }
}

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func TestInput(t *testing.T) {
  f, err := os.Open("day01_input.txt")
  check(err)
  defer f.Close()

  reader := bufio.NewReader(f)
  b4, err := reader.Peek(5)
  check(err)
 
  v := Captcha(reader)
  fmt.Println("Day 1 Result", v);
}
