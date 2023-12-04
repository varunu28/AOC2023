package day01

import (
	"unicode"
)

var digitMapping = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

// TrebuchetOnlyDigit finds the sum of calibration where calibration is represented in numeric form in the given input
func TrebuchetOnlyDigit(input []string) int {
	sum := 0
	for _, inp := range input {
		sum += getDigitBasedCalibration(inp)
	}
	return sum
}

// TrebuchetDigitAndLetter finds the sum of calibration where calibration is represented both in numeric & text form in the given input
func TrebuchetDigitAndLetter(input []string) int {
	sum := 0
	for _, inp := range input {
		sum += getDigitAndLetterBasedCalibration(inp)
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

func getDigitAndLetterBasedCalibration(input string) int {
	firstDigit, lastDigit := getFirstAndLastDigit(input)
	return firstDigit*10 + lastDigit
}

func getFirstAndLastDigit(input string) (int, int) {
	firstDigit := -1
	lastDigit := -1
	for i := range input {
		digit := -1
		if unicode.IsDigit(rune(input[i])) {
			digit = int(rune(input[i]) - '0')
		}
		if i+3 <= len(input) {
			val, ok := digitMapping[input[i:i+3]]
			if ok {
				digit = val
			}
		}
		if i+4 <= len(input) {
			val, ok := digitMapping[input[i:i+4]]
			if ok {
				digit = val
			}
		}
		if i+5 <= len(input) {
			val, ok := digitMapping[input[i:i+5]]
			if ok {
				digit = val
			}
		}
		if digit != -1 {
			if firstDigit == -1 {
				firstDigit = digit
			} else {
				lastDigit = digit
			}
		}
	}
	if lastDigit == -1 {
		lastDigit = firstDigit
	}
	return firstDigit, lastDigit
}
