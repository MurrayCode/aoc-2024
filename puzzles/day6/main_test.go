package main

import "testing"

func TestPart1(t *testing.T) {
	input := `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

	exp := 41
	res := Part1(input)
	if res != exp {
		t.Errorf("Got %d, expected %d", res, exp)
	}
}
