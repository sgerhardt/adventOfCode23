package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Hardcoded filename
	filename := "day7/cards.txt"

	part1(filename)
	part2(filename)
}

func part1(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var hands []hand
	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.Split(line, " ")
		rawHand, rawBet := lineSplit[0], lineSplit[1]
		bet, _ := strconv.Atoi(rawBet)
		cards := parseHand(rawHand)
		h := hand{
			cards: cards,
			bid:   bet,
		}
		hands = append(hands, h)
	}
	fmt.Println(calculateWinnings(orderHands(hands)))
}

func parseHand(s string) [5]int {
	cards := [5]int{}
	for idx, c := range s {
		switch c {
		case 'A':
			cards[idx] = ace
		case 'K':
			cards[idx] = king
		case 'Q':
			cards[idx] = queen
		case 'J':
			cards[idx] = jack
		case 'T':
			cards[idx] = ten
		case '9':
			cards[idx] = nine
		case '8':
			cards[idx] = eight
		case '7':
			cards[idx] = seven
		case '6':
			cards[idx] = six
		case '5':
			cards[idx] = five
		case '4':
			cards[idx] = four
		case '3':
			cards[idx] = three
		case '2':
			cards[idx] = two
		}

	}
	return cards
}

var (
	// these are vars as J can change value ordering in part 2
	two   = 0
	three = 1
	four  = 2
	five  = 3
	six   = 4
	seven = 5
	eight = 6
	nine  = 7
	ten   = 8
	jack  = 9
	queen = 10
	king  = 11
	ace   = 12
)

const (
	highCard int = iota
	onePair
	twoPair
	threeOfAKind
	fullHouse
	fourOfAKind
	fiveOfAKind
)

type hand struct {
	cards [5]int
	bid   int
}

func (h *hand) calcStrength() int {
	cardToOccurrence := map[int]int{}
	for _, card := range h.cards {
		cardToOccurrence[card]++
	}

	hasThreeOfAKind, hasOnePair := false, false
	for _, occurrences := range cardToOccurrence {
		if occurrences == 5 {
			return fiveOfAKind
		}
		if occurrences == 4 {
			return fourOfAKind
		}
		if occurrences == 3 {
			hasThreeOfAKind = true
			if hasOnePair {
				return fullHouse
			}
		}
		if occurrences == 2 && hasOnePair {
			return twoPair
		}
		if occurrences == 2 {
			hasOnePair = true
		}
	}

	if hasThreeOfAKind && hasOnePair {
		return fullHouse
	}

	if hasThreeOfAKind {
		return threeOfAKind
	}
	if hasOnePair {
		return onePair
	}

	return highCard
}

func orderHands(hands []hand) []hand {
	orderHandsByType(hands)

	// for each type, order those hands
	return orderHandsByCardForType(hands)
}

func orderHandsByType(hands []hand) []hand {
	// first get hands in order of base strength
	sort.Sort(ByType(hands))
	return hands
}

func orderHandsByCardForType(hands []hand) []hand {
	var highCardHands []hand
	var onePairHands []hand
	var twoPairHands []hand
	var threeOfAKindHands []hand
	var fullHouseHands []hand
	var fourOfAKindHands []hand
	var fiveOfAKindHands []hand

	for _, h := range hands {
		// set hands for each type
		switch h.calcStrength() {
		case highCard:
			highCardHands = append(highCardHands, h)
		case onePair:
			onePairHands = append(onePairHands, h)
		case twoPair:
			twoPairHands = append(twoPairHands, h)
		case threeOfAKind:
			threeOfAKindHands = append(threeOfAKindHands, h)
		case fullHouse:
			fullHouseHands = append(fullHouseHands, h)
		case fourOfAKind:
			fourOfAKindHands = append(fourOfAKindHands, h)
		case fiveOfAKind:
			fiveOfAKindHands = append(fiveOfAKindHands, h)
		}
	}

	sort.Sort(ByCard(highCardHands))
	sort.Sort(ByCard(onePairHands))
	sort.Sort(ByCard(twoPairHands))
	sort.Sort(ByCard(threeOfAKindHands))
	sort.Sort(ByCard(fullHouseHands))
	sort.Sort(ByCard(fourOfAKindHands))
	sort.Sort(ByCard(fiveOfAKindHands))
	sortedCards := append(highCardHands, onePairHands...)
	sortedCards = append(sortedCards, twoPairHands...)
	sortedCards = append(sortedCards, threeOfAKindHands...)
	sortedCards = append(sortedCards, fullHouseHands...)
	sortedCards = append(sortedCards, fourOfAKindHands...)
	sortedCards = append(sortedCards, fiveOfAKindHands...)

	return sortedCards
}

type ByCard []hand

func (b ByCard) Len() int { return len(b) }

func (b ByCard) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b ByCard) Less(i, j int) bool {
	for k := 0; k < 5; k++ {
		if b[i].cards[k] == b[j].cards[k] {
			continue
		} else {
			return b[i].cards[k] < b[j].cards[k]
		}
	}
	return false
}

type ByType []hand

func (b ByType) Len() int { return len(b) }

func (b ByType) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b ByType) Less(i, j int) bool {
	return b[i].calcStrength() < b[j].calcStrength()
}

func calculateWinnings(orderedHands []hand) int {
	sum := 0
	for idx, h := range orderedHands {
		sum = sum + (idx+1)*h.bid
	}
	return sum
}
