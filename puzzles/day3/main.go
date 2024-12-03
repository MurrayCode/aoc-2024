package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	read, _ := os.ReadFile("input.txt")
	fmt.Println(Part1(string(read)))
	fmt.Println(Part2(string(read)))
}

func Part2(input string) int {
	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	res := 0
	do := true
	for scanner.Scan() {
		r := regexp.MustCompile("mul\\(\\d+,\\d+\\)|do\\(\\)|don't\\(\\)")
		for _, v := range r.FindAllString(scanner.Text(), -1) {
			if v[0:3] == "mul" {
				var val1, val2 int
				fmt.Sscanf(v, "mul(%d,%d)", &val1, &val2)
				if do {
					res += val1 * val2
				}
			} else {
				do = v == "do()"
			}
		}
	}
	return res
}

func Part1(input string) int {
	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	res := 0
	for scanner.Scan() {
		r := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
		matches := r.FindAllString(scanner.Text(), -1)
		for _, v := range matches {
			var val1, val2 int
			fmt.Sscanf(v, "mul(%d,%d)", &val1, &val2)
			res += val1 * val2
		}
	}
	return res
}
