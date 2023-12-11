package day10

import "strings"

const (
	Up int = iota
	Right
	Down
	Left
)

// CountStepsForFarthestPoint returns the number of steps required to reach the farthest point in the maze
func CountStepsForFarthestPoint(input []string) int {
	graph, start := buildGraph(input)
	loop := recurse(graph, start)
	return len(loop) / 2
}

// CountTilesEnclosedByLoop returns the number of tiles enclosed by the loop
func CountTilesEnclosedByLoop(input []string) int {
	graph, start := buildGraph(input)
	loop := recurse(graph, start)
	count := 0
	for i := range graph {
		flag := false
		for j := range graph[i] {
			if !loop[[2]int{i, j}] {
				if flag {
					count++
				}
			} else if graph[i][j][0] {
				flag = !flag
			}
		}
	}
	return count
}

func buildGraph(rows []string) ([][][4]bool, [2]int) {
	var start [2]int
	graph := make([][][4]bool, len(rows))
	for i, row := range rows {
		graph[i] = make([][4]bool, len(row))
		for j, tile := range strings.Split(row, "") {
			if tile == "S" {
				start[0], start[1] = i, j
			}
			graph[i][j] = tileType(tile)
		}
	}
	return graph, start
}

func tileType(tile string) [4]bool {
	switch tile {
	case "|":
		return [4]bool{true, false, true, false}
	case "-":
		return [4]bool{false, true, false, true}
	case "L":
		return [4]bool{true, true, false, false}
	case "J":
		return [4]bool{true, false, false, true}
	case "7":
		return [4]bool{false, false, true, true}
	case "F":
		return [4]bool{false, true, true, false}
	default:
		return [4]bool{false, false, false, false}
	}
}

func recurse(graph [][][4]bool, start [2]int) map[[2]int]bool {
	loop := make(map[[2]int]bool)
	for _, tile := range []string{"J", "|", "-", "L", "7", "F"} {
		graph[start[0]][start[1]] = tileType(tile)
		if l, found := findLoop(graph, start); found {
			loop = l
			break
		}
	}
	return loop
}

func findLoop(graph [][][4]bool, start [2]int) (map[[2]int]bool, bool) {
	row, col := start[0], start[1]
	var dir int
	for i, val := range graph[row][col] {
		if val {
			dir = i
			break
		}
	}
	visited := make(map[[2]int]bool)
	for {
		if _, exists := visited[[2]int{row, col}]; exists {
			return visited, true
		}
		visited[[2]int{row, col}] = true
		from := 0
		switch dir {
		case Up:
			row--
			from = Down
		case Right:
			col++
			from = Left
		case Down:
			row++
			from = Up
		case Left:
			col--
			from = Right
		default:
			panic("Invalid direction")
		}
		if !graph[row][col][from] {
			return nil, false
		}
		for i := 0; i < 4; i++ {
			if i != from && graph[row][col][i] {
				dir = i
				break
			}
		}
	}
}
