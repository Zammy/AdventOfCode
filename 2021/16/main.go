package main

import (
	"fmt"
	"strconv"
	"strings"

	"evgeni.com/util"
)

const (
	LITERAL = 4
)

var sumVersions int

func main() {
	input, err := util.ReadLines("input")
	if err != nil {
		panic(err)
	}

	transmision := convertToBinary(input[0])
	var pointer int
	for pointer < len(transmision) {
		parsePacket(transmision, &pointer)
		forwardWhileZero(transmision, &pointer)
	}

	fmt.Println("Sum of all versions is: ", sumVersions)
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

func parsePacket(transmision string, pointer *int) {
	pckVersion, pckType := parseVersionAndPacketType(transmision, pointer)
	if pckType == LITERAL {
		literlValue := parseLiteralPacket(transmision, pointer)
		fmt.Println(literlValue)
	} else {
		parseOperatorPacket(transmision, pointer)
	}
	sumVersions += pckVersion
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

func parseOperatorPacket(transmision string, pointer *int) {
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
			parsePacket(transmision, pointer)
		}
	} else {
		numberSubPackets, err := strconv.ParseInt(transmision[*pointer:*pointer+11], 2, 64)
		*pointer += 11
		if err != nil {
			panic(err)
		}
		for i := 0; i < int(numberSubPackets); i++ {
			parsePacket(transmision, pointer)
		}
	}
}
