package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"sort"
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

func readInputFile() string {
	dat, err := ioutil.ReadFile("./input.txt")
	check(err)
	return string(dat)
}

// we need to parse date and sort by date first
type logLine struct {
	timestamp string
	log       string
}

type entry struct {
	timestamp string
	guardID   int
	isAsleep  bool // default false when guard sleeps.
}

// TODO: Optimization, read file per line instead of store whole file in memory.
func parseInput(input string) []logLine {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1] // remove last empty line
	logs := make([]logLine, len(lines))
	for i, line := range lines {
		start := strings.Index(line, "[") + 1
		end := strings.LastIndex(line, "]")
		date := line[start:end]
		logs[i] = logLine{date, line[end+2:]}
	}
	return logs
}

var beginsShiftExpression = regexp.MustCompile(`\d+`)

func createSchedule(lines []logLine) []entry {
	entries := make([]entry, len(lines))
	var currentGuardID int
	for i, line := range lines {
		match := beginsShiftExpression.FindAllString(line.log, -1)
		if len(match) > 0 {
			guardID, err := strconv.Atoi(match[0])
			check(err)
			currentGuardID = guardID
		}
		fmt.Printf("blaat: [%v]\n", line.log)

		// TODO: check inclusivity
		if line.log == "falls asleep" {
			entries[i] = entry{timestamp: line.timestamp, guardID: currentGuardID, isAsleep: true}
		} else {
			// wakes up or starts shift
			entries[i] = entry{timestamp: line.timestamp, guardID: currentGuardID, isAsleep: false}
		}
	}
	return entries
}

func main() {
	input := readInputFile()
	lines := parseInput(input)
	sort.Slice(lines, func(i, j int) bool {
		return strings.Compare(lines[i].timestamp, lines[j].timestamp) < 0
	})

	schedule := createSchedule(lines)

	for i, line := range lines {
		fmt.Printf("Line: %v, schedule: %+v\n", line, schedule[i])
	}
}
