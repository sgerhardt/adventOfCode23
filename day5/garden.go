package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Hardcoded filename
	filename := "day5/garden.txt"

	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()
	garden(file)

}

type gardenMap struct {
	destinationRangeStart int
	sourceRangeStart      int
	rangeLength           int
}

func garden(file *os.File) {
	scanner := bufio.NewScanner(file)

	var seeds []int
	var seedToSoilMapping, soilToFertilizerMapping, fertilizerToWaterMapping, waterToLightMap, lightToTemperatureMap, temperatureToHumidityMap, humidityToLocationMap []gardenMap
	for scanner.Scan() {
		//line := scanner.Text()
		if seeds == nil {
			seeds = parseSeeds(scanner)
		}
		if seedToSoilMapping == nil {
			seedToSoilMapping = parseMap("seed-to-soil map:", scanner)
		}
		if soilToFertilizerMapping == nil {
			soilToFertilizerMapping = parseMap("soil-to-fertilizer map:", scanner)
		}
		if fertilizerToWaterMapping == nil {
			fertilizerToWaterMapping = parseMap("fertilizer-to-water map:", scanner)
		}
		if waterToLightMap == nil {
			waterToLightMap = parseMap("water-to-light map:", scanner)
		}
		if lightToTemperatureMap == nil {
			lightToTemperatureMap = parseMap("light-to-temperature map:", scanner)
		}
		if temperatureToHumidityMap == nil {
			temperatureToHumidityMap = parseMap("temperature-to-humidity map:", scanner)
		}
		if humidityToLocationMap == nil {
			humidityToLocationMap = parseMap("humidity-to-location map:", scanner)
		}
	}

	part1(seeds, seedToSoilMapping, soilToFertilizerMapping, fertilizerToWaterMapping, waterToLightMap, lightToTemperatureMap, temperatureToHumidityMap, humidityToLocationMap)
}

func part1(seeds []int, seedToSoilMapping []gardenMap, soilToFertilizerMapping []gardenMap, fertilizerToWaterMapping []gardenMap, waterToLightMap []gardenMap, lightToTemperatureMap []gardenMap, temperatureToHumidityMap []gardenMap, humidityToLocationMap []gardenMap) {
	lowest := math.MaxInt64
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
		if loc < lowest {
			lowest = loc
		}
		fmt.Printf("seed %d at location: %d\n", seed, loc)
	}
	fmt.Printf("lowest is seed: %d", lowest)
}

//func part2(seeds []int, seedToSoilMapping []gardenMap, soilToFertilizerMapping []gardenMap, fertilizerToWaterMapping []gardenMap, waterToLightMap []gardenMap, lightToTemperatureMap []gardenMap, temperatureToHumidityMap []gardenMap, humidityToLocationMap []gardenMap) {
//	lowest := math.MaxInt64
//	for _, seed := range seedRanges {
//		loc := findLowestLocationNumberForSeedRange(
//			seed,
//			seedToSoilMapping,
//			soilToFertilizerMapping,
//			fertilizerToWaterMapping,
//			waterToLightMap,
//			lightToTemperatureMap,
//			temperatureToHumidityMap,
//			humidityToLocationMap,
//		)
//		if loc < lowest {
//			lowest = loc
//		}
//		fmt.Printf("seed %d at location: %d\n", seed, loc)
//	}
//	fmt.Printf("lowest is seed: %d", lowest)
//}

func findLowestLocationNumberForSeedRange(
	seed int,
	seedToSoilMapping,
	soilToFertilizerMapping,
	fertilizerToWaterMapping,
	waterToLightMapping,
	lightToTemperatureMapping,
	temperatureToHumidityMapping,
	humidityToLocationMapping []gardenMap) int {

	soil := mapNumber(seed, seedToSoilMapping)
	fertilizer := mapNumber(soil, soilToFertilizerMapping)
	water := mapNumber(fertilizer, fertilizerToWaterMapping)
	light := mapNumber(water, waterToLightMapping)
	temp := mapNumber(light, lightToTemperatureMapping)
	humidity := mapNumber(temp, temperatureToHumidityMapping)
	location := mapNumber(humidity, humidityToLocationMapping)

	return location
}

func findLowestLocationNumberForSeed(
	seed int,
	seedToSoilMapping,
	soilToFertilizerMapping,
	fertilizerToWaterMapping,
	waterToLightMapping,
	lightToTemperatureMapping,
	temperatureToHumidityMapping,
	humidityToLocationMapping []gardenMap) int {

	soil := mapNumber(seed, seedToSoilMapping)
	fertilizer := mapNumber(soil, soilToFertilizerMapping)
	water := mapNumber(fertilizer, fertilizerToWaterMapping)
	light := mapNumber(water, waterToLightMapping)
	temp := mapNumber(light, lightToTemperatureMapping)
	humidity := mapNumber(temp, temperatureToHumidityMapping)
	location := mapNumber(humidity, humidityToLocationMapping)

	return location
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

func parseSeeds(scanner *bufio.Scanner) []int {
	var seeds []int
	line := scanner.Text()
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
	return seeds
}

func mapNumber(number int, maps []gardenMap) int {
	for _, m := range maps {
		if m.sourceRangeStart <= number && number < m.sourceRangeStart+m.rangeLength {
			return m.destinationRangeStart + (number - m.sourceRangeStart)
		}
	}
	return number
}
