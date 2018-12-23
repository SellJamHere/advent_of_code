package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func main() {
	// Part 1

	// Fuel cell at     3,5, grid serial number  8: power level 4.
	fmt.Println(calculatePowerLevel(point{3, 5}, 8))
	// Fuel cell at  122,79, grid serial number 57: power level -5.
	fmt.Println(calculatePowerLevel(point{122, 79}, 57))
	// Fuel cell at 217,196, grid serial number 39: power level  0.
	fmt.Println(calculatePowerLevel(point{217, 196}, 39))
	// Fuel cell at 101,153, grid serial number 71: power level  4.
	fmt.Println(calculatePowerLevel(point{101, 153}, 71))

	// For grid serial number 18: 33,45 (29)
	fmt.Println(findFuelRegionWithSize(nil, 18, point{3, 3}))
	// For grid serial number 42: 21,61 (30)
	fmt.Println(findFuelRegionWithSize(nil, 42, point{3, 3}))

	// Your puzzle input is 7165.
	fmt.Println(findFuelRegionWithSize(nil, 7165, point{3, 3}))

	// Part 2
	// For grid serial number 18, the largest total square (with a total power of 113) is 16x16 and has a top-left corner of 90,269, so its identifier is 90,269,16.
	fmt.Println(findFuelRegion(18))
	// For grid serial number 42, the largest total square (with a total power of 119) is 12x12 and has a top-left corner of 232,251, so its identifier is 232,251,12.
	fmt.Println(findFuelRegion(42))

	// Your puzzle input is 7165.
	start := time.Now()
	var coordinate point
	var size, power int
	coordinate, size, power = findFuelRegion(7165)
	fmt.Println("serial:", coordinate, size, power)
	fmt.Println(time.Since(start))

	start = time.Now()
	coordinate, size, power = findFuelRegionParallel(7165)
	fmt.Println("serial:", coordinate, size, power)
	fmt.Println(time.Since(start))
}

type point struct {
	x int
	y int
}

func calculatePowerLevel(coordinate point, serial int) int {
	rackID := coordinate.x + 10
	power := rackID * coordinate.y
	power += serial
	power *= rackID

	hundredsDigit := (power / 100) % 10

	return hundredsDigit - 5
}

func populateGrid(serial int) [][]int {
	grid := make([][]int, 300)
	for i := range grid {
		grid[i] = make([]int, 300)
		for j := range grid[i] {
			grid[i][j] = calculatePowerLevel(point{i + 1, j + 1}, serial)
		}
	}

	return grid
}

func calculateGridWindow(grid [][]int, start, size point) int {
	totalPower := 0
	for i := start.x; i < start.x+size.x; i++ {
		for j := start.y; j < start.y+size.y; j++ {
			totalPower += grid[i][j]
		}
	}

	return totalPower
}

func findOptimalGridRegion(grid [][]int, regionSize point) (point, int) {
	largestPower := -1 * math.MaxInt32
	largestPoint := point{}
	for i := 0; i < len(grid)-regionSize.x; i++ {
		for j := 0; j < len(grid[i])-regionSize.y; j++ {
			currentPoint := point{i, j}
			power := calculateGridWindow(grid, currentPoint, regionSize)
			if power > largestPower {
				largestPower = power
				largestPoint = currentPoint
			}
		}
	}

	return point{
		x: largestPoint.x + 1,
		y: largestPoint.y + 1,
	}, largestPower
}

func findFuelRegionWithSize(grid [][]int, serial int, size point) (point, int) {
	if grid == nil {
		grid = populateGrid(serial)
	}

	return findOptimalGridRegion(grid, size)
}

func findFuelRegion(serial int) (point, int, int) {
	grid := populateGrid(serial)

	largestPower := -1 * math.MaxInt32
	largestPoint := point{}
	largestSize := 0
	for i := 1; i <= 300; i++ {
		currentPoint, currentPower := findOptimalGridRegion(grid, point{i, i})
		if currentPower > largestPower {
			largestPower = currentPower
			largestPoint = currentPoint
			largestSize = i
		}
	}

	return largestPoint, largestSize, largestPower
}

type region struct {
	coordinate point
	size       point
	power      int
}

func regionWorker(grid [][]int, jobs <-chan region, results chan<- region) {
	for r := range jobs {
		currentPoint, currentPower := findOptimalGridRegion(grid, r.size)
		r.coordinate = currentPoint
		r.power = currentPower

		results <- r
	}
}

func findFuelRegionParallel(serial int) (point, int, int) {
	grid := populateGrid(serial)

	jobs := make(chan region)
	done := make(chan struct{})
	results := make(chan region)

	for w := 1; w < runtime.NumCPU()-1; w++ {
		go regionWorker(grid, jobs, results)
	}

	largestPower := -1 * math.MaxInt32
	largestPoint := point{}
	largestSize := 0
	go func() {
		for r := range results {
			if r.power > largestPower {
				largestPower = r.power
				largestPoint = r.coordinate
				largestSize = r.size.x
			}
		}

		done <- struct{}{}
	}()

	for i := 1; i <= 300; i++ {
		jobs <- region{
			size: point{i, i},
		}
	}
	close(jobs)

	return largestPoint, largestSize, largestPower
}
