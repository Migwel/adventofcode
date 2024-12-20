package day20

import (
	"math"
	"strings"
)

type CellType int

const (
	EMPTY  CellType = iota
	WALL   CellType = iota
	START  CellType = iota
	FINISH CellType = iota
)

type Position struct {
	x, y int
}

type Cell struct {
	position Position
	cellType CellType
	cost     int
}

type Maze struct {
	height, width int
	start, finish Position
	cells         [][]Cell
}

func computeNumberUsefulCheats(input string, timeSaved int) int {
	initialMaze := parseInput(input)
	var timesToFinish []int
	timesToFinish = append(timesToFinish, computeTimeToFinish(&initialMaze))
	for y := 1; y < len(initialMaze.cells)-1; y++ {
		row := initialMaze.cells[y]
		for x := 1; x < len(row)-1; x++ {
			cell := row[x]
			if cell.cellType != WALL {
				continue
			}
			maze := copyMaze(initialMaze)
			maze.cells[y][x].cellType = EMPTY
			timeToFinish := computeTimeToFinish(&maze)
			timesToFinish = append(timesToFinish, timeToFinish)
		}
	}
	nbUsefulCheats := countUsefulCheats(timesToFinish, timeSaved)
	return nbUsefulCheats
}

func computeNumberUsefulBiggerCheats(input string, timeSaved int) int {
	initialMaze := parseInput(input)
	computeTimeToFinish(&initialMaze)
	cheats := findBiggerCheats(initialMaze)
	count := 0
	for _, cheat := range cheats {
		if cheat >= timeSaved {
			count += 1
		}
	}
	return count
}

func findBiggerCheats(maze Maze) []int {
	var timesSaved []int
	for fromY, fromRow := range maze.cells {
		for fromX, fromCell := range fromRow {
			for toY, toRow := range maze.cells {
				for toX, toCell := range toRow {
					if toCell.cellType == WALL || fromCell.cellType == WALL {
						continue
					}
					if fromX == toX && fromY == toY {
						continue
					}
					if fromCell.cost > toCell.cost {
						continue
					}
					distance := computeDistance(Position{fromX, fromY}, Position{toX, toY})
					if distance > 20 {
						continue
					}
					if toCell.cost-fromCell.cost <= distance {
						continue
					}
					timeSaved := toCell.cost - fromCell.cost - distance
					timesSaved = append(timesSaved, timeSaved)
				}
			}
		}
	}
	return timesSaved
}

func computeDistance(position, otherPosition Position) int {
	return int(math.Abs(float64(position.x-otherPosition.x))) + int(math.Abs(float64(position.y-otherPosition.y)))
}

func countUsefulCheats(timesToFinish []int, timeSaved int) int {
	timeWithoutCheat := timesToFinish[0]
	count := 0
	for i := 1; i < len(timesToFinish); i++ {
		if timeWithoutCheat-timesToFinish[i] >= timeSaved {
			count += 1
		}
	}
	return count
}

func computeTimeToFinish(maze *Maze) int {
	computeCostsRecursively(maze, maze.start, 0)
	return maze.cells[maze.finish.y][maze.finish.x].cost
}

func computeCostsRecursively(maze *Maze, currentPosition Position, currentCost int) {
	currentCell := maze.cells[currentPosition.y][currentPosition.x]
	maze.cells[currentPosition.y][currentPosition.x].cost = currentCost
	if currentCell.cellType == FINISH {
		return
	}
	nextCost := currentCost + 1
	if currentPosition.y > 0 {
		nextCell := maze.cells[currentPosition.y-1][currentPosition.x]
		if nextCell.cellType != WALL && nextCell.cost > nextCost {
			computeCostsRecursively(maze, nextCell.position, nextCost)
		}
	}
	if currentPosition.y < maze.height-1 {
		nextCell := maze.cells[currentPosition.y+1][currentPosition.x]
		if nextCell.cellType != WALL && nextCell.cost > nextCost {
			computeCostsRecursively(maze, nextCell.position, nextCost)
		}
	}
	if currentPosition.x > 0 {
		nextCell := maze.cells[currentPosition.y][currentPosition.x-1]
		if nextCell.cellType != WALL && nextCell.cost > nextCost {
			computeCostsRecursively(maze, nextCell.position, nextCost)
		}
	}
	if currentPosition.x < maze.width-1 {
		nextCell := maze.cells[currentPosition.y][currentPosition.x+1]
		if nextCell.cellType != WALL && nextCell.cost > nextCost {
			computeCostsRecursively(maze, nextCell.position, nextCost)
		}
	}
}

func copyMaze(initialMaze Maze) Maze {
	cells := make([][]Cell, len(initialMaze.cells))
	for y, row := range initialMaze.cells {
		rowCells := make([]Cell, len(row))
		for x, cell := range row {
			rowCells[x] = Cell{Position{x, y}, cell.cellType, 999999999}
		}
		cells[y] = rowCells
	}
	return Maze{initialMaze.height, initialMaze.width, initialMaze.start, initialMaze.finish, cells}
}

func parseInput(input string) Maze {
	rows := strings.Split(input, "\n")
	cells := make([][]Cell, len(rows))
	var startPosition, finishPosition Position
	for y, row := range rows {
		rowCells := make([]Cell, len(row))
		for x, val := range row {
			var cellType CellType
			switch val {
			case '.':
				cellType = EMPTY
			case '#':
				cellType = WALL
			case 'S':
				cellType = START
				startPosition = Position{x, y}
			case 'E':
				cellType = FINISH
				finishPosition = Position{x, y}
			}
			rowCells[x] = Cell{Position{x, y}, cellType, 999999999}
		}
		cells[y] = rowCells
	}
	return Maze{len(cells), len(cells[0]), startPosition, finishPosition, cells}
}
