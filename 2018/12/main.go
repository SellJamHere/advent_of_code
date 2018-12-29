package main

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
	"time"
)

func main() {
	// Part 1
	p1State := agePlants(p1, 20)
	fmt.Println(string(p1State.state))
	fmt.Println(sumPlants(p1State))

	inputState := agePlants(input, 20)
	fmt.Println(string(inputState.state))
	fmt.Println(sumPlants(inputState))

	// Part 2
	part2State := agePlants(input, 1)
	fmt.Println(sumPlants(part2State))
	start := time.Now()
	part2State = agePlants(input, 50000000000)
	fmt.Println(string(part2State.state))
	fmt.Println(sumPlants(part2State))
	fmt.Println(time.Since(start))

	// running the above for 1s creates the follow (generate, sum) pairs
	// 1000, 51883
	// 2000, 102883
	// 3000, 153883
	// 4000, 204883

	// 204883 - 153883 = 51000
	// 153883 - 102883 = 51000
	// 102883 - 51883 = 51000

	// (51000 / 1000) * (50000000000)
	// 2549999999949

	// ((50000000000 - 4000) * (51000 / 1000)) + 1975
	// 2549999797975

	// 2550000000883
}

type rule struct {
	pattern []byte
	plant   []byte
	regex   *regexp.Regexp
}

type plants struct {
	state []byte
	index int
}

func parseInput(input string) (plants, []rule) {
	lines := strings.Split(input, "\n")

	initialState := strings.TrimSpace(lines[0][14:])

	rules := []rule{}
	for _, r := range lines[2:] {
		parts := strings.Split(r, " => ")
		pattern := strings.TrimSpace(parts[0])
		plant := strings.TrimSpace(parts[1])

		r := rule{
			pattern: []byte(pattern),
			plant:   []byte(plant),
			regex:   regexp.MustCompile(regexp.QuoteMeta(pattern)),
		}
		rules = append(rules, r)
	}

	p := plants{
		state: []byte(initialState),
		index: 0,
	}

	return p, rules
}

func padLeft(initial plants, pad byte, length int) plants {
	// Find number of 'pad' on left side of string
	var count int
	for count = 0; count < len(initial.state); count++ {
		if initial.state[count] != pad {
			break
		}
	}

	if count > length {
		offset := count - length
		initial.index -= offset
		initial.state = initial.state[offset:]
	}

	// Pad
	var bb bytes.Buffer
	var i int
	for i = 0; i < length-count; i++ {
		bb.WriteByte(pad)
	}
	bb.Write(initial.state)

	initial.state = bb.Bytes()
	initial.index += i

	return initial
}

func padRight(initial plants, pad byte, length int) plants {
	// Find number of 'pad' on right side of string
	var count int
	for i := len(initial.state) - 1; i > len(initial.state)-length; i-- {
		if initial.state[i] != pad {
			break
		}

		count++
	}

	// Pad
	var bb bytes.Buffer
	bb.Write(initial.state)

	var i int
	for i = 0; i < length-count-1; i++ {
		bb.WriteByte(pad)
	}

	initial.state = bb.Bytes()

	return initial
}

func applyRules(initial plants, rules []rule) plants {
	// Pad left side of pots to account for potential negative indices
	initial = padLeft(initial, []byte(".")[0], 5)
	initial = padRight(initial, []byte(".")[0], 5)

	final := make([]byte, len(initial.state))
	for i := range final {
		final[i] = []byte(".")[0]
	}

	for _, r := range rules {
		// fmt.Println("rule:", r)
		i := 0
		// get all overlapping matches (regexp "Find Alls" ignore overlaps)
		for {
			match := r.regex.FindIndex(initial.state[i:])
			if match == nil {
				break
			}

			// fmt.Println("i:", i, "match:", match)
			// fmt.Println(string(final))
			final[i+match[0]+2] = []byte(r.plant)[0]
			// fmt.Println(string(final))

			i += match[0] + 1
		}
	}

	initial.state = final

	return initial
}

func sumPlants(state plants) int {
	sum := 0
	for i, val := range state.state {
		index := i - state.index
		if string(val) == "#" {
			sum += index
		}
	}

	return sum
}

func agePlants(input string, days int) plants {
	state, rules := parseInput(input)

	start := time.Now()
	for i := 1; i < days+1; i++ {
		state = applyRules(state, rules)

		if i%1000 == 0 {
			fmt.Println(string(state.state), sumPlants(state))
			fmt.Println(i, time.Since(start))
		}
	}

	return state
}

var p1 = `initial state: #..#.#..##......###...###

...## => #
..#.. => #
.#... => #
.#.#. => #
.#.## => #
.##.. => #
.#### => #
#.#.# => #
#.### => #
##.#. => #
##.## => #
###.. => #
###.# => #
####. => #`

var input = `initial state: ..##.#######...##.###...#..#.#.#..#.##.#.##....####..........#..#.######..####.#.#..###.##..##..#..#

#..#. => .
..#.. => .
..#.# => #
##.#. => .
.#... => #
#.... => .
##### => #
.#.## => .
#.#.. => .
#.### => #
.##.. => #
##... => .
#...# => #
####. => #
#.#.# => .
#..## => .
.#### => .
...## => .
..### => #
.#..# => .
##..# => #
.#.#. => .
..##. => .
###.. => .
###.# => #
#.##. => #
..... => .
.##.# => #
....# => .
##.## => #
...#. => #
.###. => .`
