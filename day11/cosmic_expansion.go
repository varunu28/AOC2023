package day11

import (
	"slices"
	"strings"
)

type point struct {
	x int
	y int
}

func CalculateSumOfDistances(input []string) int {
	return calculateSumOfDistances(input, 2)
}

func CalculateSumOfDistancesForOlderGalaxies(input []string) int {
	return calculateSumOfDistances(input, 1000000)
}

func calculateSumOfDistances(input []string, expansionFactor int) int {
	grid := make([][]string, 0)
	for _, line := range input {
		grid = append(grid, strings.Split(line, ""))
	}

	emptyRows, emptyCols := findEmptyRowsAndCols(grid)
	points := findGalaxyPositions(grid)

	sum := 0
	for i := range points {
		pt := points[i]
		for j := i + 1; j < len(points); j++ {
			currPt := points[j]
			for r := min(pt.x, currPt.x); r < max(pt.x, currPt.x); r++ {
				if slices.Contains(emptyRows, r) {
					sum += expansionFactor
				} else {
					sum += 1
				}
			}
			for c := min(pt.y, currPt.y); c < max(pt.y, currPt.y); c++ {
				if slices.Contains(emptyCols, c) {
					sum += expansionFactor
				} else {
					sum += 1
				}
			}
		}
	}
	return sum
}

func findEmptyRowsAndCols(grid [][]string) ([]int, []int) {
	emptyRows := make([]int, 0)
	for i, row := range grid {
		empty := true
		for _, col := range row {
			if col != "." {
				empty = false
				break
			}
		}
		if empty {
			emptyRows = append(emptyRows, i)
		}
	}
	emptyCols := make([]int, 0)
	for i := range grid[0] {
		empty := true
		for _, row := range grid {
			if row[i] != "." {
				empty = false
				break
			}
		}
		if empty {
			emptyCols = append(emptyCols, i)
		}
	}
	return emptyRows, emptyCols
}

func findGalaxyPositions(grid [][]string) []point {
	points := make([]point, 0)
	for i, row := range grid {
		for j, col := range row {
			if col == "#" {
				points = append(points, point{i, j})
			}
		}
	}
	return points
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
