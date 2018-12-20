package main

import (
	"fmt"
	"time"
)

func main() {
	// Part 1

	// 9 players; last marble is worth 25 points: high score is 32
	fmt.Println(playGame(9, 32))
	// 10 players; last marble is worth 1618 points: high score is 8317
	fmt.Println(playGame(10, 1618))
	// 13 players; last marble is worth 7999 points: high score is 146373
	fmt.Println(playGame(13, 7999))
	// 17 players; last marble is worth 1104 points: high score is 2764
	fmt.Println(playGame(17, 1104))
	// 21 players; last marble is worth 6111 points: high score is 54718
	fmt.Println(playGame(21, 6111))
	// 30 players; last marble is worth 5807 points: high score is 37305
	fmt.Println(playGame(30, 5807))

	// input
	// 430 players; last marble is worth 71588 points
	start := time.Now()
	fmt.Println(playGame(430, 71588))
	fmt.Println("part 1:", time.Since(start))

	// Part 2

	// 9 players; last marble is worth 25 points: high score is 32
	fmt.Println(playGameLinkedList(9, 32))
	// 10 players; last marble is worth 1618 points: high score is 8317
	fmt.Println(playGameLinkedList(10, 1618))
	// 13 players; last marble is worth 7999 points: high score is 146373
	fmt.Println(playGameLinkedList(13, 7999))
	// 17 players; last marble is worth 1104 points: high score is 2764
	fmt.Println(playGameLinkedList(17, 1104))
	// 21 players; last marble is worth 6111 points: high score is 54718
	fmt.Println(playGameLinkedList(21, 6111))
	// 30 players; last marble is worth 5807 points: high score is 37305
	fmt.Println(playGameLinkedList(30, 5807))

	// input
	// 430 players; last marble is worth 71588 points
	start = time.Now()
	fmt.Println(playGameLinkedList(430, 71588*100))
	fmt.Println("part 2:", time.Since(start))
}

func playGame(playerCount, finalMarble int) int {
	scores := make([]int, playerCount)

	circle := []int{0}
	currentIndex := 1

	for i := 1; i <= finalMarble; i++ {
		if i%23 != 0 {
			insertPosition := currentIndex + 2
			for insertPosition > len(circle) {
				insertPosition = insertPosition - len(circle)
			}

			// insert
			circle = append(circle, 0) // add element to ensure slice is large enough
			copy(circle[insertPosition+1:], circle[insertPosition:])
			circle[insertPosition] = i

			currentIndex = insertPosition
		} else {
			// special case
			currentPlayer := i % playerCount

			scores[currentPlayer] += i

			removeIndex := currentIndex - 7
			if removeIndex < 0 {
				removeIndex = len(circle) + removeIndex
			}

			scores[currentPlayer] += circle[removeIndex]
			// delete
			circle = append(circle[:removeIndex], circle[removeIndex+1:]...)

			currentIndex = removeIndex
		}
	}

	highScore := 0
	for _, score := range scores {
		if score > highScore {
			highScore = score
		}
	}

	return highScore
}

type marble struct {
	value int
	prev  *marble
	next  *marble
}

func (m *marble) insert(value int) *marble {
	newMarble := &marble{
		value: value,
		prev:  m,
		next:  m.next,
	}

	m.next.prev = newMarble
	m.next = newMarble

	return newMarble
}

func (m *marble) delete() *marble {
	m.prev.next = m.next
	m.next.prev = m.prev

	return m.next
}

func (m *marble) print() {
	line := fmt.Sprintf("%d ", m.value)
	ptr := m.next
	for ptr != m {
		line += fmt.Sprintf("%d ", ptr.value)
		ptr = ptr.next
	}

	fmt.Println(line)
}

func playGameLinkedList(playerCount, finalMarble int) int {

	scores := make([]int, playerCount)

	currentMarble := &marble{0, nil, nil}
	currentMarble.prev = currentMarble
	currentMarble.next = currentMarble

	for i := 1; i <= finalMarble; i++ {
		if i%23 != 0 {
			currentMarble = currentMarble.next.insert(i)
		} else {
			// special case
			removedMarble := currentMarble.prev.prev.prev.prev.prev.prev.prev
			currentPlayer := i % playerCount
			scores[currentPlayer] += i + removedMarble.value

			currentMarble = removedMarble.delete()
		}
	}

	highScore := 0
	for _, score := range scores {
		if score > highScore {
			highScore = score
		}
	}

	return highScore
}
