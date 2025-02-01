package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// coordinate pair
type coord struct {
	x int
	y int
}

var stepDirection = []coord{
	{-1, -1}, {0, -1}, {1, -1},
	{-1, 0}, {1, 0},
	{-1, 1}, {0, 1}, {1, 1}}

func main() {
	
	rawText := parseInput("input.txt")
	grid := toGrid(rawText)
	total := wordSearch(grid, "XMAS")

	// I got this one first try!!! 🥹
	fmt.Println(total)
}

func parseInput(fileName string) string {
	// open file
	f, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	return string(f)
}

func toGrid(rawText string) []string {
	return strings.Split(rawText, "\n")
}

func wordSearch(grid []string, word string) int {
	total := 0

	for y, line := range grid {
		for x := range line {
			for _, step := range stepDirection {
				if evalPath(grid, x, y, step, word) {
					total++
				}
			}
		}
	}

	return total
}

// evaluate if the word is along the path executed by taking each step
func evalPath(grid []string, xStart int, yStart int, step coord, word string) bool {
	yEnd := yStart + (step.y * (len(word) - 1))
	if yEnd < 0 || yEnd > len(grid)-1 {
		// OOB check for Y
		return false
	}

	xEnd := xStart + (step.x * (len(word) - 1))
	if xEnd < 0 || xEnd > len(grid[yEnd])-1 {
		// OOB check for X
		return false
	}

	for i, w := range word {
		x := xStart + (step.x * i)
		y := yStart + (step.y * i)

		if rune(grid[y][x]) != w {
			// ran into character that isn't in the search word, fail fast
			return false
		}
	}
	return true
}
