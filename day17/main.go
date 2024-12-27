package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/shoccho/aoc2024/utils"
)

type machine struct {
	a int
	b int
	c int

	instructions []int
	ip           int
	outputBuffer []int
}

func (m *machine) run() {
	for m.ip < len(m.instructions)-1 {
		opcode := m.instructions[m.ip]
		m.execute(opcode, m.instructions[m.ip+1])
		if opcode != 3 {
			m.ip += 2
		}
	}
}

func (m *machine) execute(opcode, operand int) {
	combo := operand
	if operand > 3 {
		if operand == 4 {
			combo = m.a
		} else if operand == 5 {
			combo = m.b
		} else if operand == 6 {
			combo = m.c
		}
	}

	switch opcode {
	case 0:
		//adv
		// fmt.Println("calling adv")
		m.a /= int(math.Pow(2, float64(combo)))
	case 1:
		//bxl
		// fmt.Println("calling bxl")
		m.b = m.b ^ operand

	case 2:
		//bst
		// fmt.Println("calling bst")
		m.b = combo % 8
	case 3:
		//jnz
		// fmt.Println("calling jnz")
		if m.a != 0 {
			m.ip = operand
		} else {
			m.ip += 2
		}
	case 4:
		//bxc
		// fmt.Println("calling bxc")
		m.b = m.b ^ m.c
	case 5:
		//out
		// fmt.Println("calling out")
		m.outputBuffer = append(m.outputBuffer, combo%8)
	case 6:
		//bdv
		// fmt.Println("calling bdv")
		m.b = m.a / int(math.Pow(2, float64(combo)))
	case 7:
		//cdv
		// fmt.Println("calling cdv")
		m.c = m.a / int(math.Pow(2, float64(combo)))
	default:
		/* code */
		return
	}
}

func createMachine(filename string) machine {
	lines := utils.ReadLines(filename)
	var m machine
	for _, line := range lines {
		if line == "" {
			continue
		}
		if line[0] == 'R' {
			parts := strings.Split(line, " ")
			regName := parts[1]
			num, err := strconv.Atoi(parts[2])
			if err != nil {
				panic(err)
			}
			if regName == "A:" {
				m.a = num
			} else if regName == "B:" {
				m.b = num
			} else if regName == "C:" {
				m.c = num
			}
		} else if line[0] == 'P' {
			parts := strings.Split(line, " ")
			nums := strings.Split(parts[1], ",")
			for _, n := range nums {
				num, err := strconv.Atoi(n)
				if err != nil {
					panic(err)
				}
				m.instructions = append(m.instructions, num)
			}
		}
	}
	return m
}

func part1(filename string) {
	m := createMachine(filename)
	m.run()
	for i := 0; i < len(m.outputBuffer); i++ {
		if i > 0 {
			fmt.Print(",")
		}
		fmt.Print(m.outputBuffer[i])

	}
	fmt.Println()

}

func part2(filename string) {
	computer := createMachine(filename)
	program := computer.instructions
	var findA func(a, ci int, possible []int) int

	findA = func(a, ci int, possible []int) int {
		for n := range 8 {
			a2 := (a << 3) | n

			var cp machine
			cp.b = computer.b
			cp.c = computer.c
			cp.a = a2
			cp.instructions = program
			cp.run()
			output := cp.outputBuffer
			if slices.Equal(output, program[len(program)-ci:]) {
				if slices.Equal(output, program) {
					possible = append(possible, a2)
				} else {
					p := findA(a2, ci+1, possible)
					if p > 0 {
						possible = append(possible, p)
					}
				}
			}
		}
		if len(possible) > 0 {
			low := math.MaxInt
			for _, r := range possible {
				if low > r {
					low = r
				}
			}
			return low
		}
		return 0
	}

	possible := []int{}
	fmt.Println(findA(0, 1, possible))
}

func test() {
	{
		var x machine
		x.c = 9
		x.instructions = []int{2, 6}
		x.run()
		if x.b != 1 {
			panic("wrong in test 1 ")
		}
	}
	{
		var x machine
		x.a = 10
		x.instructions = []int{5, 0, 5, 1, 5, 4}
		x.run()
		if !slices.Equal(x.outputBuffer, []int{0, 1, 2}) {
			panic("test2")
		}
	}
	{
		var x machine
		x.a = 2024
		x.instructions = []int{0, 1, 5, 4, 3, 0}
		x.run()
		if !slices.Equal(x.outputBuffer, []int{4, 2, 5, 6, 7, 7, 7, 7, 3, 1, 0}) && x.a != 0 {
			panic("test3")
		}
	}
	{
		var x machine
		x.b = 29
		x.instructions = []int{1, 7}
		x.run()
		if x.b != 26 {
			panic("test4")
		}
	}
	{
		var x machine
		x.b = 2024
		x.c = 43690
		x.instructions = []int{4, 0}
		x.run()
		if x.b != 44354 {
			panic("test5")
		}
	}

}

func main() {
	test()
	part1("sample")
	part1("input")
	part2("sample")
	part2("input")

}
