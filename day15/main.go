package main

import (
	"fmt"
	"strings"

	"github.com/shoccho/aoc2024/utils"
)

type pair struct {
	i int
	j int
}

func getDir(r rune) (pair, bool) {
	plsLetMeMakeGlobalMapsGolang := map[rune]pair{
		'^': {-1, 0},
		'v': {1, 0},
		'<': {0, -1},
		'>': {0, 1},
	}
	p, ok := plsLetMeMakeGlobalMapsGolang[r]
	return p, ok
}

func part1(filename string) int64 {
	input := utils.ReadFile(filename)
	parts := strings.Split(input, utils.GetEndl()+utils.GetEndl())
	grid_input := strings.Split(parts[0], utils.GetEndl())
	ops := parts[1]
	n := len(grid_input)
	m := len(grid_input[0])
	grid := make([][]int, n)
	for i := range n {
		grid[i] = make([]int, m)
	}
	pi, pj := 0, 0
	for i := range n {
		for j := range m {
			if grid_input[i][j] == '#' {
				grid[i][j] = 1
			} else if grid_input[i][j] == 'O' {
				grid[i][j] = 2

			} else if grid_input[i][j] == '@' {
				pi = i
				pj = j
			}
		}
	}

	for _, op := range ops {

		dir, ok := getDir(op)
		if !ok {
			continue

		}
		if pi+dir.i >= 0 && pi+dir.i < n && pj+dir.j >= 0 && pj+dir.j < m && grid[pi+dir.i][pj+dir.j] != 1 {
			if grid[pi+dir.i][pj+dir.j] == 0 {
				pi += dir.i
				pj += dir.j
			} else {
				ti := pi + dir.i
				tj := pj + dir.j
				st := []pair{}

				visited := make(map[pair]bool)
				if grid[ti][tj] == 2 || grid[ti][tj] == 3 {

					st = append(st, pair{ti, tj})
					visited[pair{ti, tj}] = true

				}
				anyblocked := false
				sti := 0
				for sti < len(st) {
					top := st[sti]

					if grid[top.i+dir.i][top.j+dir.j] == 1 {
						anyblocked = true
						break
					} else if grid[top.i+dir.i][top.j+dir.j] >= 2 {
						if !visited[pair{top.i + dir.i, top.j + dir.j}] {
							st = append(st, pair{top.i + dir.i, top.j + dir.j})
							visited[pair{top.i + dir.i, top.j + dir.j}] = true
						}
					}
					sti++
				}

				if !anyblocked {
					for i := sti - 1; i >= 0; i-- {
						gp := st[i]
						v := grid[gp.i][gp.j]
						nv := pair{gp.i + dir.i, gp.j + dir.j}
						grid[nv.i][nv.j] = v
						grid[gp.i][gp.j] = 0
					}
					pi += dir.i
					pj += dir.j

				}

			}

		}
	}

	return totalDistance(grid)
}
func part2(filename string) int64 {
	input := utils.ReadFile(filename)
	parts := strings.Split(input, utils.GetEndl()+utils.GetEndl())
	grid_input := strings.Split(parts[0], utils.GetEndl())
	ops := parts[1]
	n := len(grid_input)
	m := len(grid_input[0])
	grid := make([][]int, n)
	for i := range n {
		grid[i] = make([]int, 2*m)
	}
	pi, pj := 0, 0
	for i := range n {
		for j := range m {
			if grid_input[i][j] == '#' {
				grid[i][j*2] = 1
				grid[i][1+(j*2)] = 1
			} else if grid_input[i][j] == 'O' {
				grid[i][j*2] = 2
				grid[i][1+(j*2)] = 3
			} else if grid_input[i][j] == '@' {
				pi = i
				pj = j * 2
			}
		}
	}
	m = m * 2

	for _, op := range ops {

		dir, ok := getDir(op)
		if !ok {
			continue
		}
		if pi+dir.i >= 0 && pi+dir.i < n && pj+dir.j >= 0 && pj+dir.j < m && grid[pi+dir.i][pj+dir.j] != 1 {
			if grid[pi+dir.i][pj+dir.j] == 0 {
				pi += dir.i
				pj += dir.j
			} else {
				ti := pi + dir.i
				tj := pj + dir.j
				st := []pair{}

				visited := make(map[pair]bool)
				if grid[ti][tj] == 2 || grid[ti][tj] == 3 {
					if grid[ti][tj] == 3 {
						st = append(st, pair{ti, tj})
						visited[pair{ti, tj}] = true
					} else {
						st = append(st, pair{ti, tj})
						visited[pair{ti, tj}] = true
					}
				}
				anyblocked := false
				sti := 0
				for sti < len(st) {
					top := st[sti]

					if grid[top.i][top.j] == 2 {
						if !visited[pair{top.i, top.j + 1}] {
							st = append(st, pair{top.i, top.j + 1})
							visited[pair{top.i, top.j + 1}] = true
						}
					}
					if grid[top.i][top.j] == 3 {
						if !visited[pair{top.i, top.j - 1}] {
							st = append(st, pair{top.i, top.j - 1})
							visited[pair{top.i, top.j - 1}] = true
						}
					}
					if grid[top.i+dir.i][top.j+dir.j] == 1 {
						anyblocked = true
						break
					} else if grid[top.i+dir.i][top.j+dir.j] >= 2 {
						if !visited[pair{top.i + dir.i, top.j + dir.j}] {
							st = append(st, pair{top.i + dir.i, top.j + dir.j})
							visited[pair{top.i + dir.i, top.j + dir.j}] = true
						}
					}
					sti++
				}

				if !anyblocked {
					for i := sti - 1; i >= 0; i-- {
						gp := st[i]
						v := grid[gp.i][gp.j]
						nv := pair{gp.i + dir.i, gp.j + dir.j}
						grid[nv.i][nv.j] = v
						grid[gp.i][gp.j] = 0
					}
					pi += dir.i
					pj += dir.j

				}

			}

		}
	}

	return totalDistance(grid)
}

func totalDistance(grid [][]int) int64 {
	var sum int64 = 0

	for i, g := range grid {
		for j, c := range g {
			if c == 2 {
				sum += int64((i * 100) + j)
			}
		}
	}
	return sum
}

func main() {

	fmt.Println("Part1 sample: ", part1("sample"))
	fmt.Println("Part1 input: ", part1("input"))

	fmt.Println("Part2 sample: ", part2("sample"))
	fmt.Println("Part2 input: ", part2("input"))
}
