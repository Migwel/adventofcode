package day2

import (
	"math"
)

func checkSafety(data [][]int) int {
	nbSafe := 0
	for i := 0; i < len(data); i++ {
		if isSafe(data[i]) {
			nbSafe += 1
		}
	}
	return nbSafe
}

func checkSafetyWithDampener(data [][]int) int {
	nbSafe := 0
	for i := 0; i < len(data); i++ {
		reports := data[i]
		if isSafe(reports) {
			nbSafe += 1
			continue
		}
		for j := 0; j < len(reports); j++ {
			updatedReport := append([]int{}, reports...)
			updatedReport = append(updatedReport[:j], updatedReport[j+1:]...)
			if isSafe(updatedReport) {
				nbSafe += 1
				break
			}
		}
	}
	return nbSafe
}

func isSafe(reports []int) bool {
	prevLevel := reports[0]
	ascending := reports[0] < reports[1]
	for i := 1; i < len(reports); i++ {
		currLevel := reports[i]
		diffVal := (int)(math.Abs((float64)(currLevel - prevLevel)))
		if diffVal < 1 || diffVal > 3 {
			return false
		}
		if (reports[i-1] < reports[i]) != ascending {
			return false
		}
		prevLevel = currLevel
	}
	return true
}
