package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

	exp := 18
	res := Part1(input)
	if res != exp {
		t.Errorf("got %v, expected %v", res, exp)
	}
}
func TestPart2(t *testing.T) {
	input := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

	exp := 9
	res := Part2(input)
	if res != exp {
		t.Errorf("got %v, expected %v", res, exp)
	}
}
