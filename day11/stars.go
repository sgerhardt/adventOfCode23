package main

import (
	"bufio"
	"fmt"
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
		addEmptyRow(input, emptyRow+i)
	}

	for _, emptyCol := range emptyColIndexes {
		addEmptyCol(input, emptyCol)
	}

	for _, rows := range input {
		for _, val := range rows {
			print(val)
		}
		println()
	}

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
