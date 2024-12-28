package main

import (
	"fmt"
	"strings"

	"github.com/shoccho/aoc2024/utils"
)

func howManyWays(list []string, target string) int {
	memo := make(map[string]int)

	var dfs func(target string) int
	dfs = func(target string) int {
		if target == "" {
			return 1
		}

		if val, found := memo[target]; found {
			return val
		}

		count := 0

		for _, s := range list {
			if len(target) >= len(s) && target[:len(s)] == s {
				count += dfs(target[len(s):])
			}
		}

		memo[target] = count

		return count
	}

	return dfs(target)
}

func part1(filename string) int {
	res := 0
	lines := utils.ReadLines(filename)
	have := strings.Split(lines[0], ", ")
	for i := 2; i < len(lines); i++ {
		if lines[i] == "" {
			continue
		}
		count := howManyWays(have, lines[i])
		if count > 0 {
			res++
		}
	}
	return res
}

func part2(filename string) int {
	res := 0
	lines := utils.ReadLines(filename)
	have := strings.Split(lines[0], ", ")
	for i := 2; i < len(lines); i++ {
		if lines[i] == "" {
			continue
		}
		p := howManyWays(have, lines[i])
		if p != -1 {
			res += p
		}
	}
	return res
}

func main() {
	fmt.Println("Part 1 sample: ", part1("sample"))
	fmt.Println("Part 1 input: ", part1("input"))

	fmt.Println("Part 2 sample: ", part2("sample"))
	fmt.Println("Part 2 input: ", part2("input"))
}
