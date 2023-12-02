package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1()
}

var blockLimits = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

var validGameIds []int
var sum int

func part1() {
	rawData, _ := os.ReadFile("input.txt")
	games := strings.Split(strings.ReplaceAll(string(rawData), "\r\n", "\n"), "\n")

	for _, d := range games {
		game := strings.Split(d, ":")
		if game[0] == "" {
			//TODO: result
			// fmt.Println(validGameIds)
			for _, g := range validGameIds {
				sum += g
			}
			fmt.Print("Part1: ")
			fmt.Println(sum)
			return
		}
		rounds := strings.Split(game[1], ";")
		roundIsPossible := true
		for _, r := range rounds {
			blocks := strings.Split(r, ",")
			for _, b := range blocks {
				block := strings.Split(strings.TrimSpace(b), " ")
				blockValue, _ := strconv.Atoi(block[0])
				blockKind := block[1]

				if blockValue > blockLimits[blockKind] {
					roundIsPossible = false
					break
				}
			}
			if roundIsPossible == false {
				break
			}
		}

		if roundIsPossible {
			gameid, _ := strconv.Atoi(strings.Split(game[0], " ")[1])
			validGameIds = append(validGameIds, gameid)
		}
	}

}
