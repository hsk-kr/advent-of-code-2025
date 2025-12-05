package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadFile(path string) []string {
	data, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	return strings.Split(string(data), ",")
}

func GetRange(rangeStr string) (int64, int64) {
	nums := strings.Split(strings.TrimSpace(rangeStr), "-")

	minNum, _ := strconv.ParseInt(nums[0], 10, 64)
	maxNum, _ := strconv.ParseInt(nums[1], 10, 64)

	return minNum, maxNum
}

func calcNumInvalidId(minNum int64, maxNum int64) int64 {
	var sumInvalidIds int64 = 0

	for num := minNum; num <= maxNum; num++ {
		if !isValidId(strconv.FormatInt(num, 10)) {
			sumInvalidIds += num
		}
	}

	return sumInvalidIds
}

func isValidId(num string) bool {
	halfLen := len(num) / 2
	if len(num)%2 == 1 {
		return true
	}

	for i := range halfLen {
		if num[i] != num[halfLen+i] {
			return true
		}
	}

	return false
}

func calcNumInvalidIdPart2(minNum int64, maxNum int64) int64 {
	var sumInvalidIds int64 = 0

	for num := minNum; num <= maxNum; num++ {
		if !IsValidIdPart2(strconv.FormatInt(num, 10)) {
			sumInvalidIds += num
		}
	}

	return sumInvalidIds
}

func IsValidIdPart2(num string) bool {
	halfLen := len(num) / 2
	for n := halfLen; n >= 1; n-- {
		if len(num)%n != 0 {
			continue
		}

		valid := false
		for i := 0; i < n; i++ {
			if valid {
				break
			}
			for j := i + n; j < len(num); j += n {
				if num[i] != num[j] {
					valid = true
					break
				}
			}
		}
		if !valid {
			return false
		}
	}

	return true
}

func part1() {
	input := ReadFile("./input")
	var answer int64 = 0

	for _, rangePair := range input {
		minNum, maxNum := GetRange(rangePair)
		answer += calcNumInvalidId(minNum, maxNum)
	}

	fmt.Println(answer)
}

func part2() {
	input := ReadFile("./input")
	var answer int64 = 0

	for _, rangePair := range input {
		minNum, maxNum := GetRange(rangePair)
		answer += calcNumInvalidIdPart2(minNum, maxNum)
	}

	fmt.Println(answer)
}

func main() {
	part1()
	part2()
}
