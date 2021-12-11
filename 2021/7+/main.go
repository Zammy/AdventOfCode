package main

import (
	"fmt"
	"math"
	"strings"

	"evgeni.com/util"
)

var rnge int = 10

func main() {
	input, err := util.ReadLines("input")
	if err != nil {
		panic(err)
	}
	crabPositions := parseInput(input[0])

	avrPos, minPos, maxPos := calcAvrgMinMax(crabPositions)
	rnge = (maxPos - minPos) / 4
	bestDelta := sumDistanceTo(crabPositions, avrPos)
	to := avrPos + rnge/2
	for i := avrPos - rnge/2; i < to; i++ {
		delta := sumDistanceTo(crabPositions, i)
		if delta < bestDelta {
			avrPos = i
			bestDelta = delta
		}
	}

	fmt.Printf("Fuel: %v at horz pos: %v", bestDelta, avrPos)
}

func calcAvrgMinMax(input []int) (avrg, min, max int) {
	sum := 0
	min = math.MaxInt
	max = math.MinInt
	for _, v := range input {
		sum += v
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	avrg = sum / len(input)
	return
}

func sumDistanceTo(input []int, ref int) int {
	sum := 0
	for _, v := range input {
		diff := abs(v - ref)
		for ; diff > 0; diff-- {
			sum += diff
		}
	}
	return sum
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func parseInput(input string) []int {
	split := strings.Split(input, ",")
	result := make([]int, len(split))
	for i, v := range split {
		result[i] = util.ParseInt(v)
	}
	return result
}
