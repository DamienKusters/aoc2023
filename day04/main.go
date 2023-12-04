package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	part1()
}

var sum = 0

func part1() {
	rawData, _ := os.ReadFile("input.txt")
	cards := strings.Split(strings.ReplaceAll(string(rawData), "\r\n", "\n"), "\n")

	for _, card := range cards {
		rawCard := strings.Split(strings.Split(card, "Card")[1], ":")
		rawCard[0] = strings.TrimSpace(rawCard[0])

		a := strings.ReplaceAll(rawCard[1], "  ", " ")
		winningYour := strings.Split(a, "|")
		winning := winningYour[0]
		your := strings.TrimSpace(winningYour[1])

		points := 0
		for _, number := range strings.Split(your, " ") {
			if strings.Contains(winning, " "+number+" ") {
				if points == 0 {
					points = 1
				} else {
					points = points + points
				}
			}
		}
		sum += points
	}
	fmt.Print("Part1: ")
	fmt.Println(sum)
}
