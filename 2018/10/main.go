package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	trackLights(parseInput(p1))

	trackLights(parseInput(input))
}

type light struct {
	i          int
	coordinate pair
	velocity   pair
}

type pair struct {
	x int
	y int
}

func parseInput(input string) []light {
	lights := []light{}
	for i, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, "> ")

		// position
		positionParts := strings.Split(parts[0], "<")
		positionParts = strings.Split(positionParts[1], ">")
		positionParts = strings.Split(positionParts[0], ",")

		positionX, _ := strconv.Atoi(strings.TrimSpace(positionParts[0]))
		positionY, _ := strconv.Atoi(strings.TrimSpace(positionParts[1]))

		// velocity
		velocityParts := strings.Split(parts[1], "<")
		velocityParts = strings.Split(velocityParts[1], ">")
		velocityParts = strings.Split(velocityParts[0], ",")

		velocityX, _ := strconv.Atoi(strings.TrimSpace(parts[1][10:12]))
		velocityY, _ := strconv.Atoi(strings.TrimSpace(parts[1][14:16]))

		lights = append(lights, light{
			i: i,
			coordinate: pair{
				x: positionX,
				y: positionY,
			},
			velocity: pair{
				x: velocityX,
				y: velocityY,
			},
		})
	}

	return lights
}

func findCoordinateEdges(lights []light) (int, int, int, int) {
	lowestX := math.MaxInt32
	largestX := -1 * math.MaxInt32
	lowestY := math.MaxInt32
	largestY := -1 * math.MaxInt32

	for _, light := range lights {
		if light.coordinate.x < lowestX {
			lowestX = light.coordinate.x
		}
		if light.coordinate.x > largestX {
			largestX = light.coordinate.x
		}
		if light.coordinate.y < lowestY {
			lowestY = light.coordinate.y
		}
		if light.coordinate.y > largestY {
			largestY = light.coordinate.y
		}
	}

	return lowestX, largestX, lowestY, largestY
}

func printLights(lights []light) (int, int) {
	lowestX, largestX, lowestY, largestY := findCoordinateEdges(lights)

	xDelta := largestX - lowestX + 1
	yDelta := largestY - lowestY + 1

	sky := make([][]string, yDelta)
	for i := 0; i < yDelta; i++ {
		sky[i] = make([]string, xDelta)
		for j := 0; j < xDelta; j++ {
			sky[i][j] = "."
		}
	}

	for _, light := range lights {
		x := light.coordinate.x - lowestX
		y := light.coordinate.y - lowestY
		sky[y][x] = "#"
	}

	printSky(sky)

	return len(sky), len(sky[0])
}

func printSky(sky [][]string) {
	for i := 0; i < len(sky); i++ {
		line := ""
		for j := 0; j < len(sky[i]); j++ {
			line += sky[i][j]
		}

		fmt.Println(line)
	}
}

func moveLights(lights []light) {
	for i, light := range lights {
		light.coordinate.x += light.velocity.x
		light.coordinate.y += light.velocity.y

		lights[i] = light
	}
}

func moveLightsBackwards(lights []light) {
	for i, light := range lights {
		light.coordinate.x -= light.velocity.x
		light.coordinate.y -= light.velocity.y

		lights[i] = light
	}
}

func trackLights(lights []light) {
	smallestSky := math.MaxInt64
	i := 0
	for {
		lowestX, largestX, lowestY, largestY := findCoordinateEdges(lights)
		xDelta := largestX - lowestX + 1
		yDelta := largestY - lowestY + 1

		area := xDelta * yDelta

		if area < smallestSky {
			smallestSky = area
		} else {
			// reverse one step
			fmt.Println("seconds:", i-1)
			moveLightsBackwards(lights)
			printLights(lights)
			break
		}

		moveLights(lights)
		i++
	}
}

var p1 = `position=< 9,  1> velocity=< 0,  2>
position=< 7,  0> velocity=<-1,  0>
position=< 3, -2> velocity=<-1,  1>
position=< 6, 10> velocity=<-2, -1>
position=< 2, -4> velocity=< 2,  2>
position=<-6, 10> velocity=< 2, -2>
position=< 1,  8> velocity=< 1, -1>
position=< 1,  7> velocity=< 1,  0>
position=<-3, 11> velocity=< 1, -2>
position=< 7,  6> velocity=<-1, -1>
position=<-2,  3> velocity=< 1,  0>
position=<-4,  3> velocity=< 2,  0>
position=<10, -3> velocity=<-1,  1>
position=< 5, 11> velocity=< 1, -2>
position=< 4,  7> velocity=< 0, -1>
position=< 8, -2> velocity=< 0,  1>
position=<15,  0> velocity=<-2,  0>
position=< 1,  6> velocity=< 1,  0>
position=< 8,  9> velocity=< 0, -1>
position=< 3,  3> velocity=<-1,  1>
position=< 0,  5> velocity=< 0, -1>
position=<-2,  2> velocity=< 2,  0>
position=< 5, -2> velocity=< 1,  2>
position=< 1,  4> velocity=< 2,  1>
position=<-2,  7> velocity=< 2, -2>
position=< 3,  6> velocity=<-1, -1>
position=< 5,  0> velocity=< 1,  0>
position=<-6,  0> velocity=< 2,  0>
position=< 5,  9> velocity=< 1, -2>
position=<14,  7> velocity=<-2,  0>
position=<-3,  6> velocity=< 2, -1>`

var input = `position=<-42417,  32097> velocity=< 4, -3>
position=<-10502, -10533> velocity=< 1,  1>
position=<-53094,  32093> velocity=< 5, -3>
position=<-53090, -21188> velocity=< 5,  2>
position=< 53486,  21441> velocity=<-5, -2>
position=<-21142, -42496> velocity=< 2,  4>
position=<-42422,  32088> velocity=< 4, -3>
position=< 42778,  10784> velocity=<-4, -1>
position=< 10826,  42748> velocity=<-1, -4>
position=<-10449,  53401> velocity=< 1, -5>
position=<-42453, -21187> velocity=< 4,  2>
position=< 32154, -31839> velocity=<-3,  3>
position=<-42434, -53156> velocity=< 4,  5>
position=<-21149,  32097> velocity=< 2, -3>
position=<-10497,  53409> velocity=< 1, -5>
position=<-42438, -53151> velocity=< 4,  5>
position=< 21490,  10778> velocity=<-2, -1>
position=< 42831,  42753> velocity=<-4, -4>
position=< 53474, -42501> velocity=<-5,  4>
position=<-42436, -21192> velocity=< 4,  2>
position=< 53450, -53154> velocity=<-5,  5>
position=< 32118, -21185> velocity=<-3,  2>
position=<-21158, -10533> velocity=< 2,  1>
position=< 10842,  42749> velocity=<-1, -4>
position=< 53467,  21436> velocity=<-5, -2>
position=< 32130, -53160> velocity=<-3,  5>
position=< 32119, -53151> velocity=<-3,  5>
position=< 32143, -21183> velocity=<-3,  2>
position=<-10486,  10781> velocity=< 1, -1>
position=<-53098,  21433> velocity=< 5, -2>
position=<-10465, -42503> velocity=< 1,  4>
position=<-53083,  10776> velocity=< 5, -1>
position=< 32173, -53151> velocity=<-3,  5>
position=<-31782,  21436> velocity=< 3, -2>
position=< 32143,  10782> velocity=<-3, -1>
position=< 21487, -42500> velocity=<-2,  4>
position=<-53082,  42752> velocity=< 5, -4>
position=< 42775, -21183> velocity=<-4,  2>
position=< 53431, -53160> velocity=<-5,  5>
position=<-53078,  10784> velocity=< 5, -1>
position=<-10505,  53408> velocity=< 1, -5>
position=< 21507, -53151> velocity=<-2,  5>
position=< 42814, -53160> velocity=<-4,  5>
position=< 53427,  53402> velocity=<-5, -5>
position=<-21164,  10780> velocity=< 2, -1>
position=< 21461,  53404> velocity=<-2, -5>
position=< 32173,  21436> velocity=<-3, -2>
position=<-53086,  21438> velocity=< 5, -2>
position=< 32146, -42503> velocity=<-3,  4>
position=< 32159,  21439> velocity=<-3, -2>
position=< 53442,  42751> velocity=<-5, -4>
position=<-31817,  42752> velocity=< 3, -4>
position=< 53459, -21188> velocity=<-5,  2>
position=< 21458, -21191> velocity=<-2,  2>
position=<-21134,  53402> velocity=< 2, -5>
position=<-10481, -10534> velocity=< 1,  1>
position=<-21137, -53158> velocity=< 2,  5>
position=<-21110, -31839> velocity=< 2,  3>
position=<-21153,  21441> velocity=< 2, -2>
position=<-31766, -42495> velocity=< 3,  4>
position=< 42799, -53155> velocity=<-4,  5>
position=<-42449,  42752> velocity=< 4, -4>
position=<-53102, -10531> velocity=< 5,  1>
position=<-21149, -10527> velocity=< 2,  1>
position=<-31801,  53401> velocity=< 3, -5>
position=<-10462, -31842> velocity=< 1,  3>
position=<-31806,  42751> velocity=< 3, -4>
position=<-10502,  10777> velocity=< 1, -1>
position=<-53129,  10785> velocity=< 5, -1>
position=< 10803, -21186> velocity=<-1,  2>
position=< 42810, -31847> velocity=<-4,  3>
position=< 53450, -21189> velocity=<-5,  2>
position=< 53434,  21433> velocity=<-5, -2>
position=<-21126,  32094> velocity=< 2, -3>
position=< 53446, -53157> velocity=<-5,  5>
position=< 21518, -31839> velocity=<-2,  3>
position=<-21146, -53157> velocity=< 2,  5>
position=< 32135,  32089> velocity=<-3, -3>
position=<-53094, -21188> velocity=< 5,  2>
position=< 32170, -53152> velocity=<-3,  5>
position=<-42476,  21437> velocity=< 4, -2>
position=<-42446, -31840> velocity=< 4,  3>
position=< 53471, -42503> velocity=<-5,  4>
position=< 10850,  42752> velocity=<-1, -4>
position=<-21141, -10531> velocity=< 2,  1>
position=<-10508,  53404> velocity=< 1, -5>
position=< 42818, -21190> velocity=<-4,  2>
position=< 21483, -42503> velocity=<-2,  4>
position=<-10482, -42503> velocity=< 1,  4>
position=<-53089, -31847> velocity=< 5,  3>
position=< 32162,  32092> velocity=<-3, -3>
position=<-31763,  42753> velocity=< 3, -4>
position=<-10462, -10535> velocity=< 1,  1>
position=< 42831,  21433> velocity=<-4, -2>
position=< 42819, -42504> velocity=<-4,  4>
position=< 53426, -21192> velocity=<-5,  2>
position=< 21475, -42495> velocity=<-2,  4>
position=<-10462,  32092> velocity=< 1, -3>
position=< 42802, -10533> velocity=<-4,  1>
position=< 32146,  32088> velocity=<-3, -3>
position=< 10831, -21184> velocity=<-1,  2>
position=< 53466,  42752> velocity=<-5, -4>
position=< 21516, -42499> velocity=<-2,  4>
position=< 32159, -53152> velocity=<-3,  5>
position=<-53123,  32097> velocity=< 5, -3>
position=<-53126,  21434> velocity=< 5, -2>
position=<-31790, -21189> velocity=< 3,  2>
position=<-31789, -31843> velocity=< 3,  3>
position=<-10502, -21185> velocity=< 1,  2>
position=<-53124,  21441> velocity=< 5, -2>
position=< 53486,  10779> velocity=<-5, -1>
position=< 42805, -31846> velocity=<-4,  3>
position=<-42470,  53406> velocity=< 4, -5>
position=< 53469, -42495> velocity=<-5,  4>
position=< 53450,  42751> velocity=<-5, -4>
position=< 42822, -31843> velocity=<-4,  3>
position=< 32133,  21436> velocity=<-3, -2>
position=< 21487,  21437> velocity=<-2, -2>
position=< 32149, -31841> velocity=<-3,  3>
position=<-31778, -53151> velocity=< 3,  5>
position=< 32175, -42504> velocity=<-3,  4>
position=< 10834,  42750> velocity=<-1, -4>
position=< 10855,  10777> velocity=<-1, -1>
position=< 21492,  21438> velocity=<-2, -2>
position=<-21131,  21434> velocity=< 2, -2>
position=< 53455,  32091> velocity=<-5, -3>
position=<-31786,  53408> velocity=< 3, -5>
position=< 10829,  10781> velocity=<-1, -1>
position=< 10831, -10529> velocity=<-1,  1>
position=<-42460,  53405> velocity=< 4, -5>
position=< 10813, -31839> velocity=<-1,  3>
position=<-53076,  21432> velocity=< 5, -2>
position=<-10452, -42499> velocity=< 1,  4>
position=< 21487, -10534> velocity=<-2,  1>
position=<-53082,  10781> velocity=< 5, -1>
position=< 10839, -10527> velocity=<-1,  1>
position=< 42770,  21441> velocity=<-4, -2>
position=< 53466, -21185> velocity=<-5,  2>
position=< 10812,  42753> velocity=<-1, -4>
position=< 32156,  21432> velocity=<-3, -2>
position=< 10803,  53407> velocity=<-1, -5>
position=< 53430,  42747> velocity=<-5, -4>
position=<-53114,  32088> velocity=< 5, -3>
position=< 32156,  42748> velocity=<-3, -4>
position=< 42802, -42499> velocity=<-4,  4>
position=<-53110, -42500> velocity=< 5,  4>
position=< 21474, -31840> velocity=<-2,  3>
position=<-21166,  21432> velocity=< 2, -2>
position=<-10506,  42746> velocity=< 1, -4>
position=<-31819, -42499> velocity=< 3,  4>
position=<-31769, -10528> velocity=< 3,  1>
position=< 53426,  32088> velocity=<-5, -3>
position=<-21150, -10529> velocity=< 2,  1>
position=<-53086,  53408> velocity=< 5, -5>
position=<-31782,  32095> velocity=< 3, -3>
position=<-42449,  42747> velocity=< 4, -4>
position=<-53098, -31847> velocity=< 5,  3>
position=< 32162, -21185> velocity=<-3,  2>
position=<-53110, -42497> velocity=< 5,  4>
position=< 53455, -10532> velocity=<-5,  1>
position=< 32157, -42504> velocity=<-3,  4>
position=< 42802, -53155> velocity=<-4,  5>
position=<-53106,  42749> velocity=< 5, -4>
position=<-31781, -31848> velocity=< 3,  3>
position=<-53100,  53403> velocity=< 5, -5>
position=<-10508, -21187> velocity=< 1,  2>
position=<-42477,  53407> velocity=< 4, -5>
position=<-31795, -42504> velocity=< 3,  4>
position=<-53077, -10527> velocity=< 5,  1>
position=< 32133,  32097> velocity=<-3, -3>
position=<-53090,  53409> velocity=< 5, -5>
position=< 10847,  42749> velocity=<-1, -4>
position=< 21466, -21192> velocity=<-2,  2>
position=< 21476, -21192> velocity=<-2,  2>
position=<-10462,  21439> velocity=< 1, -2>
position=< 53430, -53153> velocity=<-5,  5>
position=< 42831, -10534> velocity=<-4,  1>
position=<-42435,  21432> velocity=< 4, -2>
position=< 32131,  53400> velocity=<-3, -5>
position=<-10481, -31844> velocity=< 1,  3>
position=<-10478,  53404> velocity=< 1, -5>
position=<-42438, -53160> velocity=< 4,  5>
position=< 32170, -53153> velocity=<-3,  5>
position=< 42778,  32092> velocity=<-4, -3>
position=<-42427,  42753> velocity=< 4, -4>
position=< 21476, -21183> velocity=<-2,  2>
position=<-21126, -21191> velocity=< 2,  2>
position=<-21118, -31842> velocity=< 2,  3>
position=<-10505, -31848> velocity=< 1,  3>
position=<-42421, -10536> velocity=< 4,  1>
position=<-31779,  21436> velocity=< 3, -2>
position=<-53106,  21437> velocity=< 5, -2>
position=< 42807, -53151> velocity=<-4,  5>
position=< 10810,  10778> velocity=<-1, -1>
position=< 32154, -10534> velocity=<-3,  1>
position=< 42814, -53160> velocity=<-4,  5>
position=< 53430,  21439> velocity=<-5, -2>
position=<-42462,  32097> velocity=< 4, -3>
position=< 21506,  10777> velocity=<-2, -1>
position=<-53092, -42495> velocity=< 5,  4>
position=<-53081, -10528> velocity=< 5,  1>
position=< 53471, -53153> velocity=<-5,  5>
position=<-53106, -21187> velocity=< 5,  2>
position=< 53450,  21437> velocity=<-5, -2>
position=< 21515, -10527> velocity=<-2,  1>
position=< 42791, -31846> velocity=<-4,  3>
position=< 21463,  21433> velocity=<-2, -2>
position=< 53485,  10776> velocity=<-5, -1>
position=< 42790, -31848> velocity=<-4,  3>
position=<-21122, -21192> velocity=< 2,  2>
position=< 21511,  53406> velocity=<-2, -5>
position=<-10505, -21192> velocity=< 1,  2>
position=<-21163,  21436> velocity=< 2, -2>
position=< 53426, -42496> velocity=<-5,  4>
position=< 32162,  10781> velocity=<-3, -1>
position=<-42470, -10527> velocity=< 4,  1>
position=<-10458,  32093> velocity=< 1, -3>
position=<-21105,  42745> velocity=< 2, -4>
position=< 53479, -31840> velocity=<-5,  3>
position=<-31805,  42744> velocity=< 3, -4>
position=< 21498,  32094> velocity=<-2, -3>
position=<-10486,  32094> velocity=< 1, -3>
position=< 21484,  10781> velocity=<-2, -1>
position=< 32140,  32088> velocity=<-3, -3>
position=<-21130,  10777> velocity=< 2, -1>
position=< 42794, -10533> velocity=<-4,  1>
position=< 21498, -21187> velocity=<-2,  2>
position=<-21116, -31848> velocity=< 2,  3>
position=< 53450, -31844> velocity=<-5,  3>
position=< 32142,  42745> velocity=<-3, -4>
position=<-31782, -53152> velocity=< 3,  5>
position=<-53091, -10527> velocity=< 5,  1>
position=< 42798,  10777> velocity=<-4, -1>
position=<-42422,  42744> velocity=< 4, -4>
position=< 10855, -10529> velocity=<-1,  1>
position=< 21490,  32094> velocity=<-2, -3>
position=<-42437, -42504> velocity=< 4,  4>
position=< 53483, -21192> velocity=<-5,  2>
position=< 53466,  53406> velocity=<-5, -5>
position=<-10458, -31848> velocity=< 1,  3>
position=< 42771,  32090> velocity=<-4, -3>
position=<-31780,  53404> velocity=< 3, -5>
position=<-53115, -10536> velocity=< 5,  1>
position=<-10449,  10776> velocity=< 1, -1>
position=<-31772,  32097> velocity=< 3, -3>
position=< 21515, -21186> velocity=<-2,  2>
position=< 42815,  42749> velocity=<-4, -4>
position=<-21158,  21437> velocity=< 2, -2>
position=< 42774, -31846> velocity=<-4,  3>
position=<-53105,  53405> velocity=< 5, -5>
position=< 21478, -10527> velocity=<-2,  1>
position=<-31781, -21183> velocity=< 3,  2>
position=< 53450,  32093> velocity=<-5, -3>
position=< 10810, -42499> velocity=<-1,  4>
position=< 32132,  10781> velocity=<-3, -1>
position=< 32138,  10785> velocity=<-3, -1>
position=< 10863, -21190> velocity=<-1,  2>
position=< 10859,  53409> velocity=<-1, -5>
position=<-53109, -42499> velocity=< 5,  4>
position=< 32132,  53400> velocity=<-3, -5>
position=< 53434,  32088> velocity=<-5, -3>
position=< 53479, -10530> velocity=<-5,  1>
position=< 53487,  21432> velocity=<-5, -2>
position=< 53427,  10779> velocity=<-5, -1>
position=<-10465,  21434> velocity=< 1, -2>
position=<-21149,  21432> velocity=< 2, -2>
position=< 10823, -42504> velocity=<-1,  4>
position=<-42446,  10783> velocity=< 4, -1>
position=<-31771,  42753> velocity=< 3, -4>
position=< 21501,  10780> velocity=<-2, -1>
position=<-53081, -31847> velocity=< 5,  3>
position=< 53475,  53400> velocity=<-5, -5>
position=<-53074, -53160> velocity=< 5,  5>
position=<-53126,  21441> velocity=< 5, -2>
position=< 10823, -21190> velocity=<-1,  2>
position=< 53485, -53160> velocity=<-5,  5>
position=<-21108,  10785> velocity=< 2, -1>
position=<-42457,  32097> velocity=< 4, -3>
position=<-10486, -53158> velocity=< 1,  5>
position=< 53466,  21432> velocity=<-5, -2>
position=< 10834,  53409> velocity=<-1, -5>
position=< 21493,  32090> velocity=<-2, -3>
position=<-31798, -21183> velocity=< 3,  2>
position=<-42429, -10536> velocity=< 4,  1>
position=<-53131,  10781> velocity=< 5, -1>
position=< 53459,  53404> velocity=<-5, -5>
position=< 53469,  53404> velocity=<-5, -5>
position=< 42828,  53409> velocity=<-4, -5>
position=<-21139, -53160> velocity=< 2,  5>
position=<-31769, -42499> velocity=< 3,  4>
position=< 42821,  10781> velocity=<-4, -1>
position=< 21490,  21434> velocity=<-2, -2>
position=< 21459, -53158> velocity=<-2,  5>
position=< 10842, -10533> velocity=<-1,  1>
position=< 42778,  32097> velocity=<-4, -3>
position=< 32131,  42750> velocity=<-3, -4>
position=< 42821, -10531> velocity=<-4,  1>
position=< 42775, -53159> velocity=<-4,  5>
position=< 42820, -42495> velocity=<-4,  4>
position=< 53468,  53404> velocity=<-5, -5>
position=<-31818,  42746> velocity=< 3, -4>
position=<-42422, -31840> velocity=< 4,  3>
position=< 21479, -21183> velocity=<-2,  2>
position=<-42454,  21435> velocity=< 4, -2>
position=< 10847, -53152> velocity=<-1,  5>
position=<-31813,  53409> velocity=< 3, -5>
position=< 10847,  21435> velocity=<-1, -2>
position=<-53081,  32094> velocity=< 5, -3>
position=<-42473,  32096> velocity=< 4, -3>
position=< 10807, -42503> velocity=<-1,  4>
position=< 10862,  32097> velocity=<-1, -3>
position=< 53469,  10785> velocity=<-5, -1>
position=< 21503, -42498> velocity=<-2,  4>
position=< 53430,  21435> velocity=<-5, -2>
position=< 53427, -31841> velocity=<-5,  3>
position=< 42799, -42495> velocity=<-4,  4>
position=<-21107, -31839> velocity=< 2,  3>
position=< 21511, -53151> velocity=<-2,  5>
position=<-42433, -10530> velocity=< 4,  1>
position=<-31777, -31846> velocity=< 3,  3>
position=<-21141, -21191> velocity=< 2,  2>
position=<-53100,  32094> velocity=< 5, -3>
position=< 21514,  42753> velocity=<-2, -4>
position=<-42466,  10785> velocity=< 4, -1>
position=<-31790,  21441> velocity=< 3, -2>
position=<-42433,  32095> velocity=< 4, -3>
position=<-42430, -42497> velocity=< 4,  4>
position=< 21503, -31840> velocity=<-2,  3>
position=< 21475, -53154> velocity=<-2,  5>
position=<-42449,  32095> velocity=< 4, -3>
position=<-53094,  10784> velocity=< 5, -1>
position=< 21495, -21192> velocity=<-2,  2>
position=< 53434, -21191> velocity=<-5,  2>
position=< 32143, -21189> velocity=<-3,  2>
position=<-21140, -53155> velocity=< 2,  5>
position=< 21495,  42744> velocity=<-2, -4>
position=< 21502,  32097> velocity=<-2, -3>
position=<-53094, -42495> velocity=< 5,  4>
position=<-42474,  53406> velocity=< 4, -5>
position=<-21141, -10535> velocity=< 2,  1>
position=< 42788, -31843> velocity=<-4,  3>
position=<-21139, -21192> velocity=< 2,  2>
position=<-10493,  10782> velocity=< 1, -1>`
