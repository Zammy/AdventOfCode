package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"evgeni.com/util"
)

func main() {
	input, err := util.ReadLines("input")
	if err != nil {
		panic(err)
	}

	boards := make([]*bingoBoard, 0)
	numbersInStrings := strings.Split(input[0], ",")
	numbers := make([]int, len(numbersInStrings))
	for i, numStr := range numbersInStrings {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			panic(err)
		}
		numbers[i] = num
	}

	boardsInput := input[2:]
	for i := 5; i <= len(boardsInput); i += 6 {
		boardData := boardsInput[i-5 : i]
		boards = append(boards, createBoard(boardData))
	}

	for i := 5; i <= len(numbers); i++ {
		currentNumbers := numbers[0:i]

		for _, board := range boards {
			if board.Check(currentNumbers) {
				sumAllNotMarked := board.SumAllNotMarked(currentNumbers)
				score := sumAllNotMarked * currentNumbers[len(currentNumbers)-1]
				fmt.Printf("Final score: %v", score)
				os.Exit(0)
			}
		}
	}
}

type bingoBoard struct {
	board [5][5]int
}

func createBoard(input []string) *bingoBoard {
	var b bingoBoard
	for rowCount, rowStr := range input {
		row := parseRow(rowStr)
		for i := 0; i < len(row); i++ {
			b.board[rowCount][i] = row[i]
		}
	}
	return &b
}

func (b *bingoBoard) Check(numbers []int) bool {
	bingo := false
	for x := 0; x < 5; x++ {
		bingo = true
		for y := 0; y < 5; y++ {
			if !arrayContains(numbers, b.board[x][y]) {
				bingo = false
				break
			}
		}
		if bingo {
			return true
		}
	}
	for x := 0; x < 5; x++ {
		bingo = true
		for y := 0; y < 5; y++ {
			if !arrayContains(numbers, b.board[y][x]) {
				bingo = false
				break
			}
		}
		if bingo {
			return true
		}
	}
	return bingo
}

func (b *bingoBoard) String() string {
	var builder strings.Builder
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			builder.WriteString(strconv.Itoa(b.board[x][y]))
			builder.WriteString(" ")
		}
		builder.WriteString("\n")
	}
	return builder.String()
}

func (b *bingoBoard) SumAllNotMarked(numbers []int) int {
	sum := 0
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			num := b.board[x][y]
			if !arrayContains(numbers, num) {
				sum += num
			}
		}
	}
	return sum
}

func arrayContains(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func parseRow(input string) []int {
	result := []int{}
	for i := 0; i < len(input); i += 3 {
		numStr := fmt.Sprintf("%v%v", string(input[i]), string(input[i+1]))
		numStr = strings.Trim(numStr, " ")
		num, error := strconv.Atoi(numStr)
		if error != nil {
			panic(error)
		}
		result = append(result, num)
	}
	return result
}
