package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Hardcoded filename
	filename := "day5/test.txt"

	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()
	part1(file)

}

type gardenMap struct {
	destinationRangeStart int
	sourceRangeStart      int
	rangeLength           int
}

func part1(file *os.File) {
	scanner := bufio.NewScanner(file)

	var seeds []int
	var seedToSoilMapping, soilToFertilizerMapping, fertilizerToWaterMapping, waterToLightMap, lightToTemperatureMap, temperatureToHumidityMap, humidityToLocationMap []gardenMap
	for scanner.Scan() {
		line := scanner.Text()
		parseSeeds(line, seeds)
		seedToSoilMapping = parseMap("seed-to-soil map:", scanner)
		soilToFertilizerMapping = parseMap("soil-to-fertilizer map:", scanner)
		fertilizerToWaterMapping = parseMap("fertilizer-to-water map:", scanner)
		waterToLightMap = parseMap("water-to-light map:", scanner)
		lightToTemperatureMap = parseMap("light-to-temperature map:", scanner)
		temperatureToHumidityMap = parseMap("temperature-to-humidity map:", scanner)
		humidityToLocationMap = parseMap("humidity-to-location map:", scanner)
	}

	for _, seed := range seeds {
		loc := findLowestLocationNumberForSeed(
			seed,
			seedToSoilMapping,
			soilToFertilizerMapping,
			fertilizerToWaterMapping,
			waterToLightMap,
			lightToTemperatureMap,
			temperatureToHumidityMap,
			humidityToLocationMap,
		)
		fmt.Printf("seed %d at location: %d\n", seed, loc)
	}

}

func findLowestLocationNumberForSeed(
	seed int,
	seedToSoilMapping,
	soilToFertilizerMapping,
	fertilizerToWaterMapping,
	waterToLightMap,
	lightToTemperatureMap,
	temperatureToHumidityMap,
	humidityToLocationMap []gardenMap) int {
	return 0
}

func parseMap(mapName string, scanner *bufio.Scanner) []gardenMap {
	var gardenMaps []gardenMap
	if strings.HasPrefix(scanner.Text(), mapName) {
		for scanner.Scan() && scanner.Text() != "" {
			line := scanner.Text()
			numStrs := strings.Split(line, " ")
			var numbers []int
			for _, num := range numStrs {
				n, err := strconv.Atoi(num)
				if err != nil {
					continue
				}
				numbers = append(numbers, n)
			}
			gardenMaps = append(gardenMaps, gardenMap{
				destinationRangeStart: numbers[0],
				sourceRangeStart:      numbers[1],
				rangeLength:           numbers[2],
			})
		}
	}
	return gardenMaps
}

func parseSeeds(line string, seeds []int) {
	if strings.Contains(line, "seeds:") {
		startIdx := strings.Index(line, ":") + 2
		seedStrs := strings.Split(line[startIdx:], " ")
		for _, seedStr := range seedStrs {
			s := strings.TrimSpace(seedStr)
			seedNum, err := strconv.Atoi(s)
			if err == nil {
				seeds = append(seeds, seedNum)
			}
		}
		println(seeds)
	}
}
