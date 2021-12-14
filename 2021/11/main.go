package main

import (
	"fmt"
	"strconv"
	"strings"

	"evgeni.com/util"
)

const STEPS = 100
const FLASH_ENG_LVL = 10

var directions = [][]int{
	{0, -1},
	{-1, -1},
	{-1, 0},
	{-1, +1},
	{0, +1},
	{1, +1},
	{1, 0},
	{1, -1},
}

func main() {
	input, err := util.ReadLines("input")
	if err != nil {
		panic(err)
	}

	octos := parseInput(input)
	flashes := 0
	print(octos)
	fmt.Println()
	for i := 0; i < STEPS; i++ {
		for y := 0; y < len(octos); y++ {
			for x := 0; x < len(octos[y]); x++ {
				flashes += increaseEnergy(octos, y, x)
			}
		}

		for y := 0; y < len(octos); y++ {
			for x := 0; x < len(octos[y]); x++ {
				if octos[y][x] >= FLASH_ENG_LVL {
					octos[y][x] = 0
				}
			}
		}
	}

	print(octos)
	fmt.Println()
	fmt.Println("Flashes: ", flashes)
}

func parseInput(input []string) [][]int {
	result := make([][]int, len(input[0]))
	for y, row := range input {
		result[y] = make([]int, len(row))

		for x, r := range row {
			digit, err := strconv.Atoi(string(r))
			if err != nil {
				panic("Failed to parse")
			}

			result[y][x] = digit
		}
	}

	return result
}

func increaseEnergy(octos [][]int, y int, x int) int {
	flashes := 0
	if isSafe(octos, y, x) {
		octos[y][x]++
		if octos[y][x] == FLASH_ENG_LVL {
			flashes += flash(octos, y, x)
		}
	}
	return flashes
}

func flash(octos [][]int, y int, x int) int {
	// fmt.Println("FLASH!")
	// print(octos)
	flashes := 1
	for _, dir := range directions {
		flashes += increaseEnergy(octos, y+dir[0], x+dir[1])
	}
	return flashes
}

func isSafe(octos [][]int, y int, x int) bool {
	return (y >= 0) && (x >= 0) && (y < len(octos)) && (x < len(octos[y]))
}

func print(octos [][]int) {
	var builder strings.Builder
	builder.WriteString("=====================\n")
	for y := 0; y < len(octos); y++ {
		for x := 0; x < len(octos[y]); x++ {
			if octos[y][x] <= 9 {
				builder.WriteString(fmt.Sprint(octos[y][x]))
			} else {
				builder.WriteString("@")
			}
		}
		builder.WriteRune('\n')
	}
	fmt.Print(builder.String())
}
