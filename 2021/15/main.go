package main

import (
	"fmt"
	"math"

	"evgeni.com/util"
)

func main() {
	input, err := util.ReadLines("input")
	if err != nil {
		panic(err)
	}

	var risk [][]int
	var accrisk [][]int
	for y, row := range input {
		risk = append(risk, []int{})
		accrisk = append(accrisk, []int{})
		for _, r := range row {
			risk[y] = append(risk[y], util.ParseInt(string(r)))
			accrisk[y] = append(accrisk[y], math.MaxInt)
		}
	}

	stack := []point{{Y: len(risk) - 1, X: len(risk[0]) - 1}}
	accrisk[stack[0].Y][stack[0].X] = risk[stack[0].Y][stack[0].X]
	for len(stack) > 0 {
		currentPoint := stack[0]
		stack = stack[1:]
		currentRisk := accrisk[currentPoint.Y][currentPoint.X]

		for _, dir := range directions {
			newPoint := dir.Add(currentPoint)
			if newPoint.Y == len(risk) || newPoint.X == len(risk[0]) || newPoint.X < 0 || newPoint.Y < 0 {
				continue
			}
			newRisk := currentRisk + risk[newPoint.Y][newPoint.X]
			if newRisk < accrisk[newPoint.Y][newPoint.X] {
				accrisk[newPoint.Y][newPoint.X] = newRisk
				stack = append(stack, newPoint)
			}
		}
	}

	fmt.Println(accrisk[0][0] - risk[0][0])

}

type point struct {
	X, Y int
}

func (p point) Add(other point) point {
	return point{X: p.X + other.X, Y: p.Y + other.Y}
}

var directions = []point{
	{-1, 0},
	{+1, 0},
	{0, +1},
	{0, -1},
}
