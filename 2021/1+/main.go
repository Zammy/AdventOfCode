package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"evgeni.com/util"
)

func sum_previous_three(data []string) int {
	sum := 0
	for _, v := range data {
		depth, err := strconv.Atoi(strings.Trim(v, " "))
		if err != nil {
			panic(err)
		}
		sum += depth
	}
	return sum
}

func main() {
	input, err := util.ReadLines("input")
	if err != nil {
		panic(err)
	}

	previous := math.MaxInt
	increases_count := 0

	for i := 3; i <= len(input); i++ {
		sum := sum_previous_three(input[i-3 : i])
		if sum > previous {
			increases_count++
		}
		previous = sum
	}

	fmt.Printf("Depth increases %v times", increases_count)
}
