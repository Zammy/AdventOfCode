package main

import (
	"fmt"
	"strings"

	"evgeni.com/util"
)

const DAYS_TO_RUN = 80
const MOTHER_RESET_TIMER = 6
const OFFSPRING_TIMER = 8

func main() {
	input, err := util.ReadLines("input")
	if err != nil {
		panic(err)
	}

	lanternfishes := parseInput(input[0])

	for i := 0; i < DAYS_TO_RUN; i++ {
		ageFish(lanternfishes)
		lanternfishes = giveBirth(lanternfishes)
	}

	fmt.Printf("%d of fish after %d", len(lanternfishes), DAYS_TO_RUN)
}

func parseInput(input string) []int {
	split := strings.Split(input, ",")
	result := make([]int, len(split))
	for i, v := range split {
		result[i] = util.ParseInt(v)
	}
	return result
}

func ageFish(fishes []int) {
	for i := 0; i < len(fishes); i++ {
		fishes[i] -= 1
	}
}

func giveBirth(fishes []int) []int {
	newBirths := 0
	for i := 0; i < len(fishes); i++ {
		if fishes[i] < 0 {
			fishes[i] = MOTHER_RESET_TIMER
			newBirths++
		}
	}
	for i := 0; i < newBirths; i++ {
		fishes = append(fishes, OFFSPRING_TIMER)
	}
	return fishes
}
