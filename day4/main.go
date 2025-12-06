package main

import (
	"fmt"
	"os"
	"strings"
)

func ReadFile(path string) [][]rune {
	data, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	runes := make([][]rune, len(lines))
	for i, line := range lines {
		runes[i] = []rune(line)
	}

	return runes
}

func CanAccess(m [][]rune, minx, miny, maxx, maxy int) bool {
	if minx < 0 {
		minx = 0
	}
	if miny < 0 {
		miny = 0
	}
	if maxx > len(m[0])-1 {
		maxx = len(m[0]) - 1
	}
	if maxy > len(m)-1 {
		maxy = len(m) - 1
	}

	numPapers := 0

	for y := miny; y <= maxy; y++ {
		for x := minx; x <= maxx; x++ {
			if m[y][x] == '@' {
				numPapers++
			}
		}
	}

	return (numPapers - 1) < 4
}

func part1() {
	input := ReadFile("./input")
	answer := 0

	for y, v := range input {
		for x := range v {
			if input[y][x] != '@' {
				continue
			}

			if CanAccess(input, x-1, y-1, x+1, y+1) {
				answer++
			}
		}
	}

	fmt.Println(answer)
}

func part2() {
	input := ReadFile("./input")
	answer := 0

	var numAccessPaper int
	for answer == 0 || numAccessPaper != 0 {
		numAccessPaper = 0
		pos := make(map[int][]int)
		for y, v := range input {

			for x := range v {
				if input[y][x] != '@' {
					continue
				}

				if CanAccess(input, x-1, y-1, x+1, y+1) {
					pos[y*10+x] = []int{y, x}
					numAccessPaper++
				}
			}
		}

		for _, p := range pos {
			input[p[0]][p[1]] = '.'
		}

		answer += numAccessPaper
	}

	fmt.Println(answer)
}

func main() {
	part1()
	part2()
}
