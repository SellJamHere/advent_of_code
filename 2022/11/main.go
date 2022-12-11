package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	monkeyInputs := strings.Split(puzzleInput, "\n\n")

	// part 1
	var monkeys []Monkey
	for _, input := range monkeyInputs {
		monkeys = append(monkeys, NewMonkey(input))
	}

	for i := 0; i < 20; i++ {
		for _, monkey := range monkeys {
			monkey.TakeTurn(monkeys)
		}
	}

	var counts []int
	for _, monkey := range monkeys {
		counts = append(counts, *monkey.inspectCount)
	}
	sort.Ints(counts)

	fmt.Printf("part 1: monkey business: %d\n", counts[len(counts)-2]*counts[len(counts)-1])

	// part 1
	monkeys = nil
	for _, input := range monkeyInputs {
		monkeys = append(monkeys, NewMonkey(input))
	}

	for i := 0; i < 10000; i++ {
		for _, monkey := range monkeys {
			monkey.TakeTurn2(monkeys)
		}
	}

	counts = nil
	for _, monkey := range monkeys {
		counts = append(counts, *monkey.inspectCount)
	}
	sort.Ints(counts)

	fmt.Printf("part 2: monkey business: %d\n", counts[len(counts)-2]*counts[len(counts)-1])
}

type Monkey struct {
	id           int
	itemQueue    *Queue
	inspectCount *int

	operationStr string

	testDivisor int
	trueMonkey  int
	falseMonkey int
}

func NewMonkey(in string) Monkey {
	lines := strings.Split(in, "\n")
	idStr := strings.ReplaceAll(strings.ReplaceAll(lines[0], "Monkey ", ""), ":", "")
	id, _ := strconv.Atoi(strings.TrimSpace(idStr))

	var q Queue
	itemStrs := strings.Split(strings.ReplaceAll(strings.TrimSpace(lines[1]), "Starting items: ", ""), ",")
	for _, itemStr := range itemStrs {
		item, _ := strconv.Atoi(strings.TrimSpace(itemStr))
		newItem := &Node{item: item}
		q.Queue(newItem)
	}

	operationStr := strings.ReplaceAll(strings.TrimSpace(lines[2]), "Operation: new = ", "")

	divisorStr := strings.ReplaceAll(strings.TrimSpace(lines[3]), "Test: divisible by ", "")
	divisor, _ := strconv.Atoi(strings.TrimSpace(divisorStr))

	trueMonkeyStr := strings.ReplaceAll(strings.TrimSpace(lines[4]), "If true: throw to monkey ", "")
	trueMonkey, _ := strconv.Atoi(strings.TrimSpace(trueMonkeyStr))

	falseMonkeyStr := strings.ReplaceAll(strings.TrimSpace(lines[5]), "If false: throw to monkey ", "")
	falseMonkey, _ := strconv.Atoi(strings.TrimSpace(falseMonkeyStr))

	inspect := 0
	return Monkey{
		id:           id,
		itemQueue:    &q,
		inspectCount: &inspect,
		operationStr: operationStr,
		testDivisor:  divisor,
		trueMonkey:   trueMonkey,
		falseMonkey:  falseMonkey,
	}
}

func (m *Monkey) TakeTurn(monkeys []Monkey) {
	for m.itemQueue.head != nil {
		val := m.itemQueue.Dequeue()

		result := m.Inspect(val.item)
		result /= 3
		val.item = result

		if result%m.testDivisor == 0 {
			monkeys[m.trueMonkey].itemQueue.Queue(val)
		} else {
			monkeys[m.falseMonkey].itemQueue.Queue(val)
		}
	}
}

func (m *Monkey) Inspect(val int) int {
	*(m.inspectCount)++

	parts := strings.Split(m.operationStr, " ")

	operator := parts[1]
	operandStr := parts[2]

	var operand int
	if operandStr == "old" {
		operand = val
	} else {
		operand, _ = strconv.Atoi(operandStr)
	}

	switch operator {
	case "+":
		return val + operand
	case "*":
		return val * operand
	}

	return -1
}

// calculate lcm of the monkey test divisors, and mod by that to reduce
// the result. This ensures the item value is never larger than the lcm
// of the divisors. Any value reduced by the LCM mod is still divisible
// by the same values, since the LCM itself is divisible by the values.
//
// 23, 19, 13, 17 -> LCM = 96577
// 100000 % 23 = 19
//
// 100000 % 96577 = 3423
// 3423 % 23 = 19

func (m *Monkey) TakeTurn2(monkeys []Monkey) {
	var mods []int
	for _, monkey := range monkeys {
		mods = append(mods, monkey.testDivisor)
	}

	lcm := lcm(mods...)

	for m.itemQueue.head != nil {
		val := m.itemQueue.Dequeue()

		result := m.Inspect(val.item)
		result %= lcm
		val.item = result

		if result%m.testDivisor == 0 {
			monkeys[m.trueMonkey].itemQueue.Queue(val)
		} else {
			monkeys[m.falseMonkey].itemQueue.Queue(val)
		}
	}
}

// old * old overflows
// need to get the appropriate "test" mod without overflowing first
// reduce the values proportionally using a common denominator, then
// multiply, then mod. The mod value should be the same before and after
// dividing by the common denominator.
//
// 28 * 16 = 448 (448 % 7 = 0)
// (7*4) * (4*4) = (112*4)
// 7 * 4 * 4 = 112 (112 % 7 = 0)
//
// Scrapping this idea because I can't think of an analogue for addition.
// See LCM solution above.

func (m *Monkey) InspectUnused(val int) int {
	*(m.inspectCount)++

	parts := strings.Split(m.operationStr, " ")

	operator := parts[1]
	operandStr := parts[2]

	var operand int
	if operandStr == "old" {
		operand = val
	} else {
		operand, _ = strconv.Atoi(operandStr)
	}

	switch operator {
	case "+":
		return val + operand
	case "*":
		denom := gcd(val, operand)
		return val / denom * operand
	}

	return -1
}

func gcd(a, b int) int {
	for b != 0 {
		tmp := b
		b = a % b
		a = tmp
	}

	return a
}

func lcm(vals ...int) int {
	a := vals[0]
	b := vals[1]
	vals = vals[2:]
	tmp := a * b / gcd(a, b)

	for i := 0; i < len(vals); i++ {
		tmp = lcm(tmp, vals[i])
	}

	return tmp
}

type Node struct {
	item int
	next *Node
}

type Queue struct {
	head *Node
	len  int
}

func (q *Queue) Queue(n *Node) {
	q.len++

	if q.head == nil {
		q.head = n
		return
	}

	node := q.head
	for node.next != nil {
		node = node.next
	}

	node.next = n
}

func (q *Queue) Dequeue() *Node {
	q.len--
	node := q.head
	q.head = q.head.next
	node.next = nil

	return node
}

const puzzleInput1 = `Monkey 0:
  Starting items: 79, 98
  Operation: new = old * 19
  Test: divisible by 23
    If true: throw to monkey 2
    If false: throw to monkey 3

Monkey 1:
  Starting items: 54, 65, 75, 74
  Operation: new = old + 6
  Test: divisible by 19
    If true: throw to monkey 2
    If false: throw to monkey 0

Monkey 2:
  Starting items: 79, 60, 97
  Operation: new = old * old
  Test: divisible by 13
    If true: throw to monkey 1
    If false: throw to monkey 3

Monkey 3:
  Starting items: 74
  Operation: new = old + 3
  Test: divisible by 17
    If true: throw to monkey 0
    If false: throw to monkey 1`

const puzzleInput = `Monkey 0:
  Starting items: 72, 64, 51, 57, 93, 97, 68
  Operation: new = old * 19
  Test: divisible by 17
    If true: throw to monkey 4
    If false: throw to monkey 7

Monkey 1:
  Starting items: 62
  Operation: new = old * 11
  Test: divisible by 3
    If true: throw to monkey 3
    If false: throw to monkey 2

Monkey 2:
  Starting items: 57, 94, 69, 79, 72
  Operation: new = old + 6
  Test: divisible by 19
    If true: throw to monkey 0
    If false: throw to monkey 4

Monkey 3:
  Starting items: 80, 64, 92, 93, 64, 56
  Operation: new = old + 5
  Test: divisible by 7
    If true: throw to monkey 2
    If false: throw to monkey 0

Monkey 4:
  Starting items: 70, 88, 95, 99, 78, 72, 65, 94
  Operation: new = old + 7
  Test: divisible by 2
    If true: throw to monkey 7
    If false: throw to monkey 5

Monkey 5:
  Starting items: 57, 95, 81, 61
  Operation: new = old * old
  Test: divisible by 5
    If true: throw to monkey 1
    If false: throw to monkey 6

Monkey 6:
  Starting items: 79, 99
  Operation: new = old + 2
  Test: divisible by 11
    If true: throw to monkey 3
    If false: throw to monkey 1

Monkey 7:
  Starting items: 68, 98, 62
  Operation: new = old + 3
  Test: divisible by 13
    If true: throw to monkey 5
    If false: throw to monkey 6`
