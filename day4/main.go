package main

import (
	"fmt"

	"github.com/shoccho/aoc2024/utils"
)

func part1() {
	lines := utils.ReadLines("input")
	count := 0
	n := len(lines)
	m := len(lines[0])

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if lines[i][j] == 'X' {
				if j < m-3 {
					if lines[i][j+1] == 'M' &&
						lines[i][j+2] == 'A' &&
						lines[i][j+3] == 'S' {
						count++
					}
				}
				if j >= 3 {
					if lines[i][j-1] == 'M' &&
						lines[i][j-2] == 'A' &&
						lines[i][j-3] == 'S' {
						count++
					}
				}
				if i < n-3 {
					if lines[i+1][j] == 'M' &&
						lines[i+2][j] == 'A' &&
						lines[i+3][j] == 'S' {
						count++
					}
				}
				if i >= 3 {
					if lines[i-1][j] == 'M' &&
						lines[i-2][j] == 'A' &&
						lines[i-3][j] == 'S' {
						count++
					}
				}
				if i >= 3 && j >= 3 {
					if lines[i-1][j-1] == 'M' &&
						lines[i-2][j-2] == 'A' &&
						lines[i-3][j-3] == 'S' {
						count++
					}
				}
				if i >= 3 && j < m-3 {
					if lines[i-1][j+1] == 'M' &&
						lines[i-2][j+2] == 'A' &&
						lines[i-3][j+3] == 'S' {
						count++
					}
				}
				if i < n-3 && j < m-3 {
					if lines[i+1][j+1] == 'M' &&
						lines[i+2][j+2] == 'A' &&
						lines[i+3][j+3] == 'S' {
						count++
					}
				}
				if i < n-3 && j >= 3 {
					if lines[i+1][j-1] == 'M' &&
						lines[i+2][j-2] == 'A' &&
						lines[i+3][j-3] == 'S' {
						count++
					}
				}
			}
		}
	}
	fmt.Println("part 1", count)
}

func part2() {
	lines := utils.ReadLines("input")
	count := 0
	n := len(lines)
	m := len(lines[0])

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if lines[i][j] == 'A' {
				if i >= 1 && j >= 1 && i < n-1 && j < m-1 {
					if lines[i-1][j-1] == 'M' &&
						lines[i-1][j+1] == 'M' &&
						lines[i+1][j+1] == 'S' &&
						lines[i+1][j-1] == 'S' {
						count++
					}
					if lines[i-1][j-1] == 'S' &&
						lines[i-1][j+1] == 'S' &&
						lines[i+1][j+1] == 'M' &&
						lines[i+1][j-1] == 'M' {
						count++
					}
					if lines[i-1][j-1] == 'S' &&
						lines[i-1][j+1] == 'M' &&
						lines[i+1][j+1] == 'M' &&
						lines[i+1][j-1] == 'S' {
						count++
					}
					if lines[i-1][j-1] == 'M' &&
						lines[i-1][j+1] == 'S' &&
						lines[i+1][j+1] == 'S' &&
						lines[i+1][j-1] == 'M' {
						count++
					}
				}
			}
		}
	}
	fmt.Println("part 2", count)
}

func main() {
	part1()
	part2()
}
