package day1

import (
	"math"
	"sort"
)

func totalDistance(list1 []int, list2 []int) int {
	sort.Ints(list1)
	sort.Ints(list2)
	sum := 0
	for i := 0; i < len(list1); i++ {
		sum += (int)(math.Abs((float64)(list1[i] - list2[i])))
	}
	return sum
}

func calculateSimilarity(list1 []int, list2 []int) int {
	frequency := calculateFrequency(list2)
	similarity := 0
	for i := 0; i < len(list1); i++ {
		similarity += list1[i] * frequency[list1[i]]
	}
	return similarity
}

func calculateFrequency(list []int) map[int]int {
	frequency := make(map[int]int)
	for i := 0; i < len(list); i++ {
		listVal := list[i]
		existingFrequency := frequency[listVal]
		if existingFrequency != 0 {
			frequency[listVal] = existingFrequency + 1
		} else {
			frequency[listVal] = 1
		}
	}
	return frequency
}
