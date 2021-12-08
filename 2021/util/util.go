package util

import (
	"bufio"
	"os"
	"strconv"
)

func ReadLines(filePath string) ([]string, error) {
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0666)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	res := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		res = append(res, line)
	}
	if err = scanner.Err(); err != nil {
		return nil, err
	}
	return res, nil
}

func SliceCountTrue(slice []bool) int {
	count := 0
	for _, v := range slice {
		if v {
			count += 1
		}
	}
	return count
}

func ParseInt(input string) int {
	v, err := strconv.ParseInt(input, 0, 0)
	if err != nil {
		panic(err)
	}
	return int(v)
}
