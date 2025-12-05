package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func ReadFile(path string) []string {
	data, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	return strings.Split(string(data), "\n")
}

func Turn(currPos int, operation string) (int, int) {
	if len(operation) == 0 {
		return currPos, 0
	}

	dir := operation[0]
	num, _ := strconv.Atoi(operation[1:])
	var val int

	if dir == 'L' {
		val = currPos - num
	} else {
		val = currPos + num
	}

	nextPos := val % 100
	if nextPos < 0 {
		nextPos = 100 + nextPos
	}
	numClicks := int(math.Abs(float64(val / 100)))

	if currPos > 0 && val < 0 || val == 0 {
		numClicks++
	}

	return nextPos, numClicks
}

func part1() {
	input := ReadFile("./input")
	cur := 50
	answer := 0

	for _, oper := range input {
		newCur, _ := Turn(cur, oper)
		cur = newCur

		if cur == 0 {
			answer++
		}
	}

	fmt.Println(answer)
}

func part2() {
	input := ReadFile("./input")
	cur := 50
	answer := 0

	for _, oper := range input {
		newCur, numClicks := Turn(cur, oper)
		cur = newCur
		answer += numClicks
	}

	fmt.Println(answer)
}

func main() {
	part1()
	part2()
}
