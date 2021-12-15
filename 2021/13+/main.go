package main

import (
	"fmt"
	"math"
	"strings"

	"evgeni.com/util"
)

func main() {
	input, err := util.ReadLines("input")
	if err != nil {
		panic(err)
	}

	dots, folds := parseInput(input)
	paper := paper{}
	paper.InitWithSize(findBiggest(dots))
	for _, dot := range dots {
		paper.AddDotAt(dot[0], dot[1])
	}

	for _, fold := range folds {
		paper.Fold(fold)
	}

	fmt.Println(paper.String())
}

type fold struct {
	axis  rune
	value int
}

func parseInput(input []string) ([][]int, []fold) {
	dots := [][]int{}
	folds := []fold{}
	for _, row := range input {
		split := strings.Split(row, ",")
		if len(split) > 1 {
			x := util.ParseInt(split[0])
			y := util.ParseInt(split[1])
			dots = append(dots, []int{x, y})
		} else {
			split := strings.Split(row, "=")
			if len(split) > 1 {
				fold := fold{
					axis:  rune(split[0][len(split[0])-1]),
					value: util.ParseInt(split[1]),
				}
				folds = append(folds, fold)
			}
		}
	}
	return dots, folds
}

type paper struct {
	dots [][]bool
}

func (p *paper) InitWithSize(x int, y int) {
	p.dots = make([][]bool, y+1)
	for i := 0; i < y+1; i++ {
		p.dots[i] = make([]bool, x+1)
	}
}

func (p *paper) AddDotAt(x int, y int) {
	p.dots[y][x] = true
}

func (p *paper) String() string {
	var sb strings.Builder
	for y := 0; y < len(p.dots); y++ {
		for x := 0; x < len(p.dots[y]); x++ {
			r := '.'
			if p.dots[y][x] {
				r = '#'
			}
			sb.WriteRune(r)
		}
		sb.WriteRune('\n')
	}
	sb.WriteRune('\n')
	return sb.String()
}

func (p *paper) Fold(fold fold) {
	if fold.axis == 'y' {
		for y := fold.value + 1; y < len(p.dots); y++ {
			for x := 0; x < len(p.dots[y]); x++ {
				foldedY := fold.value - (y - fold.value)
				p.dots[foldedY][x] = p.dots[foldedY][x] || p.dots[y][x]
			}
		}
		p.dots = p.dots[:fold.value]
	} else {
		for y := 0; y < len(p.dots); y++ {
			for x := fold.value + 1; x < len(p.dots[y]); x++ {
				foldedX := fold.value - (x - fold.value)
				p.dots[y][foldedX] = p.dots[y][foldedX] || p.dots[y][x]
			}
			p.dots[y] = p.dots[y][:fold.value]
		}
	}
}

func (p *paper) CountDots() int {
	count := 0
	for y := 0; y < len(p.dots); y++ {
		for x := 0; x < len(p.dots[y]); x++ {
			if p.dots[y][x] {
				count++
			}
		}
	}
	return count
}

func findBiggest(dots [][]int) (int, int) {
	biggestX, biggsetY := math.MinInt, math.MinInt
	for _, dot := range dots {
		if dot[0] > biggestX {
			biggestX = dot[0]
		}
		if dot[1] > biggsetY {
			biggsetY = dot[1]
		}
	}
	return biggestX, biggsetY
}
