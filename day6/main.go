package main

import (
	"fmt"

	"github.com/shoccho/aoc2024/utils"
)

func part1() {
	lines := utils.ReadLines("input")
	n := len(lines)
	m := len(lines[0])
	mappu := make([][]uint8, n)
	for i := 0; i < n; i++ {
		mappu[i] = make([]uint8, m)
	}
	sx, sy := 0, 0
	for i, line := range lines {
		for j, char := range line {
			if char == '#' {
				mappu[i][j] = 2
			} else if char == '^' {
				mappu[i][j] = 1
				sx = j
				sy = i
			}
		}
	}
	count := 0
	sc := 0
	cdir := "up"
	cx, cy := sx, sy
	outside := false
	for {
		if cdir == "up" {
			for cy >= 0 && mappu[cy][cx] != 2 {
				mappu[cy][cx] = 1
				cy--
				sc++
			}
			if cy < 0 {
				outside = true
				break
			} else {
				cdir = "right"
				cy++
			}
		} else if cdir == "down" {
			for cy < n && mappu[cy][cx] != 2 {
				mappu[cy][cx] = 1
				cy++
				sc++
			}
			if cy >= n {
				outside = true
				break
			} else {
				cdir = "left"
				cy--
			}
		} else if cdir == "left" {
			for cx >= 0 && mappu[cy][cx] != 2 {
				mappu[cy][cx] = 1
				cx--
				sc++
			}
			if cx < 0 {
				outside = true
				break
			} else {
				cdir = "up"
				cx++
			}
		} else if cdir == "right" {
			for cx < m && mappu[cy][cx] != 2 {
				mappu[cy][cx] = 1
				cx++
				sc++
			}
			if cx >= m {
				outside = true
				break
			} else {
				cdir = "down"
				cx--
			}
		}
		if sc > m*n || outside {
			break
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if mappu[i][j] == 1 {
				count++
			}
		}
	}
	fmt.Println("part 1:", count)
}

func part2() {
	lines := utils.ReadLines("input")
	n := len(lines)
	m := len(lines[0])
	mappu := make([][]uint8, n)
	for i := 0; i < n; i++ {
		mappu[i] = make([]uint8, m)
	}
	sx, sy := 0, 0
	for i, line := range lines {
		for j, char := range line {
			if char == '#' {
				mappu[i][j] = 2
			} else if char == '^' {
				mappu[i][j] = 1
				sx = j
				sy = i
			}
		}
	}
	count := 0

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if mappu[i][j] != 0 {
				continue
			}
			mappu[i][j] = 2
			sc := 0

			cx, cy := sx, sy
			cdir := "up"
			outside := false
			for {
				if cdir == "up" {
					for cy >= 0 && mappu[cy][cx] != 2 {
						cy--
						sc++
					}
					if cy < 0 {
						outside = true
						break
					} else {
						cdir = "right"
						cy++
					}
				} else if cdir == "down" {
					for cy < n && mappu[cy][cx] != 2 {
						cy++
						sc++
					}
					if cy >= n {
						outside = true
						break
					} else {
						cdir = "left"
						cy--
					}
				} else if cdir == "left" {
					for cx >= 0 && mappu[cy][cx] != 2 {
						cx--
						sc++
					}
					if cx < 0 {
						outside = true
						break
					} else {
						cdir = "up"
						cx++
					}
				} else if cdir == "right" {
					for cx < m && mappu[cy][cx] != 2 {
						cx++
						sc++
					}
					if cx >= m {
						outside = true
						break
					} else {
						cdir = "down"
						cx--
					}
				}
				if sc > m*n || outside {
					if sc > m*n {
						count++
					}
					break
				}
			}
			mappu[i][j] = 0
		}
	}

	fmt.Println("part 2:", count)
}

func main() {
	part1()
	part2()
}
