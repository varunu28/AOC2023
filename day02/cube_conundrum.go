package day02

import (
	"strconv"
	"strings"
)

type draw struct {
	color string
	count int
}

type game struct {
	id    int
	draws []draw
}

// GetValidGamesIdSum returns the sum of game ids which are valid based upon the rules mentioned by the elf
func GetValidGamesIdSum(input []string) int {
	sum := 0
	for _, inp := range input {
		currGame := parseGame(inp)
		valid := true
		for _, draw := range currGame.draws {
			if (draw.color == "red" && draw.count > 12) ||
				(draw.color == "green" && draw.count > 13) ||
				(draw.color == "blue" && draw.count > 14) {
				valid = false
				break
			}
		}
		if valid {
			sum += currGame.id
		}
	}
	return sum
}

// GetSumOfPowerOfGame returns the sum of powers of each game based upon the criteria described by the elf
func GetSumOfPowerOfGame(input []string) int {
	sum := 0
	for _, inp := range input {
		currGame := parseGame(inp)
		minValueMap := make(map[string]int)
		for _, draw := range currGame.draws {
			val, ok := minValueMap[draw.color]
			if !ok || (ok && val < draw.count) {
				minValueMap[draw.color] = draw.count
			}
		}
		pow := 1
		for _, v := range minValueMap {
			pow *= v
		}
		sum += pow
	}
	return sum
}

func parseGame(input string) game {
	splits := strings.Split(input, ":")
	gameId, _ := strconv.Atoi(strings.Split(splits[0], " ")[1])
	subsets := strings.Split(strings.Trim(splits[1], " "), ";")
	drawList := make([]draw, 0)
	for _, st := range subsets {
		draws := strings.Split(st, ",")
		for _, dw := range draws {
			split := strings.Split(strings.Trim(dw, " "), " ")
			count, _ := strconv.Atoi(split[0])
			color := split[1]
			drawList = append(drawList, draw{
				color: color,
				count: count,
			})
		}
	}
	return game{
		id:    gameId,
		draws: drawList,
	}
}
