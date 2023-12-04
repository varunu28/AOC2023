package day04

import (
	"strconv"
	"strings"
)

type Card struct {
	id           int
	yourCards    []string
	winningCards map[string]bool
}

func newCard(input string) *Card {
	splits := strings.Split(input, ":")
	fields := strings.Fields(splits[0])
	cardId, _ := strconv.Atoi(strings.TrimSpace(fields[1]))
	winCardSplit := strings.Split(strings.TrimSpace(splits[1]), "|")[0]
	yourCardSplit := strings.Split(strings.TrimSpace(splits[1]), "|")[1]

	winningCards := make(map[string]bool)
	for _, card := range strings.Fields(winCardSplit) {
		winningCards[strings.TrimSpace(card)] = true
	}

	yourCards := make([]string, 0)
	for _, card := range strings.Fields(yourCardSplit) {
		c := strings.TrimSpace(card)
		if c != "" {
			yourCards = append(yourCards, c)
		}
	}

	return &Card{
		id:           cardId,
		yourCards:    yourCards,
		winningCards: winningCards,
	}
}

func (card *Card) getScoreAndMatch() (int, int) {
	score := 0
	match := 0
	for _, c := range card.yourCards {
		if card.winningCards[c] {
			if score == 0 {
				score = 1
			} else {
				score *= 2
			}
			match++
		}
	}
	return score, match
}

// FindTotalPoints calculates total points for all cards that Elf has
func FindTotalPoints(input []string) int {
	total := 0
	for _, inp := range input {
		card := newCard(inp)
		score, _ := card.getScoreAndMatch()
		total += score
	}
	return total
}

// FindTotalCardsWithNewRule calculates total cards instances with the new matching rules
func FindTotalCardsWithNewRule(input []string) int {
	cardCounter := make(map[int]int)
	for _, inp := range input {
		card := newCard(inp)
		cardCounter[card.id] = 1
	}
	count := 0
	for _, inp := range input {
		card := newCard(inp)
		_, match := card.getScoreAndMatch()
		cardCount := cardCounter[card.id]
		count += cardCount
		for i := 1; i <= match; i++ {
			cardCounter[card.id+i] += cardCount
		}
	}
	return count
}
