package day16

import (
	"strings"
)

type Orientation int

const (
	NORTH Orientation = iota
	EAST  Orientation = iota
	WEST  Orientation = iota
	SOUTH Orientation = iota
)

type CellType int

const (
	WALL     CellType = iota
	EMPTY    CellType = iota
	REINDEER CellType = iota
	FINISH   CellType = iota
)

type Cell struct {
	position Position
	costs    map[Orientation]int
	cellType CellType
}

type Position struct {
	x, y int
}

type Reindeer struct {
	x, y        int
	orientation Orientation
}

type Path struct {
	pathId int
	cells  []Cell
}

var nextPathId int

func computeLowestScore(input string) int {
	maze, reindeer := parseInput(input)
	updateCosts(&maze, reindeer)
	_, _, finishCell := findFinishCell(maze)
	lowestScore := 999999999
	for _, cost := range finishCell.costs {
		if cost < lowestScore {
			lowestScore = cost
		}
	}
	return lowestScore
}

func computeCellsOnBestPaths(input string) int {
	maze, reindeer := parseInput(input)
	updateCosts(&maze, reindeer)
	x, y, finishCell := findFinishCell(maze)
	bestPaths := backtrackPaths(maze, finishCell, x, y)
	visitedCells := make(map[Position]bool)
	for _, path := range bestPaths {
		for _, cell := range path.cells {
			visitedCells[cell.position] = true
		}
	}
	return len(visitedCells)
}

func backtrackPaths(maze [][]Cell, finishCell *Cell, x, y int) []Path {
	minimalCost := 999999999
	for _, cost := range finishCell.costs {
		if cost < minimalCost {
			minimalCost = cost
		}
	}
	nextPathId = 0
	var bestPaths []Path
	for orientation, cost := range finishCell.costs {
		if cost == minimalCost {
			bestPath := Path{nextPathId, []Cell{}}
			nextPathId += 1
			bestPath.cells = append(bestPath.cells, *finishCell)
			nextX, nextY := computeNextCell(x, y, orientation)
			orientationBestPaths := backTrack(maze, nextX, nextY, cost, orientation, &bestPath)
			bestPaths = append(bestPaths, orientationBestPaths...)
		}
	}
	return bestPaths
}

func backTrack(maze [][]Cell, x, y, nextCellCost int, comingFrom Orientation, bestPath *Path) []Path {
	pathAlreadyUsed := false
	bestPath.cells = append(bestPath.cells, maze[y][x])
	nextPathId := (*bestPath).pathId + 1
	var bestPaths []Path

	if maze[y][x].cellType == REINDEER {
		return []Path{*bestPath}
	}

	for orientation, cost := range maze[y][x].costs {
		if costMatch(nextCellCost, cost, orientation, comingFrom) {
			var currentBestPath Path
			if pathAlreadyUsed {
				currentBestPathCells := make([]Cell, len(bestPath.cells))
				copy(currentBestPathCells, bestPath.cells)
				currentBestPath = Path{nextPathId, currentBestPathCells}
				nextPathId += 1
			} else {
				currentBestPath = *bestPath
				pathAlreadyUsed = true
			}
			nextX, nextY := computeNextCell(x, y, orientation)
			paths := backTrack(maze, nextX, nextY, cost, orientation, &currentBestPath)
			for _, path := range paths {
				bestPaths = append(bestPaths, path)
			}
		}
	}
	return bestPaths
}

func costMatch(nextCellCost, cost int, orientation, comingFrom Orientation) bool {
	diffCost := computeDiffCost(orientation, comingFrom)
	return nextCellCost == cost+diffCost+1
}

func computeDiffCost(orientation, comingFrom Orientation) int {
	switch comingFrom {
	case NORTH:
		switch orientation {
		case SOUTH:
			return 2000
		case EAST, WEST:
			return 1000
		}
	case EAST:
		switch orientation {
		case WEST:
			return 2000
		case NORTH, SOUTH:
			return 1000
		}
	case SOUTH:
		switch orientation {
		case NORTH:
			return 2000
		case EAST, WEST:
			return 1000
		}
	case WEST:
		switch orientation {
		case EAST:
			return 2000
		case NORTH, SOUTH:
			return 1000
		}
	}
	return 0
}

func computeNextCell(currentX, currentY int, orientation Orientation) (int, int) {
	var x, y int
	switch orientation {
	case NORTH:
		x = currentX
		y = currentY - 1
	case WEST:
		x = currentX - 1
		y = currentY
	case SOUTH:
		x = currentX
		y = currentY + 1
	case EAST:
		x = currentX + 1
		y = currentY
	}
	return x, y

}

func findFinishCell(maze [][]Cell) (int, int, *Cell) {
	for y, rows := range maze {
		for x, cell := range rows {
			if cell.cellType == FINISH {
				return x, y, &cell
			}
		}
	}
	return 0, 0, nil
}

func updateCosts(maze *[][]Cell, reindeer Reindeer) {
	startingCell := (*maze)[reindeer.y][reindeer.x]
	startingCell.costs[NORTH] = 0
	startingCell.costs[EAST] = 0
	startingCell.costs[SOUTH] = 0
	startingCell.costs[WEST] = 0
	updateCostsRecursively(maze, reindeer.x, reindeer.y, reindeer.orientation, 0, WEST)
}

func updateCostsRecursively(maze *[][]Cell, currentX, currentY int, orientation Orientation, cost int, comingFrom Orientation) {
	(*maze)[currentY][currentX].costs[comingFrom] = cost
	if (*maze)[currentY][currentX].cellType == FINISH {
		return
	}
	if currentX > 0 && (*maze)[currentY][currentX-1].cellType != WALL {
		newOrientation := WEST
		costTurning := computeCostNewOrientation(orientation, newOrientation)
		costMoving := cost + costTurning + 1
		if costMoving < (*maze)[currentY][currentX-1].costs[EAST] {
			updateCostsRecursively(maze, currentX-1, currentY, newOrientation, costMoving, EAST)
		}
	}
	if currentX < len((*maze)[0])-1 && (*maze)[currentY][currentX+1].cellType != WALL {
		newOrientation := EAST
		costTurning := computeCostNewOrientation(orientation, newOrientation)
		costMoving := cost + costTurning + 1
		if costMoving < (*maze)[currentY][currentX+1].costs[WEST] {
			updateCostsRecursively(maze, currentX+1, currentY, newOrientation, costMoving, WEST)
		}
	}
	if currentY > 0 && (*maze)[currentY-1][currentX].cellType != WALL {
		newOrientation := NORTH
		costTurning := computeCostNewOrientation(orientation, newOrientation)
		costMoving := cost + costTurning + 1
		if costMoving < (*maze)[currentY-1][currentX].costs[SOUTH] {
			updateCostsRecursively(maze, currentX, currentY-1, newOrientation, costMoving, SOUTH)
		}
	}
	if currentX < len((*maze))-1 && (*maze)[currentY+1][currentX].cellType != WALL {
		newOrientation := SOUTH
		costTurning := computeCostNewOrientation(orientation, newOrientation)
		costMoving := cost + costTurning + 1
		if costMoving < (*maze)[currentY+1][currentX].costs[NORTH] {
			updateCostsRecursively(maze, currentX, currentY+1, newOrientation, costMoving, NORTH)
		}
	}
}

func computeCostNewOrientation(currentOrientation, newOrientation Orientation) int {
	switch currentOrientation {
	case NORTH:
		switch newOrientation {
		case SOUTH:
			return 2000
		case EAST, WEST:
			return 1000
		}
	case EAST:
		switch newOrientation {
		case WEST:
			return 2000
		case NORTH, SOUTH:
			return 1000
		}
	case SOUTH:
		switch newOrientation {
		case NORTH:
			return 2000
		case EAST, WEST:
			return 1000
		}
	case WEST:
		switch newOrientation {
		case EAST:
			return 2000
		case NORTH, SOUTH:
			return 1000
		}
	}
	return 0
}

func parseInput(input string) ([][]Cell, Reindeer) {
	rows := strings.Split(input, "\n")
	maze := make([][]Cell, len(rows))
	var reindeer Reindeer
	for y, row := range rows {
		rowCells := make([]Cell, len(row))
		for x, val := range row {
			var cellType CellType
			switch val {
			case '#':
				cellType = WALL
			case '.':
				cellType = EMPTY
			case 'S':
				cellType = REINDEER
				reindeer = Reindeer{x, y, EAST}
			case 'E':
				cellType = FINISH
			}
			rowCells[x] = Cell{Position{x, y}, buildHighestCosts(), cellType}
		}
		maze[y] = rowCells
	}
	return maze, reindeer
}

func buildHighestCosts() map[Orientation]int {
	costs := make(map[Orientation]int)
	highestCost := 999999999
	costs[NORTH] = highestCost
	costs[EAST] = highestCost
	costs[SOUTH] = highestCost
	costs[WEST] = highestCost
	return costs
}
