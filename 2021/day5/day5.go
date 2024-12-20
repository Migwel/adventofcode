package day5

import (
	"strconv"
	"strings"
)

type Position struct {
	x, y int
}

type Line struct {
	fromPosition, toPosition Position
}

type Cell struct {
	position           Position
	nbOverlappingLines int
}

type Grid struct {
	cells [][]Cell
}

func ComputePointsWithOverlap(input string, considerDiagonals bool) int {
	lines, maxX, maxY := parseInput(input)
	grid := makeGrid(maxX, maxY)
	applyLines(grid, lines, considerDiagonals)
	return countPointsWithOverlappingLines(grid)
}

func countPointsWithOverlappingLines(grid Grid) int {
	count := 0
	for _, rows := range grid.cells {
		for _, cell := range rows {
			if cell.nbOverlappingLines >= 2 {
				count += 1
			}
		}
	}
	return count
}

func applyLines(grid Grid, lines []Line, considerDiagonals bool) {
	for _, line := range lines {
		if line.fromPosition.x == line.toPosition.x {
			applyVerticalLine(grid, line, line.fromPosition.x)
			continue
		}
		if line.fromPosition.y == line.toPosition.y {
			applyHorizontalLine(grid, line, line.fromPosition.y)
			continue
		}
		if considerDiagonals {
			applyDiagonalLine(grid, line)
		}
	}
}

func applyDiagonalLine(grid Grid, line Line) {
	var higherPosition, lowerPosition Position
	if line.fromPosition.y < line.toPosition.y {
		higherPosition = line.fromPosition
		lowerPosition = line.toPosition
	} else {
		higherPosition = line.toPosition
		lowerPosition = line.fromPosition
	}
	if higherPosition.x < lowerPosition.x {
		applyDiagonalLeftToRight(grid, higherPosition, lowerPosition)
	} else {
		applyDiagonalRightToLeft(grid, higherPosition, lowerPosition)
	}
}

func applyDiagonalLeftToRight(grid Grid, higherPosition, lowerPosition Position) {
	for i := 0; i <= lowerPosition.y-higherPosition.y; i++ {
		grid.cells[higherPosition.y+i][higherPosition.x+i].nbOverlappingLines += 1
	}
}

func applyDiagonalRightToLeft(grid Grid, higherPosition, lowerPosition Position) {
	for i := 0; i <= lowerPosition.y-higherPosition.y; i++ {
		grid.cells[higherPosition.y+i][higherPosition.x-i].nbOverlappingLines += 1
	}
}

func applyVerticalLine(grid Grid, line Line, x int) {
	var lowerY, higherY int
	if line.fromPosition.y > line.toPosition.y {
		lowerY = line.toPosition.y
		higherY = line.fromPosition.y
	} else {
		lowerY = line.fromPosition.y
		higherY = line.toPosition.y
	}
	for y := lowerY; y <= higherY; y++ {
		grid.cells[y][x].nbOverlappingLines += 1
	}
}

func applyHorizontalLine(grid Grid, line Line, y int) {
	var lowerX, higherX int
	if line.fromPosition.x > line.toPosition.x {
		lowerX = line.toPosition.x
		higherX = line.fromPosition.x
	} else {
		lowerX = line.fromPosition.x
		higherX = line.toPosition.x
	}
	for x := lowerX; x <= higherX; x++ {
		grid.cells[y][x].nbOverlappingLines += 1
	}
}

func makeGrid(maxX, maxY int) Grid {
	cells := make([][]Cell, maxY+1)
	for y := 0; y <= maxY; y++ {
		rowCells := make([]Cell, maxX+1)
		for x := 0; x <= maxX; x++ {
			rowCells[x] = Cell{Position{x, y}, 0}
		}
		cells[y] = rowCells
	}
	return Grid{cells}
}

func parseInput(input string) ([]Line, int, int) {
	inputLines := strings.Split(input, "\n")
	lines := make([]Line, len(inputLines))
	maxX := 0
	maxY := 0
	for idx, inputLine := range inputLines {
		inputPositions := strings.Split(inputLine, " -> ")
		fromPositionSplit := strings.Split(inputPositions[0], ",")
		toPositionSplit := strings.Split(inputPositions[1], ",")
		fromX, _ := strconv.Atoi(fromPositionSplit[0])
		fromY, _ := strconv.Atoi(fromPositionSplit[1])
		toX, _ := strconv.Atoi(toPositionSplit[0])
		toY, _ := strconv.Atoi(toPositionSplit[1])
		line := Line{Position{fromX, fromY}, Position{toX, toY}}
		lines[idx] = line

		if fromX > maxX {
			maxX = fromX
		}
		if toX > maxX {
			maxX = toX
		}
		if fromY > maxY {
			maxY = fromY
		}
		if toY > maxY {
			maxY = toY
		}
	}
	return lines, maxX, maxY
}
