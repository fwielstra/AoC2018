package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"
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

func calculateChecksum(boxIds []string) int {
	totalContainsTwo := 0
	totalContainsThree := 0

	for _, boxId := range boxIds {
		var occurrences = make(map[int32]int)

		for _, char := range boxId {
			occurrences[char] += 1
		}

		countsTwice := false
		countsThrice := false

		for _, occurs := range occurrences {
			if occurs == 2 && !countsTwice {
				totalContainsTwo++
				countsTwice = true
			}

			if occurs == 3 && !countsThrice {
				totalContainsThree++
				countsThrice = true
			}
		}
	}

	fmt.Printf("totalContainsTwo: %v, totalContainsThree: %v\n", totalContainsTwo, totalContainsThree)

	return totalContainsTwo * totalContainsThree
}

type Pair struct {
	one, other string
}

func findCommonCharacters(boxIds []string) string {

	distances := make(map[Pair]int)
	for _, boxId := range boxIds {
		for _, comparedBoxId := range boxIds {
			for i := 0; i < len(boxId) - 1; i++ {
				if boxId[i] != comparedBoxId[i] {
					distances[Pair{boxId, comparedBoxId}] += 1
				}
			}
		}
	}

	var results []string

	for pair, distance := range distances {
		// skip anything with a distance > 1
		if distance != 1 {
			continue
		}

		// if distance is 1, find the common characters
		result := ""
		for i := range pair.one {
			if pair.one[i] == pair.other[i] {
				result += string(pair.one[i])
			}
		}
		results = append(results, result)
	}


	// blunt: just assume there's only one result
	return results[0]
}

func readInputFile(filename string) string {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return string(dat)
}

func cleanInput(input []string) []string {
	var result []string
	for _, str := range input {
		if str != "" {
			result = append(result, str)
		}
	}
	return result
}


func main() {
	input := readInputFile("input.txt")
	boxIds := cleanInput(strings.Split(input, "\n"))
	logger.Printf("Read %v lines\n", len(boxIds))
	start := time.Now()
	checksum := calculateChecksum(boxIds)
	elapsed := time.Since(start)
	logger.Printf("Checksum: %v calculated in %s", checksum, elapsed)

	startFindCommon := time.Now()
	common := findCommonCharacters(boxIds)
	logger.Printf("Common ID: %v calculated in %s", common, time.Since(startFindCommon))


	fmt.Print(&buf)
}
