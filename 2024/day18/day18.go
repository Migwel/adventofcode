package day18

import (
	"strconv"
	"strings"
)

type Register int

type Instruction struct {
	opcode, operand int
}

type CellType int

const (
	EMPTY CellType = iota
	WALL  CellType = iota
)

type Position struct {
	x, y int
}

type Cell struct {
	position Position
	cellType CellType
	cost     int
}

type Grid struct {
	height, width int
	cells         [][]Cell
}

func computeMinimalNumberOfSteps(input string, height, width, nbCorruptedBytes int) int {
	corruptedBytes := parseInput(input)
	grid := buildGrid(height, width)
	applyCorruptedBytes(grid, corruptedBytes[:nbCorruptedBytes])
	updateCosts(grid)
	return grid.cells[height-1][width-1].cost
}

func findFirstCorruptingByte(input string, height, width int) Position {
	corruptedBytes := parseInput(input)
	for i := 1; i < len(corruptedBytes); i++ {
		grid := buildGrid(height, width)
		applyCorruptedBytes(grid, corruptedBytes[:i])
		updateCosts(grid)
		if grid.cells[height-1][width-1].cost == 999999999 {
			return corruptedBytes[i-1]
		}
	}
	return Position{1, 1}
}

func findShortestPath(grid Grid) []Cell {
	currentCell := grid.cells[grid.height-1][grid.width-1]
	shortestPath := []Cell{currentCell}
	for currentCell.position.x != 0 || currentCell.position.y != 0 {
		if currentCell.position.x > 0 {
			nextCell := grid.cells[currentCell.position.y][currentCell.position.x-1]
			if nextCell.cost == currentCell.cost-1 {
				shortestPath = append(shortestPath, nextCell)
				currentCell = nextCell
				continue
			}
		}
		if currentCell.position.x < grid.width-1 {
			nextCell := grid.cells[currentCell.position.y][currentCell.position.x+1]
			if nextCell.cost == currentCell.cost-1 {
				shortestPath = append(shortestPath, nextCell)
				currentCell = nextCell
				continue
			}
		}
		if currentCell.position.y > 0 {
			nextCell := grid.cells[currentCell.position.y-1][currentCell.position.x]
			if nextCell.cost == currentCell.cost-1 {
				shortestPath = append(shortestPath, nextCell)
				currentCell = nextCell
				continue
			}
		}
		if currentCell.position.y < grid.height-1 {
			nextCell := grid.cells[currentCell.position.y+1][currentCell.position.x]
			if nextCell.cost == currentCell.cost-1 {
				shortestPath = append(shortestPath, nextCell)
				currentCell = nextCell
				continue
			}
		}
	}
	return shortestPath
}

func updateCosts(grid Grid) {
	updateCostsRecursively(grid, Position{0, 0}, 0)
}

func updateCostsRecursively(grid Grid, currentPosition Position, cost int) {
	grid.cells[currentPosition.y][currentPosition.x].cost = cost

	if currentPosition.x > 0 {
		nextCell := grid.cells[currentPosition.y][currentPosition.x-1]
		if nextCell.cellType != WALL && nextCell.cost > cost+1 {
			updateCostsRecursively(grid, nextCell.position, cost+1)
		}
	}

	if currentPosition.x < grid.width-1 {
		nextCell := grid.cells[currentPosition.y][currentPosition.x+1]
		if nextCell.cellType != WALL && nextCell.cost > cost+1 {
			updateCostsRecursively(grid, nextCell.position, cost+1)
		}
	}
	if currentPosition.y > 0 {
		nextCell := grid.cells[currentPosition.y-1][currentPosition.x]
		if nextCell.cellType != WALL && nextCell.cost > cost+1 {
			updateCostsRecursively(grid, nextCell.position, cost+1)
		}
	}

	if currentPosition.y < grid.height-1 {
		nextCell := grid.cells[currentPosition.y+1][currentPosition.x]
		if nextCell.cellType != WALL && nextCell.cost > cost+1 {
			updateCostsRecursively(grid, nextCell.position, cost+1)
		}
	}
}

func applyCorruptedBytes(grid Grid, corruptedBytes []Position) {
	for _, position := range corruptedBytes {
		grid.cells[position.y][position.x].cellType = WALL
	}
}

func parseInput(input string) []Position {
	lines := strings.Split(input, "\n")
	corruptedBytes := make([]Position, len(lines))
	for idx, line := range lines {
		values := strings.Split(line, ",")
		x, _ := strconv.Atoi(values[0])
		y, _ := strconv.Atoi(values[1])
		corruptedBytes[idx] = Position{x, y}
	}
	return corruptedBytes
}

func buildGrid(height, width int) Grid {
	cells := make([][]Cell, height)
	for y := 0; y < height; y++ {
		row := make([]Cell, width)
		for x := 0; x < width; x++ {
			row[x] = Cell{Position{x, y}, EMPTY, 999999999}
		}
		cells[y] = row
	}
	return Grid{height, width, cells}
}
