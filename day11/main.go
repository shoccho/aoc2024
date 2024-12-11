package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/shoccho/aoc2024/utils"
)

func digCount(x uint64) int {
	c := 0
	for x > 0 {
		c++
		x /= 10
	}
	return c
}

func isEvenDigits(n uint64) bool {
	count := 0
	for n > 0 {
		count++
		n /= 10
	}
	return count%2 == 0
}

func splitN(n uint64) (uint64, uint64) {
	digits := 0
	temp := n
	for temp > 0 {
		digits++
		temp /= 10
	}

	half := digits / 2
	pow10 := uint64(math.Pow10(half))

	a := n % pow10
	b := n / pow10
	return a, b
}

func doDaTing(filename string, blinks int) uint64 {
	line := utils.ReadFile(filename)
	nums := strings.Split(line, " ")
	rocks := make(map[uint64]uint64)
	for _, v := range nums {
		number, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		rocks[uint64(number)] = 1
	}

	for i := 0; i < blinks; i++ {
		newRocks := make(map[uint64]uint64, len(rocks)*2)
		for rock, count := range rocks {
			if rock == 0 {
				newRocks[1] += count
			} else if isEvenDigits(uint64(rock)) {
				a, b := splitN(uint64(rock))
				newRocks[a] += count
				newRocks[b] += count
			} else {
				newRocks[uint64(rock*2024)] += count
			}
		}
		rocks = newRocks
	}

	var sum uint64 = 0
	for _, v := range rocks {
		sum += v
	}
	return sum
}

func part1(filename string) uint64 {
	return doDaTing(filename, 25)
}

func part2(filename string) uint64 {
	return doDaTing(filename, 75)
}

func main() {
	fmt.Println("part 1 : ", part1("input"))
	fmt.Println("part 2 : ", part2("input"))
}
