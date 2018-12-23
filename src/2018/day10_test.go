package days

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
	"testing"
)

var test_10 = []string{
	"position=< 9,  1> velocity=< 0,  2>",
	"position=< 7,  0> velocity=<-1,  0>",
	"position=< 3, -2> velocity=<-1,  1>",
	"position=< 6, 10> velocity=<-2, -1>",
	"position=< 2, -4> velocity=< 2,  2>",
	"position=<-6, 10> velocity=< 2, -2>",
	"position=< 1,  8> velocity=< 1, -1>",
	"position=< 1,  7> velocity=< 1,  0>",
	"position=<-3, 11> velocity=< 1, -2>",
	"position=< 7,  6> velocity=<-1, -1>",
	"position=<-2,  3> velocity=< 1,  0>",
	"position=<-4,  3> velocity=< 2,  0>",
	"position=<10, -3> velocity=<-1,  1>",
	"position=< 5, 11> velocity=< 1, -2>",
	"position=< 4,  7> velocity=< 0, -1>",
	"position=< 8, -2> velocity=< 0,  1>",
	"position=<15,  0> velocity=<-2,  0>",
	"position=< 1,  6> velocity=< 1,  0>",
	"position=< 8,  9> velocity=< 0, -1>",
	"position=< 3,  3> velocity=<-1,  1>",
	"position=< 0,  5> velocity=< 0, -1>",
	"position=<-2,  2> velocity=< 2,  0>",
	"position=< 5, -2> velocity=< 1,  2>",
	"position=< 1,  4> velocity=< 2,  1>",
	"position=<-2,  7> velocity=< 2, -2>",
	"position=< 3,  6> velocity=<-1, -1>",
	"position=< 5,  0> velocity=< 1,  0>",
	"position=<-6,  0> velocity=< 2,  0>",
	"position=< 5,  9> velocity=< 1, -2>",
	"position=<14,  7> velocity=<-2,  0>",
	"position=<-3,  6> velocity=< 2, -1>",
}

var PositionRegexp = regexp.MustCompile(`position=<\s*([0-9-]+),\s*([0-9-]+)> velocity=<\s*([0-9-]+),\s*([0-9-]+)>`)

type NightSky struct {
	lights   []Point
	velocity []Point
}

func ParseLights(input []string) *NightSky {
	ns := NightSky{}

	ns.lights = make([]Point, len(input))
	ns.velocity = make([]Point, len(input))

	for i, s := range input {
		match := PositionRegexp.FindStringSubmatch(s)
		if match == nil {
			panic(fmt.Sprintf("[%s] doesn't match expected line format", s))
		}

		ns.lights[i] = Point{ToInt(match[1]), ToInt(match[2])}
		ns.velocity[i] = Point{ToInt(match[3]), ToInt(match[4])}
	}

	return &ns
}

func MoveLight(xy Point, v Point) Point {
	return Point{xy.x + v.x, xy.y + v.y}
}

func SetBounds(max Point, min Point, xy Point) (Point, Point) {
	if max.x < xy.x {
		max.x = xy.x
	} else if xy.x < min.x {
		min.x = xy.x
	}
	if max.y < xy.y {
		max.y = xy.y
	} else if xy.y < min.y {
		min.y = xy.y
	}
	return max, min
}

func MovePoints(ns *NightSky) (Point, Point, int) {
	var max, min Point
	for i := 0; i < len(ns.lights); i++ {
		next := MoveLight(ns.lights[i], ns.velocity[i])
		max, min = SetBounds(max, min, next)
		ns.lights[i] = next
	}
	box := Point{max.x - min.x, max.y - min.y}
	return min, max, box.x * box.y
}

func PlotLights(ns *NightSky, min Point, max Point) {
	points := map[Point]bool{}
	for _, v := range ns.lights {
		points[v] = true
	}
	for y := min.y; y <= max.y; y++ {
		for x := min.x; x <= max.x; x++ {
			_, p := points[Point{x, y}]
			if p {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func Iterate(ns *NightSky) {
	var (
		min Point
		max Point
		a   int
	)
	seconds := 0
	area := 0
	for {
		min, max, a = MovePoints(ns)
		if a >= area &&  area != 0 {
			break
		}
		if a < 30000 {
			PlotLights(ns, min, max)
		}
		area = a
		seconds++
	}
	fmt.Println("Total seconds:", seconds)
}

func TestSampleData_10(t *testing.T) {
	ns := ParseLights(test_10)
	Iterate(ns)
}

func TestInput_10(t *testing.T) {
	content, err := ioutil.ReadFile("day10_input.txt")
	check(err)

	defer elapsed("TestInput_10")() // time execution

	list := strings.Split(strings.TrimSpace(string(content)), "\n")
	ns := ParseLights(list)
	Iterate(ns)

	fmt.Println("Day 10 Result", "")
}

