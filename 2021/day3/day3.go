package day3

import (
	"math"
	"strconv"
	"strings"
)

func ComputePowerConsumption(input string) int {
	binaryNumbers := parseInput(input)
	totals := make([]int, len(binaryNumbers[0]))
	for _, binaryNumber := range binaryNumbers {
		for idx, val := range binaryNumber {
			intVal, _ := strconv.Atoi(string(val))
			totals[idx] += intVal
		}
	}
	gammaRate, epsilonRate := computeRate(totals, len(binaryNumbers)/2)
	return gammaRate * epsilonRate
}

func ComputeLifeSupportRating(input string) int {
	binaryNumbers := parseInput(input)
	oxygenGeneratorRating := computeOxygenGeneratorRating(binaryNumbers)
	co2ScrubberRating := computeCO2ScrubberRating(binaryNumbers)
	return oxygenGeneratorRating * co2ScrubberRating
}

func computeOxygenGeneratorRating(binaryNumbers []string) int {
	for i := 0; i < len(binaryNumbers[0]); i++ {
		total := computeTotalCurrentIndex(binaryNumbers, i)
		if total >= int(math.Ceil(float64(len(binaryNumbers))/2.0)) {
			binaryNumbers = filterNumbers(binaryNumbers, i, "1")
		} else {
			binaryNumbers = filterNumbers(binaryNumbers, i, "0")
		}
		if len(binaryNumbers) == 1 {
			binaryNumberDecimal, _ := strconv.ParseInt(binaryNumbers[0], 2, 64)
			return int(binaryNumberDecimal)
		}
	}
	return -1
}

func computeCO2ScrubberRating(binaryNumbers []string) int {
	for i := 0; i < len(binaryNumbers[0]); i++ {
		total := computeTotalCurrentIndex(binaryNumbers, i)
		if total >= int(math.Ceil(float64(len(binaryNumbers))/2.0)) {
			binaryNumbers = filterNumbers(binaryNumbers, i, "0")
		} else {
			binaryNumbers = filterNumbers(binaryNumbers, i, "1")
		}
		if len(binaryNumbers) == 1 {
			binaryNumberDecimal, _ := strconv.ParseInt(binaryNumbers[0], 2, 64)
			return int(binaryNumberDecimal)
		}
	}
	return -1
}

func filterNumbers(binaryNumbers []string, idx int, filteredBy string) []string {
	var filteredNumbers []string
	for _, binarybinaryNumber := range binaryNumbers {
		if string(binarybinaryNumber[idx]) == filteredBy {
			filteredNumbers = append(filteredNumbers, binarybinaryNumber)
		}
	}
	return filteredNumbers
}

func computeTotalCurrentIndex(binaryNumbers []string, idx int) int {
	total := 0
	for _, binaryNumber := range binaryNumbers {
		intVal, _ := strconv.Atoi(string(binaryNumber[idx]))
		total += intVal
	}
	return total
}

func computeRate(totals []int, majorityLimit int) (int, int) {
	gammaRateStr := ""
	epsilonRateStr := ""
	for _, total := range totals {
		if total >= majorityLimit {
			gammaRateStr += "1"
			epsilonRateStr += "0"
		} else {
			gammaRateStr += "0"
			epsilonRateStr += "1"
		}
	}
	gammaRate, _ := strconv.ParseInt(gammaRateStr, 2, 64)
	epsilonRate, _ := strconv.ParseInt(epsilonRateStr, 2, 64)
	return int(gammaRate), int(epsilonRate)
}

func parseInput(input string) []string {
	lines := strings.Split(input, "\n")
	binaryNumbers := make([]string, len(lines))
	for idx, line := range lines {
		binaryNumbers[idx] = line
	}
	return binaryNumbers
}
