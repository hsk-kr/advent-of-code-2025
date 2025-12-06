package main

import (
	"fmt"
	"os"
	"strings"
)

func ReadFile(path string) []string {
	data, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	return strings.Split(string(data), "\n")
}

func part1() {
	input := ReadFile("./input")
	answer := 0

	for _, v := range input {
		v = strings.TrimSpace(v)
		if len(v) == 0 {
			break
		}

		firstJoltageIndex := 0
		secondJoltageIndex := 0

		for j := 1; j < len(v)-1; j++ {
			if v[j] > v[firstJoltageIndex] {
				firstJoltageIndex = j
			}
		}

		secondJoltageIndex = firstJoltageIndex + 1
		for j := firstJoltageIndex + 2; j < len(v); j++ {
			if v[j] > v[secondJoltageIndex] {
				secondJoltageIndex = j
			}
		}

		joltage := int((v[firstJoltageIndex]-'0'))*10 + int(v[secondJoltageIndex]-'0')
		answer += joltage
	}

	fmt.Println(answer)
}

func joltageToPower(line string, joltageIndexes [12]int) int64 {
	var power int64 = 0

	for _, joltageIndex := range joltageIndexes {
		power = power*10 + int64(line[joltageIndex]-'0')
	}

	return power
}

func CalcJoltagePart2(v string) int64 {
	var joltageIndexes [12]int
	for i := range joltageIndexes {
		if i == 0 {
			joltageIndexes[i] = 0
		} else {
			joltageIndexes[i] = joltageIndexes[i-1] + 1
		}

		for j := joltageIndexes[i] + 1; j <= len(v)-len(joltageIndexes)+i; j++ {
			if v[j] > v[joltageIndexes[i]] {
				joltageIndexes[i] = j
			}
		}
	}

	return joltageToPower(v, joltageIndexes)
}

func part2() {
	input := ReadFile("./input")
	var answer int64 = 0

	for _, v := range input {
		v = strings.TrimSpace(v)
		if len(v) == 0 {
			break
		}

		answer += CalcJoltagePart2(v)
	}

	fmt.Println(answer)
}

func main() {
	part1()
	part2()
}
