package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/shoccho/aoc2024/utils"
)

func part1() {
	lines := utils.ReadLines("input")
	table := make([][]bool, 100)
	readingTable := true
	sum := 0
	for i := 0; i < 100; i++ {
		table[i] = make([]bool, 100)
	}
	for _, line := range lines {
		if readingTable {
			if line == "" {
				readingTable = false
			} else {
				parts := strings.Split(line, "|")
				before, e := strconv.Atoi(parts[0])
				if e != nil {
					panic(e)
				}
				after, e := strconv.Atoi(parts[1])
				if e != nil {
					panic(e)
				}
				table[after][before] = true
			}
		} else {
			parts := strings.Split(line, ",")
			n := len(parts)
			valid := true
			nums := make([]int, n)
			for i := 0; i < n; i++ {
				tmp, err := strconv.Atoi(parts[i])
				if err != nil {
					panic(err)
				}
				nums[i] = tmp
			}
			for i := n - 1; i >= 0; i-- {
				for j := 0; j < i; j++ {
					if !table[nums[i]][nums[j]] {
						valid = false
						break
					}
				}
			}
			if valid {
				sum += nums[n/2]
			}
		}
	}
	fmt.Println("part 1", sum)
}

func part2() {
	lines := utils.ReadLines("input")
	table := make([][]bool, 100)
	readingTable := true
	sum := 0
	for i := 0; i < 100; i++ {
		table[i] = make([]bool, 100)
	}
	invalidLines := [][]int{}
	for _, line := range lines {
		if readingTable {
			if line == "" {
				readingTable = false
			} else {
				parts := strings.Split(line, "|")
				before, e := strconv.Atoi(parts[0])
				if e != nil {
					panic(e)
				}
				after, e := strconv.Atoi(parts[1])
				if e != nil {
					panic(e)
				}
				table[after][before] = true
			}
		} else {
			parts := strings.Split(line, ",")
			n := len(parts)
			nums := make([]int, n)
			for i := 0; i < n; i++ {
				tmp, err := strconv.Atoi(parts[i])
				if err != nil {
					panic(err)
				}
				nums[i] = tmp
			}
			valid := true
			for i := n - 1; i >= 0; i-- {
				for j := 0; j < i; j++ {
					if !table[nums[i]][nums[j]] {
						valid = false
						break
					}
				}
			}
			if !valid {
				invalidLines = append(invalidLines, nums)
			}
		}
	}

	for _, nums := range invalidLines {
		n := len(nums)
		for {
			for i := 0; i < n; i++ {
				for j := i + 1; j < n; j++ {
					if !table[nums[j]][nums[i]] {
						nums[j], nums[i] = nums[i], nums[j]
					}
				}
			}

			valid := true
			for i := n - 1; i >= 0; i-- {
				for j := 0; j < i; j++ {
					if !table[nums[i]][nums[j]] {
						valid = false
						break
					}
				}
			}
			if valid {
				break
			}
		}
		sum += nums[n/2]
	}
	fmt.Println("part 2", sum)
}

func main() {
	part1()
	part2()
}
