package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	s "github.com/murraycode/aoc-2024/pkg/slices"
)

func main() {
	read, _ := os.ReadFile("input.txt")
	fmt.Println(Part1(string(read)))
	fmt.Println(Part2(string(read)))
}

func getRules(input string) []string {
	return strings.Split(input, "|")
}

func checkRule(rule []string, updates string) bool {
	arr := strings.Split(updates, ",")
	first := false
	containsFirst := s.Contains(arr, rule[0])
	containsSecond := s.Contains(arr, rule[1])
	if !containsFirst || !containsSecond {
		return true
	}
	if containsFirst && containsSecond {
		for _, item := range arr {
			if item == rule[0] {
				first = true
			} else if item == rule[1] {
				if !first {
					return false
				}
			}
		}
	}
	return true
}

func Part1(input string) int {
	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	var rules [][]string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		rules = append(rules, getRules(line))
	}
	updates := strings.Split(input, "\n\n")[1]
	arrUpdates := strings.Split(updates, "\n")
	res := 0
	for _, update := range arrUpdates {
		ok := true
		for _, rule := range rules {
			ok = checkRule(rule, update)
			if !ok {
				break
			}
		}
		if ok {
			res += s.GetMiddleValueConvertToInt(strings.Split(update, ","))
		}
	}
	return res
}

func convertSliceStringToInts(arr []string) []int {
	res := make([]int, len(arr))
	for i, item := range arr {
		res[i], _ = strconv.Atoi(item)
	}
	return res
}

func Part2(input string) int {
	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	var rules [][]string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		rules = append(rules, getRules(line))
	}
	updates := strings.Split(input, "\n\n")[1]
	arrUpdates := strings.Split(updates, "\n")
	res := 0
	brokenUpdates := [][]int{}
	for _, update := range arrUpdates {
		ok := true
		for _, rule := range rules {
			ok = checkRule(rule, update)
			if !ok {
				break
			}
		}
		if !ok {
			bu := strings.Split(update, ",")
			brokenUpdates = append(brokenUpdates, convertSliceStringToInts(bu))
		}

		var compare = func(l, r int) int {
			for _, rule := range rules {

				one, _ := strconv.Atoi(rule[0])
				two, _ := strconv.Atoi(rule[1])

				if one == r && two == l {
					return 1
				}
				if one == l && two == r {
					return -1
				}
			}
			return 0
		}
		for _, bu := range brokenUpdates {
			if slices.IsSortedFunc(bu, compare) {
				continue
			} else {
				slices.SortFunc(bu, compare)
				res += s.GetMiddleValue(bu)
			}
		}
	}
	return res
}
