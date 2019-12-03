package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func main() {
	runCarts(p1)

	// read input from file, as saving it as a string compresses whitespace
	input, _ := ioutil.ReadFile("input.txt")
	runCarts(string(input))
}

const (
	crashed direction = -1
	up      direction = 0
	right   direction = 1
	down    direction = 2
	left    direction = 3
)

type direction int

func newDirection(cart string) direction {
	switch cart {
	case "^":
		return up
	case ">":
		return right
	case "v":
		return down
	case "<":
		return left
	case "X":
		return crashed
	default:
		return direction(-1)
	}
}

func (d direction) string() string {
	switch d {
	case up:
		return "^"
	case right:
		return ">"
	case down:
		return "v"
	case left:
		return "<"
	case crashed:
		return "X"
	}

	return ""
}

func (d direction) turn(turnDir direction) direction {
	if turnDir == left {
		d--
		if d < up {
			d = left
		}
		return d
	} else if turnDir == right {
		d++
		if d > left {
			d = up
		}

		return d
	}

	return d
}

type cart struct {
	x                 int
	y                 int
	d                 direction
	intersectionCount int
	previousTrack     string
}

type carts []cart

func (c carts) hasCrash() bool {
	for _, cart := range c {
		if cart.d == crashed {
			return true
		}
	}

	return false
}

func (c carts) crashes() [][]int {
	var crashes [][]int
	for _, cart := range c {
		if cart.d == crashed {
			crashes = append(crashes, []int{cart.x, cart.y})
		}
	}

	return crashes
}

func findCarts(cartMap []*string) carts {
	carts := carts{}
	for i, line := range cartMap {
		for j, r := range *line {
			char := string(r)
			if char == "^" || char == ">" || char == "v" || char == "<" {
				var previousTrack string
				if char == "^" || char == "v" {
					previousTrack = "|"
				} else {
					previousTrack = "-"
				}
				c := cart{
					x:             j,
					y:             i,
					d:             newDirection(char),
					previousTrack: previousTrack,
				}

				carts = append(carts, c)
			}
		}
	}

	return carts
}

func tick(cartMap []*string, carts carts) {
	// move each cart
	for i, c := range carts {
		prevX := c.x
		prevY := c.y
		prevTrack := c.previousTrack

		// look at next track piece
		var nextTrack string
		switch c.d {
		case up:
			c.y--
			nextTrack = string((*cartMap[c.y])[c.x])
		case right:
			c.x++
			nextTrack = string((*cartMap[c.y])[c.x])
		case down:
			c.y++
			nextTrack = string((*cartMap[c.y])[c.x])
		case left:
			c.x--
			nextTrack = string((*cartMap[c.y])[c.x])
		}

		c.previousTrack = nextTrack

		if nextTrack == "\\" {
			if c.d == up {
				c.d = left
			} else if c.d == right {
				c.d = down
			} else if c.d == down {
				c.d = right
			} else if c.d == left {
				c.d = up
			}
		} else if nextTrack == "/" {
			if c.d == up {
				c.d = right
			} else if c.d == right {
				c.d = up
			} else if c.d == down {
				c.d = left
			} else if c.d == left {
				c.d = down
			}
		} else if nextTrack == "+" {
			switch c.intersectionCount % 3 {
			case 0: // left
				c.d = c.d.turn(left)
			case 1: // straight
			case 2: // right
				c.d = c.d.turn(right)
			}

			c.intersectionCount++
		} else if nextTrack == "^" || nextTrack == ">" || nextTrack == "v" || nextTrack == "<" || nextTrack == "X" {
			c.d = crashed
		}
		// else, stay the course

		// update previous track
		prevLine := ((*cartMap[prevY])[:prevX] + prevTrack + (*cartMap[prevY])[prevX+1:])
		cartMap[prevY] = &prevLine

		// update current track
		newLine := ((*cartMap[c.y])[:c.x] + c.d.string() + (*cartMap[c.y])[c.x+1:])
		cartMap[c.y] = &newLine

		carts[i] = c
	}
}

func runCarts(input string) {
	tmpCartMap := strings.Split(input, "\n")
	cartMap := make([]*string, len(tmpCartMap))
	for i := range tmpCartMap {
		cartMap[i] = &tmpCartMap[i]
	}

	carts := findCarts(cartMap)

	i := 0
	for !carts.hasCrash() {
		sort.SliceStable(carts, func(i, j int) bool {
			if carts[i].y < carts[j].y {
				return true
			} else if carts[i].y == carts[j].y {
				return carts[i].x < carts[j].x
			} else {
				return false
			}
		})

		tick(cartMap, carts)
		i++
	}

	fmt.Println("Final:", i)
	crashes := carts.crashes()
	fmt.Println(crashes)

	// for _, line := range cartMap {
	// 	fmt.Println(*line)
	// }
}

var p1 = `/->-\
|   |  /----\
| /-+--+-\  |
| | |  | v  |
\-+-/  \-+--/
  \------/   `

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"os"
// 	"sort"
// )

// var width int
// var data []byte
// var carts []*complex64
// var cartData map[complex64]*Cart = make(map[complex64]*Cart)

// func read(pos complex64) byte {
// 	return data[int(real(pos))+int(imag(pos))*width]
// }

// type CartSort []*complex64

// func (c CartSort) Len() int      { return len(c) }
// func (c CartSort) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
// func (c CartSort) Less(i, j int) bool {
// 	return imag(*c[i]) < imag(*c[j]) || (imag(*c[i]) == imag(*c[j]) && real(*c[i]) < real(*c[j]))
// }

// type Cart struct {
// 	pos   complex64
// 	dir   complex64
// 	state complex64
// }

// func (c *Cart) Tick() {
// 	delete(cartData, c.pos)
// 	c.pos += c.dir
// 	if crash, exists := cartData[c.pos]; exists {
// 		fmt.Println("Crash at:", real(c.pos), imag(c.pos))
// 		delete(cartData, c.pos)
// 		crash.pos, c.pos = 0, 0
// 		return
// 	}
// 	cartData[c.pos] = c
// 	if read(c.pos) == '+' {
// 		c.dir *= c.state
// 		switch c.state {
// 		case -1i:
// 			c.state = 1
// 		case 1:
// 			c.state = 1i
// 		case 1i:
// 			c.state = -1i
// 		}
// 	} else if read(c.pos) == '/' {
// 		c.dir = complex(-imag(c.dir), -real(c.dir))
// 	} else if read(c.pos) == '\\' {
// 		c.dir = complex(imag(c.dir), real(c.dir))
// 	}
// }

// func main() {
// 	reader, err := os.Open("input.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	scanner := bufio.NewScanner(reader)
// 	for scanner.Scan() {
// 		line := scanner.Bytes()
// 		if width == 0 {
// 			width = len(line)
// 		}
// 		data = append(data, line...)
// 	}
// 	reader.Close()

// 	for i := 0; i < len(data); i++ {
// 		pos := complex(float32(i%width), float32(i/width))
// 		switch read(pos) {
// 		case '^':
// 			cartData[pos] = &Cart{pos, -1i, -1i}
// 		case '>':
// 			cartData[pos] = &Cart{pos, 1, -1i}
// 		case 'v':
// 			cartData[pos] = &Cart{pos, 1i, -1i}
// 		case '<':
// 			cartData[pos] = &Cart{pos, -1, -1i}
// 		default:
// 			continue
// 		}
// 		carts = append(carts, &cartData[pos].pos)
// 	}

// 	for len(cartData) > 1 {
// 		sort.Sort(CartSort(carts))

// 		for _, cart := range carts {
// 			if *cart != 0 {
// 				cartData[*cart].Tick()
// 			}
// 		}
// 	}

// 	for pos, _ := range cartData {
// 		fmt.Println("Last cart:", real(pos), imag(pos))
// 	}
// }
