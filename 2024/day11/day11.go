package day11

import (
	"strconv"
	"strings"
)

func countStones(input string, nbBlinks int) int {
	stoneArrangement := parseInput(input)
	for i := 0; i < nbBlinks; i++ {
		stoneArrangement = blink(stoneArrangement)
	}
	stones := 0
	for _, val := range stoneArrangement {
		stones += val
	}
	return stones
}

func blink(stonesFrequency map[int]int) map[int]int {
	newStonesFrequency := make(map[int]int)
	for stoneValue, frequency := range stonesFrequency {
		if stoneValue == 0 {
			addToFrequency(newStonesFrequency, 1, frequency)
			continue
		}
		if numberOfDigits(stoneValue)%2 == 0 {
			firstHalf, secondHalf := splitStone(stoneValue)
			addToFrequency(newStonesFrequency, firstHalf, frequency)
			addToFrequency(newStonesFrequency, secondHalf, frequency)
			continue
		}
		addToFrequency(newStonesFrequency, stoneValue*2024, frequency)
	}
	return newStonesFrequency
}

func addToFrequency(stonesFrequency map[int]int, stoneValue, frequency int) {
	existingFrequency := stonesFrequency[stoneValue]
	stonesFrequency[stoneValue] = existingFrequency + frequency
}

func numberOfDigits(stoneValue int) int {
	stoneValueStr := strconv.Itoa(stoneValue)
	return len(stoneValueStr)
}

func splitStone(stoneValue int) (int, int) {
	stoneValueStr := strconv.Itoa(stoneValue)
	firstHalf, _ := strconv.Atoi(stoneValueStr[:len(stoneValueStr)/2])
	secondHalf, _ := strconv.Atoi(stoneValueStr[len(stoneValueStr)/2:])
	return firstHalf, secondHalf

}

func parseInput(input string) map[int]int {
	values := strings.Split(input, " ")
	stonesFrequency := make(map[int]int)
	for _, val := range values {
		intVal, _ := strconv.Atoi(val)
		existingValue := stonesFrequency[intVal]
		stonesFrequency[intVal] = existingValue + 1
	}
	return stonesFrequency
}
