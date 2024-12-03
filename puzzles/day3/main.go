package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	read, _ := os.ReadFile("input.txt")
	fmt.Println(Part1(string(read)))
}

func Part1(input string) int {
	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	res := 0
	for scanner.Scan() {
		r := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
		matches := r.FindAllStringSubmatch(scanner.Text(), -1)
		for _, v := range matches {
			val1, _ := strconv.Atoi(v[1])
			val2, _ := strconv.Atoi(v[2])
			res += val1 * val2
		}
	}
	return res
}
