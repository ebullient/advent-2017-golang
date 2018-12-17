package days

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
)

type Tree struct {
	nodes       map[int]Node
	totalNodes  int
	sumMetadata int
}

type Node struct {
	id       int
	children []int
	metadata []int
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
	node.children = make([]int, numChildren)

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
		tree.sumMetadata += node.metadata[i]
	}

	// store node in tree
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
		{2, 3, 4},
		{0, 3, 33},
		{1, 1, 2},
		{0, 1, 99},
	}

	scanner := bufio.NewScanner(strings.NewReader(input))
	tree := ParseTree(scanner)
	fmt.Println(tree)

	if tree.sumMetadata != 138 {
		t.Error("Expected metadata to equal 138, got", tree.sumMetadata)
	}
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
		sum := 0
		for _, i := range n.metadata {
			sum += i
		}
		if sum != verify[i][2] {
			t.Error("Expected numChildren to equal", verify[i][2], "got", sum)
		}
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
	fmt.Println("Day 8 / Part 1 Result", tree.sumMetadata)
}

func TestSampleData_8part2(t *testing.T) {
}

func TestInput_8part2(t *testing.T) {
	//	fmt.Println("Day 8 / Part 2 Result", work.ticks, strings.Join(result, ""))
}
