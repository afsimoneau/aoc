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
	// open file
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	var parsedLeft []int
	var parsedRight []int
	for scanner.Scan() {
		// get line
		line := scanner.Text()
		fmt.Printf("line: %s\n", line)

		// split line
		splitLine := strings.Split(line, "   ")
		fmt.Printf("strings.Split(): %#v\n", splitLine)

		// parse left
		left, err := strconv.Atoi(splitLine[0])
		if err != nil {
			log.Fatal(err)
		}
		parsedLeft = append(parsedLeft, left)

		// parse right
		right, err := strconv.Atoi(splitLine[1])
		if err != nil {
			log.Fatal(err)
		}
		parsedRight = append(parsedRight, right)
	}

	// verify parsing
	if len(parsedLeft) != len(parsedRight) {
		log.Fatal("parsed slices differ in size")
	}

	//print parsed
	for i, left := range parsedLeft {
		fmt.Printf("left: %d | right: %d\n", left, parsedRight[i])
	}

	//in-place sort
	sort.Ints(parsedLeft)
	sort.Ints(parsedRight)

	//print sorted
	for i, left := range parsedLeft {
		fmt.Printf("left: %d | right: %d\n", left, parsedRight[i])
	}

	// total
	var total int
	for i, left := range parsedLeft {
		diff := left - parsedRight[i]
		if diff < 0 {
			diff *= -1
		}
		total += diff
	}

	fmt.Printf("total: %d\n", total)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
