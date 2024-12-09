package main

import (
	"fmt"

	"github.com/shoccho/aoc2024/utils"
)

type pos struct {
	i int
	j int
}

func part1(filename string) int64 {
	lines := utils.ReadLines(filename)
	n := len(lines)
	m := len(lines[n-1])
	var sum int64
	layer := make([][]bool, n)
	for i := range n {
		layer[i] = make([]bool, m)
	}
	mappu := make(map[rune][]pos)
	for i, line := range lines {
		for j, r := range line {
			if r != '.' {
				mappu[r] = append(mappu[r], pos{i, j})
			}
		}
	}
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	for _, positions := range mappu {
		plen := len(positions)
		for i := 0; i < plen; i++ {
			anchor := positions[i]
			for j := i + 1; j < plen; j++ {
				point := positions[j]
				idif := abs(point.i - anchor.i)
				jdif := abs(point.j - anchor.j)
				var np1, np2 pos
				if point.j < anchor.j {
					np1.j = point.j - jdif
					np2.j = anchor.j + jdif
				} else {
					np1.j = point.j + jdif
					np2.j = anchor.j - jdif
				}
				if point.i < anchor.i {
					np1.i = point.i - idif
					np2.i = anchor.i + idif
				} else {
					np1.i = point.i + idif
					np2.i = anchor.i - idif
				}
				if np1.i >= 0 && np1.i < n && np1.j >= 0 && np1.j < m {
					layer[np1.i][np1.j] = true
				}
				if np2.i >= 0 && np2.i < n && np2.j >= 0 && np2.j < m {
					layer[np2.i][np2.j] = true
				}
			}
		}
	}
	for i := range n {
		for j := range m {
			if layer[i][j] {
				sum++
			}
		}
	}
	return sum
}

func part2(filename string) int64 {
	lines := utils.ReadLines(filename)
	n := len(lines)
	m := len(lines[n-1])
	var sum int64
	layer := make([][]bool, n)
	for i := range n {
		layer[i] = make([]bool, m)
	}
	mappu := make(map[rune][]pos)
	for i, line := range lines {
		for j, r := range line {
			if r != '.' {
				mappu[r] = append(mappu[r], pos{i, j})
			}
		}
	}
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	for _, positions := range mappu {
		plen := len(positions)
		for i := 0; i < plen; i++ {
			anchor := positions[i]
			if plen > 1 {
				layer[anchor.i][anchor.j] = true
			}
			for j := i + 1; j < plen; j++ {
				point := positions[j]
				idif := abs(point.i - anchor.i)
				jdif := abs(point.j - anchor.j)
				var pp1, pp2 pos
				pp1 = point
				pp2 = anchor

				for {
					var np1, np2 pos
					if point.j < pp2.j {
						np1.j = pp1.j - jdif
						np2.j = pp2.j + jdif
					} else {
						np1.j = pp1.j + jdif
						np2.j = pp2.j - jdif
					}
					if pp1.i < pp2.i {
						np1.i = pp1.i - idif
						np2.i = pp2.i + idif
					} else {
						np1.i = pp1.i + idif
						np2.i = pp2.i - idif
					}

					any_valid := false
					if np1.i >= 0 && np1.i < n && np1.j >= 0 && np1.j < m {
						layer[np1.i][np1.j] = true
						any_valid = true
					}
					if np2.i >= 0 && np2.i < n && np2.j >= 0 && np2.j < m {
						layer[np2.i][np2.j] = true
						any_valid = true
					}
					pp1 = np1
					pp2 = np2
					if !any_valid {
						break
					}
				}
			}
		}
	}
	for i := range n {
		for j := range m {
			if layer[i][j] {
				sum++
			}
		}
	}
	return sum
}

func main() {
	fmt.Println(part1("input"))
	fmt.Println(part2("input"))
}
