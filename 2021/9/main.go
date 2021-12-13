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
	heightmap := parseInput(input)

	totalrisk := 0
	for y := 0; y < len(heightmap); y++ {
		for x := 0; x < len(heightmap[y]); x++ {
			digit := heightmap[y][x]
			surroundingDigits := allSurroundingDigits(heightmap, y, x)
			smallest := true
			for _, neighbour := range surroundingDigits {
				if digit >= neighbour {
					smallest = false
					break
				}
			}
			if smallest {
				totalrisk += 1 + digit
			}
		}
	}
	fmt.Println("Total risk : ", totalrisk)
}

func allSurroundingDigits(heightmap [][]int, y int, x int) []int {
	surrounding := []int{}

	if y-1 >= 0 {
		surrounding = append(surrounding, heightmap[y-1][x])
	}
	if x-1 >= 0 {
		surrounding = append(surrounding, heightmap[y][x-1])
	}
	if y+1 < len(heightmap) {
		surrounding = append(surrounding, heightmap[y+1][x])
	}
	if x+1 < len(heightmap[y]) {
		surrounding = append(surrounding, heightmap[y][x+1])
	}

	return surrounding
}

func parseInput(input []string) [][]int {
	heightmap := make([][]int, len(input))

	for i, row := range input {
		heightmap[i] = make([]int, len(row))

		for y, v := range row {
			digit, err := strconv.Atoi(string(v))
			if err != nil {
				panic(err)
			}
			heightmap[i][y] = int(digit)
		}
	}

	return heightmap
}
