package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Hardcoded filename
	filename := "day3/test.txt"

	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var schematic [][]rune

	for scanner.Scan() {
		line := scanner.Text()
		row := []rune(line)
		schematic = append(schematic, row)
	}
	partNumbers := findPartNumbers(schematic)
	sum := 0
	for _, num := range partNumbers {
		sum += num
	}
	//fmt.Printf("Part numbers: %v\n", partNumbers)
	fmt.Printf("Part numbers sum: %v\n", sum)

	gearRatios := findGearRatios(schematic)

	gearRatioSum := 0
	for _, ratio := range gearRatios {
		gearRatioSum += ratio
	}
	fmt.Printf("Gear ratios sum: %v\n", gearRatioSum)

}

var schematic [][]rune

type numberSequence struct {
	number        string
	hasPartNumber bool
	gearNumber    int

	connectedGearCell       *cell
	connectedGearCellAnchor *cell // upper right cell location

	anchorCell cell // upper right cell location
}

func (n *numberSequence) reset() {
	n.number = ""
	n.hasPartNumber = false
	n.gearNumber = 0
	n.connectedGearCell = nil
}

func findPartNumbers(s [][]rune) []int {
	schematic = s
	// if the rune is a number, visit each neighbor and check for two things:
	// 1. if the neighbor is a digit, append the digit to the current number
	// 2. if the neighbor is any symbol other than '.', mark the digit sequence to be added to the list of numbers
	var foundPartNumbers []int
	for rowIndex, row := range schematic {
		ns := &numberSequence{}
		for colIndex, c := range row {
			_, err := strconv.Atoi(string(c))
			if err != nil {
				// current number sequence has ended
				if ns.hasPartNumber {
					num, _ := strconv.Atoi(ns.number)
					foundPartNumbers = append(foundPartNumbers, num)
				}
				ns.reset()
				continue
			}
			ns.number += string(c)
			// visit neighbors of digit in all directions
			if isPartNumber(cell{rowIndex, colIndex}) {
				// add the digit's number sequence to the list of numbers
				ns.hasPartNumber = true
			}
			// if at end of row, and we have a number sequence that has a part number, add it to the list of numbers
			if colIndex == len(row)-1 && ns.hasPartNumber {
				num, _ := strconv.Atoi(ns.number)
				foundPartNumbers = append(foundPartNumbers, num)
				ns.reset()
			}
		}
	}
	return foundPartNumbers
}

func findGearRatios(s [][]rune) []int {
	previouslyCalculatedCells := map[cell]struct{}{}
	schematic = s
	var gearRatios []int
	for rowIndex, row := range s {
		ns := numberSequence{}
		for colIndex, c := range row {
			_, err := strconv.Atoi(string(c))
			if err != nil {
				// current number sequence has ended
				ns.anchorCell = cell{rowIndex, colIndex - 1}
				if ns.hasPartNumber && ns.gearNumber != 0 {
					num, _ := strconv.Atoi(ns.number)
					gearRatio := num * ns.gearNumber

					// if cells aren't already calculated, add it to the list of gear ratios
					var calcedConnectedCell bool
					if ns.connectedGearCellAnchor != nil {
						_, calcedConnectedCell = previouslyCalculatedCells[cell{ns.connectedGearCellAnchor.x, ns.connectedGearCellAnchor.y}]
					}
					_, calcedOriginalCell := previouslyCalculatedCells[ns.anchorCell]
					if !calcedConnectedCell && !calcedOriginalCell {
						gearRatios = append(gearRatios, gearRatio)
					}
					if ns.connectedGearCellAnchor != nil {
						previouslyCalculatedCells[cell{ns.connectedGearCellAnchor.x, ns.connectedGearCellAnchor.y}] = struct{}{}
					}
					previouslyCalculatedCells[ns.anchorCell] = struct{}{}
				}
				ns.reset()
				continue
			}
			ns.number += string(c)
			if g, _ := getConnectedGearCell(cell{rowIndex, colIndex}, nil); g != nil {
				// add the digit's number sequence to the list of numbers
				ns.hasPartNumber = true
				var anchor int
				ns.gearNumber, anchor = g.completeNumberSequence()
				ns.connectedGearCellAnchor = &cell{g.x, anchor}
				ns.connectedGearCell = g

				_, anotherConnectedCellErr := getConnectedGearCell(*ns.connectedGearCellAnchor, &cell{rowIndex, colIndex})
				if anotherConnectedCellErr != nil {
					ns.reset()
					continue
				}
			}
			// if at end of row, and we have a number sequence that has a gear, add it to the list of gear ratios
			if colIndex == len(row)-1 && ns.hasPartNumber && ns.gearNumber != 0 {
				ns.anchorCell = cell{rowIndex, colIndex}
				num, _ := strconv.Atoi(ns.number)
				gearRatio := num * ns.gearNumber
				// if cells aren't already calculated, add it to the list of gear ratios
				_, originalCell := previouslyCalculatedCells[cell{rowIndex, colIndex}]
				var connectedCell bool
				if ns.connectedGearCell != nil {
					_, connectedCell = previouslyCalculatedCells[cell{ns.connectedGearCellAnchor.x, ns.connectedGearCellAnchor.y}]
				}
				_, anotherConnectedCellErr := getConnectedGearCell(*ns.connectedGearCellAnchor, &cell{rowIndex, colIndex})
				if anotherConnectedCellErr != nil {
					ns.reset()
					continue
				}
				if !originalCell && !connectedCell {
					gearRatios = append(gearRatios, gearRatio)
				}
				// add cell to map of visited cells
				previouslyCalculatedCells[cell{rowIndex, colIndex}] = struct{}{}
				if ns.connectedGearCell != nil {
					previouslyCalculatedCells[cell{ns.connectedGearCellAnchor.x, ns.connectedGearCellAnchor.y}] = struct{}{}
				}
			}
		}
	}
	return gearRatios
}

func (c *cell) completeNumberSequence() (int, int) {
	base := schematic[c.x][c.y]
	sequence := string(base)
	// prefix number sequence with digits from the left until we hit a non-digit or out of bounds
	for y := c.y - 1; y >= 0; y-- {
		if schematic[c.x][y] >= '0' && schematic[c.x][y] <= '9' {
			// add digit to number sequence
			sequence = string(schematic[c.x][y]) + sequence
		} else {
			break
		}
	}

	// suffix number sequence with digits from the right
	maxY := 0
	for y := c.y + 1; y < len(schematic[0]); y++ {
		if schematic[c.x][y] >= '0' && schematic[c.x][y] <= '9' {
			// add digit to number sequence
			sequence += string(schematic[c.x][y])
			maxY = y
		} else {
			break
		}
	}

	if c.y > maxY {
		maxY = c.y
	}

	n, err := strconv.Atoi(sequence)
	if err != nil {
		n = 0
	}
	return n, maxY
}

func getConnectedGearCell(c cell, previouslyConnectedCell *cell) (*cell, error) {
	// a part number is a gear if it has a neighbor that is a "*" in any direction,
	// and that "*" has another neighbor that is a digit
	connectedCount := 0
	var connectedCell *cell
	for _, neighbor := range neighbors(c) {
		if previouslyConnectedCell != nil && neighbor.x == previouslyConnectedCell.x && neighbor.y == previouslyConnectedCell.y {
			connectedCount++
			continue
		}
		if schematic[neighbor.x][neighbor.y] == '*' {
			// see if any neighbors of the "*" are digits, excluding the current cell
			for _, neighborOfStar := range neighbors(neighbor) {
				if neighborOfStar.x == c.x && neighborOfStar.y == c.y {
					continue
				}
				_, err := strconv.Atoi(string(schematic[neighborOfStar.x][neighborOfStar.y]))
				if err == nil {
					if connectedCount == 1 {
						// we are already connected so this isn't a gear!
						return nil, errors.New("already visited cell")
					}
					connectedCount++
					connectedCell = &cell{
						x: neighborOfStar.x,
						y: neighborOfStar.y,
					}

					// TODO fix this
					return &neighborOfStar, nil
				}
			}
		}
	}
	if connectedCount > 1 {
		// we are already connected so this isn't a gear!
		return nil, errors.New("already visited cell")
	}
	return connectedCell, nil
}

func isPartNumber(c cell) bool {
	// visit neighbors of cell in all directions and return true if a non-period, non-digit symbol is found
	for _, neighbor := range neighbors(c) {
		_, err := strconv.Atoi(string(schematic[neighbor.x][neighbor.y]))
		if err == nil {
			continue
		}
		if schematic[neighbor.x][neighbor.y] != '.' {
			return true
		}
	}
	return false
}

type cell struct {
	x, y int
}

func neighbors(c cell) []cell {
	// return neighbors of cell in all directions - note that we must also bound check
	var n []cell
	for _, neighbor := range directions(c.x, c.y) {
		if neighbor.x >= 0 && neighbor.x < len(schematic) && neighbor.y >= 0 && neighbor.y < len(schematic[0]) {
			n = append(n, cell{neighbor.x, neighbor.y})
		}
	}
	return n
}

func directions(x, y int) []cell {
	return []cell{
		{x - 1, y - 1}, // upper left
		{x - 1, y},     // left
		{x - 1, y + 1}, // lower left
		{x, y - 1},     // up
		{x, y + 1},     // down
		{x + 1, y - 1}, // upper right
		{x + 1, y},     // right
		{x + 1, y + 1}, // lower right
	}
}
