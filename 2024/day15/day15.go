package day15

import (
	"fmt"
	"sort"
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

type Position struct {
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

var isSane bool

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
	isSane = true
	for idx, movement := range movements {
		applyMovementSecondWarehouse(warehouse, robot, movement)
		if isSane {
			sanityCheck(warehouse, idx)
		}
		// if idx >= 2939 && idx <= 2941 {
		// 	displayWarehouse(*warehouse, movement)
		// }
	}
}

func sanityCheck(warehouse *Warehouse, idx int) {
	countRobots := 0
	for y, row := range warehouse.cells {
		for x, cell := range row {
			if cell == LEFT_BOX {
				if warehouse.cells[y][x+1] != RIGHT_BOX {
					fmt.Printf("Left box with right box at x=%d, y=%d, idx=%d", x, y, idx)
					isSane = false
				}
			}
			if cell == RIGHT_BOX {
				if warehouse.cells[y][x-1] != LEFT_BOX {
					fmt.Printf("Right box with left box at x=%d, y=%d, idx=%d", x, y, idx)
					isSane = false
				}
			}
			if cell == ROBOT {
				countRobots += 1
			}
		}
	}
	if countRobots > 1 {
		fmt.Printf("Too many robots, idx=%d", idx)
		isSane = false
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
	for y := robot.y; y >= 0; y-- {
		cell := warehouse.cells[y][robot.x]
		if cell == WALL {
			return
		}
		if cell == RIGHT_BOX {
			adjancentCell := AdjacentCell{LEFT_BOX, robot.x - 1, y}
			if !canMoveAdjacentCellUp(warehouse, adjancentCell, &adjacentCells) {
				return
			}
			adjacentCells = append(adjacentCells, adjancentCell)
		}
		if cell == LEFT_BOX {
			adjancentCell := AdjacentCell{RIGHT_BOX, robot.x + 1, y}
			if !canMoveAdjacentCellUp(warehouse, adjancentCell, &adjacentCells) {
				return
			}
			adjacentCells = append(adjacentCells, adjancentCell)
		}
		if cell == EMPTY {
			moveUpSecondWarehouse(warehouse, robot.x, robot.y, y, adjacentCells)
			robot.y = robot.y - 1
			return
		}
	}
}

func moveUpSecondWarehouse(warehouse *Warehouse, x, fromY, toY int, adjacentCells []AdjacentCell) {
	alreadyMoved := make(map[Position]bool)
	for y := toY; y < fromY; y++ {
		warehouse.cells[y][x] = warehouse.cells[y+1][x]
		alreadyMoved[Position{x, y}] = true
	}
	warehouse.cells[fromY][x] = EMPTY
	alreadyMoved[Position{x, fromY}] = true
	moveUpAdjacentCells(warehouse, adjacentCells, alreadyMoved)
}

func moveUpAdjacentCells(warehouse *Warehouse, adjacentCells []AdjacentCell, alreadyMoved map[Position]bool) {
	cellsPerColumn := make(map[int][]int)
	for _, adjacentCell := range adjacentCells {
		currentCells := cellsPerColumn[adjacentCell.x]
		cellsPerColumn[adjacentCell.x] = append(currentCells, adjacentCell.y)
	}
	for x, cells := range cellsPerColumn {
		sort.Sort(sort.Reverse(sort.IntSlice(cells)))
		for _, cell := range cells {
			if alreadyMoved[Position{x, cell}] {
				continue
			}
			currentCell := warehouse.cells[cell][x]
			var previousCell CellType
			for y := cell; y >= 0; y-- {
				alreadyMoved[Position{x, y}] = true
				previousCell = currentCell
				currentCell = warehouse.cells[y][x]
				warehouse.cells[y][x] = previousCell
				if currentCell == EMPTY {
					break
				}
			}
			warehouse.cells[cell][x] = EMPTY
		}
	}
}

func canMoveAdjacentCellUp(warehouse *Warehouse, currentCell AdjacentCell, adjacentCells *[]AdjacentCell) bool {
	if currentCell.y == 0 {
		return false
	}
	for y := currentCell.y - 1; y >= 0; y-- {
		nextCell := warehouse.cells[y][currentCell.x]
		if nextCell == WALL {
			return false
		}
		if nextCell == EMPTY {
			return true
		}
		if nextCell == RIGHT_BOX {
			adjancentCell := AdjacentCell{LEFT_BOX, currentCell.x - 1, y}
			if !canMoveAdjacentCellUp(warehouse, adjancentCell, adjacentCells) {
				return false
			}
			*adjacentCells = append(*adjacentCells, adjancentCell)
		}
		if nextCell == LEFT_BOX {
			adjancentCell := AdjacentCell{RIGHT_BOX, currentCell.x + 1, y}
			if !canMoveAdjacentCellUp(warehouse, adjancentCell, adjacentCells) {
				return false
			}
			*adjacentCells = append(*adjacentCells, adjancentCell)
		}
	}
	return true
}

func applyDownMovementSecondWarehouse(warehouse *Warehouse, robot *Robot) {
	var adjacentCells []AdjacentCell
	for y := robot.y; y < len(warehouse.cells); y++ {
		cell := warehouse.cells[y][robot.x]
		if cell == WALL {
			return
		}
		if cell == RIGHT_BOX {
			adjancentCell := AdjacentCell{LEFT_BOX, robot.x - 1, y}
			if !canMoveAdjacentCellDown(warehouse, adjancentCell, &adjacentCells) {
				return
			}
			adjacentCells = append(adjacentCells, adjancentCell)
		}
		if cell == LEFT_BOX {
			adjancentCell := AdjacentCell{RIGHT_BOX, robot.x + 1, y}
			if !canMoveAdjacentCellDown(warehouse, adjancentCell, &adjacentCells) {
				return
			}
			adjacentCells = append(adjacentCells, adjancentCell)
		}
		if cell == EMPTY {
			moveDownSecondWarehouse(warehouse, robot.x, robot.y, y, adjacentCells)
			robot.y = robot.y + 1
			return
		}
	}
}

func moveDownSecondWarehouse(warehouse *Warehouse, x, fromY, toY int, adjacentCells []AdjacentCell) {
	alreadyMoved := make(map[Position]bool)
	for y := toY; y > fromY; y-- {
		warehouse.cells[y][x] = warehouse.cells[y-1][x]
		alreadyMoved[Position{x, y}] = true
	}
	warehouse.cells[fromY][x] = EMPTY
	alreadyMoved[Position{x, fromY}] = true
	moveDownAdjacentCells(warehouse, adjacentCells, alreadyMoved)
}

func moveDownAdjacentCells(warehouse *Warehouse, adjacentCells []AdjacentCell, alreadyMoved map[Position]bool) {
	cellsPerColumn := make(map[int][]int)
	for _, adjacentCell := range adjacentCells {
		currentCells := cellsPerColumn[adjacentCell.x]
		cellsPerColumn[adjacentCell.x] = append(currentCells, adjacentCell.y)
	}
	for x, cells := range cellsPerColumn {
		sort.Sort(sort.IntSlice(cells))
		for _, cell := range cells {
			if alreadyMoved[Position{x, cell}] {
				continue
			}
			currentCell := warehouse.cells[cell][x]
			var previousCell CellType
			for y := cell; y < len(warehouse.cells); y++ {
				alreadyMoved[Position{x, y}] = true
				previousCell = currentCell
				currentCell = warehouse.cells[y][x]
				warehouse.cells[y][x] = previousCell
				if currentCell == EMPTY {
					break
				}
			}
			warehouse.cells[cell][x] = EMPTY
		}
	}
}

func canMoveAdjacentCellDown(warehouse *Warehouse, currentCell AdjacentCell, adjacentCells *[]AdjacentCell) bool {
	if currentCell.y == len(warehouse.cells)-1 {
		return false
	}
	for y := currentCell.y + 1; y < len(warehouse.cells); y++ {
		nextCell := warehouse.cells[y][currentCell.x]
		if nextCell == WALL {
			return false
		}
		if nextCell == EMPTY {
			return true
		}
		if nextCell == RIGHT_BOX {
			adjancentCell := AdjacentCell{LEFT_BOX, currentCell.x - 1, y}
			if !canMoveAdjacentCellDown(warehouse, adjancentCell, adjacentCells) {
				return false
			}
			*adjacentCells = append(*adjacentCells, adjancentCell)
		}
		if nextCell == LEFT_BOX {
			adjancentCell := AdjacentCell{RIGHT_BOX, currentCell.x + 1, y}
			if !canMoveAdjacentCellDown(warehouse, adjancentCell, adjacentCells) {
				return false
			}
			*adjacentCells = append(*adjacentCells, adjancentCell)
		}
	}
	return true
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
