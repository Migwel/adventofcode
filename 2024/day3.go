package main

import (
	"regexp"
	"strconv"
)

func runMul(data string) int {
	sum := 0
	mulOccurrences := extractMul(data)
	for i := 0; i < len(mulOccurrences); i++ {
		nb1, _ := strconv.Atoi((string)(mulOccurrences[i][1]))
		nb2, _ := strconv.Atoi((string)(mulOccurrences[i][2]))
		sum += nb1 * nb2
	}
	return sum
}

func extractMul(data string) [][][]byte {
	re := regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)`)
	return re.FindAllSubmatch([]byte(data), -1)
}

func runMulWithConditions(data string) int {
	sum := 0
	mulEnabled := true
	mulOccurrences := extractMulWithConditions(data)
	for i := 0; i < len(mulOccurrences); i++ {
		matchedStr := (string)(mulOccurrences[i][0])
		if matchedStr == "do()" {
			mulEnabled = true
			continue
		}
		if matchedStr == "don't()" {
			mulEnabled = false
			continue
		}
		if !mulEnabled {
			continue
		}
		nb1, _ := strconv.Atoi((string)(mulOccurrences[i][3]))
		nb2, _ := strconv.Atoi((string)(mulOccurrences[i][4]))
		sum += nb1 * nb2
	}
	return sum
}

func extractMulWithConditions(data string) [][][]byte {
	re := regexp.MustCompile(`(do\(\)|don't\(\)|(mul\(([0-9]+),([0-9]+)\)))`)
	return re.FindAllSubmatch([]byte(data), -1)
}
