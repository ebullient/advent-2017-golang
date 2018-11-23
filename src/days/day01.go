package days

import (
	//  "fmt"
	"io"
)

func toInt(b byte) (int, bool) {
	if b < '0' || '9' < b {
		return 0, true
	}
	return int(b - '0'), false
}

// The captcha requires you to review a sequence of digits (your puzzle input)
// and find the sum of all digits that match the next digit in the list. The
// list is circular, so the digit after the last digit is the first digit in the list.

func Captcha(reader io.ByteScanner) int {
	var (
		sum   int  = 0
		first int  = 0
		prev  int  = 0
		init  bool = false
	)

	for {
		b, e := reader.ReadByte()
		if e == io.EOF {
			if init && prev == first {
				sum = prev + sum
			}
			break
		}

		i, skip := toInt(b)
		if skip {
			continue
		}

		if init {
			if prev == i {
				sum = prev + sum
			}
			prev = i
		} else {
			prev = i
			first = i
			init = true
		}
	}
	return sum
}

// Now, instead of considering the next digit, it wants you to consider the
// digit halfway around the circular list. That is, if your list contains 10
// items, only include a digit in your sum if the digit 10/2 = 5 steps
// forward matches it. Fortunately, your list has an even number of elements.

func HalfCaptcha(b []byte) int {
	var (
		sum    int = 0
		length int = len(b)
		j      int = length / 2
	)

	for i := 0; i < length; i++ {
		x := int(b[i] - '0')
		y := int(b[j] - '0')
		if x == y {
			sum = sum + x
		}
		//fmt.Println(length, i, j, "::", x, y,  sum)

		j++
		if j >= length {
			j = 0
		}
	}

	//fmt.Println("---")
	return sum
}
