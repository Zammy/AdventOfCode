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

	nums := parseInput(input)
	maxMgn := math.MinInt
	for i, num1 := range nums {
		for _, num2 := range nums[:i+1] {
			for y, num := range []*snailNum{num1.Add(*num2), num2.Add(*num1)} {
				num = reduce(num)
				magnitude := num.Magnitude()
				if y == 0 {
					fmt.Printf("%v + %v = %v\n", num1, num2, magnitude)
				} else {
					fmt.Printf("%v + %v = %v\n", num2, num1, magnitude)
				}
				if magnitude > maxMgn {
					maxMgn = magnitude
				}
			}
		}
	}

	fmt.Println("Max magnitude: ", maxMgn)
}

func reduce(input *snailNum) *snailNum {
	num := input.Copy()
	for reduced := false; ; {
		if num.ReduceExplode(&reduceData{}, 1) {
			reduced = true
			continue
		}
		reduced = num.ReduceSplit()
		if !reduced {
			break
		}
	}
	return num
}

type snailNum struct {
	parent, left, right *snailNum
	value               int
}

func parseInput(input []string) []*snailNum {
	var nums []*snailNum
	for _, row := range input {
		nums = append(nums, parseSnailNumber(row))
	}
	return nums
}

func parseSnailNumber(input string) *snailNum {
	var index int
	return newSnailNumFromString(input, &index)
}

func newSnailNumFromString(input string, index *int) *snailNum {
	var s snailNum
	*index++
	if input[*index] == '[' {
		s.left = newSnailNumFromString(input, index)
		s.left.parent = &s
	} else {
		s.left = newSnailNumFromNum(rune(input[*index]))
	}
	*index++
	if input[*index] != ',' {
		panic("Something hasgone wrong")
	}
	*index++
	if input[*index] == '[' {
		s.right = newSnailNumFromString(input, index)
		s.right.parent = &s
	} else {
		s.right = newSnailNumFromNum(rune(input[*index]))
	}
	*index++
	if input[*index] != ']' {
		panic("Something hasgone wrong")
	}
	return &s
}

func newSnailNumFromNum(input rune) *snailNum {
	num := util.ParseInt(string(input))
	return &snailNum{value: num}
}

func (s snailNum) Add(other snailNum) *snailNum {
	return &snailNum{left: &s, right: &other}
}

func (s *snailNum) IsLeaf() bool {
	return s.left == nil && s.right == nil
}

func (s *snailNum) String() string {
	var sb strings.Builder
	if s.IsLeaf() {
		sb.WriteString(fmt.Sprint(s.value))
	} else {
		sb.WriteRune('[')

		if s.left != nil {
			sb.WriteString(s.left.String())
		}
		sb.WriteRune(',')
		if s.right != nil {
			sb.WriteString(s.right.String())
		}
		sb.WriteRune(']')
	}

	return sb.String()
}

type reduceData struct {
	lastLeafNode      *snailNum
	addToNextLeafNode int
	exploded          bool
}

func (s *snailNum) ReduceExplode(data *reduceData, depth int) bool {
	if depth > 4 && !data.exploded && s.left.IsLeaf() && s.right.IsLeaf() {
		// fmt.Println("Exploding ", s)
		if data.lastLeafNode != nil {
			data.lastLeafNode.value += s.left.value
		}
		data.addToNextLeafNode = s.right.value
		s.left, s.right = nil, nil
		s.value = 0
		data.exploded = true
		return false
	}
	if s.left.IsLeaf() {
		if data.exploded {
			s.left.value += data.addToNextLeafNode
			return true
		} else {
			data.lastLeafNode = s.left
		}
	} else {
		if s.left.ReduceExplode(data, depth+1) {
			return true
		}
	}
	if s.right.IsLeaf() {
		if data.exploded {
			s.right.value += data.addToNextLeafNode
			return true
		} else {
			data.lastLeafNode = s.right
		}
	} else {
		if s.right.ReduceExplode(data, depth+1) {
			return true
		}
	}
	return false
}

func (s *snailNum) ReduceSplit() bool {
	if s.left.IsLeaf() {
		if s.left.value > 9 {
			// fmt.Println("Splitting  ", s.left)
			s.left.left = &snailNum{value: s.left.value / 2}
			s.left.right = &snailNum{value: int(math.Round(float64(s.left.value) / 2.0))}
			return true
		}
	} else {
		if s.left.ReduceSplit() {
			return true
		}
	}
	if s.right.IsLeaf() {
		if s.right.value > 9 {
			// fmt.Println("Splitting  ", s.right)
			s.right.left = &snailNum{value: s.right.value / 2}
			s.right.right = &snailNum{value: int(math.Round(float64(s.right.value) / 2.0))}
			return true
		}
	} else {
		if s.right.ReduceSplit() {
			return true
		}
	}
	return false
}

func (s *snailNum) Magnitude() int {
	if s.IsLeaf() {
		return s.value
	}
	return s.left.Magnitude()*3 + s.right.Magnitude()*2
}

func (s *snailNum) Copy() *snailNum {
	if s.IsLeaf() {
		return &snailNum{value: s.value}
	}
	return &snailNum{left: s.left.Copy(), right: s.right.Copy(), value: s.value}
}
