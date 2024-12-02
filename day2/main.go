package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/shoccho/aoc2024/utils"
)

func isSafe(line []string) bool {
	var prev int
	inc := false
	for i, char := range line {
		curr, err := strconv.Atoi(char)
		if err != nil {
			continue
			panic(err)
		}
		if i == 0 {
			prev = curr
			continue
		}
		if i == 1 {
			if prev < curr {
				inc = true
			} else if prev > curr {
				inc = false
			} else {
				return false
			}
		}
		if !checkMargine(curr, prev, inc) {
			return false
		}
		prev = curr
	}
	return true
}

func checkMargine(a, b int, inc bool) bool {
	if inc {
		if a <= b {
			return false
		} else if a-b > 3 {
			return false
		}
	} else {
		if a >= b {
			return false
		} else if b-a > 3 {
			return false
		}
	}
	return true
}

func part2() {
	lines := utils.ReadLines("input")
	count := 0
	for _, line := range lines {
		nums := strings.Split(line, " ")
		n := len(nums)
		safe := false

		if isSafe(nums) {
			safe = true
		} else {
			for j := 0; j < n; j++ {
				if isSafe(slices.Concat(nums[:j], nums[j+1:])) {
					safe = true
					break
				}
			}
		}
		if safe {
			count++
		}
	}
	fmt.Println("part2: ", count)
}

func part1() {
	lines := utils.ReadLines("input")
	count := 0
	for _, line := range lines {
		if isSafe(strings.Split(line, " ")) {
			count++
		}
	}
	fmt.Println("part1: ", count)
}

func main() {
	part1()
	part2()
}
