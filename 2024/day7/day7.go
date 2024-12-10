package day7

import (
	"math"
	"strconv"
	"strings"
)

func computeCalibrations(equationsStr string) int {
	equations := parseInput(equationsStr)
	sum := 0
	for _, equation := range equations {
		if isSolvable(equation) {
			sum += equation[0]
		}
	}
	return sum
}

func isSolvable(equation []int) bool {
	startPattern := (int)(math.Pow(2, (float64)(len(equation)-2)) - 1)
	for pattern := startPattern; pattern >= 0; pattern-- {
		currValue := equation[1]
		currPattern := pattern
		for i := 2; i < len(equation); i++ {
			if currPattern%2 == 0 {
				currValue += equation[i]
			} else {
				currValue *= equation[i]
			}
			currPattern = currPattern >> 1
		}
		if currValue == equation[0] {
			return true
		}
	}
	return false
}

func computeCalibrationsWithConcatenation(equationsStr string) int {
	equations := parseInput(equationsStr)
	sum := 0
	for _, equation := range equations {
		if isSolvableWithConcatenation(equation) {
			sum += equation[0]
		}
	}
	return sum
}

func isSolvableWithConcatenation(equation []int) bool {
	startPattern := (int)(math.Pow(3, (float64)(len(equation)-2)) - 1)
	for pattern := startPattern; pattern >= 0; pattern-- {
		currValue := equation[1]
		currPattern := pattern
		for i := 2; i < len(equation); i++ {
			if currPattern%3 == 0 {
				currValue += equation[i]
			} else if currPattern%3 == 1 {
				currValue *= equation[i]
			} else {
				currValue, _ = strconv.Atoi(strconv.Itoa(currValue) + strconv.Itoa(equation[i]))
			}
			currPattern = int(math.Floor(float64(currPattern) / 3))
		}
		if currValue == equation[0] {
			return true
		}
	}
	return false
}

func parseInput(equationsStr string) [][]int {
	rows := strings.Split(equationsStr, "\n")
	equations := make([][]int, len(rows))
	for i := 0; i < len(rows); i++ {
		row := rows[i]
		rowSplit := strings.Split(row, ":")
		values := strings.Split(rowSplit[1], " ")
		equation := make([]int, len(values))
		equation[0], _ = strconv.Atoi(rowSplit[0])
		for j := 1; j < len(values); j++ {
			equation[j], _ = strconv.Atoi(values[j])
		}
		equations[i] = equation
	}
	return equations
}
