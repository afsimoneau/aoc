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
	safePart1, unsafePart1 := sortReportsPart1(reports)
	safePart2, _ := sortReportsPart2(unsafePart1)
	fmt.Printf("total safe count: %d\n", len(safePart1)+len(safePart2))
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

func checkReport(report []int) bool {
	isIncreasing := false
	isDecreasing := false
	for i, value := range report {
		if i < len(report)-1 {
			// current value and next element are within bounds

			diff := report[i+1] - value

			// bounds check
			if diff > 3 || diff < -3 || diff == 0 {
				// adjacent elements outside of bounds
				return false
			}

			// within bounds, check direction of report
			if i == 0 {
				// 0'th element, set direction
				if diff > 0 {
					isIncreasing = true
				} else {
					isDecreasing = true
				}
			} else {
				// n'th element, check diff against direction
				if isIncreasing && diff < 0 {
					return false
				}
				if isDecreasing && diff > 0 {
					return false
				}
			}

			if diff <= 3 && diff >= -3 && diff != 0 {
				// adjacent elements are within bounds

			} else {
				// adjacent elements outside of bounds, unsafe
				return false
			}
		}
	}
	return true
}

func sortReportsPart1(reports [][]int) (safe [][]int, unsafe [][]int) {
	safeReports := [][]int{}
	unsafeReports := [][]int{}
	for _, report := range reports {
		if checkReport(report) {
			safeReports = append(safeReports, report)
		} else {
			unsafeReports = append(unsafeReports, report)
		}

	}
	return safeReports, unsafeReports
}

// more expensive, perform on subset of all reports
func sortReportsPart2(reports [][]int) (safe [][]int, unsafe [][]int) {
	safeReports := [][]int{}
	unsafeReports := [][]int{}
	for _, report := range reports {
		// for each report, create all possible sub-reports w/ one element removed at each position
		subReports := [][]int{}
		for i := 0; i < len(report); i++ {
			subReports = append(subReports, RemoveIndex(report, i))
		}

		anySafe := false
		for _, subReport := range subReports {
			if checkReport(subReport) {
				anySafe = true
				break
			}
		}
		if anySafe {
			safeReports = append(safeReports, report)
		} else {
			unsafeReports = append(unsafeReports, report)
		}
	}
	return safeReports, unsafeReports
}

// https://stackoverflow.com/a/57213476
func RemoveIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}
