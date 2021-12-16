package main

import (
	"strconv"
	"strings"
	"tblue-aoc-2021/utils/files"
)

func main() {
	input := files.ReadFile(15, 2021, "\n", false)
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	maze := parseInput(input)

	return findShortestPath(maze)
}

func solvePart2(input []string) int {
	maze := parseInput(input)
	biggerMaze := buildBiggerMaze(maze)

	return findShortestPath(biggerMaze)
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
