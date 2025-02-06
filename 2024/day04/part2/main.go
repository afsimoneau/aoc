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

func main() {

	rawText := parseInput("input.txt")
	grid := toGrid(rawText)
	total := shapeSearch(grid)

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

func shapeSearch(grid []string) int {
	total := 0

	for y, line := range grid {
		for x := range line {

			if evalX(grid, x, y) {
				total++
			}

		}
	}

	return total
}

func evalX(grid []string, x int, y int) bool {
	// out of bounds check

	if y-1 < 0 || y+1 > len(grid)-1 {
		return false
	}
	if x-1 < 0 || x+1 > len(grid[y])-1 {
		return false
	}

	//south-east mas
	seMas := evalPath(grid, x-1, y-1, coord{1, 1}, "MAS")

	seSam := evalPath(grid, x-1, y-1, coord{1, 1}, "SAM")

	swMas := evalPath(grid, x+1, y-1, coord{-1, 1}, "MAS")

	swSam := evalPath(grid, x+1, y-1, coord{-1, 1}, "SAM")

	
	// check if MAS or SAM (forward/backward) fits on both diagonals (southeast and southwest)
	if (seMas || seSam) && (swMas || swSam) {
		fmt.Printf("\n(%d, %d):\n%s\n%s\n%s\n", x, y, grid[y-1][x-1:x+2], grid[y][x-1:x+2], grid[y+1][x-1:x+2])
		fmt.Printf("seMas: %t\n", seMas)
		fmt.Printf("seSam: %t\n", seSam)
		fmt.Printf("swMas: %t\n", swMas)
		fmt.Printf("swSam: %t\n", swSam)
		return true
	}

	return false
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
