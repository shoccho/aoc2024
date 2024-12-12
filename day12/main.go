package main

import (
	"fmt"

	"github.com/shoccho/aoc2024/utils"
)

type pair struct {
	a int
	b int
}

func (p pair) valid(n, m int) bool {
	return p.a >= 0 && p.b >= 0 && p.a < n && p.b < m
}

func part1(filename string) int {
	lines := utils.ReadLines(filename)
	n := len(lines)
	m := len(lines[0])
	mm := []pair{}
	visited := make([]bool, n*m)
	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if visited[i*m+j] {
				continue
			}
			q := []pair{}
			q = append(q, pair{i, j})
			tt := pair{}
			for len(q) > 0 {
				pos := q[len(q)-1]
				q = q[:len(q)-1]
				if visited[pos.a*m+pos.b] {

					continue
				}
				visited[pos.a*m+pos.b] = true
				tt.a++
				tt.b += 4
				for _, dir := range dirs {
					ni, nj := pos.a+dir[0], pos.b+dir[1]
					if ni >= 0 && ni < n && nj >= 0 && nj < m && lines[ni][nj] == lines[i][j] {
						if !visited[ni*m+nj] {
							q = append(q, pair{ni, nj})
						}
						tt.b--
					}
				}
			}
			mm = append(mm, tt)
		}
	}
	sum := 0
	for _, v := range mm {
		sum += v.a * v.b
	}
	return sum
}

func part2(filename string) int {
	lines := utils.ReadLines(filename)
	n := len(lines)
	m := len(lines[0])
	mm := []pair{}
	visited := make([]bool, n*m)
	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if visited[i*m+j] {
				continue
			}
			q := []pair{}
			q = append(q, pair{i, j})
			tt := pair{}
			for len(q) > 0 {
				pos := q[len(q)-1]
				q = q[:len(q)-1]
				if visited[pos.a*m+pos.b] {

					continue
				}
				visited[pos.a*m+pos.b] = true
				tt.a++

				above := pair{pos.a - 1, pos.b}
				// aboveRight := pair{pos.a - 1, pos.b + 1}
				right := pair{pos.a, pos.b + 1}
				if (!right.valid(n, m) || lines[right.a][right.b] != lines[pos.a][pos.b]) && (!above.valid(n, m) || lines[above.a][above.b] != lines[pos.a][pos.b]) {
					tt.b++
				}

				// above := pair{pos.a - 1, pos.b}
				// aboveRight := pair{pos.a - 1, pos.b + 1}
				left := pair{pos.a, pos.b - 1}
				if (!left.valid(n, m) || lines[left.a][left.b] != lines[pos.a][pos.b]) && (!above.valid(n, m) || lines[above.a][above.b] != lines[pos.a][pos.b]) {
					tt.b++
				}
				bottom := pair{pos.a + 1, pos.b}

				if (!right.valid(n, m) || lines[right.a][right.b] != lines[pos.a][pos.b]) && (!bottom.valid(n, m) || lines[bottom.a][bottom.b] != lines[pos.a][pos.b]) {
					tt.b++
				}
				if (!left.valid(n, m) || lines[left.a][left.b] != lines[pos.a][pos.b]) && (!bottom.valid(n, m) || lines[bottom.a][bottom.b] != lines[pos.a][pos.b]) {
					tt.b++
				}

				if above.valid(n, m) && right.valid(n, m) {
					if lines[above.a][above.b] == lines[pos.a][pos.b] && lines[right.a][right.b] == lines[pos.a][pos.b] && lines[above.a][right.b] != lines[pos.a][pos.b] {
						tt.b++
					}
				}
				if above.valid(n, m) && left.valid(n, m) {
					if lines[above.a][above.b] == lines[pos.a][pos.b] && lines[left.a][left.b] == lines[pos.a][pos.b] && lines[above.a][left.b] != lines[pos.a][pos.b] {
						tt.b++
					}
				}

				if bottom.valid(n, m) && right.valid(n, m) {
					if lines[bottom.a][bottom.b] == lines[pos.a][pos.b] && lines[right.a][right.b] == lines[pos.a][pos.b] && lines[bottom.a][right.b] != lines[pos.a][pos.b] {
						tt.b++
					}
				}
				if bottom.valid(n, m) && left.valid(n, m) {
					if lines[bottom.a][bottom.b] == lines[pos.a][pos.b] && lines[left.a][left.b] == lines[pos.a][pos.b] && lines[bottom.a][left.b] != lines[pos.a][pos.b] {
						tt.b++
					}
				}
				for _, dir := range dirs {
					ni, nj := pos.a+dir[0], pos.b+dir[1]
					if ni >= 0 && ni < n && nj >= 0 && nj < m && lines[ni][nj] == lines[i][j] {
						if !visited[ni*m+nj] {
							q = append(q, pair{ni, nj})
						}
						// tt.b--
					}
				}
			}
			mm = append(mm, tt)
		}
	}
	sum := 0
	for _, v := range mm {
		// fmt.Println(v.a, v.b)
		sum += v.a * v.b
	}
	return sum
}
func main() {
	// fmt.Println("part 1", part1("sample"))
	fmt.Println("part 1", part1("input"))
	// fmt.Println("part 2: ", part2("sample"))
	fmt.Println("part 2: ", part2("input"))
}
