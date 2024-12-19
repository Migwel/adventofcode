package day19

import (
	"strings"
)

func computePossibleDesigns(input string) int {
	availablePatterns, desiredDesigns := parseInput(input)
	count := countPossibleDesigns(availablePatterns, desiredDesigns)
	return count
}

func countPossibleDesigns(availablePatterns [][]rune, desiredDesigns [][]rune) int {
	count := 0
	cache := make(map[string]int)
	for _, desiredDesign := range desiredDesigns {
		filteredAvailablePatterns := filterAvailablePatterns(availablePatterns, desiredDesign)
		if countAllDesignWays(filteredAvailablePatterns, desiredDesign, cache) != 0 {
			count += 1
		}
	}
	return count
}

func CountAllDesignWays(input string) int {
	availablePatterns, desiredDesigns := parseInput(input)
	count := 0
	cache := make(map[string]int)
	for _, desiredDesign := range desiredDesigns {
		filteredAvailablePatterns := filterAvailablePatterns(availablePatterns, desiredDesign)
		ways := countAllDesignWays(filteredAvailablePatterns, desiredDesign, cache)
		count += ways
	}
	return count
}

func countAllDesignWays(availablePatterns [][]rune, desiredDesign []rune, cache map[string]int) int {
	totalNbPossibleWays := 0
	if cachedNbPossibleWays, present := cache[string(desiredDesign)]; present {
		return cachedNbPossibleWays
	}
	for _, availablePattern := range availablePatterns {
		if len(desiredDesign) == 0 {
			return 1
		}
		if !camBeUsed(availablePattern, desiredDesign) {
			continue
		}
		nbPossibleWays := countAllDesignWays(availablePatterns, desiredDesign[len(availablePattern):], cache)
		totalNbPossibleWays += nbPossibleWays
	}
	cache[string(desiredDesign)] = totalNbPossibleWays
	return totalNbPossibleWays
}

func filterAvailablePatterns(availablePatterns [][]rune, desiredDesign []rune) [][]rune {
	var filteredAvailablePatterns [][]rune
	for _, availablePattern := range availablePatterns {
		if containsPattern(desiredDesign, availablePattern) {
			filteredAvailablePatterns = append(filteredAvailablePatterns, availablePattern)
		}
	}
	return filteredAvailablePatterns
}

func containsPattern(desiredDesign []rune, availablePattern []rune) bool {
outerLoop:
	for idx, _ := range desiredDesign {
		for i := 0; i < len(availablePattern); i++ {
			if idx+i >= len(desiredDesign) {
				return false
			}
			if desiredDesign[idx+i] != availablePattern[i] {
				continue outerLoop
			}
		}
		return true
	}
	return false
}

func isPossibleDesign(availablePatterns [][]rune, desiredDesign []rune, currentIndex int) bool {
	for _, availablePattern := range availablePatterns {
		if !camBeUsed(availablePattern, desiredDesign[currentIndex:]) {
			continue
		}
		if currentIndex+len(availablePattern) == len(desiredDesign) {
			return true
		}
		if isPossibleDesign(availablePatterns, desiredDesign, currentIndex+len(availablePattern)) {
			return true
		}
	}
	return false
}

func camBeUsed(availablePattern []rune, desiredDesign []rune) bool {
	for idx, _ := range availablePattern {
		if idx >= len(desiredDesign) {
			return false
		}
		if availablePattern[idx] != desiredDesign[idx] {
			return false
		}
	}
	return true
}

func parseInput(input string) ([][]rune, [][]rune) {
	lines := strings.Split(input, "\n")
	availablePatterns := parseAvailablePatterns(string(lines[0]))
	desiredDesigns := parseDesiredDesigns(lines[2:])
	return availablePatterns, desiredDesigns
}

func parseDesiredDesigns(lines []string) [][]rune {
	desiredDesigns := make([][]rune, len(lines))
	for idx, line := range lines {
		desiredDesign := []rune(line)
		desiredDesigns[idx] = desiredDesign
	}
	return desiredDesigns

}

func parseAvailablePatterns(input string) [][]rune {
	var availablePatterns [][]rune
	availablePatternsInput := strings.Split(input, ", ")
	for _, availablePatternInput := range availablePatternsInput {
		availablePatterns = append(availablePatterns, []rune(availablePatternInput))
	}
	return availablePatterns
}
