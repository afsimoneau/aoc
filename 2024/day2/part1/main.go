package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	reports := parseInput("input.txt")
	total := calculateSafeReports(reports)
	fmt.Printf("safe count: %d\n", total)
}

func parseInput(fileName string) [][]int {
	// open file
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file
	defer f.Close()

	// read the file line-by-line
	scanner := bufio.NewScanner(f)

	var reports [][]int
	lineNum := 0
	for scanner.Scan() {
		// get line
		line := scanner.Text()

		// split line
		splitLine := strings.Split(line, " ")

		var lineSlice []int
		for _, strVal := range splitLine {
			val, err := strconv.Atoi(strVal)
			if err != nil {
				log.Fatal(err)
			}
			lineSlice = append(lineSlice, val)
		}
		reports = append(reports, lineSlice)
		lineNum++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return reports
}

func calculateSafeReports(reports [][]int) int {
	safeCount := 0
	for _, report := range reports {
		//todo: calculate safety for a report
		// all increasing or all decreasing
		// each level differs by [1,2,3]

		isIncreasing := false
		isDecreasing := false
		for i, value := range report {
			if i < len(report)-1 {
				// current value and next element are within bounds

				diff := report[i+1] - value

				// bounds check
				if diff <= 3 && diff >= -3 && diff != 0 {
					// adjacent elements are within bounds
					if i == 0 {
						// 0'th element, set direction
						if diff > 0 {
							isIncreasing = true
						} else {
							isDecreasing = true
						}
					} else {
						// n'th element, break if diff is incorrect w/ direction
						if isIncreasing && diff < 0 {
							break
						}
						if isDecreasing && diff > 0 {
							break
						}
					}
				} else {
					// adjacent elements outside of bounds, unsafe
					break
				}
			} else {
				// current value is last element
				// made it to the end, count report as safe
				safeCount++
			}
		}
	}
	return safeCount
}
