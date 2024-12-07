package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/shoccho/aoc2024/utils"
)

func geta(x, y int) int {
	if x == 0 || y == 0 {
		return 0
	} else if x < 10 {
		return y * 10
	} else if x < 100 {
		return y * 100
	} else if x < 1000 {
		return y * 1000
	} else if x < 10000 {
		return y * 10000
	} else if x < 100000 {
		return y * 100000
	} else if x < 1000000 {
		return y * 1000000
	} else if x < 10000000 {
		return y * 10000000
	} else if x < 100000000 {
		return y * 100000000
	} else if x < 1000000000 {
		return y * 1000000000
	} else if x < 10000000000 {
		return y * 10000000000
	} else if x < 100000000000 {
		return y * 100000000000
	} else if x < 1000000000000 {
		return y * 1000000000000
	} else if x < 10000000000000 {
		return y * 10000000000000
	} else if x < 100000000000000 {
		return y * 100000000000000
	} else if x < 1000000000000000 {
		return y * 1000000000000000
	} else if x < 10000000000000000 {
		return y * 10000000000000000
	} else if x < 100000000000000000 {
		return y * 100000000000000000
	} else if x < 1000000000000000000 {
		return y * 1000000000000000000
	}
	return -1
}

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
	return geta(b, a) + b
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

func part1() int64 {
	targets, allNums := processInput("input")
	var sum int64 = 0
	var wg sync.WaitGroup
	results := make(chan int64, len(targets))

	for i, target := range targets {
		wg.Add(1)
		go func(target int, nums []int) {
			defer wg.Done()
			if try(nums, target, nums[0], 1, false) {
				results <- int64(target)
			} else {
				results <- 0
			}
		}(target, allNums[i])
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		sum += result
	}
	return sum
}

func part2() int64 {
	targets, allNums := processInput("input")
	var sum int64 = 0
	var wg sync.WaitGroup
	results := make(chan int64, len(targets))

	for i, target := range targets {
		wg.Add(1)
		go func(target int, nums []int) {
			defer wg.Done()
			if try(nums, target, nums[0], 1, true) {
				results <- int64(target)
			} else {
				results <- 0
			}
		}(target, allNums[i])
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		sum += result
	}
	return sum
}

func main() {
	fmt.Println("part 1: ", part1())
	fmt.Println("part 2: ", part2())
}
