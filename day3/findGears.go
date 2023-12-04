package main

import (
	"fmt"
	"strconv"
)

func findGears(schematic [][]rune) {

	// a map of gear to a map of number sequence keys to number sequences.
	// A number sequence key is a string of the form "row-colStart-ColEnd"
	gearToNumberSequences := map[cell]map[string]numberSequence{}

	// find a gear (*)
	for rowIndex, _ := range schematic {
		for colIndex, col := range schematic[rowIndex] {
			if col == '*' {
				potentialGearCell := cell{rowIndex, colIndex}
				// search in all directions for a gear connection
				cells := neighbors(potentialGearCell)
				numSeqMap := map[string]numberSequence{}
				for _, c := range cells {
					// check if it's a number
					_, err := strconv.Atoi(string(schematic[c.x][c.y]))
					if err == nil {
						// If it's a number, it's a potential gear connection.
						// Now complete the number sequence
						sequenceNumValue, minY, maxY := c.completeNumberSequence()
						sequenceKey := strconv.Itoa(c.x) + "-" + strconv.Itoa(minY) + "-" + strconv.Itoa(maxY)
						ns := numberSequence{
							gearNumber: sequenceNumValue,
						}

						numSeqMap[sequenceKey] = ns
						gearToNumberSequences[potentialGearCell] = numSeqMap
					}
				}
			}
		}
	}

	// Now we iterate through the map and any entry in gearToNumberSequences that has exactly two entries is a gear
	sum := 0
	for _, numSeqMap := range gearToNumberSequences {
		if len(numSeqMap) == 2 {
			// get the gear ratio by multiplying the two numbers
			num1 := 1
			num2 := 1
			idx := 0
			for _, ns := range numSeqMap {
				if idx == 0 {
					num1 = ns.gearNumber
				} else {
					num2 = ns.gearNumber
				}
				idx++
			}
			sum += num1 * num2
		}
	}
	fmt.Println(sum)
}
