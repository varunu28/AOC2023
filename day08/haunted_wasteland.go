package day08

import (
	"strings"
)

type pair struct {
	left  string
	right string
}

// CalculateTotalSteps calculates the total number of steps required to cross the desert
func CalculateTotalSteps(input []string) int {
	directions, network := parseInput(input)
	steps := 0
	idx := 0
	curr := "AAA"
	for curr != "ZZZ" {
		if directions[idx] == 'R' {
			curr = network[curr].right
		} else {
			curr = network[curr].left
		}
		steps++
		idx = (idx + 1) % len(directions)
	}
	return steps
}

// CalculateTotalStepsForGhostMap calculates the total number of steps required for a ghost to cross the desert
func CalculateTotalStepsForGhostMap(input []string) int {
	directions, network := parseInput(input)
	path := make([]string, 0)
	for k := range network {
		if strings.HasSuffix(k, "A") {
			path = append(path, k)
		}
	}
	cycles := make([]int, 0)
	for _, position := range path {
		idx := 0
		stepCount := 0
		for {
			for stepCount == 0 || !strings.HasSuffix(position, "Z") {
				stepCount++
				if directions[idx] == 'R' {
					position = network[position].right
				} else {
					position = network[position].left
				}
				idx = (idx + 1) % len(directions)
			}
			cycles = append(cycles, stepCount)
			break
		}
	}
	steps := cycles[0]
	for i := 1; i < len(cycles); i++ {
		steps = findLcm(steps, cycles[i])
	}
	return steps
}

func findLcm(numOne, numTwo int) int {
	return (numOne * numTwo) / findGcd(numOne, numTwo)
}

func findGcd(numOne, numTwo int) int {
	if numOne == 0 {
		return numTwo
	}
	return findGcd(numTwo%numOne, numOne)
}

func parseInput(input []string) (string, map[string]pair) {
	directions := input[0]
	network := make(map[string]pair)
	for i := 2; i < len(input); i++ {
		splits := strings.Split(input[i], " = ")
		node := splits[0]
		brackStart := strings.Index(splits[1], "(")
		brackEnd := strings.Index(splits[1], ")")
		conns := splits[1][brackStart+1 : brackEnd]
		left := strings.TrimSpace(strings.Split(conns, ", ")[0])
		right := strings.TrimSpace(strings.Split(conns, ", ")[1])
		network[node] = pair{left, right}
	}
	return directions, network
}
