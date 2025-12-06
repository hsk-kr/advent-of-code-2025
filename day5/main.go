package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func ReadFile(path string) []string {
	data, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	return lines
}

func isFreshId(ranges [][]uint64, num uint64) bool {
	for _, r := range ranges {
		if num < r[0] {
			break
		}

		if num <= r[1] {
			return true
		}
	}

	return false
}

func getRange(fileData []string) ([][]uint64, []string) {
	ranges := [][]uint64{}
	var ids []string

	for i, data := range fileData {
		if len(data) == 0 {
			ids = fileData[i+1:]
			break
		}

		numStrs := strings.Split(data, "-")
		minNum, _ := strconv.ParseUint(numStrs[0], 10, 64)
		maxNum, _ := strconv.ParseUint(numStrs[1], 10, 64)
		ranges = append(ranges, []uint64{minNum, maxNum})
	}

	slices.SortFunc(ranges, func(a, b []uint64) int {
		return int(a[0] - b[0])
	})

	return ranges, ids
}

func part1() {
	ranges, ids := getRange(ReadFile("./input"))
	answer := 0

	for _, id := range ids {
		num, _ := strconv.ParseUint(id, 10, 64)
		if isFreshId(ranges, num) {
			answer++
		}
	}

	fmt.Println(answer)
}

func part2() {
	ranges, _ := getRange(ReadFile("./example"))
	var answer uint64 = 0
	var maxVal uint64 = ranges[0][1]
	var minVal uint64 = ranges[0][0]

	for _, r := range ranges {
		canMerge := r[0] <= maxVal

		if canMerge {
			maxVal = r[1]
		} else {
			answer += maxVal - minVal + 1
			minVal = r[0]
			maxVal = r[1]
		}
	}

	answer += maxVal - minVal + 1
	fmt.Println(answer)
}

func main() {
	part1()
	part2()
}
