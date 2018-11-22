package days
import (
	"testing"
)

type testpair struct {
  input string
  expected int
}

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func verify(pair testpair, v int, t *testing.T) {
	if v != pair.expected {
		t.Error(
			"For", pair.input,
			"expected", pair.expected,
			"got", v,
		)
	}
}
