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
	md      int
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

func PrettyArea(g Grid) string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("(%d,%d) to (%d,%d)\n", g.xMin, g.yMin, g.xMax, g.yMax))
	for j := g.yMin; j <= g.yMax; j++ {
		for i := g.xMin; i <= g.xMax; i++ {
			p := g.plot[Point{i, j}].nearest
			if p.x < 0 {
				builder.WriteString("(.......)")
			} else {
				builder.WriteString(fmt.Sprintf("(%3d,%3d)", p.x, p.y))
			}
			builder.WriteString(" ")
		}
		builder.WriteString("\n")
	}
	return builder.String()
}

func PrettySafe(g Grid, threshold int) string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("(%d,%d) to (%d,%d)\n", g.xMin, g.yMin, g.xMax, g.yMax))
	for j := g.yMin; j <= g.yMax; j++ {
		for i := g.xMin; i <= g.xMax; i++ {
			a := g.plot[Point{i, j}].md
			if a < threshold {
				builder.WriteString("*")
			} else {
				builder.WriteString(" ")
			}
			builder.WriteString(" ")
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

var MdConflict = Point{-2, -2}

func FindMD(origin Point, grid Grid, coords []Point) {
	info := grid.plot[origin]

	if origin.x == grid.xMin ||
		origin.x == grid.xMax ||
		origin.y == grid.yMin ||
		origin.y == grid.yMax {
		info.edge = true
	}

	minDistance := 100 // arbitrary big number
	nearest := origin  // start with this

	for _, i := range coords {
		md := Abs(origin.x-i.x) + Abs(origin.y-i.y)

		// accumulate the md to all points in the input set
		info.md += md

		// find the nearest coordinate when not in the input set
		if info.name == "" {
			if md < minDistance {
				nearest = i
				minDistance = md
			} else if md == minDistance {
				nearest = MdConflict
			}
		}
	}

	// remember the nearest coordinate when not in the input set
	// (input points are excluded, as they are always shortest
	// for something else)
	if info.name == "" {
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
		grid.plot[pt] = PointInfo{fmt.Sprintf("c%d", idx), pt, 0, false}

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

	// go through points of grid, find Manhattan Distance
	for i := grid.xMin; i <= grid.xMax; i++ {
		for j := grid.yMin; j <= grid.yMax; j++ {
			FindMD(Point{i, j}, grid, input)
		}
	}

	return grid
}

func CountNearestSafe(grid Grid, threshold int) (int, int) {
	tally := map[Point][]PointInfo{}
	safeArea := 0
	for _, i := range grid.plot {
		tally[i.nearest] = append(tally[i.nearest], i)

		if i.md < threshold {
			safeArea++
		}
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

	return maxArea, safeArea
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
	maxArea, safeArea := CountNearestSafe(grid, 32)

	if maxArea != 17 {
		t.Error("Expected largest area to be 17. Got", maxArea)
	}
	if safeArea != 16 {
		t.Error("Expected largest area to be 16. Got", safeArea)
	}
}

func TestInput_6(t *testing.T) {
	content, err := ioutil.ReadFile("day06_input.txt")
	check(err)

	defer elapsed("TestInput_6")() // time execution of the rest

	list := strings.Split(string(content), "\n")
	pts := ParsePoints(list)

	grid := PlotPoints(pts)
	maxArea, safeArea := CountNearestSafe(grid, 10000)
	// fmt.Println(PrettyArea(grid))
	// fmt.Println(PrettySafe(grid, 10000))

	fmt.Println("Day 6 / Part 1 Result", maxArea)
	fmt.Println("Day 6 / Part 2 Result", safeArea)
}
