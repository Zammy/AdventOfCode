package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"evgeni.com/util"
)

const (
	SUM          = 0
	PRODUCT      = 1
	MINIMUM      = 2
	MAXIMUM      = 3
	LITERAL      = 4
	GREATERHTHAN = 5
	LESSTHAN     = 6
	EQUAL        = 7
)

func main() {
	input, err := util.ReadLines("input")
	if err != nil {
		panic(err)
	}

	transmision := convertToBinary(input[0])
	var pointer int
	var result int
	for pointer < len(transmision) {
		result = parsePacket(transmision, &pointer)
		forwardWhileZero(transmision, &pointer)
	}
	fmt.Println("Result : ", result)
}

func convertToBinary(input string) string {
	var sb strings.Builder
	for _, r := range input {
		result, err := strconv.ParseUint(string(r), 16, 64)
		if err != nil {
			panic(err)
		}
		inBinary := strconv.FormatUint(result, 2)
		for i := 0; i < 4-len(inBinary); i++ {
			sb.WriteString("0")
		}
		sb.WriteString(inBinary)
	}
	return sb.String()
}

func parsePacket(transmision string, pointer *int) int {
	_, pckType := parseVersionAndPacketType(transmision, pointer)
	if pckType == LITERAL {
		return parseLiteralPacket(transmision, pointer)
	}
	values := parseOperatorPacket(transmision, pointer)
	switch pckType {
	case SUM:
		sum := 0
		for _, v := range values {
			sum += v
		}
		return sum
	case PRODUCT:
		product := 1
		for _, v := range values {
			product *= v
		}
		return product
	case MINIMUM:
		min := math.MaxInt
		for _, v := range values {
			if v < min {
				min = v
			}
		}
		return min
	case MAXIMUM:
		max := math.MinInt
		for _, v := range values {
			if max < v {
				max = v
			}
		}
		return max
	case GREATERHTHAN:
		if values[0] > values[1] {
			return 1
		}
		return 0
	case LESSTHAN:
		if values[0] < values[1] {
			return 1
		}
		return 0
	case EQUAL:
		if values[0] == values[1] {
			return 1
		}
		return 0
	default:
		panic("Corrupted parsing. Should never reach here.")
	}
}

func parseVersionAndPacketType(transmision string, pointer *int) (int, int) {
	pckVer, err := strconv.ParseInt(transmision[*pointer:*pointer+3], 2, 64)
	if err != nil {
		panic(err)
	}
	*pointer += 3
	pckType, err := strconv.ParseInt(transmision[*pointer:*pointer+3], 2, 64)
	if err != nil {
		panic(err)
	}
	*pointer += 3
	return int(pckVer), int(pckType)
}

func parseLiteralPacket(transmision string, pointer *int) int {
	var sb strings.Builder
	for {
		binary := transmision[*pointer : *pointer+5]
		sb.WriteString(binary[1:])
		*pointer += 5
		if binary[0] == '0' {
			//last packet
			break
		}
	}

	result, err := strconv.ParseInt(sb.String(), 2, 64)
	if err != nil {
		panic(nil)
	}
	return int(result)
}

func forwardWhileZero(transmision string, pointer *int) {
	for {
		if len(transmision) <= *pointer || transmision[*pointer] != '0' {
			break
		}
		*pointer++
	}
}

func parseOperatorPacket(transmision string, pointer *int) []int {
	var values []int
	lengthTypeID := transmision[*pointer]
	*pointer++
	if lengthTypeID == '0' {
		totalLength, err := strconv.ParseInt(transmision[*pointer:*pointer+15], 2, 64)
		*pointer += 15
		targetPointer := *pointer + int(totalLength)
		if err != nil {
			panic(err)
		}
		for *pointer < targetPointer {
			values = append(values, parsePacket(transmision, pointer))
		}
	} else {
		numberSubPackets, err := strconv.ParseInt(transmision[*pointer:*pointer+11], 2, 64)
		*pointer += 11
		if err != nil {
			panic(err)
		}
		for i := 0; i < int(numberSubPackets); i++ {
			values = append(values, parsePacket(transmision, pointer))
		}
	}
	return values
}
