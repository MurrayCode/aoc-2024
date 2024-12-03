package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`
	exp := 161
	res := Part1(input)
	if res != exp {
		t.Errorf("Got %d, expected %d", res, exp)
	}
}
func TestPart2(t *testing.T) {
	input := `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))don't()?mul(2,4)do()+mul(2,2)`
	exp := 52
	res := Part2(input)
	if res != exp {
		t.Errorf("Got %d, expected %d", res, exp)
	}
}
