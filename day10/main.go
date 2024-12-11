package main

import (
	"fmt"

	"github.com/shoccho/aoc2024/utils"
)

func part1(input string) uint64 {
	var sum uint64 = 0
	lines := utils.ReadLines(input)
	matrix := make([][]byte, len(lines))
	for i, row := range lines {
		matrix[i] = []byte(row)
	}

	// Dimensions of the grid
	rowsCount := len(matrix)
	colsCount := len(matrix[0])

	// Directions for moving (up, down, left, right)
	directions := [][2]int{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
	}

	// Helper function for DFS
	var traverseFromStart func(int, int) int
	traverseFromStart = func(startRow, startCol int) int {
		visited := make(map[[2]int]bool)
		stack := [][3]int{{startRow, startCol, 0}} // (row, col, current_number)
		reachedNines := 0

		for len(stack) > 0 {
			// Pop the top of the stack
			element := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			r, c, currentNumber := element[0], element[1], element[2]

			// Check if we reached a 9
			if currentNumber == 9 {
				reachedNines++
				continue
			}

			// Visit neighbors
			for _, dir := range directions {
				nr, nc := r+dir[0], c+dir[1]
				if nr >= 0 && nr < rowsCount && nc >= 0 && nc < colsCount {
					if _, seen := visited[[2]int{nr, nc}]; !seen {
						nextNumber := currentNumber + 1
						if matrix[nr][nc] == byte('0'+nextNumber) {
							visited[[2]int{nr, nc}] = true
							stack = append(stack, [3]int{nr, nc, nextNumber})
						}
					}
				}
			}
		}
		return reachedNines
	}

	// Locate all starting positions of '0'
	// totalNines := 0
	for r := 0; r < rowsCount; r++ {
		for c := 0; c < colsCount; c++ {
			if matrix[r][c] == '0' {
				sum += uint64(traverseFromStart(r, c))
			}
		}
	}

	// return totalNines
	return sum
}

func part2(input string) uint64 {
	var sum uint64 = 0
	lines := utils.ReadLines(input)
	matrix := make([][]byte, len(lines))
	for i, row := range lines {
		matrix[i] = []byte(row)
	}
	// Dimensions of the grid
	rowsCount := len(matrix)
	colsCount := len(matrix[0])

	// Directions for moving (up, down, left, right)
	directions := [][2]int{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
	}

	// Helper function for DFS
	var traverseFromStart func(int, int) map[[2]int]int
	traverseFromStart = func(startRow, startCol int) map[[2]int]int {
		pathsToNines := make(map[[2]int]int)
		stack := [][4]int{{startRow, startCol, 0, 1}} // (row, col, current_number, path_count)
		visited := make(map[[2]int]bool)

		for len(stack) > 0 {
			// Pop the top of the stack
			element := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			r, c, currentNumber, pathCount := element[0], element[1], element[2], element[3]

			// If we reach a 9, record the number of paths to this 9
			if currentNumber == 9 {
				pathsToNines[[2]int{r, c}] += pathCount
				continue
			}

			// Visit neighbors
			for _, dir := range directions {
				nr, nc := r+dir[0], c+dir[1]
				if nr >= 0 && nr < rowsCount && nc >= 0 && nc < colsCount {
					if _, seen := visited[[2]int{nr, nc}]; !seen || matrix[nr][nc] == byte('0'+currentNumber+1) {
						nextNumber := currentNumber + 1
						if matrix[nr][nc] == byte('0'+nextNumber) {
							visited[[2]int{nr, nc}] = true
							stack = append(stack, [4]int{nr, nc, nextNumber, pathCount})
						}
					}
				}
			}
		}
		return pathsToNines
	}

	// Locate all starting positions of '0'
	// totalDistinctWays := 0
	for r := 0; r < rowsCount; r++ {
		for c := 0; c < colsCount; c++ {
			if matrix[r][c] == '0' {
				pathsToNines := traverseFromStart(r, c)
				for _, count := range pathsToNines {
					sum += uint64(count)
				}
			}
		}
	}

	return sum
}

func main() {
	fmt.Println(part1("input"))
	fmt.Println(part2("input"))
}
