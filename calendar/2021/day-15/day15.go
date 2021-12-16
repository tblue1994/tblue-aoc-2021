package main

import (
	"container/heap"
	"strconv"
	"strings"
	"tblue-aoc-2021/utils/files"
	"time"
)

// An Item is something we manage in a priority queue.
type Item struct {
	value    string // The value of the item; arbitrary.
	priority int    // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func main() {
	input := files.ReadFile(15, 2021, "\n", false)
	println(solvePart1(input))
	println(time.Now().String())
	println(solvePart2(input))
	println(time.Now().String())
}

func solvePart1(input []string) int {
	maze := parseInput(input)

	return findShortestPath(maze)
}

func solvePart2(input []string) int {
	maze := parseInput(input)
	biggerMaze := buildBiggerMaze(maze)

	return findShortestPathPQ(biggerMaze)
}

func parseInput(input []string) [][]int {
	maze := [][]int{}
	for _, val := range input {
		row := []int{}
		for _, r := range val {
			xstr := string(r) //x is rune converted to string
			xint, _ := strconv.Atoi(xstr)
			row = append(row, xint)
		}
		maze = append(maze, row)
	}
	return maze
}

func findShortestPath(maze [][]int) int {
	unvisted := map[string]int{}
	for y := range maze {
		for x := range maze[y] {
			unvisted[strconv.Itoa(x)+","+strconv.Itoa(y)] = -1
		}
	}

	cX, cY, d := 0, 0, 0

	for cX != len(maze[0])-1 || cY != len(maze)-1 {
		neighbors := [][]int{
			{cX - 1, cY},
			{cX + 1, cY},
			{cX, cY + 1},
			{cX, cY - 1},
		}
		for _, pair := range neighbors {
			nX, nY := pair[0], pair[1]
			if nX >= 0 && nX < len(maze[0]) && nY >= 0 && nY < len(maze) {
				w := maze[nY][nX]
				key := strconv.Itoa(nX) + "," + strconv.Itoa(nY)
				curVal := unvisted[key]
				if curVal == -1 || d+w < curVal {
					unvisted[key] = d + w
				}
			}
		}
		delete(unvisted, strconv.Itoa(cX)+","+strconv.Itoa(cY))
		lowest := -1
		nX, nY := 0, 0
		for k, v := range unvisted {
			if v >= 0 && (lowest == -1 || v < lowest) {
				pair := strings.Split(k, ",")
				x, _ := strconv.Atoi(pair[0])
				y, _ := strconv.Atoi(pair[1])
				nX, nY, lowest = x, y, v
			}
		}
		cX, cY, d = nX, nY, lowest
	}

	return d
}

func findShortestPathPQ(maze [][]int) int {
	seen := map[int]map[int]bool{}
	pq := PriorityQueue{}
	startItem := &Item{
		value:    "0,0",
		priority: 0,
	}
	pq.Push(startItem)
	heap.Init(&pq)

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		pair := strings.Split(item.value, ",")
		cx, _ := strconv.Atoi(pair[0])
		cy, _ := strconv.Atoi(pair[1])
		if cx == len(maze[0])-1 && cy == len(maze)-1 {
			return item.priority
		}
		if seen[cx] == nil {
			seen[cx] = map[int]bool{}
		}
		seen[cx][cy] = true

		neighbors := [][]int{
			{cx - 1, cy},
			{cx + 1, cy},
			{cx, cy + 1},
			{cx, cy - 1},
		}
		for _, pair := range neighbors {
			nX, nY := pair[0], pair[1]
			if _, found := seen[nX][nY]; found {
				continue
			}
			if nX >= 0 && nX < len(maze[0]) && nY >= 0 && nY < len(maze) {
				w := maze[nY][nX]
				key := strconv.Itoa(nX) + "," + strconv.Itoa(nY)
				found := false
				for i := range pq {
					if pq[i].value == key {
						foundItem := pq[i]
						if foundItem.priority > item.priority+w {
							pq.update(foundItem, foundItem.value, item.priority+w)
						}
						break
					}
				}
				if !found {
					nItem := &Item{
						value:    key,
						priority: item.priority + w,
					}
					heap.Push(&pq, nItem)
				}
			}
		}
	}
	return 0
}

func buildBiggerMaze(maze [][]int) [][]int {
	allSets := [][][]int{maze}
	for i := 0; i < 9; i++ {
		newMaze := [][]int{}
		for _, arr := range allSets[i] {
			newRow := []int{}
			for _, val := range arr {
				newVal := val
				if val == 9 {
					newVal = 0
				}
				newRow = append(newRow, newVal+1)
			}
			newMaze = append(newMaze, newRow)
		}
		allSets = append(allSets, newMaze)
	}
	bigMaze := [][]int{}
	for i := 0; i < 5; i++ {
		for n := range allSets[0] {
			newRow := []int{}
			newRow = append(newRow, allSets[i][n]...)
			newRow = append(newRow, allSets[i+1][n]...)
			newRow = append(newRow, allSets[i+2][n]...)
			newRow = append(newRow, allSets[i+3][n]...)
			newRow = append(newRow, allSets[i+4][n]...)
			bigMaze = append(bigMaze, newRow)
		}
	}
	return bigMaze
}
