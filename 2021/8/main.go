package main

import (
	"fmt"
	"strings"

	"evgeni.com/util"
)

func main() {
	input, err := util.ReadLines("input")
	if err != nil {
		panic(err)
	}
	output := filterOnlyOutputValue(input)
	sumAllKnown := 0
	for _, out := range output {
		for _, digit := range out {
			len := len(digit)
			if len == 2 || len == 3 || len == 4 || len == 7 {
				sumAllKnown += 1
			}
		}
	}
	fmt.Printf("Num of known digits: %v", sumAllKnown)
}

func filterOnlyOutputValue(input []string) [][]string {
	output := [][]string{}
	for _, v := range input {
		outputOnly := strings.Split(v, "|")[1]
		outputArray := strings.Split(outputOnly, " ")
		output = append(output, outputArray)
	}
	return output
}
