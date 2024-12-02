package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"

	s "github.com/murraycode/aoc-2024/pkg/slices"
)

func checker(nums []int) bool {
	safe := true
	if sort.IntsAreSorted(nums) {
		for i := 1; i < len(nums); i++ {
			if nums[i]-nums[i-1] > 3 || nums[i] == nums[i-1] {
				safe = false
				break
			}
		}
	} else {
		safe = false
	}
	return safe
}

func main() {
	read, _ := os.ReadFile("input.txt")
	fmt.Println(Part2(string(read)))
}
func Part2(input string) int {
	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	res := 0
	for scanner.Scan() {
		line := scanner.Text()
		temp := strings.Fields(line)
		nums := []int{}
		for _, item := range temp {
			num, _ := strconv.Atoi(item)
			nums = append(nums, num)
		}
		safe_a := true
		count := 0
		rev := append([]int{}, nums...)
		slices.Reverse(rev)
		for i := 0; i < len(nums); i++ {
			if !checker(nums) || !checker(rev) {
				temp_normal := append([]int{}, nums...)
				temp_reverse := append([]int{}, rev...)
				temp_normal = s.DeleteElement(temp_normal, i)
				temp_reverse = s.DeleteElement(temp_reverse, i)
				if checker(temp_normal) || checker(temp_reverse) {
					count++
				}
			}
		}
		if count != 0 {
			safe_a = false
		}
		if !safe_a {
			res++
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
		line := scanner.Text()
		temp := strings.Fields(line)
		nums := []int{}
		for _, item := range temp {
			num, _ := strconv.Atoi(item)
			nums = append(nums, num)
		}
		safe_a := true
		safe_b := true
		safe_a = checker(nums)
		slices.Reverse(nums)
		safe_b = checker(nums)
		if safe_a || safe_b {
			res++
		}

	}
	return res
}
