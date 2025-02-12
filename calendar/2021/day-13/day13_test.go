package main

import (
	"tblue-aoc-2021/utils/files"
	"testing"
)

func TestPart1SampleInput(t *testing.T) {
	input := files.ReadFile(13, 2021, "\n", true)
	want := 17
	count := solvePart1(input)
	if count != want {
		t.Fatalf(`solvePart1(input) = %v, want match for %#v`, count, want)
	}
}

func TestPart2SampleInput(t *testing.T) {
	input := files.ReadFile(13, 2021, "\n", true)
	want := 0
	count := solvePart2(input)
	if count != want {
		t.Fatalf(`solvePart2(input) = %v, want match for %#v`, count, want)
	}
}
