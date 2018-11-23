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

type Key struct {
  x, y int
}

func GetValue(grid map[Key]int, x int, y int) int {
  if ( x == 0 && y == 0 ) {
    return 1
  }
	var total int = 0
  //fmt.Println(grid[Key{x-1,y+1}], grid[Key{x,y+1}], grid[Key{x+1,y+1}])
  total = total + grid[Key{x-1, y+1}]
  total = total + grid[Key{x,   y+1}]
  total = total + grid[Key{x+1, y+1}]
	//fmt.Println(grid[Key{x-1,y}], "X", grid[Key{x+1,y}])
  total = total + grid[Key{x-1, y}]
  total = total + grid[Key{x+1, y}]
	//fmt.Println(grid[Key{x-1,y-1}], grid[Key{x,y-1}], grid[Key{x+1,y-1}])
  total = total + grid[Key{x-1, y-1}]
  total = total + grid[Key{x,   y-1}]
  total = total + grid[Key{x+1, y-1}]
  //fmt.Printf("Value {%d,%d} = %d\n",x,y,total)
	return total
}

func pretty(key Key) string {
  switch key {
    case Key{1,0}: return "\u2192"
    case Key{0,1}: return "\u2191"
    case Key{-1,0}: return "\u2190"
    case Key{0,-1}: return "\u2193"
  }
  return ""
}

func NextKey(grid map[Key]int, key Key, direction Key) (Key,Key) {
  key.x = key.x + direction.x
  key.y = key.y + direction.y
  //fmt.Println("Current ", pretty(direction), key)

  switch direction {
    case Key{1,0}:
      if ( grid[Key{key.x,key.y+1}] == 0 ) {
				direction.x = 0
				direction.y = 1
			}
			break
    case Key{0,1}:
      if ( grid[Key{key.x-1,key.y}] == 0 ) {
				direction.x = -1
				direction.y = 0
      }
      break
    case Key{-1,0}:
      if ( grid[Key{key.x,key.y-1}] == 0 ) {
				direction.x = 0
				direction.y = -1
			}
      break
    case Key{0,-1}:
      if ( grid[Key{key.x+1,key.y}] == 0 ) {
				direction.x = 1
				direction.y = 0
			}
      break
    default:
  }
  return key, direction
}

func SpiralSum(index int, target int) int {
	grid := make(map[Key]int)
  var n int = 1
  var key = Key {0, 0}
  var direction = Key {1, 0}

  for {
    grid[key] = GetValue(grid, key.x, key.y)
    if ( index > 0 && n >= index ) {
			break // test values given index
    }
    if ( target > 0 && target < grid[key] ) {
      break // value we're looking for
    }
    n++
    key,direction = NextKey(grid, key, direction)
  }
  return grid[key]
}
