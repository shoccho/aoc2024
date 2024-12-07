package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/shoccho/aoc2024/utils"
)

func try(numbers []int, target int, current int, index int, part2 bool) bool {
	if index == len(numbers) {
		return current == target
	}
	if try(numbers, target, current+numbers[index], index+1, part2) {
		return true
	}
	if try(numbers, target, current*numbers[index], index+1, part2) {
		return true
	}
	if part2 {
		merged := merge(current, numbers[index])
		if try(numbers, target, merged, index+1, part2) {
			return true
		}
	}

	return false
}

func merge(a, b int) int {
	l := int(math.Log10(float64(b)))

	return a*int(math.Pow10(l+1)) + b
}

func processInput(fileName string) ([]int, [][]int) {
	lines := utils.ReadLines(fileName)
	targets := make([]int, len(lines))
	allNums := make([][]int, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, ":")
		target, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		numParts := strings.Split(parts[1][1:], " ")
		nums := make([]int, len(numParts))
		for i, num := range numParts {
			tmp, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			nums[i] = tmp
		}
		targets[i] = target
		allNums[i] = nums
	}
	return targets, allNums
}

func part1() {
	targets, allnums := processInput("input")
	var sum int64 = 0
	for i, target := range targets {
		nums := allnums[i]
		if try(nums, target, nums[0], 1, false) {
			sum += int64(target)
		}
	}
	fmt.Println("part 1: ", sum)
}

func part2() {
	targets, allnums := processInput("input")
	var sum int64 = 0
	for i, target := range targets {
		nums := allnums[i]
		if try(nums, target, nums[0], 1, true) {
			sum += int64(target)
		}
	}
	fmt.Println("part 2: ", sum)
}

func main() {
	part1()
	part2()
}
