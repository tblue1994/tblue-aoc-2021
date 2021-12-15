package main

import (
	"strings"
	"tblue-aoc-2021/utils/files"
)

func main() {
	input := files.ReadFile(14, 2021, "\n", false)
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	startString, chainMap := parseInput(input)
	endString := processStringForTimes(startString, chainMap, 10)
	for k, v := range endString {
		println(k, v)
	}
	return processEndString(endString)
}

func solvePart2(input []string) int {
	startString, chainMap := parseInput(input)
	endString := processStringForTimes(startString, chainMap, 40)
	return processEndString(endString)
}

func parseInput(input []string) (string, map[string]string) {
	startString := input[0]
	chainMap := map[string]string{}
	for i := 2; i < len(input); i++ {
		s := strings.Split(input[i], " -> ")
		chainMap[s[0]] = s[1]
	}
	return startString, chainMap
}

func processStringForTimes(startString string, chainMap map[string]string, times int) map[string]int {
	flatMap := map[string]int{}
	countMap := map[string]int{}
	for i, val := range startString {
		countMap[string(val)]++
		if i+1 != len(startString) {
			flatMap[string(val)+string(startString[i+1])]++
		}
	}
	for i := 0; i < times; i++ {
		copyFlatMap := map[string]int{}
		for k, v := range flatMap {
			newChar := chainMap[k]
			countMap[newChar] += v
			copyFlatMap[string(k[0])+newChar] += v
			copyFlatMap[newChar+string(k[1])] += v
		}
		flatMap = copyFlatMap
	}
	return countMap
}

func processEndString(endStringMap map[string]int) int {
	max := 0
	min := -1
	for _, v := range endStringMap {
		if v > max {
			max = v
		}
		if v < min || min == -1 {
			min = v
		}
	}
	return max - min
}
