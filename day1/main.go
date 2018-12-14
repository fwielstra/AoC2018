package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var (
	buf    bytes.Buffer
	logger = log.New(&buf, "logger: ", log.Lshortfile)
)

func check(err error) {
	if err != nil {
		logger.Panic(err)
	}
}

func calculateFrequency(deltas []int) int {
	result := 0
	for _, delta := range deltas {
		result += delta
	}
	return result
}

func parseInput(input string) []int {
	var result = []int{}
	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}
		value, err := strconv.Atoi(line)
		check(err)
		result = append(result, value)
	}

	return result
}

func readInputFile(filename string) string {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return string(dat)
}

func main() {
	logger.Print("AoC day 1")
	logger.Print("Reading input file...")
	input := readInputFile("input.txt")
	parsedInput := parseInput(input)
	logger.Printf("Read %d frequency deltas", len(parsedInput))
	result := calculateFrequency(parsedInput)
	logger.Printf("Result: %v", result)
	fmt.Print(&buf)
}
