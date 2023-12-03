package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	part1()
}

var sum int

func part1() {
	regex := regexp.MustCompile("[0-9]+")
	rawData, _ := os.ReadFile("input.txt")
	rows := strings.Split(strings.ReplaceAll(string(rawData), "\r\n", "\n"), "\n")

	for rowIdx, row := range rows {
		rowRanges := regex.FindAllStringIndex(row, -1)
		for _, rowRange := range rowRanges {
			num, _ := strconv.Atoi(row[rowRange[0]:rowRange[1]])
			if scanAdjacent(rowIdx, rows, rowRange, rowIdx > 0, true, rowIdx < len(row)-1) {
				sum += num
			}
		}
	}
	fmt.Print("Part1: ")
	fmt.Println(sum)
}

func scanAdjacent(rowIdx int, rows []string, rowRange []int, top bool, mid bool, bot bool) bool {
	ignoredCharacters := regexp.MustCompile(`[0-9]|\.`)
	var scanRows []string
	if top {
		scanRows = append(scanRows, rows[rowIdx-1])
	}
	if mid {
		scanRows = append(scanRows, rows[rowIdx])
	}
	if bot {
		scanRows = append(scanRows, rows[rowIdx+1])
	}

	if rowRange[0] != 0 {
		rowRange[0] -= 1
	}
	if rowRange[1] < len(rows[0]) {
		rowRange[1] += 1
	}
	for _, row := range scanRows {
		for i := rowRange[0]; i < rowRange[1]; i++ {

			if ignoredCharacters.MatchString(string(row[i])) == false {
				return true
			}
		}
	}
	return false
}
