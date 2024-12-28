package main

import (
	"fmt"
	"math"

	"github.com/shoccho/aoc2024/utils"
)

func getStart(grid []string) (int, int) {
	si, sj := 0, 0
	for i, line := range grid {
		for j, char := range line {
			if char == 'S' {
				si, sj = i, j
				break
			}
		}
	}
	return si, sj
}

func makeDistMap(n, m int) [][]int {

	dists := make([][]int, n)
	for i := range n {
		dists[i] = make([]int, m)
		for j := range m {
			dists[i][j] = -1
		}
	}
	return dists
}

func buildDistMap(n, m, si, sj int, grid []string) [][]int {
	dists := makeDistMap(n, m)
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	ci, cj := si, sj
	dists[ci][cj] = 0
	for grid[ci][cj] != 'E' {
		for _, dir := range dirs {
			ni := ci + dir[0]
			nj := cj + dir[1]
			if ni < 0 || ni >= n || nj < 0 || nj >= m {
				continue
			}
			if grid[ni][nj] == '#' || dists[ni][nj] != -1 {
				continue
			}
			dists[ni][nj] = dists[ci][cj] + 1
			ci = ni
			cj = nj
		}
	}
	return dists
}

func part1(filename string) int {
	grid := utils.ReadLines(filename)
	n := len(grid)
	m := len(grid[0])
	si, sj := getStart(grid)

	count := 0

	dists := buildDistMap(n, m, si, sj, grid)

	cheatingDirs := [][]int{{2, 0}, {1, 1}, {0, 2}, {-1, 1}}

	for r := range n {
		for c := range m {
			if grid[r][c] == '#' {
				continue
			}
			for _, dir := range cheatingDirs {
				nr, nc := r+dir[0], c+dir[1]
				if nr < 0 || nr >= n || nc < 0 || nc >= m {
					continue
				}
				if grid[nr][nc] == '#' {
					continue
				}
				if math.Abs(float64(dists[r][c]-dists[nr][nc])) >= 102 {
					count++
				}
			}

		}
	}

	return count
}
func part2(filename string) int {
	grid := utils.ReadLines(filename)
	n := len(grid)
	m := len(grid[0])
	si, sj := getStart(grid)

	count := 0

	dists := buildDistMap(n, m, si, sj, grid)

	for r := range n {
		for c := range m {
			if grid[r][c] == '#' {
				continue
			}
			for radius := 2; radius < 21; radius++ {
				for dr := range radius + 1 {
					dc := radius - dr
					nps := [][]int{
						{r + dr, c + dc},
						{r + dr, c - dc},
						{r - dr, c + dc},
						{r - dr, c - dc},
					}
					visited := make([]bool, n*m)
					for _, dir := range nps {
						nr, nc := dir[0], dir[1]
						if nr < 0 || nr >= n || nc < 0 || nc >= m {
							continue
						}
						if grid[nr][nc] == '#' || visited[nr*m+nc] {
							continue
						}
						if dists[r][c]-dists[nr][nc] >= 100+radius {
							count++
						}
						visited[nr*m+nc] = true
					}

				}
			}
		}
	}
	return count
}

func main() {
	fmt.Println("part 1 input: ", part1("input"))
	fmt.Println("part 2 input: ", part2("input"))

}
