package day09

import (
	"strconv"
	"strings"
)

// AnalyzeOasisReport calculates the sum of next values in the report by analyzing its history.
func AnalyzeOasisReport(input []string) int {
	sum := 0
	for _, line := range input {
		numbers := parseNumbers(line)
		sum += calculateNextValue(numbers)
	}
	return sum
}

// AnalyzeOasisReportWithExtrapolation calculates the sum of previous values in the report by analyzing its history.
func AnalyzeOasisReportWithExtrapolation(input []string) int {
	sum := 0
	for _, line := range input {
		numbers := parseNumbers(line)
		sum += calculatePrevValue(numbers)
	}
	return sum
}

func parseNumbers(line string) []int {
	splits := strings.Split(line, " ")
	numbers := make([]int, len(splits))
	for i, split := range splits {
		numbers[i], _ = strconv.Atoi(strings.TrimSpace(split))
	}
	return numbers
}

func calculateNextValue(numbers []int) int {
	history := buildHistory(numbers)
	diff := 0
	for i := len(history) - 1; i >= 0; i-- {
		curr := history[i]
		diff = curr[len(curr)-1] + diff
	}
	return diff
}

func calculatePrevValue(numbers []int) int {
	history := buildHistory(numbers)
	diff := 0
	for i := len(history) - 1; i >= 0; i-- {
		curr := history[i]
		diff = curr[0] - diff
	}
	return diff
}

func buildHistory(numbers []int) [][]int {
	history := make([][]int, 0)
	history = append(history, numbers)
	for {
		prev := history[len(history)-1]
		next := make([]int, len(prev)-1)
		allZero := true
		for i := 1; i < len(prev); i++ {
			next[i-1] = prev[i] - prev[i-1]
			if next[i-1] != 0 {
				allZero = false
			}
		}
		if allZero {
			break
		}
		history = append(history, next)
	}
	return history
}
