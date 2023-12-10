package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part2(filename string) {
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

		sumOfValueHistories += predictPreviousValue(valueHistory)
	}
	fmt.Printf("sumOfValueHistories: %v\n", sumOfValueHistories)
}

func predictPreviousValue(valueHistory []int) int {
	sequencesForHistory = append(sequencesForHistory, valueHistory)

	recurseSequences(valueHistory)

	// add a zero to the beginning of your sequence of zeroes
	sequencesForHistory[len(sequencesForHistory)-1] = append([]int{0}, sequencesForHistory[len(sequencesForHistory)-1]...)
	for i := len(sequencesForHistory) - 2; i >= 0; i-- {
		// Fill in new first values for each previous sequence.
		// The placeholder needs to be the result of subtracting the value to its right by the value below it
		sequencesForHistory[i] = append([]int{sequencesForHistory[i][0] - sequencesForHistory[i+1][0]}, sequencesForHistory[i]...)
	}

	return sequencesForHistory[0][0]
}
