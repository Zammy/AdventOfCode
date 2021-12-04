package main

import (
	"fmt"
	"math"
	"strconv"

	"evgeni.com/util"
)

func main() {
	input, err := util.ReadLines("input")
	if err != nil {
		panic(err)
	}

	previous := math.MaxInt
	increases_count := 0

	for _, depthStr := range input {
		depth, err := strconv.Atoi(depthStr)
		if err != nil {
			panic(err)
		}
		if depth > previous {
			increases_count++
		}
		previous = depth
	}

	fmt.Printf("Depth increases %v times", increases_count)
}
