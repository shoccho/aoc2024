package main

import (
	"fmt"
	"slices"

	"github.com/shoccho/aoc2024/utils"
)

func part1(filename string) int64 {
	line := utils.ReadFile(filename)
	nums := []int{}
	c := 0
	for i, n := range line {
		num := n - '0'

		if i%2 == 0 {
			for range num {
				nums = append(nums, c)
			}
			c++
		} else {
			for range num {
				nums = append(nums, -1)
			}
		}
	}
	di := 0
	for i, n := range nums {
		if n == -1 {
			di = i
			break
		}
	}
	n := len(nums)
	for i := n - 1; i > 0; i-- {
		if nums[i] == -1 {
			continue
		}
		if di >= i {
			break
		}
		nums[di] = nums[i]
		nums[i] = -1
		for j := di + 1; j <= i; j++ {
			if nums[j] == -1 {
				di = j
				break
			}
		}

	}
	var sum int64 = 0
	for i := 0; i < n; i++ {
		if nums[i] == -1 {
			continue
		}
		sum += int64(i * nums[i])
	}
	return sum
}

type file struct {
	value int
	len   int
}

func part2(filename string) uint64 {
	line := utils.ReadFile(filename)
	disk := []file{}
	c := 0
	for i, n := range line {
		num := int(n - '0')
		if i%2 == 0 {
			disk = append(disk, file{c, num})
			c++
		} else {
			disk = append(disk, file{-1, num})
		}
	}

	n := len(disk)

	done := []int{}
	for i := n - 1; i >= 0; i-- {
		if slices.Index(done, disk[i].value) != -1 {
			continue
		}
		done = append(done, disk[i].value)

		if disk[i].value == -1 {
			continue
		}
		for j := 0; j < i; j++ {

			if disk[j].value == -1 && disk[j].len == disk[i].len {
				disk[j], disk[i] = disk[i], disk[j]
				break
			} else if disk[j].value == -1 && disk[j].len > disk[i].len {
				diff := disk[j].len - disk[i].len
				disk[j].value = disk[i].value
				disk[j].len = disk[i].len
				disk[i].value = -1
				if disk[j+1].value == -1 {
					disk[j+1].len += diff
				} else {
					disk = slices.Insert(disk, j+1, file{-1, diff})
				}
				break
			}
		}
	}
	idx := 0
	var sum uint64 = 0
	idx = 0
	for i := 0; i < len(disk); i++ {
		if disk[i].value == -1 {
			idx += disk[i].len
			continue
		}
		for j := 0; j < disk[i].len; j++ {
			sum += uint64((idx) * disk[i].value)
			idx++
		}
	}
	return sum

}

func main() {
	fmt.Println(part1("input"))
	fmt.Println(part2("input"))
}
