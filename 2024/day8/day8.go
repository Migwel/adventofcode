package day8

import (
	"strings"
)

type Position struct {
	x int
	y int
}

func computeAntinodes(cityStr string) int {
	city := parseInput(cityStr)
	antennas := findAntennas(city)
	antinodes := findAntinodes(antennas, len(city)-1, len(city[0])-1)
	return len(antinodes)
}

func computeResonantAntinodes(cityStr string) int {
	city := parseInput(cityStr)
	antennas := findAntennas(city)
	antinodes := findResonantAntinodes(antennas, len(city)-1, len(city[0])-1)
	return len(antinodes)
}

func findAntinodes(antennas map[rune][]Position, maxY, maxX int) map[Position]bool {
	antinodes := make(map[Position]bool)
	for _, positions := range antennas {
		for i := 0; i < len(positions); i++ {
			pos1 := positions[i]
			for j := 0; j < len(positions); j++ {
				if i == j {
					continue
				}
				pos2 := positions[j]
				x := computeCoordinate(pos1.x, pos2.x, maxX, 1)
				y := computeCoordinate(pos1.y, pos2.y, maxY, 1)
				if x == -1 || y == -1 {
					continue
				}
				pos := Position{x, y}
				antinodes[pos] = true
			}
		}
	}
	return antinodes
}

func findResonantAntinodes(antennas map[rune][]Position, maxY, maxX int) map[Position]bool {
	antinodes := make(map[Position]bool)
	for _, positions := range antennas {
		for i := 0; i < len(positions); i++ {
			pos1 := positions[i]
			for j := 0; j < len(positions); j++ {
				if i == j {
					continue
				}
				pos2 := positions[j]
				for multiplier := 0; ; multiplier++ {
					x := computeCoordinate(pos1.x, pos2.x, maxX, multiplier)
					y := computeCoordinate(pos1.y, pos2.y, maxY, multiplier)
					if x == -1 || y == -1 {
						break
					}
					pos := Position{x, y}
					antinodes[pos] = true
				}
			}
		}
	}
	return antinodes
}

func computeCoordinate(val1, val2, maxVal, multiplier int) int {
	var val int
	if val1 < val2 {
		val = val1 - multiplier*(val2-val1)
	} else {
		val = val1 + multiplier*(val1-val2)
	}

	if val < 0 {
		return -1
	}
	if val > maxVal {
		return -1
	}
	return val
}

func findAntennas(city [][]rune) map[rune][]Position {
	antennas := make(map[rune][]Position)
	for y := 0; y < len(city); y++ {
		cityRow := city[y]
		for x := 0; x < len(cityRow); x++ {
			val := cityRow[x]
			if cityRow[x] != '.' {
				existingAntennas := antennas[val]
				if len(existingAntennas) == 0 {
					antennas[val] = []Position{Position{x, y}}
				} else {
					antennas[val] = append(existingAntennas, Position{x, y})
				}
			}
		}
	}
	return antennas
}

func parseInput(cityStr string) [][]rune {
	rows := strings.Split(cityStr, "\n")
	city := make([][]rune, len(rows))
	for idx, row := range rows {
		cityRow := make([]rune, len(row))
		for rowIdx, runeVal := range row {
			cityRow[rowIdx] = runeVal
		}
		city[idx] = cityRow
	}
	return city
}
