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

	return totalContainsTwo * totalContainsThree
}

type Pair struct {
	one, other string
}

func findCommonCharacters(boxIds []string) string {

	distances := make(map[Pair]int)
	for _, boxId := range boxIds {
	Blaat:
		for _, comparedBoxId := range boxIds {
			pair := Pair{boxId, comparedBoxId}
			inversePair := Pair{comparedBoxId, boxId}
			if _, ok := distances[inversePair]; ok {
				// this pair has already been checked
				continue
			}

			for i := 0; i < len(boxId)-1; i++ {
				if boxId[i] != comparedBoxId[i] {
					distances[pair] += 1
					if distances[pair] > 1 {
						// this pair is no longer viable; remove.
						delete(distances, pair)
						continue Blaat
					}
				}
			}
		}
	}

	// distances should have length 1 here because we removed the ones with >1 distance

	for pair := range distances {
		// assemble the common characters from the pair
		result := ""
		for i := range pair.one {
			if pair.one[i] == pair.other[i] {
				result += string(pair.one[i])
			}
		}
		return result
	}
	return ""
}

func readInputFile(filename string) string {
	dat, err := ioutil.ReadFile(filename)
	check(err)
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
