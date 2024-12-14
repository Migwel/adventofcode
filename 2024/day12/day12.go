package day12

import (
	"strings"
)

type Position struct {
	x int
	y int
}

type Plot struct {
	regionId  rune
	position  Position
	perimeter int
}

type Region struct {
	regionId rune
	plots    []Plot
}

func computeFencingPrice(input string) int {
	garden := parseInput(input)
	regions := computeRegions(garden)
	return computePrice(regions)
}

func computeDiscountedFencingPrice(input string) int {
	garden := parseInput(input)
	regions := computeRegions(garden)
	return computeDiscountedPrice(garden, regions)
}

func computePrice(regions []Region) int {
	price := 0
	for _, region := range regions {
		price += len(region.plots) * computeTotalPerimeter(region)
	}
	return price
}

func computeDiscountedPrice(garden [][]rune, regions []Region) int {
	price := 0
	for _, region := range regions {
		price += len(region.plots) * computedNumberSides(garden, region)
	}
	return price
}

func computeTotalPerimeter(region Region) int {
	perimeter := 0
	plots := region.plots
	for _, plot := range plots {
		perimeter += plot.perimeter
	}
	return perimeter
}

func computedNumberSides(garden [][]rune, region Region) int {
	nbSides := 0
	plots := region.plots
	for _, plot := range plots {
		if hasRightSide(garden, plot) && !plotAboveHasRightSide(garden, plot) {
			nbSides++
		}
		if hasLeftSide(garden, plot) && !plotAboveHasLeftSide(garden, plot) {
			nbSides++
		}
		if hasUpperSide(garden, plot) && !leftPlotHasUpperSide(garden, plot) {
			nbSides++
		}
		if hasLowerSide(garden, plot) && !leftPlotHasLowerSide(garden, plot) {
			nbSides++
		}
	}
	return nbSides
}

func hasRightSide(garden [][]rune, plot Plot) bool {
	if plot.position.x == len(garden[0])-1 {
		return true
	}
	if garden[plot.position.y][plot.position.x+1] != plot.regionId {
		return true
	}
	return false
}

func plotAboveHasRightSide(garden [][]rune, plot Plot) bool {
	if plot.position.y == 0 {
		return false
	}
	if garden[plot.position.y-1][plot.position.x] != plot.regionId {
		return false
	}
	return hasRightSide(garden, Plot{plot.regionId, Position{plot.position.x, plot.position.y - 1}, 0})
}

func hasLeftSide(garden [][]rune, plot Plot) bool {
	if plot.position.x == 0 {
		return true
	}
	if garden[plot.position.y][plot.position.x-1] != plot.regionId {
		return true
	}
	return false
}

func plotAboveHasLeftSide(garden [][]rune, plot Plot) bool {
	if plot.position.y == 0 {
		return false
	}
	if garden[plot.position.y-1][plot.position.x] != plot.regionId {
		return false
	}
	return hasLeftSide(garden, Plot{plot.regionId, Position{plot.position.x, plot.position.y - 1}, 0})
}

func hasUpperSide(garden [][]rune, plot Plot) bool {
	if plot.position.y == 0 {
		return true
	}
	if garden[plot.position.y-1][plot.position.x] != plot.regionId {
		return true
	}
	return false
}

func leftPlotHasUpperSide(garden [][]rune, plot Plot) bool {
	if plot.position.x == 0 {
		return false
	}
	if garden[plot.position.y][plot.position.x-1] != plot.regionId {
		return false
	}
	return hasUpperSide(garden, Plot{plot.regionId, Position{plot.position.x - 1, plot.position.y}, 0})
}

func hasLowerSide(garden [][]rune, plot Plot) bool {
	if plot.position.y == len(garden)-1 {
		return true
	}
	if garden[plot.position.y+1][plot.position.x] != plot.regionId {
		return true
	}
	return false
}

func leftPlotHasLowerSide(garden [][]rune, plot Plot) bool {
	if plot.position.x == 0 {
		return false
	}
	if garden[plot.position.y][plot.position.x-1] != plot.regionId {
		return false
	}
	return hasLowerSide(garden, Plot{plot.regionId, Position{plot.position.x - 1, plot.position.y}, 0})
}

func computeRegions(garden [][]rune) []Region {
	alreadyInComputedRegion := make(map[Position]bool)
	var regions []Region
	for y, row := range garden {
		for x, val := range row {
			currentPosition := Position{x, y}
			if alreadyInComputedRegion[currentPosition] {
				continue
			}
			newRegion := computeNewRegion(garden, currentPosition, val, alreadyInComputedRegion)
			regions = append(regions, newRegion)
		}
	}
	return regions
}

func computeNewRegion(garden [][]rune, startPosition Position, val rune, alreadyInComputedRegion map[Position]bool) Region {
	newRegion := Region{val, []Plot{}}
	computeNewRegionsRecursively(garden, startPosition, &newRegion, alreadyInComputedRegion)
	return newRegion
}

func computeNewRegionsRecursively(garden [][]rune, currentPosition Position, newRegion *Region, alreadyInComputedRegion map[Position]bool) {
	perimeter := computePerimeter(garden, currentPosition, newRegion.regionId)
	currentPlot := Plot{newRegion.regionId, currentPosition, perimeter}
	newRegion.plots = append(newRegion.plots, currentPlot)
	alreadyInComputedRegion[currentPosition] = true
	var nextPositions []Position
	if currentPosition.y > 0 {
		nextPosition := Position{currentPosition.x, currentPosition.y - 1}
		if garden[nextPosition.y][nextPosition.x] == newRegion.regionId {
			nextPositions = append(nextPositions, nextPosition)
		}
	}
	if currentPosition.y < len(garden)-1 {
		nextPosition := Position{currentPosition.x, currentPosition.y + 1}
		if garden[nextPosition.y][nextPosition.x] == newRegion.regionId {
			nextPositions = append(nextPositions, nextPosition)
		}
	}
	if currentPosition.x > 0 {
		nextPosition := Position{currentPosition.x - 1, currentPosition.y}
		if garden[nextPosition.y][nextPosition.x] == newRegion.regionId {
			nextPositions = append(nextPositions, nextPosition)
		}
	}
	if currentPosition.x < len(garden[0])-1 {
		nextPosition := Position{currentPosition.x + 1, currentPosition.y}
		if garden[nextPosition.y][nextPosition.x] == newRegion.regionId {
			nextPositions = append(nextPositions, nextPosition)
		}
	}
	for _, nextPosition := range nextPositions {
		if alreadyInComputedRegion[nextPosition] {
			continue
		}
		computeNewRegionsRecursively(garden, nextPosition, newRegion, alreadyInComputedRegion)
	}
}

func computePerimeter(garden [][]rune, currentPosition Position, regionId rune) int {
	perimeter := 0
	if currentPosition.y > 0 {
		if garden[currentPosition.y-1][currentPosition.x] != regionId {
			perimeter++
		}
	} else {
		perimeter++
	}
	if currentPosition.y < len(garden)-1 {
		if garden[currentPosition.y+1][currentPosition.x] != regionId {
			perimeter++
		}
	} else {
		perimeter++
	}
	if currentPosition.x > 0 {
		if garden[currentPosition.y][currentPosition.x-1] != regionId {
			perimeter++
		}
	} else {
		perimeter++
	}
	if currentPosition.x < len(garden[0])-1 {
		if garden[currentPosition.y][currentPosition.x+1] != regionId {
			perimeter++
		}
	} else {
		perimeter++
	}
	return perimeter
}

func parseInput(input string) [][]rune {
	rows := strings.Split(input, "\n")
	garden := make([][]rune, len(rows))
	for idx, row := range rows {
		gardenRow := make([]rune, len(row))
		for rowIdx, val := range row {
			gardenRow[rowIdx] = val
		}
		garden[idx] = gardenRow
	}
	return garden
}
