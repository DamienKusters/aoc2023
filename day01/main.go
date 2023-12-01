package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/dlclark/regexp2"
)

func main() {
	part1()
	part2()
}

func part1() {
	regex := regexp.MustCompile("[0-9]")
	rawData, _ := os.ReadFile("input.txt")

	data := string(rawData)
	split := strings.Fields(data)

	sum := 0
	for _, word := range split {
		numbers := regex.FindAllString(word, -1)
		number, _ := strconv.Atoi(numbers[0] + numbers[len(numbers)-1])
		sum += number
	}
	fmt.Print("Part1: ")
	fmt.Println(sum)
}

func part2() {
	regex := regexp2.MustCompile("[0-9]|(?=(one|two|three|four|five|six|seven|eight|nine))", 0)
	rawData, _ := os.ReadFile("input.txt")

	data := string(rawData)
	split := strings.Fields(data)

	sum := 0
	for _, word := range split {
		numbers := regexp2FindAllString(regex, word)

		firstNumber := getNumericValue(numbers[0])
		lastNumber := getNumericValue(numbers[len(numbers)-1])

		combined, _ := strconv.Atoi(strconv.Itoa(firstNumber) + strconv.Itoa(lastNumber))
		sum += combined
	}
	fmt.Print("Part2: ")
	fmt.Println(sum)
}

func getNumericValue(input string) int {
	valuemap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	digitRegex := regexp.MustCompile("[0-9]")

	output := 0
	if digitRegex.MatchString(input) {
		output, _ = strconv.Atoi(input)
	} else {
		output = valuemap[input]
	}
	return output
}

func regexp2FindAllString(re *regexp2.Regexp, s string) []string {
	var matches []string
	m, _ := re.FindStringMatch(s)
	for m != nil {
		matches = append(matches, m.String())
		m, _ = re.FindNextMatch(m)
	}
	return matches
}
