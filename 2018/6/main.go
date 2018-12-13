package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(part1(p1))
	fmt.Println(part1(input))

	fmt.Println(part2(p1, 32))
	fmt.Println(part2(input, 10000))
}

type point struct {
	x int
	y int
}

func part1(input string) int {
	points := parseInput(input)
	largestDimensions := findLargestDimensions(points)

	// x, y, value is index of
	grid := createGrid(largestDimensions)

	fillGrid(points, grid)

	blobs := countAreaAndInfinity(grid)

	return findLargestArea(blobs)
}

func part2(input string, maxDistance int) int {
	points := parseInput(input)
	largestDimensions := findLargestDimensions(points)

	// x, y, value is index of
	grid := createGrid(largestDimensions)

	return countRegions(points, grid, maxDistance)
}

func parseInput(input string) []point {
	points := []point{}
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, ", ")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		points = append(points, point{x, y})
	}

	return points
}

func createGrid(size point) [][]int {
	grid := make([][]int, size.x+1)
	for i := range grid {
		grid[i] = make([]int, size.y+1)
	}

	return grid
}

func printGrid(grid [][]int) {
	for _, line := range grid {
		fmt.Println(line)
	}
}

func findLargestDimensions(points []point) point {
	x := 0
	y := 0
	for _, point := range points {
		if point.x > x {
			x = point.x
		}
		if point.y > y {
			y = point.y
		}
	}

	return point{x, y}
}

func fillGrid(points []point, grid [][]int) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			// map of distance to list of indicies
			distances := map[int][]int{}
			for k, point := range points {
				grid[point.x][point.y] = k
				distance := int(math.Abs(float64((i - point.x))) + math.Abs(float64((j - point.y))))
				distances[distance] = append(distances[distance], k)
			}

			// get shortest distance
			shortestDistance := math.MaxInt32
			for distance := range distances {
				if distance < shortestDistance && distance != 0 {
					shortestDistance = distance
				}
			}

			if len(distances[shortestDistance]) == 1 {
				grid[i][j] = distances[shortestDistance][0]
			} else {
				grid[i][j] = -9
			}
		}
	}
}

type blob struct {
	index int
	area  int
	edge  bool
}

func countAreaAndInfinity(grid [][]int) []blob {
	blobMap := map[int]blob{}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] < 0 {
				continue
			}

			var b blob
			var ok bool
			if b, ok = blobMap[grid[i][j]]; !ok {
				b = blob{
					index: grid[i][j],
				}
			}

			b.area++
			if !b.edge && (i == 0 || i == len(grid)-1 || j == 0 || j == len(grid[i])-1) {
				b.edge = true
			}

			blobMap[grid[i][j]] = b
		}
	}

	blobs := []blob{}
	for _, blob := range blobMap {
		blobs = append(blobs, blob)
	}

	return blobs
}

func findLargestArea(blobs []blob) int {
	largestBlobArea := 0
	for _, blob := range blobs {
		if !blob.edge && blob.area > largestBlobArea {
			largestBlobArea = blob.area
		}
	}

	return largestBlobArea
}

func countRegions(points []point, grid [][]int, maxDistance int) int {
	regionCount := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			// distance
			distanceSum := 0
			for _, point := range points {
				distanceSum += int(math.Abs(float64((i - point.x))) + math.Abs(float64((j - point.y))))
			}

			if distanceSum < maxDistance {
				regionCount++
			}
		}
	}

	return regionCount
}

var p1 = `1, 1
1, 6
8, 3
3, 4
5, 5
8, 9`

var input = `156, 193
81, 315
50, 197
84, 234
124, 162
339, 345
259, 146
240, 350
97, 310
202, 119
188, 331
199, 211
117, 348
350, 169
131, 355
71, 107
214, 232
312, 282
131, 108
224, 103
83, 122
352, 142
208, 203
319, 217
224, 207
327, 174
89, 332
254, 181
113, 117
120, 161
322, 43
115, 226
324, 222
151, 240
248, 184
207, 136
41, 169
63, 78
286, 43
84, 222
81, 167
128, 192
127, 346
213, 102
313, 319
207, 134
154, 253
50, 313
160, 330
332, 163`
