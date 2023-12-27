package main

import (
	"bufio"
	"fmt"
	"github.com/beefsack/go-astar"
	"log"
	"os"
	"strconv"
)

func main() {
	part1("day17/test.txt")
	distance := optimalPath()
	fmt.Println(distance)
	//part2(filename)
}

var input [][]*Tile

func part1(filename string) {
	// Hardcoded filename
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	idx := 0
	for scanner.Scan() {
		input = append(input, []*Tile{})
		line := scanner.Text()
		for col, digit := range line {
			d, _ := strconv.Atoi(string(digit))
			input[idx] = append(input[idx], &Tile{
				Row:   idx,
				Col:   col,
				Value: d,
			})
		}
		idx++
	}
}

func optimalPath() int {
	startTile := input[0][0]
	endTile := input[len(input)-1][len(input[0])-1]
	_, distance, found := astar.Path(startTile, endTile)
	if !found {
		log.Println("Could not find path")
	}
	return int(distance)
}

func (t *Tile) Up() *Tile {
	if t.Row == 0 {
		return nil
	}
	return input[t.Row-1][t.Col]
}

func (t *Tile) Down() *Tile {
	if t.Row == len(input)-1 {
		return nil
	}
	return input[t.Row+1][t.Col]
}

func (t *Tile) Left() *Tile {
	if t.Col == 0 {
		return nil
	}
	return input[t.Row][t.Col-1]
}

func (t *Tile) Right() *Tile {
	if t.Col == len(input[0])-1 {
		return nil
	}
	return input[t.Row][t.Col+1]
}

type Tile struct {
	Row   int
	Col   int
	Value int
}

func (t *Tile) PathNeighbors() []astar.Pather {
	directions := []func() *Tile{t.Up, t.Down, t.Left, t.Right}
	neighbors := make([]astar.Pather, 0, len(directions))

	for _, dir := range directions {
		if neighbor := dir(); neighbor != nil {
			neighbors = append(neighbors, neighbor)
		}
	}

	return neighbors
}

func (t *Tile) PathNeighborCost(to astar.Pather) float64 {
	toT := to.(*Tile)
	return float64(toT.Value)
}

func (t *Tile) PathEstimatedCost(to astar.Pather) float64 {
	// calculate the manhattan distance
	toT := to.(*Tile)
	absX := toT.Col - t.Col
	if absX < 0 {
		absX = -absX
	}
	absY := toT.Row - t.Row
	if absY < 0 {
		absY = -absY
	}
	return float64(absX + absY)
}
