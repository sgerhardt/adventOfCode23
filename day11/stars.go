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
	filename := "day11/stars.txt"
	//part1(filename)
	part2(filename)
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
	for i := 0; i < len(galaxies); i++ {
		// start at i+1 to avoid double counting
		for j := i + 1; j < len(galaxies); j++ {
			g := galaxies[i]
			otherGalaxy := galaxies[j]
			distance := calcDistance(g.pos, otherGalaxy.pos)
			sum += distance
			fmt.Printf("Between galaxy %d and galaxy %d from galaxy: %d \n", g.number, otherGalaxy.number, distance)
		}
	}
	fmt.Println("Sum of distances:", sum)
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
	galaxyNum := 1
	for row := 0; row < len(input); row++ {
		for col := 0; col < len(input[0]); col++ {
			if input[row][col] == "#" {
				// we found a galaxy
				galaxies = append(galaxies, galaxy{
					number: galaxyNum,
					pos: position{
						row: row,
						col: col,
					},
				})
				galaxyNum++
			}
		}
	}
	return galaxies
}

func calcDistance(position1, position2 position) int {
	rise := int(math.Abs(float64(position1.row - position2.row)))
	run := int(math.Abs(float64(position1.col - position2.col)))

	if run == 0 {
		return rise
	} else if rise == 0 {
		return run
	}
	// for each rise, we add two to the part2 until we reach the rise, then we add one to the part2 until we hit the target
	return 2*rise + (run - rise)
}

func part2Old(filename string) {
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

	for i, emptyRow := range emptyRowIndexes {
		input = addMultipleEmptyRows(input, emptyRow+(99*i), 99)
	}

	for i, emptyCol := range emptyColIndexes {
		input = addMultipleEmptyCols(input, emptyCol+(99*i), 99)
	}
	//for _, rows := range input {
	//	for _, val := range rows {
	//		print(val)
	//	}
	//	println()
	//}

	galaxies := gatherGalaxies(input)
	sum := 0
	for i := 0; i < len(galaxies); i++ {
		// start at i+1 to avoid double counting
		for j := i + 1; j < len(galaxies); j++ {
			g := galaxies[i]
			otherGalaxy := galaxies[j]
			distance := calcDistance(g.pos, otherGalaxy.pos)
			sum += distance
			fmt.Printf("Between galaxy %d and galaxy %d from galaxy: %d \n", g.number, otherGalaxy.number, distance)
		}
	}
	fmt.Println("Sum of distances:", sum)
}

func addMultipleEmptyRows(input [][]string, idx int, num int) [][]string {
	for i := 0; i < num; i++ {
		input = addEmptyRow(input, idx+i)
	}
	return input
}

func addMultipleEmptyCols(input [][]string, idx int, num int) [][]string {
	for i := 0; i < num; i++ {
		input = addEmptyCol(input, idx+i)
	}
	return input
}
