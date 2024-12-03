package main

import (
	"fmt"

	"github.com/shoccho/aoc2024/utils"
)

func part1() {
	str := utils.ReadFile("input")
	sum := 0
	n := len(str)
	for i := 0; i < n-3; i++ {
		if str[i] == 'm' && str[i+1] == 'u' && str[i+2] == 'l' && str[i+3] == '(' {
			first := -1
			second := -1
			tmp := 0
			i += 4
			ti := 0
			for i < n {
				if str[i] >= '0' && str[i] <= '9' {
					tmp = tmp*10 + int(str[i]-'0')
					ti++
					if ti > 3 {
						break
					}
				} else if str[i] == ',' && first == -1 {
					first = tmp
					tmp = 0
					ti = 0
				} else if str[i] == ')' && first != -1 {
					second = tmp
					tmp = 0
					break
				} else {
					break
				}
				i++
			}
			if first != -1 && second != -1 {
				sum += first * second
			}
		}
	}
	fmt.Println("part 1 ", sum)
}

func main() {
	part1()
}
