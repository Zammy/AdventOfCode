package main

import (
	"fmt"
	"sort"

	"evgeni.com/util"
)

func main() {
	input, err := util.ReadLines("input")
	if err != nil {
		panic(err)
	}

	scoring := map[rune]int{
		'(': 1,
		'[': 2,
		'{': 3,
		'<': 4,
	}

	scores := []int{}
	for _, row := range input {
		illegal, stack := findIlligalCharacter(row)
		if illegal == 0 {
			score := 0
			for i := len(stack) - 1; i >= 0; i-- {
				score *= 5
				score += scoring[stack[i]]
			}
			scores = append(scores, score)
		}
	}
	sort.Ints(scores)

	fmt.Println("Total score: ", scores[len(scores)/2])
}

func findIlligalCharacter(row string) (rune, []rune) {
	stack := []rune{}
	for _, r := range row {
		if isClosing(r) {
			prevOpenning := stack[len(stack)-1]
			if matchOpening(r) == prevOpenning {
				stack = stack[:len(stack)-1]
			} else {
				return r, stack
			}
		} else {
			stack = append(stack, r)
		}
	}

	return 0, stack
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
