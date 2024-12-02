package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	read, _ := os.ReadFile("input.txt")
	fmt.Println(Part2(string(read)))
}

func Part1(input string) int {
	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	side_a := []int{}
	side_b := []int{}
	dif := 0
	for scanner.Scan() {
		line := scanner.Text()
		a, _ := strconv.Atoi(strings.Split(line, "  ")[0])
		b, _ := strconv.Atoi(strings.Trim(strings.Split(line, "  ")[1], " "))
		side_a = append(side_a, a)
		side_b = append(side_b, b)
	}
	slices.Sort(side_a)
	slices.Sort(side_b)
	for i := 0; i < len(side_a); i++ {
		if side_a[i] < side_b[i] {
			dif += side_b[i] - side_a[i]
		} else if side_a[i] > side_b[i] {
			dif += side_a[i] - side_b[i]
		}
	}
	return dif
}

func Part2(input string) int {
	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	side_a := []int{}
	side_b := []int{}
	dict := make(map[int]int)
	ans := 0
	for scanner.Scan() {
		line := scanner.Text()
		a, _ := strconv.Atoi(strings.Split(line, "  ")[0])
		b, _ := strconv.Atoi(strings.Trim(strings.Split(line, "  ")[1], " "))
		side_a = append(side_a, a)
		side_b = append(side_b, b)
	}

	for _, item := range side_b {
		dict[item] = dict[item] + 1
	}

	for _, item := range side_a {
		if entry, ok := dict[item]; ok {
			ans += item * entry
		} else {
			ans += 0
		}
	}
	return ans
}
