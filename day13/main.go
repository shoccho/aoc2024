package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/shoccho/aoc2024/utils"
)

type pair struct {
	x, y int64
}
type machine struct {
	t pair
	a pair
	b pair
}

func processMachine(m machine) int64 {
	ax, bx, tx := m.a.x, m.b.x, m.t.x
	ay, by, ty := m.a.y, m.b.y, m.t.y
	b := (tx*ay - ty*ax) / (ay*bx - by*ax)
	a := (tx*by - ty*bx) / (by*ax - bx*ay)
	if ax*a+bx*b == tx && ay*a+by*b == ty {
		return 3*a + b
	}
	return 0
}

func parseInput(filename string, miscalc bool) []machine {
	machines := []machine{}
	lines := utils.ReadLines(filename)
	cm := machine{}
	for _, line := range lines {
		if line == "" {
			machines = append(machines, cm)
			cm = machine{}
			continue
		}
		parts := strings.Split(line, " ")
		if parts[0] == "Button" {
			xstr := parts[2][2 : len(parts[2])-1]
			ystr := parts[3][2:]
			x, err := strconv.Atoi(xstr)
			if err != nil {
				panic(err)
			}
			y, err := strconv.Atoi(ystr)
			if err != nil {
				panic(err)
			}
			if parts[1] == "A:" {
				cm.a.x = int64(x)
				cm.a.y = int64(y)
			} else if parts[1] == "B:" {
				cm.b.x = int64(x)
				cm.b.y = int64(y)
			}
		} else if parts[0] == "Prize:" {
			xstr := parts[1][2 : len(parts[1])-1]
			ystr := parts[2][2:]
			x, err := strconv.Atoi(xstr)
			if err != nil {
				panic(err)
			}
			y, err := strconv.Atoi(ystr)
			if err != nil {
				panic(err)
			}
			if miscalc {
				cm.t.x = int64(x) + 10000000000000
				cm.t.y = int64(y) + 10000000000000
			} else {
				cm.t.x = int64(x)
				cm.t.y = int64(y)
			}
		}
	}
	return machines
}
func part1(filename string) int64 {
	machines := parseInput(filename, false)
	var sum int64
	for _, m := range machines {
		sum += processMachine(m)
	}
	return sum
}

func part2(filename string) int64 {
	machines := parseInput(filename, true)
	var sum int64
	for _, m := range machines {
		sum += processMachine(m)
	}
	return sum
}

func main() {
	fmt.Println("part 1: ", part1("input"))
	fmt.Println("part 2: ", part2("input"))
}
