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
	filename := "day4/test.txt"

	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()
	//part1(file)
	part2(file)
}

func part1(file *os.File) {
	scanner := bufio.NewScanner(file)
	sum := 0
	gameNum := 0
	for scanner.Scan() {
		gameNum++
		line := scanner.Text()

		winStart := strings.Index(line, ":") + 2
		winEnd := strings.Index(line, "|") - 1

		winners := strings.Split(line[winStart:winEnd], " ")
		winnerSet := map[string]struct{}{}
		for _, winner := range winners {
			if winner == "" {
				continue
			}
			w := strings.TrimSpace(winner)
			winnerSet[w] = struct{}{}
		}

		cardStart := strings.Index(line, "|") + 2
		cardNumbers := strings.Split(line[cardStart:], " ")
		cardSet := map[string]struct{}{}
		for _, cardNumber := range cardNumbers {
			if cardNumber == "" {
				continue
			}
			c := strings.TrimSpace(cardNumber)
			cardSet[c] = struct{}{}
		}

		winnings := calcWinnings(winnerSet, cardSet)
		sum += winnings
		println("Game " + strconv.Itoa(gameNum) + " winnings: " + strconv.Itoa(winnings))
	}
	println(sum)
}

func calcWinnings(winningNumbers, cardNumbers map[string]struct{}) int {
	total := 0
	for cardNumber := range cardNumbers {
		if _, ok := winningNumbers[cardNumber]; ok {
			total++
		}
	}

	winnings := int(math.Pow(2, float64(total-1)))

	return winnings
}

func part2(file *os.File) {

	cardCounts := map[string]int{}
	scanner := bufio.NewScanner(file)
	totalScratchCards := 0
	gameNum := 0

	scratchcardSet := map[*scratchcard]struct{}{}
	for scanner.Scan() {
		gameNum++
		line := scanner.Text()

		winStart := strings.Index(line, ":") + 2
		winEnd := strings.Index(line, "|") - 1

		winners := strings.Split(line[winStart:winEnd], " ")
		winnerSet := map[string]struct{}{}
		for _, winner := range winners {
			if winner == "" {
				continue
			}
			w := strings.TrimSpace(winner)
			winnerSet[w] = struct{}{}
		}

		cardStart := strings.Index(line, "|") + 2
		cardNumbers := strings.Split(line[cardStart:], " ")
		cardSet := map[string]struct{}{}
		for _, cardNumber := range cardNumbers {
			if cardNumber == "" {
				continue
			}
			c := strings.TrimSpace(cardNumber)
			cardSet[c] = struct{}{}
		}
		scratchcardSet[&scratchcard{gameNumber: gameNum, winningNumbers: winnerSet, cardNumbers: cardSet}] = struct{}{}
	}

	for s := range scratchcardSet {
		wonCards := getScratchcardsForWin(*s)
		for _, wonCard := range wonCards {
			// we need to rerun the game for each card won
			newCards := getScratchcardsForWin(wonCard)
			cardCounts[strconv.Itoa(wonCard.gameNumber)]++
			println(newCards)
		}
		totalScratchCards += len(wonCards)
	}
	println("Total Cards Processed " + ": " + strconv.Itoa(totalScratchCards))
}

type scratchcard struct {
	gameNumber     int
	winningNumbers map[string]struct{}
	cardNumbers    map[string]struct{}
}

// getScratchcardsForWin returns a list of scratchcards that are the reward for matching numbers on an existing card.
func getScratchcardsForWin(card scratchcard) []scratchcard {
	total := 0
	var wonCards []scratchcard
	for cardNumber := range card.cardNumbers {
		if _, ok := card.winningNumbers[cardNumber]; ok {
			total++
		}
	}

	for i := card.gameNumber; i < card.gameNumber+total; i++ {
		wonCards = append(wonCards, scratchcard{gameNumber: i + 1})
	}

	return wonCards
}
