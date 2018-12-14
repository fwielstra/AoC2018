package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
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

type Coordinate struct {
	left int
	top int
}

type Claim struct {
	id string
	pos Coordinate
	width int
	height int
}

var claimStringExpression = regexp.MustCompile(`[-]{0,1}[\d]*[\.]{0,1}[\d]+`)

func toInt(str string) int {
	val, err := strconv.Atoi(str)
	check(err)
	return val
}

func parseClaimString(claimString string) Claim {
	matches := claimStringExpression.FindAllString(claimString, -1)
	//fmt.Printf("string: %v matches: %v\n", claimString, matches);
	return Claim{
		id: matches[0],
		pos: Coordinate{toInt(matches[1]), toInt(matches[2])},
		width: toInt(matches[3]),
		height: toInt(matches[4]),
	}
}

func countOverlappingInches(claims []Claim) int {
	var overlaps = make(map[Coordinate]bool)

	iterations := 0

	for _, claim := range claims {
		fmt.Printf("Checking claim %v\n...", claim.id)
		for _, otherClaim := range claims {
			for x := claim.pos.left; x <= claim.pos.left + claim.width; x++ {
				for y := claim.pos.top; y <= claim.pos.top + claim.height; y++ {
					for x2 := otherClaim.pos.left; x2 <= otherClaim.pos.left + otherClaim.width; x2++ {
						for y2 := otherClaim.pos.top; y2 <= otherClaim.pos.top + otherClaim.height; y2++ {
							iterations++
							if x == x2 && y == y2 {
								overlaps[Coordinate{x, y}] = true
							}
						}
					}
				}
			}
		}
	}

	fmt.Printf("Iterations: %d\n", iterations)

	return len(overlaps)
}

func main() {
	input := readInputFile("input.txt")
	claimStrings := cleanInput(strings.Split(input, "\n"))
	logger.Printf("Read %v lines\n", len(claimStrings))

	var claims []Claim
	startParse := time.Now()
	for _, claimString := range claimStrings {
		claims = append(claims, parseClaimString(claimString))
	}

	logger.Printf("Parsed %d claims in %s\n", len(claims), time.Since(startParse))

	startCalculate := time.Now()
	overlaps := countOverlappingInches(claims)
	fmt.Printf("overlaps: %v in %s\n", overlaps, time.Since(startCalculate))

	//307.629.966.736
	fmt.Print(&buf)
}
