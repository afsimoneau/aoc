package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	left, right := parseInput("input.txt")
	total := calculateTotal(left, right)
	fmt.Printf("total: %d\n", total)
}

func parseInput(fileName string) ([]int, []int) {
	// open file
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file
	defer f.Close()

	// read the file line-by-line
	scanner := bufio.NewScanner(f)

	var left []int
	var right []int
	for scanner.Scan() {
		// get line
		line := scanner.Text()

		// split line
		splitLine := strings.Split(line, "   ")

		// parse left
		pLeft, err := strconv.Atoi(splitLine[0])
		if err != nil {
			log.Fatal(err)
		}
		left = append(left, pLeft)

		// parse pRight
		pRight, err := strconv.Atoi(splitLine[1])
		if err != nil {
			log.Fatal(err)
		}
		right = append(right, pRight)
	}

	// verify size
	if len(left) != len(right) {
		log.Fatal("parsed slices differ in size")
	}

	// in-place sort
	sort.Ints(left)
	sort.Ints(right)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return left, right
}

func calculateTotal(leftSlice []int, rightSlice []int) int {
	total := 0

	for _, left := range leftSlice {
		// number of times left appears in right
		count := 0
		for _, right := range rightSlice {
			if left == right {
				count++
			}
			if right > left {
				break
			}
		}
		total += left * count

	}

	return total
}
