package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadFile(path string) [][]string {
	data, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	fields := make([][]string, len(lines))

	for i, line := range lines {
		fields[i] = strings.Fields(line)
	}

	return fields
}

func Calculate(oper string, acc, num int64) int64 {
	switch oper {
	case "+":
		return acc + num
	case "*":
		return acc * num
	}

	panic("wrong operation")
}

func part1() {
	data := ReadFile("./input")
	var answer int64

	for x := range len(data[0]) {
		acc, _ := strconv.ParseInt(data[0][x], 10, 64)

		for y := range len(data) - 2 {
			y = y + 1
			oper := data[len(data)-1][x]

			num, _ := strconv.ParseInt(data[y][x], 10, 64)

			acc = Calculate(oper, acc, num)
		}
		answer += acc
	}

	fmt.Println(answer)
}

func main() {
	part1()
}
