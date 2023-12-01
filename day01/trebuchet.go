package day01

import (
	"unicode"
)

var threeLetterDigit = map[string]int{
	"one": 1,
	"two": 2,
	"six": 6,
}

var fourLetterDigit = map[string]int{
	"four": 4,
	"five": 5,
	"nine": 9,
}

var fiveLetterDigit = map[string]int{
	"three": 3,
	"seven": 7,
	"eight": 8,
}

func TrebuchetOnlyDigit(input []string) int {
	sum := 0
	for _, inp := range input {
		sum += getDigitBasedCalibration(inp)
	}
	return sum
}

func getDigitBasedCalibration(input string) int {
	firstDigit := -1
	lastDigit := -1
	for _, c := range input {
		if unicode.IsDigit(c) {
			firstDigit = int(c - '0')
			break
		}
	}
	for i := len(input) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(input[i])) {
			lastDigit = int(rune(input[i]) - '0')
			break
		}
	}
	return firstDigit*10 + lastDigit
}

func TrebuchetDigitAndLetter(input []string) int {
	sum := 0
	for _, inp := range input {
		sum += getDigitAndLetterBasedCalibration(inp)
	}
	return sum
}

func getDigitAndLetterBasedCalibration(input string) int {
	return getFirstDigitForCalibration(input)*10 + getLastDigitForCalibration(input)
}

func getFirstDigitForCalibration(input string) int {
	firstDigit := -1
	for i := range input {
		if unicode.IsDigit(rune(input[i])) {
			firstDigit = int(rune(input[i]) - '0')
			break
		}
		if i+3 < len(input) {
			val, ok := threeLetterDigit[input[i:i+3]]
			if ok {
				firstDigit = val
				break
			}
		}
		if i+4 < len(input) {
			val, ok := fourLetterDigit[input[i:i+4]]
			if ok {
				firstDigit = val
				break
			}
		}
		if i+5 < len(input) {
			val, ok := fiveLetterDigit[input[i:i+5]]
			if ok {
				firstDigit = val
				break
			}
		}
	}
	return firstDigit
}

func getLastDigitForCalibration(input string) int {
	lastDigit := -1
	for i := len(input) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(input[i])) {
			lastDigit = int(rune(input[i]) - '0')
			break
		}
		if i-3 >= 0 {
			val, ok := threeLetterDigit[input[i-2:i+1]]
			if ok {
				lastDigit = val
				break
			}
		}
		if i-4 >= 0 {
			val, ok := fourLetterDigit[input[i-3:i+1]]
			if ok {
				lastDigit = val
				break
			}
		}
		if i-5 >= 0 {
			val, ok := fiveLetterDigit[input[i-4:i+1]]
			if ok {
				lastDigit = val
				break
			}
		}
	}
	return lastDigit
}