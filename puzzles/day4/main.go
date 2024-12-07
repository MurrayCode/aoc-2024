package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func checkCoords(x, y int, grid [][]rune) rune {
	if y < 0 || y >= len(grid) || x < 0 || x >= len(grid[y]) {
		return ' '
	}
	return grid[y][x]
}
func isXmas(x, y, dx, dy int, grid [][]rune) bool {
	return checkCoords(x, y, grid) == 'X' && checkCoords(x+dx, y+dy, grid) == 'M' && checkCoords(x+2*dx, y+2*dy, grid) == 'A' && checkCoords(x+3*dx, y+3*dy, grid) == 'S'
}

func isCrossMas(x, y int, grid [][]rune) bool {
	return checkCoords(x, y, grid) == 'A' && (checkCoords(x-1, y-1, grid) == 'M' && checkCoords(x+1, y-1, grid) == 'M' && checkCoords(x-1, y+1, grid) == 'S' && checkCoords(x+1, y+1, grid) == 'S' ||
		checkCoords(x-1, y+1, grid) == 'M' && checkCoords(x+1, y+1, grid) == 'M' && checkCoords(x-1, y-1, grid) == 'S' && checkCoords(x+1, y-1, grid) == 'S' ||
		checkCoords(x-1, y-1, grid) == 'M' && checkCoords(x-1, y+1, grid) == 'M' && checkCoords(x+1, y-1, grid) == 'S' && checkCoords(x+1, y+1, grid) == 'S' ||
		checkCoords(x+1, y-1, grid) == 'M' && checkCoords(x+1, y+1, grid) == 'M' && checkCoords(x-1, y-1, grid) == 'S' && checkCoords(x-1, y+1, grid) == 'S')
}

func Part1(input string) int {

	var grid [][]rune
	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	xmases := 0
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}
	for y, row := range grid {
		for x := range row {
			for _, dir := range [][]int{{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}} {
				if isXmas(x, y, dir[0], dir[1], grid) {
					xmases++
				}
			}
		}
	}
	return xmases
}
func Part2(input string) int {

	var grid [][]rune
	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	xmases := 0
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}
	for y, row := range grid {
		for x := range row {
			if isCrossMas(x, y, grid) {
				xmases++
			}
		}
	}
	return xmases
}

func main() {
	read, _ := os.ReadFile("input.txt")
	fmt.Println(Part1(string(read)))
	fmt.Println(Part2(string(read)))
}
