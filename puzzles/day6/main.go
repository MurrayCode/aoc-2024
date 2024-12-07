package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	read, _ := os.ReadFile("input.txt")
	fmt.Println(Part1(string(read)))
}

func checkTurn(x, y int, grid [][]rune) bool {
	if grid[y][x] == '#' {
		return true
	}
	return false
}
func findStart(grid [][]rune) (int, int) {
	for y, row := range grid {
		for x := range row {
			if grid[y][x] == '^' {
				return y, x
			}
		}
	}
	return 0, 0
}

func Move(x, y int, dir string, grid [][]rune) [][]rune {
	nWin := y == 0
	sWin := y == len(grid)-1
	eWin := x == len(grid[y])-1
	wWin := x == 0
	grid[y][x] = 'X'
	if dir == "n" {
		if nWin {
			return grid
		}
		if checkTurn(x, y-1, grid) {
			return Move(x, y, "e", grid)
		} else {
			return Move(x, y-1, "n", grid)
		}
	} else if dir == "e" {
		if eWin {
			return grid
		}
		if checkTurn(x+1, y, grid) {
			return Move(x, y, "s", grid)
		} else {
			return Move(x+1, y, "e", grid)
		}
	} else if dir == "s" {
		if sWin {
			return grid
		}
		if checkTurn(x, y+1, grid) {
			return Move(x, y, "w", grid)
		} else {
			return Move(x, y+1, "s", grid)
		}
	} else if dir == "w" {
		if wWin {
			return grid
		}
		if checkTurn(x-1, y, grid) {
			return Move(x, y, "n", grid)
		} else {
			return Move(x-1, y, "w", grid)
		}
	}
	return grid
}

func Part1(input string) int {
	var grid [][]rune
	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}
	startY, startX := findStart(grid)
	res := Move(startX, startY, "n", grid)
	count := 0
	for _, row := range res {
		for _, r := range row {
			if r == 'X' {
				count++
			}
		}
	}
	return count
}
