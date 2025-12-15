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

	lines := strings.Split(string(data), "\n")
	fields := make([][]rune, len(lines))

	for i, line := range lines {
		fields[i] = []rune(line)
	}

	fields = fields[0 : len(fields)-1]

	return fields
}

func Calculate(oper rune, acc, num int64) int64 {
	switch oper {
	case '+':
		return acc + num
	case '*':
		return acc * num
	}

	panic("wrong operation")
}

// doesn't work after the chnages for part2 but can see the solution the commit before this
func part1() {
	// data := ReadFile("./input")
	// var answer int64
	//
	// for x := range len(data[0]) {
	// 	acc, _ := strconv.ParseInt(string(data[0][x]), 10, 64)
	//
	// 	for y := range len(data) - 2 {
	// 		y = y + 1
	// 		oper := data[len(data)-1][x]
	//
	// 		num, _ := strconv.ParseInt(string(data[y][x]), 10, 64)
	//
	// 		acc = Calculate(oper, acc, num)
	// 	}
	// 	answer += acc
	// }
	//
	// fmt.Println(answer)
}

func part2() {
	data := ReadFile("./input")
	var answer int64
	var operator rune
	var sum int64

	for x := range len(data[0]) {
		lastRow := data[len(data)-1][x]
		if lastRow == '+' || lastRow == '*' {
			operator = lastRow

			if lastRow == '+' {
				sum = 0
			} else {
				sum = 1
			}
		}

		var n int64

		for y := range len(data) - 1 {
			if data[y][x] == ' ' {
				continue
			}

			n = n*10 + int64(data[y][x]-'0')
		}

		if n == 0 {
			answer += sum
			continue
		}

		sum = Calculate(operator, sum, n)
	}

	answer += sum
	fmt.Println(answer)
}

func main() {
	part1()
	part2()
}
