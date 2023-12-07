package day06

import (
	"strconv"
	"strings"
)

// FindProductOfNumberOfWinningWays calculates product of number of ways to win the race for each boat
func FindProductOfNumberOfWinningWays(input []string) int {
	times := strings.Fields(strings.Split(input[0], ":")[1])
	recordDistances := strings.Fields(strings.Split(input[1], ":")[1])
	prod := 1
	for i := range recordDistances {
		time, _ := strconv.Atoi(times[i])
		distance, _ := strconv.Atoi(recordDistances[i])
		prod *= binarySearchWays(0, time, distance)
	}
	return prod
}

// FindNumberOfWinningWaysForSingleRace calculates number of ways to win the race by correctly formatting the input
func FindNumberOfWinningWaysForSingleRace(input []string) int {
	times := strings.Fields(strings.Split(input[0], ":")[1])
	recordDistances := strings.Fields(strings.Split(input[1], ":")[1])
	timeStr := ""
	distanceStr := ""
	for i := range recordDistances {
		timeStr = timeStr + times[i]
		distanceStr = distanceStr + recordDistances[i]
	}
	time, _ := strconv.Atoi(timeStr)
	distance, _ := strconv.Atoi(distanceStr)
	return binarySearchWays(0, time, distance)
}

func binarySearchWays(start, end, distance int) int {
	// Binary search for the minimum number of times a boat is pressed to win the race
	minTime := end
	startTime := start
	endTime := end
	for startTime <= endTime {
		mid := (startTime + endTime) / 2
		currDistance := (end - mid) * mid
		if currDistance > distance {
			endTime = mid - 1
			if mid < minTime {
				minTime = mid
			}
		} else {
			startTime = mid + 1
		}
	}
	// Binary search for the maximum number of times a boat is pressed to win the race
	maxTime := start
	startTime = start
	endTime = end
	for startTime <= endTime {
		mid := (startTime + endTime) / 2
		currDistance := (end - mid) * mid
		if currDistance > distance {
			startTime = mid + 1
			if mid > maxTime {
				maxTime = mid
			}
		} else {
			endTime = mid - 1
		}
	}
	return maxTime - minTime + 1
}
