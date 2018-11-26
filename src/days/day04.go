package days

import (
	//	"fmt"
	"sort"
	"strings"
)

func SortString(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool { return r[i] < r[j] })
	//	fmt.Println(s, string(r))
	return string(r)
}

func Passphrase(input string, sorted bool) bool {
	if len(input) == 0 {
		return false
	}

	unique := make(map[string]bool)
	words := strings.Fields(input)

	for _, word := range words {
		if sorted {
			word = SortString(word)
		}
		_, present := unique[word]
		if present {
			//fmt.Println("FAIL: ", input)
			return false
		}
		unique[word] = true
	}

	return true
}
