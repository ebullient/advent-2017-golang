package days

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
)

type Tree struct {
	nodes      map[int]Node
	totalNodes int
	total      int
}

type Node struct {
	id       int
	children map[int]int
	metadata []int
	value    int
}

func PrettyTree(tree *Tree) {
	for k, n := range tree.nodes {
		fmt.Println(k, ": ", n.children, ", ", n.metadata)
	}
}

func ReadInt(scanner *bufio.Scanner) int {
	if scanner.Scan() {
		return ToInt(scanner.Text())
	}
	panic("Unable to read int")
}

func ReadNode(tree *Tree, scanner *bufio.Scanner) Node {
	var i int

	node := Node{}
	node.id = tree.totalNodes
	tree.totalNodes++

	// HEADER: num children, num metadata
	numChildren := ReadInt(scanner)
	node.children = map[int]int{}

	numMetadata := ReadInt(scanner)
	node.metadata = make([]int, numMetadata)

	// Child nodes (nested)
	for i = 0; i < numChildren; i++ {
		child := ReadNode(tree, scanner)
		node.children[i] = child.id
	}

	// Metadata for this node
	for i = 0; i < numMetadata; i++ {
		node.metadata[i] = ReadInt(scanner)

		// Part 1: all metadata
		tree.total += node.metadata[i]

		// Part 2
		if numChildren == 0 {
			node.value += node.metadata[i]
		} else {
			x := node.metadata[i] - 1
			if x >= 0 && x < numChildren {
				id := node.children[x]
				value := tree.nodes[id].value
				node.value += value
			}
		}
	}

	// store node in tree (extraneous: test/debug)
	tree.nodes[node.id] = node
	return node
}

func ParseTree(scanner *bufio.Scanner) *Tree {
	scanner.Split(bufio.ScanWords)

	tree := Tree{}
	tree.nodes = map[int]Node{}
	ReadNode(&tree, scanner)

	return &tree
}

func TestSampleData_8part1(t *testing.T) {
	input := "2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2"
	verify := [][]int{
		{2, 3, 4, 66},  // A
		{0, 3, 33, 33}, // B
		{1, 1, 2, 0},   // C
		{0, 1, 99, 99}, // D
	}

	scanner := bufio.NewScanner(strings.NewReader(input))
	tree := ParseTree(scanner)
	fmt.Println(tree)

	if tree.totalNodes != 4 {
		t.Error("Expected 4 nodes, got", tree.totalNodes)
	}
	for i, n := range tree.nodes {
		if len(n.children) != verify[i][0] {
			t.Error("Expected", verify[i][0], "children, got", len(n.children))
		}
		if len(n.metadata) != verify[i][1] {
			t.Error("Expected", verify[i][1], "metadata, got", len(n.metadata))
		}
		// Part 1 sub-parts
		sum := 0
		for _, i := range n.metadata {
			sum += i
		}
		if sum != verify[i][2] {
			t.Error("Expected numChildren to equal", verify[i][2], "got", sum)
		}
		// Part 2
		if n.value != verify[i][3] {
			t.Error("Expected sum of node metadata to equal", verify[i][3], "got", n.value)
		}
	}
	// Part 1 total
	if tree.total != 138 {
		t.Error("Expected metadata to equal 138, got", tree.total)
	}
}

func TestInput_8part1(t *testing.T) {
	f, err := os.Open("day08_input.txt")
	check(err)
	defer f.Close()

	reader := bufio.NewReader(f)
	_, err = reader.Peek(2)
	check(err)

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	defer elapsed("TestInput_8part1")() // time execution of the rest

	tree := ParseTree(scanner)
	fmt.Println("Day 8 / Part 1 Result", tree.total)
	fmt.Println("Day 8 / Part 2 Result", tree.nodes[0].value)
}
