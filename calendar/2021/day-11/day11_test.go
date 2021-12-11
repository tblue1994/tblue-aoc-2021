package main

import (
	"tblue-aoc-2021/utils/files"
	"testing"
)

func TestPart1SampleInput(t *testing.T) {
	input := files.ReadFile(11, 2021, "\n", true)
	want := 1656
	count := solvePart1(input)
	if count != want {
		t.Fatalf(`solvePart1(input) = %v, want match for %#v`, count, want)
	}
}

func TestPart2SampleInput(t *testing.T) {
	input := files.ReadFile(11, 2021, "\n", true)
	want := 195
	count := solvePart2(input)
	if count != want {
		t.Fatalf(`solvePart2(input) = %v, want match for %#v`, count, want)
	}
}
