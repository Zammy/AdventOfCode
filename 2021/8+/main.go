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
	digits, output := parseInput(input)

	sum := 0
	for i, outputRow := range digits {
		var display sevenSegDisplay

		one := filterWithLen(outputRow, 2)[0]
		display.vertUpRight = one
		display.vertDownRight = one

		seven := filterWithLen(outputRow, 3)[0]
		display.horzUp = digitWithoutSegments(seven, one)

		four := filterWithLen(outputRow, 4)[0]
		display.vertUpLeft = digitWithoutSegments(four, one)
		display.horzMid = display.vertUpLeft

		allKnownSegments := fmt.Sprintf("%v%v%v", one, display.horzUp, display.vertUpLeft)
		allDigitsWithSigSegments := filterWithLen(outputRow, 6)
		for i, digit := range allDigitsWithSigSegments {
			seg := digitWithoutSegments(digit, allKnownSegments)
			if len(seg) == 1 {
				display.horzDown = seg
				display.nine = digit
				allDigitsWithSigSegments = util.SliceRemoveAtIndex(allDigitsWithSigSegments, i)
				break
			}
		}
		//either 0 or 6, if we remove 9 from them we will get down left horz seg
		display.vertDownLeft = digitWithoutSegments(allDigitsWithSigSegments[0], display.nine)
		for _, digit := range allDigitsWithSigSegments {
			count := countSegmentsInDigit(digit, one)
			if count == 2 {
				display.zero = digit

				possibleSegs := display.horzMid
				for _, seg := range possibleSegs {
					if countSegmentInDigit(display.zero, seg) == 1 {
						display.vertUpLeft = string(seg)
					} else {
						display.horzMid = string(seg)
					}
				}
			} else {
				display.six = digit

				for _, seg := range one {
					if countSegmentInDigit(display.six, seg) == 1 {
						display.vertDownRight = string(seg)
					} else {
						display.vertUpRight = string(seg)
					}
				}
			}
		}

		sum += display.ParseOutput(output[i])
	}

	fmt.Println("Sum is ", sum)
}

func parseInput(input []string) (digits [][]string, output [][]string) {
	digits = [][]string{}
	output = [][]string{}
	for _, v := range input {
		split := strings.Split(v, "|")
		digitsArray := strings.Split(split[0], " ")
		digitsArray = util.SliceRemoveEmpty(digitsArray)
		digits = append(digits, digitsArray)
		outputArray := strings.Split(split[1], " ")
		outputArray = util.SliceRemoveEmpty(outputArray)
		output = append(output, outputArray)
	}
	return
}

type sevenSegDisplay struct {
	horzUp, horzMid, horzDown, vertDownLeft, vertDownRight, vertUpLeft, vertUpRight string

	zero, six, nine string
}

func (d *sevenSegDisplay) String() string {
	var stringBuilder strings.Builder
	stringBuilder.WriteString(fmt.Sprintf("[    %v    ]\n", d.horzUp))
	stringBuilder.WriteString(fmt.Sprintf("[%v] - [%v]\n", d.vertUpLeft, d.vertUpRight))
	stringBuilder.WriteString(fmt.Sprintf("[    %v    ]\n", d.horzMid))
	stringBuilder.WriteString(fmt.Sprintf("[%v] - [%v]\n", d.vertDownLeft, d.vertDownRight))
	stringBuilder.WriteString(fmt.Sprintf("[    %v    ]\n", d.horzDown))
	return stringBuilder.String()
}

func (d *sevenSegDisplay) ParseOutput(output []string) int {
	result := 0
	for i, digit := range output {
		integer := d.ParseDigit(digit)
		result += integer * int(math.Pow10(len(output)-i-1))
	}
	return result
}

func (d *sevenSegDisplay) ParseDigit(digit string) int {
	length := len(digit)
	switch length {
	case 2:
		return 1
	case 3:
		return 7
	case 4:
		return 4
	case 5:
		downLeft := countSegmentInDigit(digit, rune(d.vertDownLeft[0]))
		if downLeft == 1 {
			return 2
		}
		upLeft := countSegmentInDigit(digit, rune(d.vertUpLeft[0]))
		if upLeft == 1 {
			return 5
		}
		if downLeft == 0 && upLeft == 0 {
			return 3
		}
		panic("Should never reach here")
	case 6:
		if countSegmentsInDigit(digit, d.zero) == 6 {
			return 0
		}
		if countSegmentsInDigit(digit, d.six) == 6 {
			return 6
		}
		if countSegmentsInDigit(digit, d.nine) == 6 {
			return 9
		}
		panic("Should never reach here")
	case 7:
		return 8
	default:
		panic("Should never reach here")
	}
}

func filterWithLen(input []string, length int) []string {
	result := []string{}
	for _, v := range input {
		if len(v) == length {
			result = append(result, v)
		}
	}
	return result
}

func digitWithoutSegments(digit string, segments string) string {
	for _, v := range segments {
		digit = strings.ReplaceAll(digit, string(v), "")
	}
	return digit
}

func countSegmentsInDigit(digit string, segments string) int {
	count := 0
	for _, seg := range segments {
		if strings.Contains(digit, string(seg)) {
			count += 1
		}
	}
	return count
}

func countSegmentInDigit(digit string, segment rune) int {
	count := 0
	if strings.Contains(digit, string(segment)) {
		count += 1
	}
	return count
}
