package main

import (
	"fmt"
	"math"
	"strings"

	"evgeni.com/util"
)

const STEPS = 10

func main() {
	input, err := util.ReadLines("input")
	if err != nil {
		panic(err)
	}

	polymer := input[0]

	var rules []pairInsertion
	for _, pairData := range input[2:] {
		split := strings.Split(strings.ReplaceAll(pairData, " ", ""), "->")
		replaceStr := fmt.Sprintf("%v%v", split[1], string(split[0][1]))
		rules = append(rules, pairInsertion{pattern: split[0], replace: replaceStr})
	}

	for step := 1; step <= STEPS; step++ {
		var sb strings.Builder
		sb.WriteString(string(polymer[0]))
		for i := 1; i < len(polymer); i++ {
			pair := fmt.Sprintf("%s%s", string(polymer[i-1]), string(polymer[i]))
			for _, rule := range rules {
				if pair == rule.pattern {
					pair = rule.replace
					break
				}
			}
			sb.WriteString(pair)
		}
		polymer = sb.String()
		// fmt.Println("After step (", step, ")")
		// fmt.Println(polymer)
	}

	least, most := findLeastAndMostSymbolCount(polymer)
	fmt.Printf("Answer is %v", most-least)
}

type pairInsertion struct {
	pattern string
	replace string
}

func findLeastAndMostSymbolCount(polymer string) (least int, most int) {
	runes := make(map[rune]int)
	for _, s := range polymer {
		runes[s] = runes[s] + 1
	}

	small := math.MaxInt
	big := math.MinInt
	for _, count := range runes {
		if count < small {
			small = count
		}
		if count > big {
			big = count
		}
	}
	least = small
	most = big
	return
}
