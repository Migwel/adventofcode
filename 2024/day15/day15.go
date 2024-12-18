package day15

import (
	"fmt"
	"strings"
)

type CellType int

const (
	WALL      CellType = iota
	ROBOT     CellType = iota
	BOX       CellType = iota
	EMPTY     CellType = iota
	LEFT_BOX  CellType = iota
	RIGHT_BOX CellType = iota
)

type AdjacentCell struct {
	cellType CellType
	x        int
	y        int
}

type Movement int

const (
	UP    Movement = iota
	RIGHT Movement = iota
	DOWN  Movement = iota
	LEFT  Movement = iota
)

type Warehouse struct {
	cells [][]CellType
}

type Robot struct {
	x, y int
}

func computeGPSCoordinatesSum(input string) int {
	warehouse, robot, movements := parseInput(input)
	applyMovements(&warehouse, &robot, movements)
	return computeSum(warehouse)
}

func computeSecondWarehouseGPSCoordinatesSum(input string) int {
	warehouse, robot, movements := parseInputSecondWarehouse(input)
	applyMovementsSecondWarehouse(&warehouse, &robot, movements)
	return computeSecondWarehouseSum(warehouse)
}

func computeSum(warehouse Warehouse) int {
	sum := 0
	for y, row := range warehouse.cells {
		for x, cell := range row {
			if cell != BOX {
				continue
			}
			sum += 100*y + x
		}
	}
	return sum
}

func computeSecondWarehouseSum(warehouse Warehouse) int {
	sum := 0
	for y, row := range warehouse.cells {
		for x, cell := range row {
			if cell != LEFT_BOX {
				continue
			}
			sum += 100*y + x
		}
	}
	return sum
}

func applyMovements(warehouse *Warehouse, robot *Robot, movements []Movement) {
	for _, movement := range movements {
		applyMovement(warehouse, robot, movement)
	}
}

func applyMovement(warehouse *Warehouse, robot *Robot, movement Movement) {
	switch movement {
	case UP:
		applyUpMovement(warehouse, robot)
	case RIGHT:
		applyRightMovement(warehouse, robot)
	case DOWN:
		applyDownMovement(warehouse, robot)
	case LEFT:
		applyLeftMovement(warehouse, robot)
	}
}

func applyMovementsSecondWarehouse(warehouse *Warehouse, robot *Robot, movements []Movement) {
	for _, movement := range movements {
		applyMovementSecondWarehouse(warehouse, robot, movement)
	}
}

func displayWarehouse(warehouse Warehouse, movement Movement) {
	switch movement {
	case UP:
		fmt.Println("UP")
	case RIGHT:
		fmt.Println("RIGHT")
	case DOWN:
		fmt.Println("DOWN")
	case LEFT:
		fmt.Println("LEFT")
	}
	for _, row := range warehouse.cells {
		for _, cell := range row {
			if cell == WALL {
				fmt.Print("#")
			}
			if cell == EMPTY {
				fmt.Print(".")
			}
			if cell == LEFT_BOX {
				fmt.Print("[")
			}
			if cell == RIGHT_BOX {
				fmt.Print("]")
			}
			if cell == ROBOT {
				fmt.Print("@")
			}
		}
		fmt.Println()
	}
}

func applyMovementSecondWarehouse(warehouse *Warehouse, robot *Robot, movement Movement) {
	switch movement {
	case UP:
		applyUpMovementSecondWarehouse(warehouse, robot)
	case RIGHT:
		applyRightMovement(warehouse, robot)
	case DOWN:
		applyDownMovementSecondWarehouse(warehouse, robot)
	case LEFT:
		applyLeftMovement(warehouse, robot)
	}
}

func applyUpMovement(warehouse *Warehouse, robot *Robot) {
	for y := robot.y; y >= 0; y-- {
		cell := warehouse.cells[y][robot.x]
		if cell == WALL {
			return
		}
		if cell == EMPTY {
			moveUp(warehouse, robot.x, robot.y, y)
			robot.y = robot.y - 1
			return
		}
	}
}

func moveUp(warehouse *Warehouse, x, fromY, toY int) {
	for y := toY; y < fromY; y++ {
		warehouse.cells[y][x] = warehouse.cells[y+1][x]
	}
	warehouse.cells[fromY][x] = EMPTY
}

func applyUpMovementSecondWarehouse(warehouse *Warehouse, robot *Robot) {
	var adjacentCells []AdjacentCell
	coveredColumns := make(map[int]bool)
	coveredColumns[robot.x] = true
	for y := robot.y; y >= 0; y-- {
		cell := warehouse.cells[y][robot.x]
		if cell == WALL {
			return
		}
		if cell == RIGHT_BOX {
			if coveredColumns[robot.x-1] {
				continue
			}
			if !canMoveAdjacentCellUp(warehouse, y, &adjacentCells, AdjacentCell{LEFT_BOX, robot.x - 1, y}, coveredColumns) {
				return
			}
			adjacentCells = append(adjacentCells, AdjacentCell{LEFT_BOX, robot.x - 1, y})
			coveredColumns[robot.x-1] = true
			continue
		}
		if cell == LEFT_BOX {
			if coveredColumns[robot.x+1] {
				continue
			}
			if !canMoveAdjacentCellUp(warehouse, y, &adjacentCells, AdjacentCell{RIGHT_BOX, robot.x + 1, y}, coveredColumns) {
				return
			}
			adjacentCells = append(adjacentCells, AdjacentCell{RIGHT_BOX, robot.x + 1, y})
			coveredColumns[robot.x+1] = true
			continue
		}
		if cell == EMPTY {
			moveUpSecondWarehouse(warehouse, robot.x, robot.y, y, adjacentCells)
			robot.y = robot.y - 1
			return
		}
	}
}

func canMoveAdjacentCellUp(warehouse *Warehouse, currentY int, adjacentCells *[]AdjacentCell, adjacentCell AdjacentCell, coveredColumns map[int]bool) bool {
	coveredColumns[adjacentCell.x] = true
	for y := currentY; y >= 0; y-- {
		nextCell := warehouse.cells[y][adjacentCell.x]
		if nextCell == WALL {
			return false
		}
		if nextCell == RIGHT_BOX {
			if coveredColumns[adjacentCell.x-1] {
				continue
			}
			if !canMoveAdjacentCellUp(warehouse, y, adjacentCells, AdjacentCell{LEFT_BOX, adjacentCell.x - 1, y}, coveredColumns) {
				return false
			}
			*adjacentCells = append(*adjacentCells, AdjacentCell{LEFT_BOX, adjacentCell.x - 1, y})
			continue
		}
		if nextCell == LEFT_BOX {
			if coveredColumns[adjacentCell.x+1] {
				continue
			}
			if !canMoveAdjacentCellUp(warehouse, y, adjacentCells, AdjacentCell{RIGHT_BOX, adjacentCell.x + 1, y}, coveredColumns) {
				return false
			}
			*adjacentCells = append(*adjacentCells, AdjacentCell{RIGHT_BOX, adjacentCell.x + 1, y})
			continue
		}
		if nextCell == EMPTY {
			return true
		}
	}
	return false
}

func canMoveAdjacentCellDown(warehouse *Warehouse, currentY int, adjacentCells *[]AdjacentCell, adjacentCell AdjacentCell, coveredColumns map[int]bool) bool {
	coveredColumns[adjacentCell.x] = true
	for y := currentY; y < len(warehouse.cells); y++ {
		nextCell := warehouse.cells[y][adjacentCell.x]
		if nextCell == WALL {
			return false
		}
		if nextCell == RIGHT_BOX {
			if coveredColumns[adjacentCell.x-1] {
				continue
			}
			if !canMoveAdjacentCellDown(warehouse, y, adjacentCells, AdjacentCell{LEFT_BOX, adjacentCell.x - 1, y}, coveredColumns) {
				return false
			}
			*adjacentCells = append(*adjacentCells, AdjacentCell{LEFT_BOX, adjacentCell.x - 1, y})
			continue
		}
		if nextCell == LEFT_BOX {
			if coveredColumns[adjacentCell.x+1] {
				continue
			}
			if !canMoveAdjacentCellDown(warehouse, y, adjacentCells, AdjacentCell{RIGHT_BOX, adjacentCell.x + 1, y}, coveredColumns) {
				return false
			}
			*adjacentCells = append(*adjacentCells, AdjacentCell{RIGHT_BOX, adjacentCell.x + 1, y})
			continue
		}
		if nextCell == EMPTY {
			return true
		}
	}
	return false
}

func canMoveAdjacentCells(warehouse Warehouse, y int, adjacentCells *[]AdjacentCell, coveredColumns map[int]bool) bool {
	for _, adjacentCell := range *adjacentCells {
		nextCell := warehouse.cells[y][adjacentCell.x]
		if nextCell == WALL {
			return false
		}
		if nextCell == RIGHT_BOX {
			if coveredColumns[adjacentCell.x-1] {
				continue
			}
			*adjacentCells = append(*adjacentCells, AdjacentCell{LEFT_BOX, adjacentCell.x - 1, y})
			continue
		}
		if nextCell == LEFT_BOX {
			if coveredColumns[adjacentCell.x+1] {
				continue
			}
			*adjacentCells = append(*adjacentCells, AdjacentCell{RIGHT_BOX, adjacentCell.x + 1, y})
			continue
		}
	}
	return true
}

func moveUpSecondWarehouse(warehouse *Warehouse, x, fromY, toY int, adjacentCells []AdjacentCell) {
	for y := toY; y < fromY; y++ {
		warehouse.cells[y][x] = warehouse.cells[y+1][x]
	}
	warehouse.cells[fromY][x] = EMPTY
	for _, adjacentCell := range adjacentCells {
		currentCell := warehouse.cells[adjacentCell.y][adjacentCell.x]
		var previousCell CellType
		for y := adjacentCell.y - 1; ; y-- {
			previousCell = currentCell
			currentCell = warehouse.cells[y][adjacentCell.x]
			warehouse.cells[y][adjacentCell.x] = previousCell
			if currentCell == EMPTY {
				break
			}
		}
		warehouse.cells[adjacentCell.y][adjacentCell.x] = EMPTY
	}
}

func applyRightMovement(warehouse *Warehouse, robot *Robot) {
	for x := robot.x; x < len(warehouse.cells[0]); x++ {
		cell := warehouse.cells[robot.y][x]
		if cell == WALL {
			return
		}
		if cell == EMPTY {
			moveRight(warehouse, robot.y, robot.x, x)
			robot.x = robot.x + 1
			return
		}
	}
}

func moveRight(warehouse *Warehouse, y, fromX, toX int) {
	for x := toX; x > fromX; x-- {
		warehouse.cells[y][x] = warehouse.cells[y][x-1]
	}
	warehouse.cells[y][fromX] = EMPTY
}

func applyDownMovement(warehouse *Warehouse, robot *Robot) {
	for y := robot.y; y < len(warehouse.cells); y++ {
		cell := warehouse.cells[y][robot.x]
		if cell == WALL {
			return
		}
		if cell == EMPTY {
			moveDown(warehouse, robot.x, robot.y, y)
			robot.y = robot.y + 1
			return
		}
	}
}

func moveDown(warehouse *Warehouse, x, fromY, toY int) {
	for y := toY; y > fromY; y-- {
		warehouse.cells[y][x] = warehouse.cells[y-1][x]
	}
	warehouse.cells[fromY][x] = EMPTY
}

func applyDownMovementSecondWarehouse(warehouse *Warehouse, robot *Robot) {
	var adjacentCells []AdjacentCell
	coveredColumns := make(map[int]bool)
	coveredColumns[robot.x] = true
	for y := robot.y; y < len(warehouse.cells); y++ {
		cell := warehouse.cells[y][robot.x]
		if cell == WALL {
			return
		}
		if cell == RIGHT_BOX {
			if coveredColumns[robot.x-1] {
				continue
			}
			if !canMoveAdjacentCellDown(warehouse, y, &adjacentCells, AdjacentCell{LEFT_BOX, robot.x - 1, y}, coveredColumns) {
				return
			}
			adjacentCells = append(adjacentCells, AdjacentCell{LEFT_BOX, robot.x - 1, y})
			coveredColumns[robot.x-1] = true
			continue
		}
		if cell == LEFT_BOX {
			if coveredColumns[robot.x+1] {
				continue
			}
			if !canMoveAdjacentCellDown(warehouse, y, &adjacentCells, AdjacentCell{RIGHT_BOX, robot.x + 1, y}, coveredColumns) {
				return
			}
			coveredColumns[robot.x+1] = true
			adjacentCells = append(adjacentCells, AdjacentCell{RIGHT_BOX, robot.x + 1, y})
			continue
		}
		if cell == EMPTY {
			moveDownSecondWarehouse(warehouse, robot.x, robot.y, y, adjacentCells)
			robot.y = robot.y + 1
			return
		}
	}
}

func moveDownSecondWarehouse(warehouse *Warehouse, x, fromY, toY int, adjacentCells []AdjacentCell) {
	for y := toY; y > fromY; y-- {
		warehouse.cells[y][x] = warehouse.cells[y-1][x]
	}
	warehouse.cells[fromY][x] = EMPTY
	for _, adjacentCell := range adjacentCells {
		currentCell := warehouse.cells[adjacentCell.y][adjacentCell.x]
		var previousCell CellType
		for y := adjacentCell.y + 1; ; y++ {
			previousCell = currentCell
			currentCell = warehouse.cells[y][adjacentCell.x]
			warehouse.cells[y][adjacentCell.x] = previousCell
			if currentCell == EMPTY {
				break
			}
		}
		warehouse.cells[adjacentCell.y][adjacentCell.x] = EMPTY
	}
}

func applyLeftMovement(warehouse *Warehouse, robot *Robot) {
	for x := robot.x; x >= 0; x-- {
		cell := warehouse.cells[robot.y][x]
		if cell == WALL {
			return
		}
		if cell == EMPTY {
			moveLeft(warehouse, robot.y, robot.x, x)
			robot.x = robot.x - 1
			return
		}
	}
}

func moveLeft(warehouse *Warehouse, y, fromX, toX int) {
	for x := toX; x < fromX; x++ {
		warehouse.cells[y][x] = warehouse.cells[y][x+1]
	}
	warehouse.cells[y][fromX] = EMPTY
}

func parseInput(input string) (Warehouse, Robot, []Movement) {
	rows := strings.Split(input, "\n")
	parsingWarehouse := true
	var cells [][]CellType
	var movements []Movement
	var robot Robot
	for y, row := range rows {
		if len(row) == 0 {
			parsingWarehouse = false
			continue
		}
		if parsingWarehouse {
			rowCells := parseWarehouse(row)
			for x, cell := range rowCells {
				if cell == ROBOT {
					robot = Robot{x, y}
				}
			}
			cells = append(cells, rowCells)
		} else {
			parseMovements(row, &movements)
		}
	}
	return Warehouse{cells}, robot, movements
}

func parseMovements(row string, movement *[]Movement) {
	for _, val := range row {
		if val == '^' {
			*movement = append(*movement, UP)
		} else if val == '>' {
			*movement = append(*movement, RIGHT)
		} else if val == '<' {
			*movement = append(*movement, LEFT)
		} else if val == 'v' {
			*movement = append(*movement, DOWN)
		}
	}
}

func parseWarehouse(row string) []CellType {
	rowCells := make([]CellType, len(row))
	for x, val := range row {
		if val == '#' {
			rowCells[x] = WALL
		} else if val == '.' {
			rowCells[x] = EMPTY
		} else if val == 'O' {
			rowCells[x] = BOX
		} else if val == '@' {
			rowCells[x] = ROBOT
		}
	}
	return rowCells
}

func parseInputSecondWarehouse(input string) (Warehouse, Robot, []Movement) {
	rows := strings.Split(input, "\n")
	parsingWarehouse := true
	var cells [][]CellType
	var movements []Movement
	var robot Robot
	for y, row := range rows {
		if len(row) == 0 {
			parsingWarehouse = false
			continue
		}
		if parsingWarehouse {
			rowCells := parseWarehouseSecondWarehouse(row)
			for x, cell := range rowCells {
				if cell == ROBOT {
					robot = Robot{x, y}
				}
			}
			cells = append(cells, rowCells)
		} else {
			parseMovements(row, &movements)
		}
	}
	return Warehouse{cells}, robot, movements
}

func parseWarehouseSecondWarehouse(row string) []CellType {
	rowCells := make([]CellType, len(row)*2)
	for x, val := range row {
		if val == '#' {
			rowCells[2*x] = WALL
			rowCells[2*x+1] = WALL
		} else if val == '.' {
			rowCells[2*x] = EMPTY
			rowCells[2*x+1] = EMPTY
		} else if val == 'O' {
			rowCells[2*x] = LEFT_BOX
			rowCells[2*x+1] = RIGHT_BOX
		} else if val == '@' {
			rowCells[2*x] = ROBOT
			rowCells[2*x+1] = EMPTY
		}
	}
	return rowCells
}
