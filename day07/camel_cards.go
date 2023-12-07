package day07

import (
	"cmp"
	"fmt"
	"sort"
)

const (
	FIVE_OF_A_KIND  = 7
	FOUR_OF_A_KIND  = 6
	FULL_HOUSE      = 5
	THREE_OF_A_KIND = 4
	TWO_PAIRS       = 3
	ONE_PAIR        = 2
	HIGH_CARD       = 1
)

type card struct {
	value string
	bid   int
}

func (c *card) cardType() int {
	frequency := make(map[rune]int)
	for _, r := range c.value {
		frequency[r]++
	}
	switch len(frequency) {
	case 1:
		return FIVE_OF_A_KIND
	case 2:
		for _, v := range frequency {
			if v == 4 {
				return FOUR_OF_A_KIND
			}
		}
		return FULL_HOUSE
	case 3:
		for _, v := range frequency {
			if v == 3 {
				return THREE_OF_A_KIND
			}
		}
		return TWO_PAIRS
	case 4:
		return ONE_PAIR
	default:
		return HIGH_CARD
	}
}

func (c *card) customCardType() int {
	frequency := make(map[rune]int)
	for _, r := range c.value {
		frequency[r]++
	}
	switch len(frequency) {
	case 1:
		return FIVE_OF_A_KIND
	case 2:
		_, ok := frequency['J']
		if ok {
			return FIVE_OF_A_KIND
		}
		for _, v := range frequency {
			if v == 4 {
				return FOUR_OF_A_KIND
			}
		}
		return FULL_HOUSE
	case 3:
		if frequency['J'] > 0 {
			if frequency['J'] == 2 || frequency['J'] == 3 {
				return FOUR_OF_A_KIND
			}
			if frequency['J'] == 1 {
				for _, v := range frequency {
					if v == 3 {
						return FOUR_OF_A_KIND
					}
					if v == 2 {
						return FULL_HOUSE
					}
				}
			}
		}
		for _, v := range frequency {
			if v == 3 {
				return THREE_OF_A_KIND
			}
		}
		return TWO_PAIRS
	case 4:
		_, ok := frequency['J']
		if ok {
			return THREE_OF_A_KIND
		}
		return ONE_PAIR
	default:
		if frequency['J'] == 1 {
			return ONE_PAIR
		}
		return HIGH_CARD
	}
}

// CalculateTotalWinnings calculates the total winnings for a set of cards against their respective bids
func CalculateTotalWinnings(input []string) int {
	cards := make([]card, len(input))
	for i, line := range input {
		var c card
		fmt.Sscanf(line, "%s %d", &c.value, &c.bid)
		cards[i] = c
	}

	cardLabelMapping := map[rune]int{
		'A': 13,
		'K': 12,
		'Q': 11,
		'J': 10,
		'T': 9,
		'9': 8,
		'8': 7,
		'7': 6,
		'6': 5,
		'5': 4,
		'4': 3,
		'3': 2,
		'2': 1,
	}
	cardCmp := func(a, b card) int {
		c := cmp.Compare(a.cardType(), b.cardType())
		if c != 0 {
			return c
		}
		for i := 0; i < len(a.value); i++ {
			c = cmp.Compare(cardLabelMapping[rune(a.value[i])], cardLabelMapping[rune(b.value[i])])
			if c != 0 {
				return c
			}
		}
		return 0
	}
	sort.Slice(cards, func(i, j int) bool {
		return cardCmp(cards[i], cards[j]) < 0
	})

	winnings := 0
	for i := len(cards) - 1; i >= 0; i-- {
		winnings += cards[i].bid * (i + 1)
	}
	return winnings
}

// CalculateTotalWinningsForModifiedDeck calculates the total winnings for a custom set of cards against their respective bids
func CalculateTotalWinningsForModifiedDeck(input []string) int {
	cards := make([]card, len(input))
	for i, line := range input {
		var c card
		fmt.Sscanf(line, "%s %d", &c.value, &c.bid)
		cards[i] = c
	}

	customCardLabelMapping := map[rune]int{
		'A': 13,
		'K': 12,
		'Q': 11,
		'J': 0,
		'T': 9,
		'9': 8,
		'8': 7,
		'7': 6,
		'6': 5,
		'5': 4,
		'4': 3,
		'3': 2,
		'2': 1,
	}
	cardCmp := func(a, b card) int {
		c := cmp.Compare(a.customCardType(), b.customCardType())
		if c != 0 {
			return c
		}
		for i := 0; i < len(a.value); i++ {
			c = cmp.Compare(customCardLabelMapping[rune(a.value[i])], customCardLabelMapping[rune(b.value[i])])
			if c != 0 {
				return c
			}
		}
		return 0
	}
	sort.Slice(cards, func(i, j int) bool {
		return cardCmp(cards[i], cards[j]) < 0
	})

	winnings := 0
	for i := len(cards) - 1; i >= 0; i-- {
		winnings += cards[i].bid * (i + 1)
	}
	return winnings
}
