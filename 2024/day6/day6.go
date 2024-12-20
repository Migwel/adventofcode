package day6

import (
	"strings"
)

type ORIENTATION int

const (
	NORTH ORIENTATION = 1
	EAST  ORIENTATION = 2
	SOUTH ORIENTATION = 3
	WEST  ORIENTATION = 4
)

func countDistrinctPosition(carteInput string) int {
	carte, orientation, startPosX, startPosY := parseMap(carteInput)
	browseCarte(carte, orientation, startPosX, startPosY)
	return countVisitedCells(carte)
}

func findLoopyPositions(carteInput string) int {
	origCarte, orientation, startPosX, startPosY := parseMap(carteInput)
	sum := 0
	for i := 0; i < len(origCarte); i++ {
		row := origCarte[i]
		for j := 0; j < len(row); j++ {
			if row[j] != '.' {
				continue
			}
			carte, _, _, _ := parseMap(carteInput)
			carte[i][j] = '#'
			if isLoop(carte, orientation, startPosX, startPosY, len(origCarte)*len(row)+5) {
				sum += 1
			}
		}
	}
	return sum
}

func isLoop(carte [][]rune, orientation ORIENTATION, startPosX, startPosY, remainingSteps int) bool {
	if remainingSteps < 0 {
		return true
	}
	switch orientation {
	case NORTH:
		{
			for y := startPosY; y >= 0; y-- {
				if carte[y][startPosX] == rune(NORTH) {
					return true
				}
				if carte[y][startPosX] == '#' {
					return isLoop(carte, EAST, startPosX, y+1, remainingSteps)
				}
				remainingSteps -= 1
			}
		}
	case EAST:
		{
			for x := startPosX; x < len(carte[0]); x++ {
				if carte[startPosY][x] == rune(EAST) {
					return true
				}
				if carte[startPosY][x] == '#' {
					return isLoop(carte, SOUTH, x-1, startPosY, remainingSteps)
				}
				remainingSteps -= 1
			}
		}
	case SOUTH:
		{
			for y := startPosY; y < len(carte); y++ {
				if carte[y][startPosX] == rune(SOUTH) {
					return true
				}
				if carte[y][startPosX] == '#' {
					return isLoop(carte, WEST, startPosX, y-1, remainingSteps)
				}
				remainingSteps -= 1
			}
		}
	case WEST:
		{
			for x := startPosX; x >= 0; x-- {
				if carte[startPosY][x] == '#' {
					return isLoop(carte, NORTH, x+1, startPosY, remainingSteps)
				}
				carte[startPosY][x] = 'X'
				remainingSteps -= 1
			}
		}
	}
	return false
}

func countVisitedCells(carte [][]rune) int {
	sum := 0
	for i := 0; i < len(carte); i++ {
		row := carte[i]
		for j := 0; j < len(row); j++ {
			if row[j] == 'X' {
				sum += 1
			}
		}
	}
	return sum
}

func browseCarte(carte [][]rune, orientation ORIENTATION, startPosX, startPosY int) {
	switch orientation {
	case NORTH:
		{
			for y := startPosY; y >= 0; y-- {
				if carte[y][startPosX] == '#' {
					browseCarte(carte, EAST, startPosX, y+1)
					return
				}
				carte[y][startPosX] = 'X'
			}
		}
	case EAST:
		{
			for x := startPosX; x < len(carte[0]); x++ {
				if carte[startPosY][x] == '#' {
					browseCarte(carte, SOUTH, x-1, startPosY)
					return
				}
				carte[startPosY][x] = 'X'
			}
		}
	case SOUTH:
		{
			for y := startPosY; y < len(carte); y++ {
				if carte[y][startPosX] == '#' {
					browseCarte(carte, WEST, startPosX, y-1)
					return
				}
				carte[y][startPosX] = 'X'
			}
		}
	case WEST:
		{
			for x := startPosX; x >= 0; x-- {
				if carte[startPosY][x] == '#' {
					browseCarte(carte, NORTH, x+1, startPosY)
					return
				}
				carte[startPosY][x] = 'X'
			}
		}
	}
}

func parseMap(carteInput string) ([][]rune, ORIENTATION, int, int) {
	rows := strings.Split(carteInput, "\n")
	carte := make([][]rune, len(rows))
	var orientation ORIENTATION
	var startPosX, startPosY int
	for i, row := range rows {
		carteRow := make([]rune, len(row))
		for j, val := range row {
			if val == '^' {
				startPosX, startPosY = j, i
				orientation = NORTH
			}
			if val == '>' {
				startPosX, startPosY = j, i
				orientation = EAST
			}
			if val == 'v' {
				startPosX, startPosY = j, i
				orientation = SOUTH
			}
			if val == '<' {
				startPosX, startPosY = j, i
				orientation = WEST
			}
			carteRow[j] = val
		}
		carte[i] = carteRow
	}
	return carte, orientation, startPosX, startPosY
}
