package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/shoccho/aoc2024/utils"
)

type pos struct {
	i  int
	j  int
	sc int
}

func bfs(grid [][]bool, n int) int {
	dirs := []pos{{-1, 0, 0}, {1, 0, 0}, {0, -1, 0}, {0, 1, 0}}
	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dist[i][j] = math.MaxInt
		}
	}

	dist[0][0] = 0
	queue := []pos{{0, 0, 0}}

	for len(queue) > 0 {

		curr := queue[0]
		queue = queue[1:]
		for _, dir := range dirs {
			ni, nj := curr.i+dir.i, curr.j+dir.j
			if ni >= 0 && ni < n && nj >= 0 && nj < n && !grid[ni][nj] {
				if curr.sc+1 < dist[ni][nj] {
					dist[ni][nj] = curr.sc + 1
					queue = append(queue, pos{ni, nj, dist[ni][nj]})
				}
			}
		}
	}

	return dist[n-1][n-1]
}

func part1(filename string, n, xl int) int {

	lines := utils.ReadLines(filename)
	blocked := make([][]bool, n)
	for i := range n {
		blocked[i] = make([]bool, n)
	}
	for x, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, ",")
		j, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		i, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		blocked[i][j] = true
		if x == xl {
			break
		}
	}

	return bfs(blocked, n)
}

func part2(filename string, n int) (int, int) {

	lines := utils.ReadLines(filename)
	blocked := make([][]bool, n)
	for i := range n {
		blocked[i] = make([]bool, n)
	}
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, ",")
		j, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		i, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		blocked[i][j] = true
		ret := bfs(blocked, n)
		if ret == math.MaxInt {
			return i, j
		}
	}

	return -1, -1
}

func main() {
	fmt.Println("part 1 sample: ", part1("sample", 7, 11))
	fmt.Println("part 1 input: ", part1("input", 71, 1023))

	p2x, p2y := part2("sample", 7)
	fmt.Println("part 2 sample: ", p2x, p2y)
	p2x, p2y = part2("input", 71)
	fmt.Println("part 2 input: ", p2x, p2y)
}
