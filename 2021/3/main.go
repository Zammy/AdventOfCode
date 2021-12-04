package main

import (
	"fmt"
	"strconv"
	"strings"

	"evgeni.com/util"
)

func convert(input []int, converter func(int) byte) int {
	var strBuilder strings.Builder
	for _, v := range input {
		strBuilder.WriteByte(converter(v))
	}
	result, _ := strconv.ParseInt(strBuilder.String(), 2, 32)
	return int(result)
}

func mostSignificantBitConverter(count int) byte {
	if count > 0 {
		return '1'
	} else {
		return '0'
	}
}

func leastSignificantBitConverter(count int) byte {
	if count > 0 {
		return '0'
	} else {
		return '1'
	}
}

func main() {
	input, err := util.ReadLines("input")
	if err != nil {
		panic(err)
	}

	bitCounter := make([]int, len(input[0]))
	for _, reading := range input {
		for i := 0; i < len(bitCounter); i++ {
			if reading[i] == '1' {
				bitCounter[i] += 1
			} else {
				bitCounter[i] -= 1
			}
		}
	}

	gamma := convert(bitCounter, mostSignificantBitConverter)
	epsilon := convert(bitCounter, leastSignificantBitConverter)

	fmt.Printf("gamma:%v epsilon:%v -> Answer:%v", gamma, epsilon, gamma*epsilon)
}
