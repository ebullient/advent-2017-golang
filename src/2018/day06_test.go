package days

import (
	"fmt"
	"io/ioutil"
	//	"reflect"
	//	"regexp"
	"strconv"
	"strings"
	"testing"
)

type Point struct {
	x, y int
}

type PointInfo struct {
	name    string
	nearest Point
	edge    bool
}

func (p PointInfo) String() string {
	if p.nearest.x < 0 {
		return fmt.Sprintf("[n:%s, o: xxxx, e: %t]", p.name, p.edge)
	}
	return fmt.Sprintf("[n:%s, o: %d-%d, e: %t]", p.name, p.nearest.x, p.nearest.y, p.edge)
}

type Grid struct {
	plot map[Point]PointInfo
	xMax int
	xMin int
	yMax int
	yMin int
}

func (g Grid) String() string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("(%d,%d) to (%d,%d)\n", g.xMin, g.yMin, g.xMax, g.yMax))
	for j := g.yMin; j <= g.yMax; j++ {
		for i := g.xMin; i <= g.xMax; i++ {
			p := g.plot[Point{i, j}].nearest
			if p.x < 0 {
				builder.WriteString("(...)")
			} else {
				builder.WriteString(fmt.Sprintf("(%d,%d)", p.x, p.y))
			}
			builder.WriteString("\t")
		}
		builder.WriteString("\n")
	}
	return builder.String()
}

var test_6part1 = []Point{
	{1, 1},
	{1, 6},
	{8, 3},
	{3, 4},
	{5, 5},
	{8, 9},
}

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

var MD_CONFLICT = Point{-2, -2}

func FindNearest(origin Point, grid Grid, coords []Point) {
	info := grid.plot[origin]

	if origin.x == grid.xMin ||
		origin.x == grid.xMax ||
		origin.y == grid.yMin ||
		origin.y == grid.yMax {
		info.edge = true
	}

	if info.name == "" { // don't bother with coordinates in the input set
		minDistance := 100 // arbitrary big number
		nearest := origin  // arbitrary not 0,0

		for _, i := range coords {
			md := Abs(origin.x-i.x) + Abs(origin.y-i.y)
			if md < minDistance {
				nearest = i
				minDistance = md
			} else if md == minDistance {
				nearest = MD_CONFLICT
			}
		}

		// remember the nearest coordinate
		info.nearest = nearest
	}

	grid.plot[origin] = info
}

func PlotPoints(input []Point) Grid {
	grid := Grid{}
	grid.plot = map[Point]PointInfo{}
	grid.xMax = input[0].x
	grid.xMin = input[0].x
	grid.yMax = input[0].y
	grid.yMin = input[0].y

	for idx, pt := range input {
		grid.plot[pt] = PointInfo{fmt.Sprintf("c%d", idx), pt, false}

		// Define the edges of the grid
		if pt.x < grid.xMin {
			grid.xMin = pt.x
		}
		if pt.x > grid.xMax {
			grid.xMax = pt.x
		}
		if pt.y < grid.yMin {
			grid.yMin = pt.y
		}
		if pt.y > grid.yMax {
			grid.yMax = pt.y
		}
	}

	// go through points of grid, find nearest
	for i := grid.xMin; i <= grid.xMax; i++ {
		for j := grid.yMin; j <= grid.yMax; j++ {
			FindNearest(Point{i, j}, grid, input)
		}
	}

	return grid
}

func CountNearest(grid Grid) int {
	tally := map[Point][]PointInfo{}
	for _, i := range grid.plot {
		tally[i.nearest] = append(tally[i.nearest], i)
	}

	maxArea := 0
	for _, pts := range tally {
		area := len(pts)
		for _, pt := range pts {
			if pt.edge {
				area = 0 // infinite
				break
			}
		}
		if area > maxArea {
			maxArea = area
		}
	}

	return maxArea
}

func ParsePoints(list []string) []Point {
	var (
		i, x, y int
		e       error
	)
	pts := []Point{}

	for _, s := range list {
		i = strings.Index(s, ",")
		if i > 0 {
			x, e = strconv.Atoi(strings.TrimSpace(s[0:i]))
			y, e = strconv.Atoi(strings.TrimSpace(s[i+1:]))
			check(e)
			pts = append(pts, Point{x, y})
		}
	}

	return pts
}

func TestSampleData_6part1(t *testing.T) {
	grid := PlotPoints(test_6part1)
	largest := CountNearest(grid)
	//	fmt.Println(grid)
	//	fmt.Println(largest)

	if largest != 17 {
		t.Error("Expected largest area to be 17. Got", largest)
	}
}

func TestInput_6(t *testing.T) {
	content, err := ioutil.ReadFile("day06_input.txt")
	check(err)

	defer elapsed("TestInput_6")() // time execution of the rest

	list := strings.Split(string(content), "\n")
	pts := ParsePoints(list)

	grid := PlotPoints(pts)
	largest := CountNearest(grid)
	//	fmt.Println(grid)
	//	fmt.Println(largest)

	fmt.Println("Day 6 / Part 1 Result", largest)
}
