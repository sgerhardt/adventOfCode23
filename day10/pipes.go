package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Hardcoded filename
	filename := "day10/pipes.txt"

	part1(filename)
	//part2(filename)
}

func part1(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var input [][]rune
	idx := 0
	for scanner.Scan() {
		input = append(input, []rune{})
		line := scanner.Text()
		//print(line)
		for _, char := range line {
			input[idx] = append(input[idx], char)
		}
		//println(input[idx])
		idx++
	}
	fmt.Println(pipeFarthestPoint(input))
}

func pipeFarthestPoint(input [][]rune) int {
	startRow, startCol := findStartingPoint(input)

	distance := traversePipe(input, startRow, startCol)
	fmt.Println(distance)
	// find the starting point, "S", and follow the pipe until it reaches back to itself.
	// The farthest distance from the start is the pipe length / 2.
	return distance / 2
}

func traversePipe(input [][]rune, row, col int) int {

	// search around the starting point for a connecting pipe
	startPosition := []int{row, col}

	prevPosition := []int{-1, -1}
	currentPosition := []int{row, col}

	distanceTravelled := 0

	foundStart := distanceTravelled > 0 && startPosition[0] == currentPosition[0] && startPosition[1] == currentPosition[1]
	for !foundStart {
		for idx, directionVals := range directions(currentPosition[0], currentPosition[1]) {
			if foundStart {
				// we've reached the end of the pipe
				return distanceTravelled
			}
			if directionVals[0] < 0 || directionVals[1] < 0 {
				continue
			} else if directionVals[0] >= len(input) || directionVals[1] >= len(input[directionVals[0]]) {
				continue
			}

			checkingPosition := input[directionVals[0]][directionVals[1]]

			//TODO need to check for allowed directions based on current position

			//validOptions := validConnectingPipeOptions(checkingPosition, directionVals[0]-currentPosition[0], directionVals[1]-currentPosition[1])
			if idx == 0 { // north
				// check above for potential connection pipes (|, 7, F)
				if checkingPosition == 'S' && !isPrev(prevPosition, []int{directionVals[0], directionVals[1]}) {
					// we've looped back
					prevPosition = currentPosition
					currentPosition = directionVals
					distanceTravelled++
					foundStart = true
					break
				}
				if checkingPosition == '|' && !isPrev(prevPosition, []int{directionVals[0], directionVals[1]}) {
					prevPosition = currentPosition
					currentPosition = directionVals
					distanceTravelled++
					break
				} else if checkingPosition == '7' && !isPrev(prevPosition, []int{directionVals[0], directionVals[1]}) {
					prevPosition = currentPosition
					currentPosition = directionVals
					distanceTravelled++
					break
				} else if checkingPosition == 'F' && !isPrev(prevPosition, []int{directionVals[0], directionVals[1]}) {
					prevPosition = currentPosition
					currentPosition = directionVals
					distanceTravelled++
					break
				}
			} else if idx == 1 { // south
				// check below for potential connection pipes (|, L, J)
				if checkingPosition == 'S' && !isPrev(prevPosition, []int{directionVals[0], directionVals[1]}) {
					// we've looped back
					prevPosition = currentPosition
					currentPosition = directionVals
					distanceTravelled++
					foundStart = true
					break
				}
				if checkingPosition == '|' && !isPrev(prevPosition, []int{directionVals[0], directionVals[1]}) {
					prevPosition = currentPosition
					currentPosition = directionVals
					distanceTravelled++
					break
				} else if checkingPosition == 'L' && !isPrev(prevPosition, []int{directionVals[0], directionVals[1]}) {
					prevPosition = currentPosition
					currentPosition = directionVals
					distanceTravelled++
					break
				} else if checkingPosition == 'J' && !isPrev(prevPosition, []int{directionVals[0], directionVals[1]}) {
					prevPosition = currentPosition
					currentPosition = directionVals
					distanceTravelled++
					break
				}
			} else if idx == 2 {
				// check west for potential connection pipes (-, F, L)
				if checkingPosition == 'S' && !isPrev(prevPosition, []int{directionVals[0], directionVals[1]}) {
					// we've looped back
					prevPosition = currentPosition
					currentPosition = directionVals
					distanceTravelled++
					foundStart = true
					break
				}
				if checkingPosition == '-' && !isPrev(prevPosition, []int{directionVals[0], directionVals[1]}) {
					prevPosition = currentPosition
					currentPosition = directionVals
					distanceTravelled++
					break
				} else if checkingPosition == 'F' && !isPrev(prevPosition, []int{directionVals[0], directionVals[1]}) {
					prevPosition = currentPosition
					currentPosition = directionVals
					distanceTravelled++
					break
				} else if checkingPosition == 'L' && !isPrev(prevPosition, []int{directionVals[0], directionVals[1]}) {
					prevPosition = currentPosition
					currentPosition = directionVals
					distanceTravelled++
					break
				}
			} else if idx == 3 {
				// check east for potential connection pipes (-, 7, J)
				if checkingPosition == 'S' && !isPrev(prevPosition, []int{directionVals[0], directionVals[1]}) {
					// we've looped back
					prevPosition = currentPosition
					currentPosition = directionVals
					distanceTravelled++
					foundStart = true
					break
				}

				if checkingPosition == '-' && !isPrev(prevPosition, []int{directionVals[0], directionVals[1]}) {
					prevPosition = currentPosition
					currentPosition = directionVals
					distanceTravelled++
					break
				} else if checkingPosition == '7' && !isPrev(prevPosition, []int{directionVals[0], directionVals[1]}) {
					prevPosition = currentPosition
					currentPosition = directionVals
					distanceTravelled++
					break
				} else if checkingPosition == 'J' && !isPrev(prevPosition, []int{directionVals[0], directionVals[1]}) {
					prevPosition = currentPosition
					currentPosition = directionVals
					distanceTravelled++
					break
				}
				prevPosition = currentPosition
			}

		}
	}

	return distanceTravelled
}

func validConnectingPipeOptions(currentPipeType rune, adjoiningPipePositionRow, adjoiningPipePositionCol int) []rune {
	if currentPipeType == '|' {
		if adjoiningPipePositionRow == -1 && adjoiningPipePositionCol == 0 {
			return []rune{'|', 'F', '7'}
		} else if adjoiningPipePositionRow == 1 && adjoiningPipePositionCol == 0 {
			return []rune{'|', 'L', 'J'}
		} else {
			return []rune{}
		}
	} else if currentPipeType == '-' {
		if adjoiningPipePositionRow == 0 && adjoiningPipePositionCol == -1 {
			return []rune{'-', 'F', 'L'}
		} else if adjoiningPipePositionRow == 1 && adjoiningPipePositionCol == 1 {
			return []rune{'-', '7', 'J'}
		} else {
			return []rune{}
		}
	} else if currentPipeType == '7' {
		if adjoiningPipePositionRow == 0 && adjoiningPipePositionCol == -1 {
			return []rune{'-', 'F', 'L'}
		} else if adjoiningPipePositionRow == 1 && adjoiningPipePositionCol == 0 {
			return []rune{'L', '|', 'J'}
		} else {
			return []rune{}
		}
	} else if currentPipeType == 'F' {
		if adjoiningPipePositionRow == 0 && adjoiningPipePositionCol == 1 {
			return []rune{'-', 'J', '7'}
		} else if adjoiningPipePositionRow == 1 && adjoiningPipePositionCol == 0 {
			return []rune{'L', '|', 'J'}
		} else {
			return []rune{}
		}
	} else if currentPipeType == 'L' {
		if adjoiningPipePositionRow == -1 && adjoiningPipePositionCol == 0 {
			return []rune{'|', 'F', '7'}
		} else if adjoiningPipePositionRow == 0 && adjoiningPipePositionCol == 1 {
			return []rune{'-', '7', 'J'}
		} else {
			return []rune{}
		}
	} else if currentPipeType == 'J' {
		if adjoiningPipePositionRow == 0 && adjoiningPipePositionCol == -1 {
			return []rune{'-', 'F', 'L'}
		} else if adjoiningPipePositionRow == -1 && adjoiningPipePositionCol == 0 {
			return []rune{'|', 'F', '7'}
		} else {
			return []rune{}
		}
	}
	return []rune{}
}

func isPrev(prev, current []int) bool {
	return prev[0] == current[0] && prev[1] == current[1]
}

func directions(x, y int) [][]int {
	return [][]int{
		{x - 1, y}, // north
		{x + 1, y}, // south
		{x, y - 1}, // west
		{x, y + 1}, // east
	}
}

func findStartingPoint(input [][]rune) (int, int) {
	for row, line := range input {
		for col, char := range line {
			if char == 'S' {
				println("Found starting point at ", row, col)
				return row, col
			}
		}
	}
	return -1, -1
}
