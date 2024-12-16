package day14

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
)

type Robot struct {
	x         int
	y         int
	xVelocity int
	yVelocity int
}

type Tile struct {
	robots []Robot
}

type Area struct {
	tiles  [][]Tile
	height int
	width  int
}

func computeSafetyFactor(input string, height, width, nbIterations int) int {
	area := initBathroom(input, height, width)
	for i := 0; i < nbIterations; i++ {
		area = nextSecond(area)
	}
	return computeScore(area.tiles)
}

func computeTimeToFindChristmasTree(input string, height, width int) int {
	area := initBathroom(input, height, width)
	for i := 0; i < 100000; i++ {
		area = nextSecond(area)
		if isChristmasTree(area) {
			return i + 1
		}
	}
	return -1
}

func isChristmasTree(area Area) bool {
	for _, row := range area.tiles {
		for _, tile := range row {
			if len(tile.robots) > 1 {
				return false
			}
		}
	}
	displayArea(area)
	return true
}

func displayArea(area Area) {
	for _, row := range area.tiles {
		for _, tile := range row {
			fmt.Print(len(tile.robots))
		}
		fmt.Println()
	}
}

func computeScore(tiles [][]Tile) int {
	halfHeight := (int)(math.Floor((float64)(len(tiles)))) / 2
	halfWidth := (int)(math.Floor((float64)(len(tiles[0])))) / 2
	nbTopLeftQuadrant := 0
	nbTopRightQuadrant := 0
	nbBottomLeftQuadrant := 0
	nbBottomRightQuadrant := 0
	for y, row := range tiles {
		for x, tile := range row {
			nbRobots := len(tile.robots)
			if nbRobots == 0 {
				continue
			}
			if y < halfHeight {
				if x < halfWidth {
					nbTopLeftQuadrant += nbRobots
				}
				if x > halfWidth {
					nbTopRightQuadrant += nbRobots
				}
			}
			if y > halfHeight {
				if x < halfWidth {
					nbBottomLeftQuadrant += nbRobots
				}
				if x > halfWidth {
					nbBottomRightQuadrant += nbRobots
				}
			}
		}
	}
	return nbTopLeftQuadrant * nbTopRightQuadrant * nbBottomLeftQuadrant * nbBottomRightQuadrant
}

func nextSecond(area Area) Area {
	tiles := initTiles(area.height, area.width)
	for _, row := range area.tiles {
		for _, tile := range row {
			for _, robot := range tile.robots {
				newRobot := nextRobot(robot, area.height, area.width)
				tiles[newRobot.y][newRobot.x].robots = append(tiles[newRobot.y][newRobot.x].robots, newRobot)
			}
		}
	}
	return Area{tiles, area.height, area.width}
}

func nextRobot(robot Robot, height, width int) Robot {
	x := absoluteModulo(robot.x+robot.xVelocity, width)
	y := absoluteModulo(robot.y+robot.yVelocity, height)
	return Robot{x, y, robot.xVelocity, robot.yVelocity}
}

func absoluteModulo(val, modulo int) int {
	return ((val % modulo) + modulo) % modulo
}

func initBathroom(input string, height, width int) Area {
	tiles := initTiles(height, width)
	robots := parseInput(input)
	for _, robot := range robots {
		tiles[robot.y][robot.x].robots = append(tiles[robot.y][robot.x].robots, robot)
	}
	return Area{tiles, height, width}

}

func initTiles(height, width int) [][]Tile {
	tiles := make([][]Tile, height)
	for i := 0; i < height; i++ {
		tilesRow := make([]Tile, width)
		tiles[i] = tilesRow
	}
	return tiles
}

func parseInput(input string) []Robot {
	re := regexp.MustCompile(`p=([0-9]+),([0-9]+) v=(-?[0-9]+),(-?[0-9]+)`)
	matches := re.FindAllSubmatch([]byte(input), -1)
	var robots []Robot
	for _, match := range matches {
		x, _ := strconv.Atoi((string)(match[1]))
		y, _ := strconv.Atoi((string)(match[2]))
		xVelocity, _ := strconv.Atoi((string)(match[3]))
		yVelocity, _ := strconv.Atoi((string)(match[4]))
		robot := Robot{x, y, xVelocity, yVelocity}
		robots = append(robots, robot)
	}
	return robots
}
