package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var numberWords map[string]string

func init() {
	numberWords = map[string]string{
		"zero": "0", "one": "1", "two": "2", "three": "3", "four": "4",
		"five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9",
	}
}

func main() {
	// Hardcoded filename
	filename := "day1/calibration.txt"

	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		combined := CalcSum(line)
		sum += combined
	}
	println(sum)
}

func CalcSum(line string) int {
	numberMap := findAllNumbers(line, numberWords)
	// get the min and max index for the numbers
	minIndex := math.MaxInt64
	maxIndex := int64(-1)
	minDigit := ""
	maxDigit := ""
	for index, num := range numberMap {
		if index < int64(minIndex) {
			minIndex = int(index)
			minDigit = num
		}
		if index > maxIndex {
			maxIndex = index
			maxDigit = num
		}
	}

	combined, _ := strconv.Atoi(minDigit + maxDigit)
	return combined
}

func findAllNumbers(line string, numberWords map[string]string) map[int64]string {
	indexToNum := map[int64]string{}
	// find all digits in the line and populate a map with the index of the digit
	for idx, char := range line {
		if char >= '0' && char <= '9' {
			indexToNum[int64(idx)] = string(char)
		}
	}

	// find all spelt out words from numberwords that may be in the line and add their index to numToIndex
	for word, digit := range numberWords {
		idx := strings.Index(line, word)
		if idx > -1 {
			indexToNum[int64(idx)] = digit
		}
		lastIdx := strings.LastIndex(line, word)
		if lastIdx > -1 {
			indexToNum[int64(lastIdx)] = digit
		}
	}

	return indexToNum
}
