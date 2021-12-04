package main

import (
	"fmt"
	"strconv"
	"strings"

	"evgeni.com/util"
)

type direction int

const (
	forward = 0
	up      = 1
	down    = 2
)

func (d direction) String() string {
	switch d {
	case forward:
		return "forward"
	case up:
		return "up"
	case down:
		return "down"
	default:
		panic("Direction not supported")
	}
}

func ToDirection(s string) direction {
	switch s {
	case "forward":
		return forward
	case "up":
		return up
	case "down":
		return down
	default:
		return -1
	}
}

type command struct {
	Direction direction
	Amount    int
}

func (c command) String() string {
	return fmt.Sprintf("{%v : %v}", c.Direction, c.Amount)
}

func extractCommand(data string) *command {
	var result command
	split := strings.Split(data, " ")
	result.Direction = ToDirection(split[0])
	result.Amount, _ = strconv.Atoi(strings.Trim(split[1], " "))
	return &result
}

func main() {
	input, err := util.ReadLines("input")
	if err != nil {
		panic(err)
	}

	aim := 0
	depth := 0
	horzpos := 0

	for _, v := range input {
		command := extractCommand(v)
		if command.Direction == forward {
			horzpos += command.Amount
			depth += command.Amount * aim
		} else if command.Direction == up {
			aim -= command.Amount
		} else if command.Direction == down {
			aim += command.Amount
		}
	}

	fmt.Printf("Pos:%v Depth:%v -> Answer:%v", horzpos, depth, horzpos*depth)
}
