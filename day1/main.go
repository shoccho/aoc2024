package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readFile(fileName string) string {
	b, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func dist(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func part2() {
	input := readFile("./input")
	lines := strings.Split(input, "\n")
	n := len(lines)
	left, right := make([]int, n), make([]int, 100000) // a hashmap would have been the correct choice
	for i, line := range lines {
		nums := strings.Split(line, "   ")
		if len(nums) != 2 {
			continue
		}
		parsed, err := strconv.Atoi(nums[0])
		if err != nil {
			panic(err)
		}
		left[i] = parsed

		parsed, err = strconv.Atoi(nums[1])
		if err != nil {
			panic(err)
		}
		right[parsed]++
	}
	sum := 0
	for i := 0; i < n; i++ {
		sum += right[left[i]] * left[i]
	}
	fmt.Println(sum)
}

func part1() {
	input := readFile("./input")
	lines := strings.Split(input, "\n")
	n := len(lines)
	left, right := make([]int, n), make([]int, n)

	for i, line := range lines {

		nums := strings.Split(line, "   ")

		if len(nums) != 2 {
			continue
		}
		parsed, err := strconv.Atoi(nums[0])
		if err != nil {
			panic(err)
		}
		left[i] = parsed

		parsed, err = strconv.Atoi(nums[1])
		if err != nil {
			panic(err)
		}
		right[i] = parsed
	}
	sort.Ints(left)
	sort.Ints(right)
	sum := 0
	for i := 0; i < n; i++ {
		sum += dist(left[i], right[i])
	}
	fmt.Println(sum)
}

func main() {
	fmt.Println("Part 1:")
	part1()
	fmt.Println("Part 2:")
	part2()
}
