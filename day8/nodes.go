package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Hardcoded filename
	filename := "day8/nodes.txt"

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

	lineCount := 0
	var instructions []rune
	var nodes []*node
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if lineCount == 0 {
			instructions = []rune(line)
		} else {
			name := line[0:3]

			leftNodeStart := strings.Index(line, "(") + 1
			leftNodeEnd := strings.Index(line, ",")
			leftNodeName := line[leftNodeStart:leftNodeEnd]

			rightNodeStart := strings.Index(line, ", ") + 2
			rightNodeEnd := strings.Index(line, ")")
			rightNodeName := line[rightNodeStart:rightNodeEnd]
			nodes = append(nodes, &node{
				name: name,
				leftNode: &node{
					name:      leftNodeName,
					leftNode:  nil,
					rightNode: nil,
				},
				rightNode: &node{
					name:      rightNodeName,
					leftNode:  nil,
					rightNode: nil,
				},
			})
		}
		//fmt.Println(line)
		lineCount++
	}

	// now that the nodes are labelled with their names, do a second pass to link them together
	for _, n := range nodes {
		// find the left and right node link
		for _, n2 := range nodes {
			if n.leftNode.name == n2.name {
				n.leftNode = n2
			}
			if n.rightNode.name == n2.name {
				n.rightNode = n2
			}
		}
		//fmt.Println(n.name + ":" + n.leftNode.name + "," + n.rightNode.name)
	}

	findStepsRequired(instructions, nodes)

}

type node struct {
	previouslyVisited       bool
	prevVisitInstructionIdx int
	name                    string
	leftNode                *node
	rightNode               *node
}

func findStepsRequired(instructions []rune, nodes []*node) (int, int) {
	escaped := false
	var currentNode *node
	// only start when we find "AAA"
	for _, n := range nodes {
		if n.name == "AAA" {
			currentNode = n
		}
	}

	stepCount := 0
	instructionIterations := 0
	for !escaped {
		// continue following instructions until we hit our target node of ZZZ
		for idx := 0; idx < len(instructions); idx++ {

			if currentNode.name == "ZZZ" {
				escaped = true
				break
			}
			//idx, instructionIterations, stepCount = skipAhead(instructions, currentNode, idx, instructionIterations, stepCount)

			currentNode.previouslyVisited = true
			currentNode.prevVisitInstructionIdx = idx

			if instructions[idx] == 'L' {
				if currentNode.leftNode.name == "ZZZ" {
					escaped = true
					//break
				}
				currentNode = currentNode.leftNode
			} else if instructions[idx] == 'R' {
				if currentNode.rightNode.name == "ZZZ" {
					escaped = true
					//break
				}
				currentNode = currentNode.rightNode
			} else {
				panic("shouldn't be here")
			}
			stepCount++
			//fmt.Println(fmt.Sprintf("%d", stepCount) + " steps")
		}
		instructionIterations++
		fmt.Println(fmt.Sprintf("%d", instructionIterations) + " instructionIterations")

	}

	fmt.Println("Completed in " + fmt.Sprintf("%d", stepCount) + " steps")
	return stepCount, stepCount
}

func skipAhead(instructions []rune, currentNode *node, idx int, instructionIterations int, stepCount int) (int, int, int) {
	if instructionIterations == 0 {
		return idx, instructionIterations, stepCount
	}

	if currentNode.previouslyVisited && (currentNode.prevVisitInstructionIdx == idx) {
		// we've been here before, and it's heading the same direction, so we skip ahead
		fmt.Println("previously visited " + currentNode.name)

		println(idx)
		// calculate the steps we've taken in this iteration
		stepsTakenInThisInstructionIteration := idx + 1
		println(stepsTakenInThisInstructionIteration)
		// skip ahead to as many steps as it took us to get here before
		stepCount += stepsTakenInThisInstructionIteration
		// skip to the right idx of the loop
		remainder := stepCount % len(instructions)
		idx = remainder
		instructionIterations++

		println(instructionIterations)
	}
	return idx, instructionIterations, stepCount
}
