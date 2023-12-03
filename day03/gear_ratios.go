package day03

import (
	"fmt"
	"unicode"
)

// GetSumOfPartNumbers returns the sum of all part numbers in the input.
func GetSumOfPartNumbers(input []string) int {
	sum := 0
	for i := 0; i < len(input); i++ {
		row := input[i]
		for j := 0; j < len(row); j++ {
			if unicode.IsDigit(rune(row[j])) {
				// parse the number & use the endIdx to find the length of parsed number
				val, endIdx := parseDigit(row, j)
				// find details about closest symbol to the number. We only need the symbol for this problem
				closestSymbol, _, _ := getClosestSymbolDetails(input, i, j, endIdx)
				if closestSymbol != '.' {
					sum += val
				}
				j = endIdx
			}
		}
	}
	return sum
}

// GetSumOfGearRatios returns the sum of all gear ratios i.e. 2 numbers surrounding a valid symbol in the input.
func GetSumOfGearRatios(input []string) int {
	symbolToNumMapping := make(map[string][]int)
	for i := 0; i < len(input); i++ {
		row := input[i]
		for j := 0; j < len(row); j++ {
			if unicode.IsDigit(rune(row[j])) {
				// parse the number & use the endIdx to find the length of parsed number
				val, endIdx := parseDigit(row, j)
				// find details about closest symbol to the number. We use symbol coordinates to form the key
				closestSymbol, rowId, colId := getClosestSymbolDetails(input, i, j, endIdx)
				if closestSymbol != '.' {
					key := fmt.Sprintf("%d|%d", rowId, colId)
					symbolToNumMapping[key] = append(symbolToNumMapping[key], val)
				}
				j = endIdx
			}
		}
	}
	sum := 0
	for _, v := range symbolToNumMapping {
		if len(v) == 2 {
			sum += v[0] * v[1]
		}
	}
	return sum
}

func parseDigit(row string, index int) (int, int) {
	val := 0
	for ; index < len(row); index++ {
		if !unicode.IsDigit(rune(row[index])) {
			break
		}
		val = val*10 + int(row[index]-'0')
	}
	return val, index
}

func getClosestSymbolDetails(input []string, row, col, endIdx int) (rune, int, int) {
	dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}, {-1, 1}, {1, 1}, {1, -1}, {-1, -1}}
	for i := col; i < endIdx; i++ {
		for _, dir := range dirs {
			newRow := row + dir[0]
			newCol := i + dir[1]
			if newRow < 0 || newRow >= len(input) || newCol < 0 || newCol >= len(input[0]) {
				continue
			}
			if !unicode.IsDigit(rune(input[newRow][newCol])) && input[newRow][newCol] != '.' {
				return rune(input[newRow][newCol]), newRow, newCol
			}
		}
	}
	return '.', -1, -1
}
