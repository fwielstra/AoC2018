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

func readInputFile() string {
	dat, err := ioutil.ReadFile("./input.txt")
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

type coordinate struct {
	left int
	top  int
}

type claim struct {
	id     string
	pos    coordinate
	width  int
	height int
}

type inch struct {
	overlaps int
	claims   map[string]claim
}

var claimStringExpression = regexp.MustCompile(`[-]{0,1}[\d]*[\.]{0,1}[\d]+`)

func toInt(str string) int {
	val, err := strconv.Atoi(str)
	check(err)
	return val
}

func parseClaimString(claimString string) claim {
	matches := claimStringExpression.FindAllString(claimString, -1)
	//fmt.Printf("string: %v matches: %v\n", claimString, matches);
	return claim{
		id:     matches[0],
		pos:    coordinate{toInt(matches[1]), toInt(matches[2])},
		width:  toInt(matches[3]),
		height: toInt(matches[4]),
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func createBitmap(claims []claim) [][]inch {
	height, width := 0, 0
	for _, claim := range claims {
		height = max(claim.pos.top+claim.height, height)
		width = max(claim.pos.left+claim.width, width)
	}

	var bitmap = make([][]inch, height)
	for i := range bitmap {
		bitmap[i] = make([]inch, width)
		for j := range bitmap[i] {
			bitmap[i][j].claims = make(map[string]claim)
		}
	}

	for _, claim := range claims {
		for y := claim.pos.top; y <= claim.pos.top+claim.height-1; y++ {
			for x := claim.pos.left; x <= claim.pos.left+claim.width-1; x++ {
				bitmap[y][x].overlaps++
				bitmap[y][x].claims[claim.id] = claim
			}
		}
	}

	return bitmap
}

func countOverlappingInches(bitmap [][]inch) int {
	result := 0
	for y := range bitmap {
		for x := range bitmap[y] {
			if bitmap[y][x].overlaps > 1 {
				result++
			}
		}
	}
	return result
}

// TODO: Should find an intersection where both the overlap count is 1 and the whole claim is used. urgh.
func findNotOverlappingClaimID(bitmap [][]inch) string {
	for y := range bitmap {
		for x := range bitmap[y] {
			if len(bitmap[y][x].claims) == 1 {
				fmt.Printf("Bitmap %d, %d has only one claim %+v\n", y, x, bitmap[y][x].claims)
			}
		}
	}

	return "blaat"
}

func main() {
	input := readInputFile()
	claimStrings := cleanInput(strings.Split(input, "\n"))
	logger.Printf("Read %v lines\n", len(claimStrings))
	var claims []claim
	startParse := time.Now()
	for _, claimString := range claimStrings {
		claims = append(claims, parseClaimString(claimString))
	}

	logger.Printf("Parsed %d claims in %s\n", len(claims), time.Since(startParse))

	startCalculate := time.Now()

	bitmap := createBitmap(claims)

	// 115304
	overlaps := countOverlappingInches(bitmap)
	fmt.Printf("overlaps: %v in %s\n", overlaps, time.Since(startCalculate))
	notOverlapping := findNotOverlappingClaimID(bitmap)
	fmt.Printf("Not overlapping claim ID: %s\n", notOverlapping)

	fmt.Print(&buf)
}
