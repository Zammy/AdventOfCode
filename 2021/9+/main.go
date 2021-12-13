package main

import (
	"fmt"
	"sort"
	"strconv"

	"evgeni.com/util"
)

type point struct {
	X, Y int
}

func main() {
	input, err := util.ReadLines("input")
	if err != nil {
		panic(err)
	}
	heightmap, visited := parseInput(input)
	lowpoints := []point{}

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
				lowpoints = append(lowpoints, point{Y: y, X: x})
			}
		}
	}

	basinSizes := []int{}
	for _, basinLowPoint := range lowpoints {
		size := chartBasin(heightmap, visited, basinLowPoint)
		basinSizes = append(basinSizes, size)
	}

	sort.Ints(basinSizes)
	biggestBasins := basinSizes[len(basinSizes)-3:]

	fmt.Println("Answer ", biggestBasins[0]*biggestBasins[1]*biggestBasins[2])
}

func chartBasin(heightmap [][]int, visited [][]bool, p point) int {
	if p.X < 0 || p.Y < 0 || p.Y >= len(heightmap) || p.X >= len(heightmap[p.Y]) {
		return 0
	}
	if heightmap[p.Y][p.X] == 9 || visited[p.Y][p.X] {
		return 0
	}

	visited[p.Y][p.X] = true

	return 1 +
		chartBasin(heightmap, visited, point{Y: p.Y - 1, X: p.X}) +
		chartBasin(heightmap, visited, point{Y: p.Y + 1, X: p.X}) +
		chartBasin(heightmap, visited, point{Y: p.Y, X: p.X - 1}) +
		chartBasin(heightmap, visited, point{Y: p.Y, X: p.X + 1})
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

func parseInput(input []string) ([][]int, [][]bool) {
	heightmap := make([][]int, len(input))
	visited := make([][]bool, len(input))

	for i, row := range input {
		heightmap[i] = make([]int, len(row))
		visited[i] = make([]bool, len(row))

		for y, v := range row {
			digit, err := strconv.Atoi(string(v))
			if err != nil {
				panic(err)
			}
			heightmap[i][y] = int(digit)
		}
	}

	return heightmap, visited
}
