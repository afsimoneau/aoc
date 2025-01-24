package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	content := parseInput("input.txt")
	total := calcMuls(parseMuls(StripDoDont(content)))

	fmt.Printf("total: %d\n", total)
}

func parseInput(fileName string) string {
	// open file
	f, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	return string(f)
}

func parseMuls(content string) [][]string {
	r := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	validMuls := r.FindAllStringSubmatch(content, -1)

	return validMuls
}

func calcMuls(matchedMuls [][]string) int {
	total := 0
	for _, mul := range matchedMuls {
		num1, err := strconv.Atoi(mul[1])
		if err != nil {
			log.Fatal(err)
		}
		num2, err := strconv.Atoi(mul[2])
		if err != nil {
			log.Fatal(err)
		}
		
		total += num1 * num2
	}
	return total
}

// https://www.meetgor.com/aoc-2024-day-3/
func StripDoDont(line string) string {
	result := ""
	enabled := true
	dontOffset := len("don't()")
	doOffset := len("do()")

	for len(line) > 0 {
		dontIndex := strings.Index(line, "don't()")
		doIndex := strings.Index(line, "do()")

		if dontIndex == -1 && doIndex == -1 {
			if enabled {
				result += line
			}
			break
		}

		if dontIndex != -1 && (doIndex == -1 || dontIndex < doIndex) {
			if enabled {
				result += line[:dontIndex]
			}
			enabled = false
			line = line[dontIndex+dontOffset:]
		} else {
			if enabled {
				result += line[:doIndex]
			}
			enabled = true
			line = line[doIndex+doOffset:]
		}
	}

	return result
}
