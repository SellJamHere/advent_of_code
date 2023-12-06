package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	races := parseInput(puzzleInput)

	errorMargin := 1
	for _, race := range races {
		winningCount := countWinningRoutes(race[0], race[1])
		errorMargin *= winningCount
	}

	fmt.Println("part 1 error margin:", errorMargin)

	time, distance := parseInputs2(puzzleInput)

	winningCount := countWinningRoutes(time, distance)

	fmt.Println("part 2 win count:", winningCount)
}

func race(waitTime, totalTime int) int {
	return waitTime * (totalTime - waitTime)
}

// binary search the fastest route
func countWinningRoutes(totalTime, winningDistance int) int {
	winningCount := 0
	for waitingTime := 0; waitingTime <= totalTime; waitingTime++ {
		distance := race(waitingTime, totalTime)
		if distance > winningDistance {
			winningCount++
		}
	}

	return winningCount
}

func parseInput(s string) [][]int {
	var races [][]int
	lines := strings.Split(s, "\n")
	rawTimes := strings.Split(strings.TrimSpace(strings.Split(lines[0], ":")[1]), " ")
	rawDistances := strings.Split(strings.TrimSpace(strings.Split(lines[1], ":")[1]), " ")

	var times []int
	var distances []int

	for _, time := range rawTimes {
		if time == "" {
			continue
		}

		times = append(times, mustAtoi(time))
	}

	for _, distance := range rawDistances {
		if distance == "" {
			continue
		}

		distances = append(distances, mustAtoi(distance))
	}

	for i := 0; i < len(times); i++ {
		races = append(races, []int{times[i], distances[i]})
	}

	return races
}

func parseInputs2(s string) (int, int) {
	lines := strings.Split(s, "\n")

	time := mustAtoi(strings.ReplaceAll(strings.TrimSpace(strings.Split(lines[0], ":")[1]), " ", ""))
	distance := mustAtoi(strings.ReplaceAll(strings.TrimSpace(strings.Split(lines[1], ":")[1]), " ", ""))

	return time, distance
}

func mustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return i
}

const samplePuzzle = `Time:      7  15   30
Distance:  9  40  200`

const puzzleInput = `Time:        48     93     85     95
Distance:   296   1928   1236   1391`
