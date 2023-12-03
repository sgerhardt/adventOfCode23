package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	filename := "day2/games.txt"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	sum := 0
	powerSum := 0
	lineNumber := 0
	for scanner.Scan() {
		lineNumber++
		setOfBagGameColors := parseLine(scanner.Text())
		allPossible := allGamesPossible(setOfBagGameColors)
		if allPossible {
			sum += lineNumber
		}

		minSet := findMinSetToMakePossible(setOfBagGameColors)
		power := minSet.red * minSet.blue * minSet.green
		powerSum += power
	}

	println("sum of game ids that are possible: ", sum)
	println("sum of power of each game with min number of cubes to make game possible: ", powerSum)
}

func allGamesPossible(setOfBagGameColors []bagColors) bool {
	possible := true
	for _, colorsInBag := range setOfBagGameColors {
		if isGamePossible(colorsInBag, gameColors{12, 14, 13}) {
			continue
		} else {
			possible = false
		}
	}
	return possible
}

func parseLine(line string) []bagColors {
	gameNumAndColors := strings.Split(line, ":")
	unparsedColors := gameNumAndColors[1]
	games := strings.Split(unparsedColors, "; ")
	var bc []bagColors
	for _, game := range games {
		red, green, blue := 0, 0, 0
		colors := strings.Split(game, ",")
		// an example of the color format is 14 red, 6 blue, 2 green and the order can vary.
		for _, color := range colors {
			// split the color into the number and the color
			colorSplit := strings.Split(strings.TrimLeft(color, " "), " ")
			num, _ := strconv.Atoi(colorSplit[0])
			switch colorSplit[1] {
			case "red":
				red = num
			case "blue":
				blue = num
			case "green":
				green = num
			}
		}
		bc = append(bc, bagColors{red, blue, green})
	}

	return bc
}

type bagColors struct {
	red   int
	blue  int
	green int
}

type gameColors struct {
	red   int
	blue  int
	green int
}

type minSetToMakeGamePossible struct {
	red   int
	blue  int
	green int
}

func isGamePossible(bagColors bagColors, gameColors gameColors) bool {
	if bagColors.red <= gameColors.red && bagColors.blue <= gameColors.blue && bagColors.green <= gameColors.green {
		return true
	}
	return false
}

func findMinSetToMakePossible(bags []bagColors) bagColors {
	red := 0
	blue := 0
	green := 0
	// iterate through the bags and find the min of each color
	for _, bag := range bags {
		if bag.red > red {
			red = bag.red
		}
		if bag.blue > blue {
			blue = bag.blue
		}
		if bag.green > green {
			green = bag.green
		}
	}
	return bagColors{
		red:   red,
		blue:  blue,
		green: green,
	}
}
