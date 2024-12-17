package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/shoccho/aoc2024/utils"
)

func part1(filename string) int64 {
	h := 103
	w := 101
	mh := 51
	mw := 50
	iters := 100
	lines := utils.ReadLines(filename)
	var q1, q2, q3, q4 int64

	for _, line := range lines {
		parts := strings.Split(line, " ")
		posParts := strings.Split(parts[0][2:], ",")
		vParts := strings.Split(parts[1][2:], ",")
		p1, err := strconv.Atoi(posParts[0])
		if err != nil {
			panic(err)
		}
		p2, err := strconv.Atoi(posParts[1])
		if err != nil {
			panic(err)
		}
		v1, err := strconv.Atoi(vParts[0])
		if err != nil {
			panic(err)
		}
		v2, err := strconv.Atoi(vParts[1])
		if err != nil {
			panic(err)
		}

		fpx := (p1 + (v1 * iters)) % (w)
		fpy := (p2 + (v2 * iters)) % (h)
		if fpx < 0 {
			fpx += w
		}
		if fpy < 0 {
			fpy += h
		}

		if fpx == mw || fpy == mh {
			continue
		}
		left := false
		if fpx < w/2 {
			left = true
		}
		top := false
		if fpy < h/2 {
			top = true
		}
		if top && left {
			q1++
		} else if top && !left {
			q2++
		} else if !top && left {
			q3++
		} else if !top && !left {
			q4++
		}
	}
	return int64(q1 * q2 * q3 * q4)
}

func part2(filename string) int64 {
	h := 103
	w := 101
	lines := utils.ReadLines(filename)
	type pair struct {
		x int
		y int
	}
	type machine struct {
		pos pair
		vel pair
	}

	machines := make([]machine, len(lines))

	for i, line := range lines {
		parts := strings.Split(line, " ")
		posParts := strings.Split(parts[0][2:], ",")
		vParts := strings.Split(parts[1][2:], ",")
		p1, err := strconv.Atoi(posParts[0])
		if err != nil {
			panic(err)
		}
		p2, err := strconv.Atoi(posParts[1])
		if err != nil {
			panic(err)
		}
		v1, err := strconv.Atoi(vParts[0])
		if err != nil {
			panic(err)
		}
		v2, err := strconv.Atoi(vParts[1])
		if err != nil {
			panic(err)
		}

		machines[i] = machine{pair{p1, p2}, pair{v1, v2}}
	}
	mxi := 0
	for t := 1; t <= h*w; t++ {
		next := make(map[pair]bool)
		count := 0
		for _, m := range machines {
			fpx := (m.pos.x + (m.vel.x * t)) % (w)
			fpy := (m.pos.y + (m.vel.y * t)) % (h)
			if fpx < 0 {
				fpx += w
			}
			if fpy < 0 {
				fpy += h
			}
			if next[pair{fpx, fpy}] {
				count++
			}
			for dx := -1; dx <= 1; dx++ {
				for dy := -1; dy <= 1; dy++ {
					next[pair{fpx + dx, fpy + dy}] = true
				}
			}
		}
		if count > 256 {
			return int64(t)
		}
	}

	return int64(mxi)

}

func main() {
	fmt.Println(part1("input"))
	fmt.Println(part2("input"))
}
