package day9

import (
	"strconv"
)

func computeChecksum(input string) int {
	compactedBlocks := computeCompactedBlocks(input)
	blocks := expandBlocks(compactedBlocks)
	moveBlocks(blocks)
	return computeBlocksChecksum(blocks)
}

func computeDefragmentedChecksum(input string) int {
	compactedBlocks := computeCompactedBlocks(input)
	blocks := expandBlocks(compactedBlocks)
	moveDefragmentedBlocks(blocks)
	return computeBlocksChecksum(blocks)
}

func computeBlocksChecksum(blocks []int) int {
	checkSum := 0
	for idx, val := range blocks {
		if val == -1 {
			continue
		}
		checkSum += idx * val
	}
	return checkSum
}

func expandBlocks(compactedBlocks []int) []int {
	size := 0
	for _, val := range compactedBlocks {
		size += val
	}
	blocks := make([]int, size)
	id := 0
	blockIdx := 0
	for idx, val := range compactedBlocks {
		var blockValue int
		if idx%2 == 0 {
			blockValue = id
			id += 1
		} else {
			blockValue = -1
		}
		for i := 0; i < val; i++ {
			blocks[blockIdx] = blockValue
			blockIdx += 1
		}
	}
	return blocks
}

func computeCompactedBlocks(input string) []int {
	blocks := make([]int, len(input))
	for idx, val := range input {
		blocks[idx], _ = strconv.Atoi((string)(val))
	}
	return blocks
}

func moveBlocks(blocks []int) {
	freeSpaceIdx := 0
	for i := len(blocks) - 1; i > freeSpaceIdx; i-- {
		currentId := blocks[i]
		if currentId == -1 {
			continue
		}
		freeSpaceIdx = findFreeSpaceIdx(blocks, freeSpaceIdx)
		if freeSpaceIdx == -1 || freeSpaceIdx >= i {
			return
		}
		blocks[freeSpaceIdx] = currentId
		blocks[i] = -1
		freeSpaceIdx++
	}
}

func moveDefragmentedBlocks(blocks []int) {
	for i := len(blocks) - 1; i >= 0; {
		currentId := blocks[i]
		if currentId == -1 {
			i -= 1
			continue
		}
		moveWholeFile(i, blocks)
		currentFileSize := findCurrentIdSize(i, blocks)
		i -= currentFileSize
	}
}

func moveWholeFile(currentFileId int, blocks []int) {
	currentFileSize := findCurrentIdSize(currentFileId, blocks)
	currentVal := blocks[currentFileId]

	for freeSpaceIdx := findFreeSpaceIdx(blocks, 0); freeSpaceIdx < currentFileId; {
		if freeSpaceIdx >= currentFileId {
			return
		}
		freeSpaceSize := findFreeSpaceSize(freeSpaceIdx, blocks)
		if freeSpaceSize < currentFileSize {
			freeSpaceIdx = findFreeSpaceIdx(blocks, freeSpaceIdx+freeSpaceSize)
			continue
		}
		for i := 0; i < currentFileSize; i++ {
			blocks[freeSpaceIdx+i] = currentVal
			blocks[currentFileId-i] = -1
		}
		return
	}

}

func findCurrentIdSize(currentId int, blocks []int) int {
	currentVal := blocks[currentId]
	currentIdSize := 0
	for i := currentId; i >= 0; i-- {
		if blocks[i] == currentVal {
			currentIdSize += 1
		} else {
			break
		}
	}
	return currentIdSize
}

func findFreeSpaceSize(freeSpaceId int, blocks []int) int {
	currentIdSize := 0
	for i := freeSpaceId; i < len(blocks); i++ {
		if blocks[i] == -1 {
			currentIdSize += 1
		} else {
			break
		}
	}
	return currentIdSize
}

func findFreeSpaceIdx(blocks []int, freeSpaceIdx int) int {
	for i := freeSpaceIdx; i < len(blocks); i++ {
		if blocks[i] == -1 {
			return i
		}
	}
	return -1
}
