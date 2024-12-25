package main

import (
	"container/heap"
	"fmt"

	"github.com/shoccho/aoc2024/utils"
)

type pos struct {
	i int
	j int
}

type vstate struct {
	pos
	dir int
}
type state struct {
	adj  vstate
	time int
}

type PriorityQueue []*state

func (pq PriorityQueue) Len() int            { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool  { return pq[i].time < pq[j].time }
func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(*state)) }
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func part1(filename string) int {
	lines := utils.ReadLines(filename)
	n := len(lines)
	m := len(lines[0])
	grid := make([][]int, n)
	for i := range n {
		grid[i] = make([]int, m)
	}
	sp := pos{0, 0}
	ep := pos{0, 0}
	for i, line := range lines {
		for j, r := range line {
			if r == '#' {
				grid[i][j] = 1
			} else if r == 'E' {
				ep.i = i
				ep.j = j
			} else if r == 'S' {
				sp.i = i
				sp.j = j
			}
		}
	}

	dirs := []pos{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &state{vstate{pos: sp, dir: 0}, 0})
	visited := make(map[vstate]int)
	low := int(^uint(0) >> 1)
	for pq.Len() > 0 {
		cur := heap.Pop(pq).(*state)
		if cur.adj.pos == ep {
			low = cur.time
			break
		}
		for _, turn := range []int{-1, 0, 1} {
			newDir := (cur.adj.dir + turn + 4) % 4
			d := 1
			adj := cur.adj
			if turn != 0 {
				d = 1000
				adj.dir = newDir
			} else {
				cd := dirs[cur.adj.dir]
				if grid[cur.adj.i+cd.i][cur.adj.j+cd.j] == 1 {
					continue
				}
				adj.i += cd.i
				adj.j += cd.j
			}
			if cur.time+d < visited[adj] || visited[adj] == 0 {
				visited[cur.adj] = cur.time + d
				heap.Push(pq, &state{adj, visited[cur.adj]})
			}

		}
	}

	return low
}

func part2(filename string) int {
	lines := utils.ReadLines(filename)
	n := len(lines)
	m := len(lines[0])
	grid := make([][]int, n)
	for i := range n {
		grid[i] = make([]int, m)
	}
	sp := pos{0, 0}
	ep := pos{0, 0}
	for i, line := range lines {
		for j, r := range line {
			if r == '#' {
				grid[i][j] = 1
			} else if r == 'E' {
				ep.i = i
				ep.j = j
			} else if r == 'S' {
				sp.i = i
				sp.j = j
			}
		}
	}

	dirs := []pos{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &state{vstate{pos: sp, dir: 0}, 0})
	visited := make(map[vstate]int)
	from := make(map[vstate][]vstate)
	low := int(^uint(0) >> 1)
	ld := 0
	for pq.Len() > 0 {
		cur := heap.Pop(pq).(*state)
		if cur.adj.pos == ep {

			if low > cur.time {
				low = cur.time
				ld = cur.adj.dir
			}
			// break
			// continue
		}
		for _, turn := range []int{-1, 0, 1} {

			d := 1
			adj := cur.adj
			if turn != 0 {
				d = 1000
				newDir := (cur.adj.dir + turn + 4) % 4
				adj.dir = newDir
			} else {
				cd := dirs[adj.dir]
				if grid[cur.adj.i+cd.i][cur.adj.j+cd.j] == 1 {
					continue
				}
				adj.i += cd.i
				adj.j += cd.j
			}
			if cur.time+d < visited[adj] || visited[adj] == 0 {
				visited[adj] = cur.time + d
				heap.Push(pq, &state{adj, visited[adj]})
				// from[adj] = append(from[adj], cur.adj)
				from[adj] = []vstate{cur.adj}
			} else if cur.time+d <= visited[adj] {
				from[adj] = append(from[adj], cur.adj)
			}

		}
	}
	gnode := make(map[vstate]bool)
	var ist vstate
	ist.i = ep.i
	ist.j = ep.j
	ist.dir = ld
	st := []vstate{ist}
	for len(st) > 0 {
		cur := st[len(st)-1]
		st = st[0 : len(st)-1]
		for _, other := range from[cur] {
			if !gnode[other] {
				st = append(st, other)
				gnode[other] = true
			}
		}
	}

	fp := make(map[pos]bool)
	for k := range gnode {
		fp[k.pos] = true
	}

	return len(fp) + 1
}

func main() {
	fmt.Println("part 1 sample: ", part1("sample"))
	fmt.Println("part 1 input: ", part1("input"))

	fmt.Println("part 2 sample: ", part2("sample"))
	fmt.Println("part 2 input: ", part2("input"))
}
