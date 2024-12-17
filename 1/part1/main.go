package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// First, radix sort the two arrays for O(n)
	// Next, walk two pointers slow through the arrays, calculating the difference one-by-one
	// Sum the deltas as we go
	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var left_parts []int
	var right_parts []int

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
		right_parts = append(right_parts, rhs)
	}

	// TODO: Switch to radix sort
	sort.Slice(left_parts, func(i, j int) bool {
		return left_parts[i] < left_parts[j]
	})
	sort.Slice(right_parts, func(i, j int) bool {
		return right_parts[i] < right_parts[j]
	})

	sum := 0
	// right_parts/right_idx should have same valules
	for idx := 0; idx < len(left_parts); idx++ {
		delta := left_parts[idx] - right_parts[idx]
		if delta < 0 {
			delta = -delta
		}

		sum += delta
	}

	println(sum)
}
