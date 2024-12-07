package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	read, _ := os.ReadFile("input.txt")
	fmt.Println(Part1(string(read)))
}

func getRules(input string) []string {
	return strings.Split(input, "|")
}

func contains(arr []string, item string) bool {
	for _, i := range arr {
		if i == item {
			return true
		}
	}
	return false
}

func checkRule(rule []string, updates string) bool {
	arr := strings.Split(updates, ",")
	first := false
	containsFirst := contains(arr, rule[0])
	containsSecond := contains(arr, rule[1])
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

func getMiddleValue(arr []string) int {
	m := len(arr) / 2
	i, _ := strconv.Atoi(arr[m])
	return i
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
			res += getMiddleValue(strings.Split(update, ","))
		}
	}
	return res
}
