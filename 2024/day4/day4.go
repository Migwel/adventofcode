package day4

import (
	"strings"
)

type DIRECTION int

const (
	NORTH     DIRECTION = 1
	NORTHEAST DIRECTION = 2
	EAST      DIRECTION = 3
	SOUTHEAST DIRECTION = 4
	SOUTH     DIRECTION = 5
	SOUTHWEST DIRECTION = 6
	WEST      DIRECTION = 7
	NORTHWEST DIRECTION = 8
)

var LETTERS = [3]rune{'M', 'A', 'S'}

func xmasCounter(data string) int {
	table := parseInput(data)
	count := 0
	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table[i]); j++ {
			if table[i][j] != 'X' {
				continue
			}
			if searchNextLetters(table, i, j, 0, NORTH) {
				count += 1
			}
			if searchNextLetters(table, i, j, 0, NORTHEAST) {
				count += 1
			}
			if searchNextLetters(table, i, j, 0, EAST) {
				count += 1
			}
			if searchNextLetters(table, i, j, 0, SOUTHEAST) {
				count += 1
			}
			if searchNextLetters(table, i, j, 0, SOUTH) {
				count += 1
			}
			if searchNextLetters(table, i, j, 0, SOUTHWEST) {
				count += 1
			}
			if searchNextLetters(table, i, j, 0, WEST) {
				count += 1
			}
			if searchNextLetters(table, i, j, 0, NORTHWEST) {
				count += 1
			}
		}
	}
	return count
}

func searchNextLetters(table [][]rune, i, j, idxLetter int, direction DIRECTION) bool {
	nextLetter, nextI, nextJ := findNextLetter(table, i, j, direction)
	if nextLetter == 0 || nextLetter != LETTERS[idxLetter] {
		return false
	}
	if idxLetter == len(LETTERS)-1 {
		return true
	}
	return searchNextLetters(table, nextI, nextJ, idxLetter+1, direction)
}

func findNextLetter(table [][]rune, i, j int, direction DIRECTION) (rune, int, int) {
	switch direction {
	case NORTH:
		if i <= 0 {
			return 0, 0, 0
		}
		return table[i-1][j], i - 1, j
	case NORTHEAST:
		if i <= 0 || j >= len(table[i])-1 {
			return 0, 0, 0
		}
		return table[i-1][j+1], i - 1, j + 1
	case EAST:
		if j >= len(table[i])-1 {
			return 0, 0, 0
		}
		return table[i][j+1], i, j + 1
	case SOUTHEAST:
		if i >= len(table)-1 || j >= len(table[i])-1 {
			return 0, 0, 0
		}
		return table[i+1][j+1], i + 1, j + 1
	case SOUTH:
		if i >= len(table)-1 {
			return 0, 0, 0
		}
		return table[i+1][j], i + 1, j
	case SOUTHWEST:
		if i >= len(table)-1 || j <= 0 {
			return 0, 0, 0
		}
		return table[i+1][j-1], i + 1, j - 1
	case WEST:
		if j <= 0 {
			return 0, 0, 0
		}
		return table[i][j-1], i, j - 1
	case NORTHWEST:
		if i <= 0 || j <= 0 {
			return 0, 0, 0
		}
		return table[i-1][j-1], i - 1, j - 1
	}

	return 0, 0, 0
}

func parseInput(data string) [][]rune {
	rows := strings.Split(data, "\n")
	table := make([][]rune, len(rows))
	for i := 0; i < len(rows); i++ {
		row := make([]rune, len(rows[i]))
		for pos, char := range rows[i] {
			row[pos] = char
		}
		table[i] = row
	}
	return table
}

func xmasCrossCounter(data string) int {
	table := parseInput(data)
	count := 0
	for i := 1; i < len(table)-1; i++ {
		for j := 1; j < len(table[i])-1; j++ {
			if table[i][j] != 'A' {
				continue
			}
			topLeft := table[i-1][j-1]
			topRight := table[i-1][j+1]
			bottomLeft := table[i+1][j-1]
			bottomRight := table[i+1][j+1]

			if ((topLeft == 'M' && bottomRight == 'S') || topLeft == 'S' && bottomRight == 'M') &&
				((topRight == 'M' && bottomLeft == 'S') || topRight == 'S' && bottomLeft == 'M') {
				count += 1
				continue
			}
		}
	}
	return count
}
