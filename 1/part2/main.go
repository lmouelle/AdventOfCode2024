package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	// As you scan the input file, count the number of rhs number occurences in a map
	// After reading input and counting, iterate through lhs array
	// As you iterate, get the count of the num from the map and multiply together,
	// summing the product

	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var left_parts []int
	right_parts := make(map[int]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		lhs, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		rhs, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		left_parts = append(left_parts, lhs)
		right_parts[rhs]++
	}
	sum := 0
	for _, v := range left_parts {
		operand := right_parts[v]
		sum += (operand * v)
	}

	println(sum)
}
