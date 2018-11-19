package days

import (
 "io"
)

// The captcha requires you to review a sequence of digits (your puzzle input)
// and find the sum of all digits that match the next digit in the list. The 
// list is circular, so the digit after the last digit is the first digit in the list.

func Captcha(reader io.ByteScanner) int {
  var (
    sum int = 0
    first int = 0
    prev int = 0
    init bool = false
  )

  for {
    b,e := reader.ReadByte()
    if e == io.EOF {
      if ( init && prev == first ) {
        sum = prev + sum
      }
      break
    }
    if ( b < '0' || '9' < b ) {
      continue
    }

    i := int(b - '0')

    if ( init ) {
      if ( 0 <= i && i <= 9 && prev == i ) {
        sum = prev + sum
      }
      prev = i
    } else {
      prev = i;
      first = i;
      init = true
    }
  }
  return sum
}
