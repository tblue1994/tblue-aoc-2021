package main

import (
	"regexp"
	"strconv"
	"strings"
	"tblue-aoc-2021/utils/files"
)

type Fold struct {
	direction  string
	coordinate int
}

func main() {
	input := files.ReadFile(13, 2021, "\n", false)
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	paper, folds := parseInput(input)

	return len(foldPaper(paper, folds[:1]))
}

func solvePart2(input []string) int {
	paper, folds := parseInput(input)
	foldedPaper := foldPaper(paper, folds)
	printPaper(foldedPaper)
	return 0
}

func parseInput(input []string) (map[string]int, []Fold) {
	paper := map[string]int{}
	folds := []Fold{}
	re := regexp.MustCompile(`fold along (?P<Dir>x|y)=(?P<Coor>\d+)`)
	for _, val := range input {
		if strings.Contains(val, ",") {
			paper[val] = 1
		}
		if strings.Contains(val, "fold along") {
			matches := re.FindStringSubmatch(val)
			direction := matches[re.SubexpIndex("Dir")]
			coor := matches[re.SubexpIndex("Coor")]
			coorInt, _ := strconv.Atoi(coor)
			folds = append(folds, Fold{
				direction:  direction,
				coordinate: coorInt,
			})

		}
	}
	return paper, folds
}

func foldPaper(paper map[string]int, folds []Fold) map[string]int {
	foldedPaper := map[string]int{}
	for k, v := range paper {
		foldedPaper[k] = v
	}

	for _, fold := range folds {
		newFoldedPaper := map[string]int{}
		for k := range foldedPaper {
			coordinates := strings.Split(k, ",")
			key := ""
			if fold.direction == "x" {
				x, _ := strconv.Atoi(coordinates[0])
				if x > fold.coordinate {
					x = fold.coordinate - (x - fold.coordinate)
				}
				key = strconv.Itoa(x) + "," + coordinates[1]
			} else {
				y, _ := strconv.Atoi(coordinates[1])
				if y > fold.coordinate {
					y = fold.coordinate - (y - fold.coordinate)
				}
				key = coordinates[0] + "," + strconv.Itoa(y)
			}
			newFoldedPaper[key] = 1
		}
		foldedPaper = newFoldedPaper
	}

	return foldedPaper
}

func printPaper(paper map[string]int) {
	a := [][]string{}
	for k := range paper {
		coordinates := strings.Split(k, ",")
		x, _ := strconv.Atoi(coordinates[0])
		y, _ := strconv.Atoi(coordinates[1])
		println(x, ",", y)
		z := len(a)
		if z-1 < y {
			for i := z; i <= y; i++ {
				println("y")
				a = append(a, []string{})
			}
		}
		for i := range a {
			p := len(a[i])
			if p-1 < x {
				for j := p; j <= x; j++ {
					println("x")
					a[i] = append(a[i], ".")
				}
			}
		}
		a[y][x] = "#"
	}
	for _, b := range a {
		println(strings.Join(b, ""))
	}

}
