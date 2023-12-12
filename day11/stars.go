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
	for scanner.Scan() {
		line := scanner.Text()
		print(line)
		if !strings.Contains(line, "#") {
			// we found an empty row
			emptyRowIndexes = append(emptyRowIndexes, idx)
		}
		idx++
	}

	// double the empty rows now...

	for idx, emptyRow := range emptyRowIndexes {
		addEmptyRow(input, emptyRow)
	}
}

func addEmptyRow(input [][]string, idx int) [][]string {

	emptyRow := make([]string, len(input[0]))
	for i := 0; i < len(input[0]); i++ {
		emptyRow[i] = "."
	}

	for i := 0; i < len(input); i++ {
		println(input[i])

		slices.Insert(input, idx, emptyRow)
	}

	return input
}
