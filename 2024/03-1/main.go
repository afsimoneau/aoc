package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type pair struct {
	num1 int
	num2 int
}

func main() {
	content := parseInput("input.txt")
	rawMuls := parseContent(content)
	for _, rawMul := range rawMuls{
		fmt.Println(rawMul)
	}
	total:= calcTotal(rawMuls)
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

func parseContent(content string) [][]string {
	r := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	validMuls := r.FindAllStringSubmatch(content, -1)

	return validMuls
}

func calcTotal(rawMuls [][]string) int {
	total := 0
	for _, mul := range rawMuls {
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
