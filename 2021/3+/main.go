package main

import (
	"fmt"
	"strconv"

	"evgeni.com/util"
)

func main() {
	input, err := util.ReadLines("input")
	if err != nil {
		panic(err)
	}

	oxygenGeneratorRating := find(input, oxygenFilterFunc)
	co2ScrubingRating := find(input, co2FilterFunc)

	fmt.Printf("oxygenGeneratorRating :%v co2ScrubingRating:%v -> Answer life support rating:%v", oxygenGeneratorRating, co2ScrubingRating, oxygenGeneratorRating*co2ScrubingRating)
}

type filterFunc func(int, byte) bool

func oxygenFilterFunc(bitsCount int, bit byte) bool {
	return (bitsCount >= 0 && bit == '1') || (bitsCount < 0 && bit == '0')
}

func co2FilterFunc(bitsCount int, bit byte) bool {
	return (bitsCount >= 0 && bit == '0') || (bitsCount < 0 && bit == '1')
}

func countBitsOnBit(input *[]string, onBit int) int {
	count := 0
	for _, binaryNum := range *input {
		if binaryNum[onBit] == '1' {
			count++
		} else {
			count--
		}
	}
	return count
}

func filterOnBit(input []string, onBit int, filter filterFunc) []string {
	result := []string{}
	bitsCount := countBitsOnBit(&input, onBit)
	for _, num := range input {
		if filter(bitsCount, num[onBit]) {
			result = append(result, num)
		}
	}
	return result
}

func find(input []string, filter filterFunc) int {
	for onBit := 0; len(input) > 1; onBit++ {
		input = filterOnBit(input, onBit, filter)
	}

	result, _ := strconv.ParseInt(input[0], 2, 32)
	return int(result)
}

// func countBits(input []string) *[]int {
// 	bitCounter := make([]int, len(input[0]))
// 	for _, reading := range input {
// 		for i := 0; i < len(bitCounter); i++ {
// 			if reading[i] == '1' {
// 				bitCounter[i] += 1
// 			} else {
// 				bitCounter[i] -= 1
// 			}
// 		}
// 	}
// 	return &bitCounter
// }
