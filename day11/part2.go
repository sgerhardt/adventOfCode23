package main

import (
	"fmt"
	"os"
	"strings"
)

type Galaxies struct {
	x []int
	y []int
}

func (g *Galaxies) append(x, y int) {
	g.x = append(g.x, x)
	g.y = append(g.y, y)
}

func makeRange[T any](length int, symbol T) []T {
	var r []T
	for i := 0; i < length; i++ {
		r = append(r, symbol)
	}
	return r
}

func replaceAtIndex[T any](input []T, replace T, index, offset int) []T {
	return append(input[:index], append([]T{replace}, input[index+offset:]...)...)
}

func expandUniverse(universe []string, galaxies Galaxies) ([]int, []int) {
	rangeX := makeRange(len(universe[0]), 1)
	rangeY := makeRange(len(universe), 1)
	for i := 0; i < len(galaxies.x); i++ {
		// binary list represents empty rows & columns
		rangeX = replaceAtIndex(rangeX, 0, galaxies.x[i], 1)
		rangeY = replaceAtIndex(rangeY, 0, galaxies.y[i], 1)
	}
	return rangeX, rangeY
}

func findAllGalaxies(universe []string) Galaxies {
	galaxies := Galaxies{}
	for row := 0; row < len(universe); row++ {
		for col := 0; col < len(universe[row]); col++ {
			if string(universe[row][col]) == "#" {
				galaxies.append(col, row)
			}
		}
	}
	return galaxies
}

func sumReduce(input []int) int {
	sum := 0
	for _, val := range input {
		sum += val
	}
	return sum
}

var M = 999_999

func getSumDistances(galaxies Galaxies, rangeX, rangeY []int) int {
	trailing := 0
	for i := 0; i < len(galaxies.y); i++ {
		for j := i; j < len(galaxies.x); j++ {
			var minX, maxX int
			if galaxies.x[i] < galaxies.x[j] {
				minX = galaxies.x[i]
				maxX = galaxies.x[j]
			} else {
				minX = galaxies.x[j]
				maxX = galaxies.x[i]
			}
			minY := galaxies.y[i]
			maxY := galaxies.y[j]
			componentX := maxX - minX + M*sumReduce(rangeX[minX:maxX])
			componentY := maxY - minY + M*sumReduce(rangeY[minY:maxY])
			trailing += componentX + componentY
		}
	}
	return trailing
}

func part2(filename string) int {
	data, _ := os.ReadFile(filename)
	universe := strings.Split(string(data), "\n")

	galaxies := findAllGalaxies(universe)
	rangeX, rangeY := expandUniverse(universe, galaxies)
	distances := getSumDistances(galaxies, rangeX, rangeY)
	fmt.Println(distances)
	return distances
}
