package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sequencesGeneratedForHistory int

var sequencesForHistory [][]int

func main() {
	// Hardcoded filename
	filename := "day9/oasis.txt"

	part1(filename)
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

	var valueHistories [][]int
	for scanner.Scan() {
		var valueHistory []int

		line := scanner.Text()

		rawVals := strings.Split(line, " ")
		for _, rawVal := range rawVals {
			val, errConv := strconv.Atoi(rawVal)
			if errConv != nil {
				continue
			}
			valueHistory = append(valueHistory, val)
		}
		valueHistories = append(valueHistories, valueHistory)
	}

	sumOfValueHistories := 0
	for _, valueHistory := range valueHistories {
		sequencesForHistory = [][]int{}

		sumOfValueHistories += predictNextValue(valueHistory)
	}
	fmt.Printf("sumOfValueHistories: %v\n", sumOfValueHistories)
}

func predictNextValue(valueHistory []int) int {
	sequencesForHistory = append(sequencesForHistory, valueHistory)

	recurseSequences(valueHistory)
	println(sequencesGeneratedForHistory)

	// first add a zero to the zero sequence, then fill in placeholders from the bottom up
	sequencesForHistory[len(sequencesForHistory)-1] = append(sequencesForHistory[len(sequencesForHistory)-1], 0)
	for i := len(sequencesForHistory) - 2; i >= 0; i-- {
		// The placeholder needs to be the result of increasing the value to its left by the value below it
		sequencesForHistory[i] = append(sequencesForHistory[i], sequencesForHistory[i][len(sequencesForHistory[i])-1]+sequencesForHistory[i+1][len(sequencesForHistory[i+1])-1])
	}

	return sequencesForHistory[0][len(sequencesForHistory[0])-1]
}

func recurseSequences(valueHistory []int) int {
	sequencesGeneratedForHistory++
	diff := calcDifferences(valueHistory)
	sequencesForHistory = append(sequencesForHistory, diff)

	// check if the new sequence is all zeroes
	allZeroes := true
	for _, val := range diff {
		if val != 0 {
			allZeroes = false
			break
		}
	}
	for !allZeroes {
		return recurseSequences(diff)
	}

	// once we have all zeros we can extrapolate the next value
	return valueHistory[len(valueHistory)-1] + diff[len(diff)-1]
}

func calcDifferences(valueHistory []int) []int {
	differences := make([]int, len(valueHistory)-1)
	for i := 0; i < len(valueHistory)-1; i++ {
		differences[i] = valueHistory[i+1] - valueHistory[i]
	}
	return differences
}
