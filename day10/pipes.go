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

const (
	north int = iota
	south int = iota
	west  int = iota
	east
)

var deadEnds map[string]bool
var restart bool

func init() {
	deadEnds = make(map[string]bool)
}

func traversePipe(input [][]rune, row, col int) int {

	// search around the starting point for a connecting pipe
	startPosition := []int{row, col}

	prevPosition := []int{-1, -1}
	currentPosition := []int{row, col}

	distanceTravelled := 0

	foundStart := distanceTravelled > 0 && startPosition[0] == currentPosition[0] && startPosition[1] == currentPosition[1]
	for !foundStart {

		if restart {
			//panic("dead end found...")
			// TODO restart at the beginning, and omit our previous start direction...
			fmt.Println("dead end found... restarting")
			// this means we've hit a dead end and need to restart from the beginning
			restart = false
			prevPosition = []int{-1, -1}
			currentPosition = []int{row, col}
			distanceTravelled = 0
		}

		if prevPosition[0] == currentPosition[0] && prevPosition[1] == currentPosition[1] {
			deadEnds[fmt.Sprintf("%d,%d", currentPosition[0], currentPosition[1])] = true
		}

		for dir, directionVals := range directions(currentPosition[0], currentPosition[1]) {
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

			//validOptions := validConnectingPipeOptions(checkingPosition, directionVals[0]-currentPosition[0], directionVals[1]-currentPosition[1])
			if dir == north {
				// check above for potential connection pipes (|, 7, F)
				if checkingPosition == 'S' && !isPrev(prevPosition, []int{directionVals[0], directionVals[1]}) &&
					isValidDirection(input, checkingPosition, dir, currentPosition) {
					fmt.Println("Moving north to S at ", directionVals[0], directionVals[1])
					// we've looped back
					prevPosition = currentPosition
					currentPosition = directionVals
					distanceTravelled++
					foundStart = true
					break
				}
				if checkingPosition == '|' && !isPrev(prevPosition, []int{directionVals[0], directionVals[1]}) &&
					isValidDirection(input, checkingPosition, dir, currentPosition) {
					fmt.Println("Moving north to | at ", directionVals[0], directionVals[1])
					prevPosition = currentPosition
					currentPosition = directionVals
					distanceTravelled++

					break
				} else if checkingPosition == '7' && !isPrev(prevPosition, []int{directionVals[0], directionVals[1]}) &&
					isValidDirection(input, checkingPosition, dir, currentPosition) {
					fmt.Println("Moving north to 7 at ", directionVals[0], directionVals[1])
					prevPosition = currentPosition
					currentPosition = directionVals
					distanceTravelled++
					break
				} else if checkingPosition == 'F' && !isPrev(prevPosition, []int{directionVals[0], directionVals[1]}) &&
					isValidDirection(input, checkingPosition, dir, currentPosition) {
					fmt.Println("Moving north to F at ", directionVals[0], directionVals[1])
					prevPosition = currentPosition
					currentPosition = directionVals
					distanceTravelled++
					break
				}
			} else if dir == south {
				// check below for potential connection pipes (|, L, J)
				if checkingPosition == 'S' && !isPrev(prevPosition, []int{directionVals[0], directionVals[1]}) &&
					isValidDirection(input, checkingPosition, dir, currentPosition) {
					fmt.Println("Moving south to S at ", directionVals[0], directionVals[1])
					// we've looped back
					prevPosition = currentPosition
					currentPosition = directionVals
					distanceTravelled++
					foundStart = true
					break
				}
				if checkingPosition == '|' &&
					!isPrev(prevPosition, []int{directionVals[0], directionVals[1]}) &&
					isValidDirection(input, checkingPosition, dir, currentPosition) {
					fmt.Println("Moving south to | at ", directionVals[0], directionVals[1])
					prevPosition = currentPosition
					currentPosition = directionVals
					distanceTravelled++
					break
				} else if checkingPosition == 'L' &&
					!isPrev(prevPosition, []int{directionVals[0], directionVals[1]}) &&
					isValidDirection(input, checkingPosition, dir, currentPosition) {
					fmt.Println("Moving south to L at ", directionVals[0], directionVals[1])
					prevPosition = currentPosition
					currentPosition = directionVals
					distanceTravelled++
					break
				} else if checkingPosition == 'J' && !isPrev(prevPosition, []int{directionVals[0], directionVals[1]}) &&
					isValidDirection(input, checkingPosition, dir, currentPosition) {
					fmt.Println("Moving south to J at ", directionVals[0], directionVals[1])
					prevPosition = currentPosition
					currentPosition = directionVals
					distanceTravelled++
					break
				}
			} else if dir == west {
				// check west for potential connection pipes (-, F, L)
				if checkingPosition == 'S' && !isPrev(prevPosition, []int{directionVals[0], directionVals[1]}) &&
					isValidDirection(input, checkingPosition, dir, currentPosition) {
					fmt.Println("Moving west to S at ", directionVals[0], directionVals[1])

					prevPosition = currentPosition
					currentPosition = directionVals
					distanceTravelled++
					foundStart = true
					break
				}
				if checkingPosition == '-' &&
					!isPrev(prevPosition, []int{directionVals[0], directionVals[1]}) &&
					isValidDirection(input, checkingPosition, dir, currentPosition) {
					fmt.Println("Moving west to - at ", directionVals[0], directionVals[1])
					prevPosition = currentPosition
					currentPosition = directionVals
					distanceTravelled++
					break
				} else if checkingPosition == 'F' && !isPrev(prevPosition, []int{directionVals[0], directionVals[1]}) &&
					isValidDirection(input, checkingPosition, dir, currentPosition) {
					fmt.Println("Moving west to F at ", directionVals[0], directionVals[1])
					prevPosition = currentPosition
					currentPosition = directionVals
					distanceTravelled++
					break
				} else if checkingPosition == 'L' && !isPrev(prevPosition, []int{directionVals[0], directionVals[1]}) &&
					isValidDirection(input, checkingPosition, dir, currentPosition) {
					fmt.Println("Moving west to L at ", directionVals[0], directionVals[1])
					prevPosition = currentPosition
					currentPosition = directionVals
					distanceTravelled++
					break
				}
			} else if dir == east {
				// check east for potential connection pipes (-, 7, J)
				if checkingPosition == 'S' && !isPrev(prevPosition, []int{directionVals[0], directionVals[1]}) &&
					isValidDirection(input, checkingPosition, dir, currentPosition) {
					fmt.Println("Moving east to S at ", directionVals[0], directionVals[1])
					// we've looped back
					prevPosition = currentPosition
					currentPosition = directionVals
					distanceTravelled++
					foundStart = true
					break
				}

				if checkingPosition == '-' && !isPrev(prevPosition, []int{directionVals[0], directionVals[1]}) &&
					isValidDirection(input, checkingPosition, dir, currentPosition) {
					fmt.Println("Moving east to - at ", directionVals[0], directionVals[1])
					prevPosition = currentPosition
					currentPosition = directionVals
					distanceTravelled++
					break
				} else if checkingPosition == '7' && !isPrev(prevPosition, []int{directionVals[0], directionVals[1]}) &&
					isValidDirection(input, checkingPosition, dir, currentPosition) {
					fmt.Println("Moving east to 7 at ", directionVals[0], directionVals[1])
					prevPosition = currentPosition
					currentPosition = directionVals
					distanceTravelled++
					break
				} else if checkingPosition == 'J' && !isPrev(prevPosition, []int{directionVals[0], directionVals[1]}) &&
					isValidDirection(input, checkingPosition, dir, currentPosition) {
					fmt.Println("Moving east to J at ", directionVals[0], directionVals[1])
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

func isValidDirection(input [][]rune, checkingPipeType rune, direction int, currentPosition []int) bool {
	currentPipeType := input[currentPosition[0]][currentPosition[1]]

	nextRow := currentPosition[0]
	nextCol := currentPosition[1]
	if direction == north {
		nextRow++
	} else if direction == south {
		nextRow--
	} else if direction == west {
		nextCol--
	} else if direction == east {
		nextCol++
	}
	_, isDeadEnd := deadEnds[fmt.Sprintf("%d,%d", nextRow, nextCol)]
	if isDeadEnd {
		return false
	}
	_, inDeadEnd := deadEnds[fmt.Sprintf("%d,%d", currentPosition[0], currentPosition[1])]
	if inDeadEnd {
		return false
	}

	if checkingPipeType == 'S' {
		return true
	}

	if currentPipeType == 'S' && checkingPipeType != '.' {
		return true
	}

	// | is a vertical pipe connecting north and south.
	if currentPipeType == '|' && checkingPipeType == '|' && direction == north {
		return true
	} else if currentPipeType == '|' && checkingPipeType == '7' && direction == north {
		return true
	} else if currentPipeType == '|' && checkingPipeType == 'F' && direction == north {
		return true
	} else if currentPipeType == '|' && checkingPipeType == '|' && direction == south {
		return true
	} else if currentPipeType == '|' && checkingPipeType == 'L' && direction == south {
		return true
	} else if currentPipeType == '|' && checkingPipeType == 'J' && direction == south {
		return true
	}

	// - is a horizontal pipe connecting east and west.
	if currentPipeType == '-' && checkingPipeType == '-' && direction == west {
		return true
	} else if currentPipeType == '-' && checkingPipeType == 'F' && direction == west {
		return true
	} else if currentPipeType == '-' && checkingPipeType == 'L' && direction == west {
		return true
	} else if currentPipeType == '-' && checkingPipeType == '-' && direction == east {
		return true
	} else if currentPipeType == '-' && checkingPipeType == '7' && direction == east {
		return true
	} else if currentPipeType == '-' && checkingPipeType == 'J' && direction == east {
		return true
	}

	// 7 is a 90-degree bend connecting south and west.
	if currentPipeType == '7' && checkingPipeType == '|' && direction == south {
		return true
	} else if currentPipeType == '7' && checkingPipeType == 'L' && direction == south {
		return true
	} else if currentPipeType == '7' && checkingPipeType == 'J' && direction == south {
		return true
	} else if currentPipeType == '7' && checkingPipeType == 'L' && direction == west {
		return true
	} else if currentPipeType == '7' && checkingPipeType == '-' && direction == west {
		return true
	} else if currentPipeType == '7' && checkingPipeType == 'F' && direction == west {
		return true
	}

	// F is a 90-degree bend connecting south and east
	if currentPipeType == 'F' && checkingPipeType == '-' && direction == east {
		return true
	} else if currentPipeType == 'F' && checkingPipeType == 'J' && direction == east {
		return true
	} else if currentPipeType == 'F' && checkingPipeType == '7' && direction == east {
		return true
	} else if currentPipeType == 'F' && checkingPipeType == '|' && direction == south {
		return true
	} else if currentPipeType == 'F' && checkingPipeType == 'L' && direction == south {
		return true
	} else if currentPipeType == 'F' && checkingPipeType == 'J' && direction == south {
		return true
	}

	// L is a 90-degree bend connecting north and east
	if currentPipeType == 'L' && checkingPipeType == '|' && direction == north {
		return true
	} else if currentPipeType == 'L' && checkingPipeType == '7' && direction == north {
		return true
	} else if currentPipeType == 'L' && checkingPipeType == 'F' && direction == north {
		return true
	} else if currentPipeType == 'L' && checkingPipeType == '-' && direction == east {
		return true
	} else if currentPipeType == 'L' && checkingPipeType == 'J' && direction == east {
		return true
	} else if currentPipeType == 'L' && checkingPipeType == '7' && direction == east {
		return true
	}

	// J is a 90-degree bend connecting north and west.
	if currentPipeType == 'J' && checkingPipeType == '|' && direction == north {
		return true
	} else if currentPipeType == 'J' && checkingPipeType == '7' && direction == north {
		return true
	} else if currentPipeType == 'J' && checkingPipeType == 'F' && direction == north {
		return true
	} else if currentPipeType == 'J' && checkingPipeType == 'F' && direction == west {
		return true
	} else if currentPipeType == 'J' && checkingPipeType == '-' && direction == west {
		return true
	} else if currentPipeType == 'J' && checkingPipeType == 'L' && direction == west {
		return true
	}

	//todo can L (north and east) connect to 7 (south and west)?

	return false
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
