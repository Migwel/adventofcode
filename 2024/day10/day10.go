package day10

import (
	"strconv"
	"strings"
)

type Position struct {
	x int
	y int
}

func computeTrailheadsScore(input string) int {
	topographicMap := parseInput(input)
	return findTrailheadsScore(topographicMap)
}

func computeTrailheadsRating(input string) int {
	topographicMap := parseInput(input)
	return findTrailheadsRatings(topographicMap)
}

func findTrailheadsScore(topographicMap [][]int) int {
	score := 0
	for y, row := range topographicMap {
		for x, val := range row {
			if val != 0 {
				continue
			}
			trails := findTrails(Position{x, y}, topographicMap)
			score += len(trails)
		}
	}
	return score
}

func findTrailheadsRatings(topographicMap [][]int) int {
	rating := 0
	for y, row := range topographicMap {
		for x, val := range row {
			if val != 0 {
				continue
			}
			rating += findDistinctTrails(Position{x, y}, topographicMap)
		}
	}
	return rating
}

func findTrails(startPosition Position, topographicMap [][]int) []Position {
	var alreadyFoundTrails []Position
	alreadyVisitedPositions := make(map[Position]bool)
	findTrailsRecursively(startPosition, topographicMap, &alreadyFoundTrails, alreadyVisitedPositions)
	return alreadyFoundTrails
}

func findDistinctTrails(startPosition Position, topographicMap [][]int) int {
	timesVisited := make(map[Position]int)
	return findDistinctTrailsRecursively(startPosition, topographicMap, timesVisited)
}

func findTrailsRecursively(currentPosition Position, topographicMap [][]int, alreadyFoundTrails *[]Position, alreadyVisitedPositions map[Position]bool) {
	currentVal := topographicMap[currentPosition.y][currentPosition.x]
	if currentVal == 9 {
		if !alreadyFoundTrail(currentPosition, alreadyFoundTrails) {
			*alreadyFoundTrails = append(*alreadyFoundTrails, currentPosition)
		}
		return
	}
	var nextPositions []Position
	if currentPosition.x > 0 {
		nextPositions = append(nextPositions, Position{currentPosition.x - 1, currentPosition.y})
	}
	if currentPosition.x < len(topographicMap[0])-1 {
		nextPositions = append(nextPositions, Position{currentPosition.x + 1, currentPosition.y})
	}
	if currentPosition.y > 0 {
		nextPositions = append(nextPositions, Position{currentPosition.x, currentPosition.y - 1})
	}
	if currentPosition.y < len(topographicMap)-1 {
		nextPositions = append(nextPositions, Position{currentPosition.x, currentPosition.y + 1})
	}

	for _, nextPosition := range nextPositions {
		nextVal := topographicMap[nextPosition.y][nextPosition.x]
		if nextVal != currentVal+1 {
			continue
		}
		if alreadyVisited(nextPosition, alreadyVisitedPositions) {
			continue
		}
		alreadyVisitedPositions[nextPosition] = true
		findTrailsRecursively(nextPosition, topographicMap, alreadyFoundTrails, alreadyVisitedPositions)
	}
}

func alreadyFoundTrail(position Position, alreadyFoundTrails *[]Position) bool {
	for _, pos := range *alreadyFoundTrails {
		if position == pos {
			return true
		}
	}
	return false
}

func findDistinctTrailsRecursively(currentPosition Position, topographicMap [][]int, timesVisited map[Position]int) int {
	currentVal := topographicMap[currentPosition.y][currentPosition.x]
	if currentVal == 9 {
		return 1
	}
	var nextPositions []Position
	if currentPosition.x > 0 {
		nextPositions = append(nextPositions, Position{currentPosition.x - 1, currentPosition.y})
	}
	if currentPosition.x < len(topographicMap[0])-1 {
		nextPositions = append(nextPositions, Position{currentPosition.x + 1, currentPosition.y})
	}
	if currentPosition.y > 0 {
		nextPositions = append(nextPositions, Position{currentPosition.x, currentPosition.y - 1})
	}
	if currentPosition.y < len(topographicMap)-1 {
		nextPositions = append(nextPositions, Position{currentPosition.x, currentPosition.y + 1})
	}
	nbPreviousVisits := 0
	for _, nextPosition := range nextPositions {
		nextVal := topographicMap[nextPosition.y][nextPosition.x]
		if nextVal != currentVal+1 {
			continue
		}
		rating := findDistinctTrailsRecursively(nextPosition, topographicMap, timesVisited)
		nbPreviousVisits += rating
	}
	timesVisited[currentPosition] = nbPreviousVisits
	return nbPreviousVisits
}

func alreadyVisited(position Position, alreadyVisitedPositions map[Position]bool) bool {
	return alreadyVisitedPositions[position]
}

func parseInput(input string) [][]int {
	rows := strings.Split(input, "\n")
	topographicMap := make([][]int, len(rows))
	for rowIdx, row := range rows {
		topographicRow := make([]int, len(row))
		for idx, val := range row {
			topographicRow[idx], _ = strconv.Atoi((string)(val))
		}
		topographicMap[rowIdx] = topographicRow
	}
	return topographicMap
}
