package day1

import (
	"strconv"
	"strings"
)

func ComputeIncreasesNumber(input string) int {
	measurements := parseInput(input)
	previousMeasurement := measurements[0]
	count := 0
	for i := 1; i < len(measurements); i++ {
		if measurements[i] > previousMeasurement {
			count += 1
		}
		previousMeasurement = measurements[i]
	}
	return count
}

func ComputeIncreasesNumberSlidingWindow(input string) int {
	measurements := parseInput(input)
	previousSlidingWindowValue := measurements[0] + measurements[1] + measurements[2]
	count := 0
	for i := 1; i < len(measurements)-2; i++ {
		currentSlidingWindowValue := measurements[i] + measurements[i+1] + measurements[i+2]
		if currentSlidingWindowValue > previousSlidingWindowValue {
			count += 1
		}
		previousSlidingWindowValue = currentSlidingWindowValue
	}
	return count
}

func parseInput(input string) []int {
	var measurements []int
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		measurement, _ := strconv.Atoi(line)
		measurements = append(measurements, measurement)
	}
	return measurements
}
