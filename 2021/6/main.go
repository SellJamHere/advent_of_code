package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// Part 1
	initialFish := strings.Split(input, ",")
	fish := make([]int, len(initialFish))
	for i, fishStr := range initialFish {
		fishVal, _ := strconv.Atoi(fishStr)
		fish[i] = fishVal
	}

	for day := 0; day < 80; day++ {
		newFishCount := 0
		for i, val := range fish {
			if val == 0 {
				fish[i] = 6
				newFishCount++
				continue
			}

			fish[i] = val - 1
		}

		for i := 0; i < newFishCount; i++ {
			fish = append(fish, 8)
		}
	}

	fmt.Printf("fish count: %d\n", len(fish))

	// Part 2
	fishMap := map[int]int{}
	for _, fishStr := range initialFish {
		fishVal, _ := strconv.Atoi(fishStr)
		fishMap[fishVal] = fishMap[fishVal] + 1
	}

	for day := 0; day < 256; day++ {
		for i := 0; i <= 8; i++ {
			fishMap[i-1] = fishMap[i]
		}

		fishMap[8] = fishMap[-1]
		fishMap[6] = fishMap[6] + fishMap[-1]
		fishMap[-1] = 0
	}

	totalFish := 0
	for _, count := range fishMap {
		totalFish += count
	}

	fmt.Printf("fish count: %d\n", totalFish)
}

//const input = `3,4,3,1,2`

const input = `4,3,4,5,2,1,1,5,5,3,3,1,5,1,4,2,2,3,1,5,1,4,1,2,3,4,1,4,1,5,2,1,1,3,3,5,1,1,1,1,4,5,1,2,1,2,1,1,1,5,3,3,1,1,1,1,2,4,2,1,2,3,2,5,3,5,3,1,5,4,5,4,4,4,1,1,2,1,3,1,1,4,2,1,2,1,2,5,4,2,4,2,2,4,2,2,5,1,2,1,2,1,4,4,4,3,2,1,2,4,3,5,1,1,3,4,2,3,3,5,3,1,4,1,1,1,1,2,3,2,1,1,5,5,1,5,2,1,4,4,4,3,2,2,1,2,1,5,1,4,4,1,1,4,1,4,2,4,3,1,4,1,4,2,1,5,1,1,1,3,2,4,1,1,4,1,4,3,1,5,3,3,3,4,1,1,3,1,3,4,1,4,5,1,4,1,2,2,1,3,3,5,3,2,5,1,1,5,1,5,1,4,4,3,1,5,5,2,2,4,1,1,2,1,2,1,4,3,5,5,2,3,4,1,4,2,4,4,1,4,1,1,4,2,4,1,2,1,1,1,1,1,1,3,1,3,3,1,1,1,1,3,2,3,5,4,2,4,3,1,5,3,1,1,1,2,1,4,4,5,1,5,1,1,1,2,2,4,1,4,5,2,4,5,2,2,2,5,4,4`
