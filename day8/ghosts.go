package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func part2(filename string) {
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

	var nodesEndingInA []*node
	for _, n := range nodes {
		if strings.HasSuffix(n.name, "A") {
			nodesEndingInA = append(nodesEndingInA, n)
		}
	}

	var stepsRequired []int
	for _, n := range nodesEndingInA {
		stepsRequired = append(stepsRequired, findStepsRequiredForGhost(instructions, n, nodes))
	}
	fmt.Println(stepsRequired)

	fmt.Println(LCM(stepsRequired[0], stepsRequired[1], stepsRequired[2], stepsRequired[3], stepsRequired[4], stepsRequired[5]))
}

func findStepsRequiredForGhost(instructions []rune, startingNode *node, nodes []*node) int {
	escaped := false
	var currentNode *node
	currentNode = startingNode

	stepCount := 0
	for !escaped {
		// continue following instructions until we hit our target node of ZZZ
		for idx := 0; idx < len(instructions); idx++ {

			if strings.HasSuffix(currentNode.name, "Z") {
				escaped = true
				break
			}

			currentNode.previouslyVisited = true
			currentNode.prevVisitInstructionIdx = idx

			if instructions[idx] == 'L' {
				if strings.HasSuffix(currentNode.leftNode.name, "Z") {
					escaped = true
					//break
				}
				currentNode = currentNode.leftNode
			} else if instructions[idx] == 'R' {
				if strings.HasSuffix(currentNode.rightNode.name, "Z") {
					escaped = true
					//break
				}
				currentNode = currentNode.rightNode
			} else {
				panic("shouldn't be here")
			}
			stepCount++
		}
	}

	fmt.Println("Completed in " + fmt.Sprintf("%d", stepCount) + " steps")
	return stepCount
}
