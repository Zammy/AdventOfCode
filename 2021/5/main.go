package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"evgeni.com/util"
)

var matcher = regexp.MustCompile(`(\d+),(\d+) -> (\d+),(\d+)`)

func main() {
	input, err := util.ReadLines("input")
	if err != nil {
		panic(err)
	}

	seabed := new(board2D)

	for i, row := range input {
		pairPoints := parsePoints(row)
		fmt.Printf("%d: %+d -> %+d\n", i, pairPoints[0], pairPoints[1])
		seabed.Draw(pairPoints)
		// fmt.Println(seabed)
	}

	overlaps2 := seabed.CountHigherThen(1)

	fmt.Printf(`Overlaps %v`, overlaps2)
}

type board2D struct {
	state [1000][1000]int
}

func (b *board2D) Draw(points [2]*point) {
	direction := points[1].SubtractNew(points[0])
	direction.Normalize()
	if direction.X != 0 && direction.Y != 0 {
		//TODO: ignore for now
		fmt.Println("Ignored")
		return
	}
	for p := *points[0]; !p.Equals(points[1]); p.Add(direction) {
		b.state[p.Y][p.X] += 1
		// fmt.Println(b)
		// fmt.Println()
	}
	b.state[points[1].Y][points[1].X] += 1
}

func (b *board2D) CountHigherThen(value int) int {
	count := 0
	for y := 0; y < len(b.state); y++ {
		for x := 0; x < len(b.state[y]); x++ {
			if b.state[y][x] > value {
				count += 1
			}
		}
	}
	return count
}

func (b *board2D) String() string {
	var builder = new(strings.Builder)
	for y := 0; y < len(b.state); y++ {
		for x := 0; x < len(b.state[y]); x++ {
			if b.state[y][x] > 0 {
				builder.WriteString(fmt.Sprint(b.state[y][x]))
			} else {
				builder.WriteRune('.')
			}
		}
		builder.WriteRune('\n')
	}
	return builder.String()
}

type point struct {
	X, Y int
}

func (p *point) Normalize() {
	if p.X > 0 {
		p.X = 1
	} else if p.X < 0 {
		p.X = -1
	}
	if p.Y > 0 {
		p.Y = 1
	} else if p.Y < 0 {
		p.Y = -1
	}
}

func (a *point) Add(b *point) {
	a.X += b.X
	a.Y += b.Y
}

func (a *point) SubtractNew(b *point) *point {
	return &point{X: a.X - b.X, Y: a.Y - b.Y}
}

func (a *point) Equals(b *point) bool {
	return a.X == b.X && a.Y == b.Y
}

func parsePoints(input string) [2]*point {
	points := [2]*point{new(point), new(point)}

	results := matcher.FindStringSubmatch(input)
	points[0].X = parseInt(results[1])
	points[0].Y = parseInt(results[2])
	points[1].X = parseInt(results[3])
	points[1].Y = parseInt(results[4])

	return points
}

func parseInt(input string) int {
	v, err := strconv.ParseInt(input, 0, 0)
	if err != nil {
		panic(err)
	}
	return int(v)
}
