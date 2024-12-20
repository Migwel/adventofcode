package day6

import (
	"strconv"
	"strings"
)

type Fish int

func ComputeNumberLanternFishes(input string, nbDays int) int {
	fishes := parseInput(input)
	fishesPerRemainingDay := initalizeFishesMap(fishes)
	for day := 0; day < nbDays; day++ {
		newFishesPerRemainingDay := make(map[int]int)
		for remainingDays, nbFishes := range fishesPerRemainingDay {
			if remainingDays > 0 {
				if val, ok := newFishesPerRemainingDay[remainingDays-1]; ok {
					newFishesPerRemainingDay[remainingDays-1] = val + nbFishes
				} else {
					newFishesPerRemainingDay[remainingDays-1] = nbFishes
				}
			} else {
				if val, ok := newFishesPerRemainingDay[6]; ok {
					newFishesPerRemainingDay[6] = val + nbFishes
				} else {
					newFishesPerRemainingDay[6] = nbFishes
				}
				newFishesPerRemainingDay[8] = nbFishes
			}
		}
		fishesPerRemainingDay = newFishesPerRemainingDay
	}
	count := 0
	for _, nbFishes := range fishesPerRemainingDay {
		count += nbFishes
	}
	return count
}

func initalizeFishesMap(fishes []int) map[int]int {
	fishesPerRemainingDay := make(map[int]int)
	for _, fish := range fishes {
		if val, ok := fishesPerRemainingDay[fish]; ok {
			fishesPerRemainingDay[fish] = val + 1
		} else {
			fishesPerRemainingDay[fish] = 1
		}
	}
	return fishesPerRemainingDay
}

func parseInput(input string) []int {
	ages := strings.Split(input, ",")
	fishes := make([]int, len(ages))
	for idx, age := range ages {
		ageValue, _ := strconv.Atoi(age)
		fishes[idx] = ageValue
	}
	return fishes
}
