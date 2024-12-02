package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := `7 6 4 2 1
85 85 82 81 79 76 74
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`
	exp := 2
	res := Part1(input)
	if res != exp {
		t.Errorf("Got %d, expected %d", res, exp)
	}
}
func TestPart2(t *testing.T) {
	input := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`
	exp := 4
	res := Part2(input)
	if res != exp {
		t.Errorf("Got %d, expected %d", res, exp)
	}
}

func TestChecker(t *testing.T) {
	test := []struct {
		input []int
		exp   bool
	}{
		{[]int{1, 2, 3, 4, 5}, true},
		{[]int{21, 22, 22, 23}, false},
		{[]int{1, 2, 3, 4, 4}, false},
		{[]int{1, 2, 5, 7, 9}, true},
	}
	for _, tc := range test {
		res := checker(tc.input)
		if res != tc.exp {
			t.Errorf("Got %t, expected %t", res, tc.exp)
		}
	}
}
