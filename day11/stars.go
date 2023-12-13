package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

func main() {
	part1("day11/test.txt")
}

func part1(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var input [][]string
	idx := 0
	var emptyRowIndexes []int
	var emptyColIndexes []int
	for scanner.Scan() {
		line := scanner.Text()
		//print(line)
		if !strings.Contains(line, "#") {
			// we found an empty row
			emptyRowIndexes = append(emptyRowIndexes, idx)
		}
		input = append(input, strings.Split(line, ""))
		idx++
	}

	// check for cols that only contain "."
	for col := 0; col < len(input[0]); col++ {
		emptyCol := true
		for row := 0; row < len(input); row++ {
			if input[row][col] != "." {
				emptyCol = false
				break
			}
		}
		if emptyCol {
			emptyColIndexes = append(emptyColIndexes, col)
		}
	}

	// double the empty rows now...
	for i, emptyRow := range emptyRowIndexes {
		input = addEmptyRow(input, emptyRow+i)
	}

	for i, emptyCol := range emptyColIndexes {
		input = addEmptyCol(input, emptyCol+i)
	}
	for _, rows := range input {
		for _, val := range rows {
			print(val)
		}
		println()
	}

	galaxies := gatherGalaxies(input)
	sum := 0
	for _, g := range galaxies {
		for _, otherGalaxy := range galaxies {
			if g == otherGalaxy {
				continue
			}
			if _, ok := galaxyPairs[otherGalaxy]; ok {
				// we already calculated this distance
				continue
			}
			distance := calcDistance(g.pos, otherGalaxy.pos)
			galaxyPairs[g] = otherGalaxy
			sum += distance
			fmt.Printf("Between galaxy %d and galaxy %d from galaxy: %d \n", g.number, otherGalaxy.number, distance)
			//fmt.Println(distance)
		}
	}
	fmt.Println("Sum of distances:", sum)
}

var galaxyPairs map[galaxy]galaxy

func init() {
	galaxyPairs = make(map[galaxy]galaxy)
}

func addEmptyRow(input [][]string, idx int) [][]string {
	emptyRow := make([]string, len(input[0]))
	for i := 0; i < len(input[0]); i++ {
		emptyRow[i] = "."
	}
	return slices.Insert(input, idx, emptyRow)
}

func addEmptyCol(input [][]string, idx int) [][]string {
	for row := 0; row < len(input); row++ {
		input[row] = slices.Insert(input[row], idx, ".")
	}
	return input
}

type position struct {
	row int
	col int
}

type galaxy struct {
	number int
	pos    position
}

func gatherGalaxies(input [][]string) []galaxy {
	// find all the galaxies
	var galaxies []galaxy
	galaxNum := 1
	for row := 0; row < len(input); row++ {
		for col := 0; col < len(input[0]); col++ {
			if input[row][col] == "#" {
				// we found a galaxy
				galaxies = append(galaxies, galaxy{
					number: galaxNum,
					pos: position{
						row: row,
						col: col,
					},
				})
				galaxNum++
			}
		}
	}
	return galaxies
}

func calcDistance(position1, position2 position) int {
	rise := int(math.Abs(float64(position1.row - position2.row)))
	run := int(math.Abs(float64(position1.col - position2.col)))

	// TODO can't use pythagorean theorem, we are supposed to count discreet steps instead
	return int(math.Sqrt(float64(run*run + rise*rise)))
}
