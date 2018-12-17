package days

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strings"
	"testing"
)

var instructions = regexp.MustCompile(`Step (\w) must be finished before step (\w) can begin`)

type InstructionGraph struct {
	Blocked map[string]map[string]bool
	Next    map[string][]string
	Final   []string
	Queue   []string
}

func CreateGraph(input []string) (*InstructionGraph, int) {
	var (
		to   string
		from string
	)
	numEdges := 0
	graph := InstructionGraph{
		Blocked: make(map[string]map[string]bool),
		Next:    make(map[string][]string),
	}

	for _, s := range input {
		match := instructions.FindStringSubmatch(s)
		from = match[1]
		to = match[2]
		graph.AddEdge(from, to)
		numEdges++
	}

	// prime queue with unblocked items
	for k := range graph.Blocked {
		graph.QueueUnblocked(k)
	}
	sort.Strings(graph.Queue)

	return &graph, numEdges
}

func (g *InstructionGraph) AddEdge(from string, to string) {
	// Make sure both nodes exist: Find the roots and leaves
	if _, ok := g.Blocked[to]; !ok {
		g.Blocked[to] = map[string]bool{}
	}
	if _, ok := g.Blocked[from]; !ok {
		g.Blocked[from] = map[string]bool{}
	}
	if _, ok := g.Next[to]; !ok {
		g.Next[to] = []string{}
	}
	if _, ok := g.Next[from]; !ok {
		g.Next[from] = []string{}
	}

	g.Blocked[to][from] = true
	g.Next[from] = append(g.Next[from], to)
}

func (g *InstructionGraph) QueueUnblocked(k string) {
	if len(g.Blocked[k]) == 0 {
		g.Queue = append(g.Queue, k)
		delete(g.Blocked, k)
	}
}

func (g *InstructionGraph) TakeFirst() (string, bool) {
	// assumes sort maintained elsewhere (e.g. Complete)
	if len(g.Queue) > 0 {
		first := g.Queue[0]
		g.Queue = g.Queue[1:]
		return first, true
	}
	return "", false
}

func (g *InstructionGraph) Complete(k string) {
	g.Final = append(g.Final, k)

	// Free up items blocked by k
	next := g.Next[k]
	for _, x := range next {
		delete(g.Blocked[x], k)
		g.QueueUnblocked(x)
	}

	// sort queue
	sort.Strings(g.Queue)
}

type Work struct {
	ticks           int
	baseDuration    int
	numInstructions int
	workers         []Worker
}

type Worker struct {
	id           string
	assigned     string
	completeTime int
}

func CreateWork(numWorkers int, baseDuration int) *Work {
	work := Work{}
	work.ticks = -1 // increment to 0 at start
	work.baseDuration = baseDuration
	work.workers = make([]Worker, numWorkers)
	for x := range work.workers {
		work.workers[x].id = fmt.Sprintf("Wkr-%d", x)
	}
	return &work
}

func (w *Work) GetTime(s string) int {
	return int(s[0]-64) + w.baseDuration
}

func offerWork(work *Work, g *InstructionGraph, w Worker) Worker {
	if w.completeTime != 0 {
		if w.completeTime == work.ticks {
			g.Complete(w.assigned)
			//fmt.Println(w.id, " DONE:", g.Final, w)
			w.assigned = ""
			w.completeTime = 0
		}
	}

	if w.completeTime == 0 {
		if step, ok := g.TakeFirst(); ok {
			w.assigned = step
			w.completeTime = work.ticks + work.GetTime(step)
			//fmt.Println(w.id, "START:", step, work.GetTime(step), g.Queue)
		} else {
			//fmt.Println(w.id, " WAIT:", w, g.Queue, g.Final)
		}
	} else {
		//fmt.Println(w.id, " HOLD:", w, g.Queue, g.Final)
	}

	return w
}

// PART TWO
func ParallelTraverse(g *InstructionGraph, w *Work) []string {
	// While there are elements remaining..
	for len(g.Final) < len(g.Next) {
		// increment tick
		w.ticks++

		// Offer work to workers
		for i, worker := range w.workers {
			w.workers[i] = offerWork(w, g, worker)
		}
	}
	return g.Final
}

// PART ONE
func SingleTraverse(g *InstructionGraph) []string {
	// While there are elements still on the queue..
	for len(g.Queue) > 0 {
		first := g.Queue[0]
		g.Queue = g.Queue[1:]
		g.Complete(first)
	}

	return g.Final
}

var day7_sampleInput = []string{
	"Step C must be finished before step A can begin.",
	"Step C must be finished before step F can begin.",
	"Step A must be finished before step B can begin.",
	"Step A must be finished before step D can begin.",
	"Step B must be finished before step E can begin.",
	"Step D must be finished before step E can begin.",
	"Step F must be finished before step E can begin.",
}

func TestSampleData_7part1(t *testing.T) {
	graph, numEdges := CreateGraph(day7_sampleInput)
	fmt.Println(graph)
	result := SingleTraverse(graph)

	if numEdges != len(day7_sampleInput) {
		t.Error("Expected", len(day7_sampleInput), "edges, got", numEdges)
	}
	if strings.Join(result, "") != "CABDFE" {
		t.Error("Expected CABDFE from test data, got", result)
	}
}

func TestSampleData_7part1_2(t *testing.T) {
	input := append(day7_sampleInput, "Step X must be finished before step B can begin.")
	graph, numEdges := CreateGraph(input)
	result := SingleTraverse(graph)

	if numEdges != len(input) {
		t.Error("Expected", len(input), "edges, got", numEdges)
	}
	if strings.Join(result, "") != "CADFXBE" {
		t.Error("Expected CADFXBE from test data, got", result)
	}
}

func TestInput_7part1(t *testing.T) {
	content, err := ioutil.ReadFile("day07_input.txt")
	check(err)

	defer elapsed("TestInput_7part1")() // time execution of the rest

	list := strings.Split(strings.TrimSpace(string(content)), "\n")
	graph, _ := CreateGraph(list)
	result := SingleTraverse(graph)

	fmt.Println("Day 7 / Part 1 Result", strings.Join(result, ""))
}

func TestSampleData_7part2(t *testing.T) {
	graph, numEdges := CreateGraph(day7_sampleInput)
	fmt.Println(graph)

	work := CreateWork(2, 0)
	fmt.Println(work)

	result := ParallelTraverse(graph, work)

	if numEdges != len(day7_sampleInput) {
		t.Error("Expected", len(day7_sampleInput), "edges, got", numEdges)
	}
	if strings.Join(result, "") != "CABFDE" {
		t.Error("Expected CABDFE from test data, got", result)
	}
	if work.ticks != 15 {
		t.Error("Expected 15 ticks, got", work.ticks)
	}
}

func TestInput_7part2(t *testing.T) {
	content, err := ioutil.ReadFile("day07_input.txt")
	check(err)

	defer elapsed("TestInput_7part2")() // time execution of the rest

	list := strings.Split(strings.TrimSpace(string(content)), "\n")
	graph, _ := CreateGraph(list)

	work := CreateWork(5, 60)
	result := ParallelTraverse(graph, work)

	fmt.Println("Day 7 / Part 2 Result", work.ticks, strings.Join(result, ""))
}
