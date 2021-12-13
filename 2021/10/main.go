package main

import (
	"fmt"

	"evgeni.com/util"
)

func main() {
	input, err := util.ReadLines("input")
	if err != nil {
		panic(err)
	}

	scoring := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}

	score := 0
	for _, row := range input {
		illegal := findIlligalCharacter(row)
		if illegal != 0 {
			score += scoring[illegal]
		}
	}
	fmt.Println("Answer ", score)
}

func findIlligalCharacter(row string) rune {
	stack := []rune{}
	for _, r := range row {
		if isClosing(r) {
			prevOpenning := stack[len(stack)-1]
			if matchOpening(r) == prevOpenning {
				stack = stack[:len(stack)-1]
			} else {
				return r
			}
		} else {
			stack = append(stack, r)
		}
	}

	return 0
}

func isClosing(r rune) bool {
	return r == '}' || r == ')' || r == '>' || r == ']'
}

func matchOpening(r rune) rune {
	switch r {
	case '}':
		return '{'
	case ')':
		return '('
	case '>':
		return '<'
	case ']':
		return '['
	default:
		panic("Should never reach here")
	}

}
