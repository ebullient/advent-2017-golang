package days

import (
	// "bufio"
	"fmt"
	// "os"
	// "strings"
	"testing"
)

type Marble struct {
	value int
	prev  *Marble
	next  *Marble
}

type MarbleGame struct {
	scores     []int
	allMarbles map[int]Marble
	current    *Marble
	lastMarble int
	highScore  int
}

func CreateGame(numPlayers int, lastMarble int) *MarbleGame {
	game := MarbleGame{}
	game.scores = make([]int, numPlayers)
	game.lastMarble = lastMarble

	// start with a perfect circle
	marble := Marble{}
	marble.prev = &marble
	marble.next = &marble

	game.current = &marble

	// we also have a straight list (exta / sanity)
	game.allMarbles = map[int]Marble{}
	game.allMarbles[0] = marble

	return &game
}

func (game *MarbleGame) Place(value int) *Marble {
	m := Marble{}
	m.value = value

	// after<->next
	after := game.current.next     // one from current
	next := game.current.next.next // two from current

	// after<->m<->next
	after.next = &m
	m.prev = after
	m.next = next
	next.prev = &m

	game.current = &m
	game.allMarbles[value] = m
	return &m
}

func (game *MarbleGame) RemoveSeventh() int {
	m := game.current
	for i := 0; i < 7; i++ {
		m = m.prev
	}

	// Remove the sevent back from current
	// Use its next as current
	prev := m.prev
	next := m.next
	game.current = next

	prev.next = next
	next.prev = prev

	delete(game.allMarbles, m.value)
	return m.value
}

func (game *MarbleGame) NextPlayer(j int) int {
	j++
	if j == len(game.scores) {
		return 0
	}
	return j
}

func (game *MarbleGame) Play() int {
	var elf int

	for i := 1; i <= game.lastMarble; i++ {
		if i%23 == 0 {
			more := game.RemoveSeventh()
			game.scores[elf] += i + more
			if game.scores[elf] > game.highScore {
				game.highScore = game.scores[elf]
			}
			//fmt.Println("23! Elf", elf, "gets", i, "+",more,"points")
		} else {
			//fmt.Println("-- new current", game.current.value, prev, next, "==", game.current)
			game.Place(i)
			//fmt.Printf("[%3d]: Place (%d)\n", elf, i)
		}
		elf = game.NextPlayer(elf)
	}
	return game.highScore
}

var test_9part1 = [][]int{
	{9, 25, 32},
	{10, 1618, 8317},
	{13, 7999, 146373},
	{17, 1104, 2764},
	{21, 6111, 54718},
	{30, 5807, 37305},
}

func TestSampleData_9part1(t *testing.T) {
	for _, sample := range test_9part1 {
		// play game with n players
		game := CreateGame(sample[0], sample[1])
		highScore := game.Play()
		if highScore != sample[2] {
			t.Error("For", sample[0], "players, expected high score of", sample[2], ", got", highScore)
		}
	}
}

func TestInput_9part1(t *testing.T) {
	defer elapsed("TestInput_9part1")() // time execution of the rest
	//418 players; last marble is worth 70769 points
	game := CreateGame(418, 70769)
	highScore := game.Play()
	fmt.Println("Day 9 / Part 1 Result", highScore)
}

func TestInput_9part(t *testing.T) {
	defer elapsed("TestInput_9part2")() // time execution of the rest

	//418 players; last marble is worth 70769 points
	game := CreateGame(418, 7076900)
	highScore := game.Play()
	fmt.Println("Day 9 / Part 2 Result", highScore)
}
