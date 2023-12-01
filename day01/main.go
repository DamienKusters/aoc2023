package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
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
	fmt.Println(sum)
}
