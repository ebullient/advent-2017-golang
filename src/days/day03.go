package days

import (
  //"fmt"
)

func calculate(index int) int {
  var x int = 1      // number of squares from zero on the "x" axis
  var ring int = 8   // number of squares in the x-th ring around the center
  var n int = 2      // value on square
  var next_n int = 2 // next value of square on x axis
  var highest int    // the highest value in the current ring

  for {
    next_n = n + ring + 1
    highest = next_n - (ring/8) - 1

    if ( index < highest ) {
      break;
    }
    n = next_n
    x++
    ring = ring + 8
  }
  // we've found which "ring" we need to be in. The minimum
  // manhattan distance for the specified index is x
  // and the max is x
  var min = x
  var max = x + ring/8
  var direction = 1

  // we're going to work backwards from the highest possible
  // adding (or subtracting) one from the manhattan distance
  // as we go between max and min
  n = highest
  x = max

  for {
		//fmt.Println(index, "|", min, max, "|", x, n)
    if ( index == n ) {
      break;
    }
    if ( x == max ) {
      direction = -1
    }
    if ( x == min ) {
      direction = 1
    }
    x = x + direction * 1
    n--
  }
  //fmt.Println("---", index, x)
  return x
}

func Spiral(index int) int {
  switch index {
    case 1: return 0
    case 2: return 1
    default: return calculate(index)
  }
}

