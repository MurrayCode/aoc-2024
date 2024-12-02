package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := `3   4
4   3
2   5
1   3
3   9
3   3`

	res := Part1(input)
	exp := 11
	if res != exp {
		t.Errorf("Got %d, expected %d", res, exp)
	}
}

func TestPart2(t *testing.T) {
	input := `3   4
4   3
2   5
1   3
3   9
3   3`

	res := Part2(input)
	exp := 31
	if res != exp {
		t.Errorf("Got %d, expected %d", res, exp)
	}
}
