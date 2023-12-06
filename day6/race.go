package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	parse()
}

type raceRecord struct {
	timeAllotted   int
	recordDistance int
}

func parse() {

	file, err := os.Open("day6/part2.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var times []int
	var distances []int
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Time") {
			unparsedTimes := line[5:]
			strTimes := strings.Split(unparsedTimes, " ")
			for _, strTime := range strTimes {
				time, strConvErr := strconv.Atoi(strTime)
				if strConvErr != nil {
					continue
				}
				times = append(times, time)
			}
		}

		if strings.HasPrefix(line, "Distance") {
			unparsedTimes := line[9:]
			strDistances := strings.Split(unparsedTimes, " ")
			for _, strDistance := range strDistances {
				distance, strConvErr := strconv.Atoi(strDistance)
				if strConvErr != nil {
					continue
				}
				distances = append(distances, distance)
			}
		}
	}

	var raceRecords []raceRecord
	for idx, time := range times {
		raceRecords = append(raceRecords, raceRecord{
			timeAllotted:   time,
			recordDistance: distances[idx],
		})
	}

	product := 1
	waysToWin := 0
	for _, record := range raceRecords {
		waysToWin = countNumWaysToWin(record.timeAllotted, record.recordDistance)
		waysToWin *= product
		product = waysToWin
	}
	println(waysToWin)
}

func countNumWaysToWin(raceDuration, currentRecord int) int {
	waysToWin := 0
	for i := 0; i < raceDuration; i++ {
		speed := calculateSpeed(i)

		// calculate with time held increasing
		if distanceTraveled(speed, raceDuration-i) > currentRecord {
			waysToWin++
		}

	}
	return waysToWin
}

func calculateSpeed(timeHeld int) int {
	return timeHeld
}

func distanceTraveled(speed, duration int) int {
	return speed * duration
}
