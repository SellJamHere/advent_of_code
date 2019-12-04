package main

import (
	"fmt"
	"strconv"
	"strings"
)

const input = `172851-675869`

// const input = `111122-111122`

func main() {
	ranges := strings.Split(input, "-")
	start, err := strconv.Atoi(ranges[0])
	if err != nil {
		panic(err)
	}

	end, err := strconv.Atoi(ranges[1])
	if err != nil {
		panic(err)
	}

	// Part 1
	count := 0
	password := start
	for password <= end {
		passwordStr := strconv.Itoa(password)
		if len(passwordStr) == 6 && hasDoubleDigits(passwordStr) && increases(passwordStr) {
			count++
		}

		password++
	}

	fmt.Printf("password count: %d\n", count)

	// Part 2
	count = 0
	password = start
	for password <= end {
		passwordStr := strconv.Itoa(password)
		if len(passwordStr) == 6 && hasOnlyDoubleDigits(passwordStr) && increases(passwordStr) {
			count++
		}

		password++
	}

	fmt.Printf("password count 2: %d\n", count)
}

func hasDoubleDigits(password string) bool {
	for i := 1; i < len(password); i++ {
		if password[i-1] == password[i] {
			return true
		}
	}
	return false
}

func hasOnlyDoubleDigits(password string) bool {
	i := 1
	matchingDigit := password[0]
	count := 1
	for i < len(password) {
		for i < len(password) && password[i] == matchingDigit {
			count++
			i++
		}

		if count == 2 {
			return true
		}

		count = 1
		if i < len(password) {
			matchingDigit = password[i]
		}
		i++
	}

	return false
}

func increases(password string) bool {
	prev := string(password[0])
	for i := 1; i < len(password); i++ {
		if prev > string(password[i]) {
			return false
		}
		prev = string(password[i])
	}
	return true
}
