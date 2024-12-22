package day21

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Action int

const (
	UP    Action = iota
	RIGHT Action = iota
	DOWN  Action = iota
	LEFT  Action = iota
	PRESS Action = iota
)

type KeyPadType int

const (
	NUMERIC     KeyPadType = iota
	DIRECTIONAL KeyPadType = iota
)

type Position struct {
	x, y int
}

type Cell struct {
	value string
	cost  int
}

type KeyPad struct {
	keypadType KeyPadType
	cells      [][]Cell
}

func ComputeSumOfComplexities(input string) int {
	codes := parseInput(input)
	numericPadCache := make(map[string]map[string][]Action)
	sum := 0
	directionalPadCache := make(map[string]map[string][]Action)
	for _, code := range codes {
		complexity := computeComplexity(code, 2, numericPadCache, directionalPadCache)
		sum += complexity
	}
	return sum
}

func ComputeSumOfComplexitiesMoreRobots(input string) int {
	codes := parseInput(input)
	numericPadCache := make(map[string]map[string][]Action)
	sum := 0
	for _, code := range codes {
		complexity := computeComplexityAdvanced(code, 25, numericPadCache)
		sum += complexity
	}
	return sum
}

func buildDirectionalPad() KeyPad {
	rows := make([][]Cell, 2)
	rows[0] = make([]Cell, 3)
	rows[1] = make([]Cell, 3)
	rows[0][0] = Cell{"", 999999999}
	rows[0][1] = Cell{"^", 999999999}
	rows[0][2] = Cell{"A", 999999999}
	rows[1][0] = Cell{"<", 999999999}
	rows[1][1] = Cell{"v", 999999999}
	rows[1][2] = Cell{">", 999999999}
	return KeyPad{DIRECTIONAL, rows}
}

func buildNumericPad() KeyPad {
	rows := make([][]Cell, 4)
	rows[0] = make([]Cell, 3)
	rows[1] = make([]Cell, 3)
	rows[2] = make([]Cell, 3)
	rows[3] = make([]Cell, 3)
	rows[0][0] = Cell{"7", 999999999}
	rows[0][1] = Cell{"8", 999999999}
	rows[0][2] = Cell{"9", 999999999}
	rows[1][0] = Cell{"4", 999999999}
	rows[1][1] = Cell{"5", 999999999}
	rows[1][2] = Cell{"6", 999999999}
	rows[2][0] = Cell{"1", 999999999}
	rows[2][1] = Cell{"2", 999999999}
	rows[2][2] = Cell{"3", 999999999}
	rows[3][0] = Cell{"", 999999999}
	rows[3][1] = Cell{"0", 999999999}
	rows[3][2] = Cell{"A", 999999999}
	return KeyPad{NUMERIC, rows}
}

func computeComplexity(code string, nbRobots int, numericPadCache map[string]map[string][]Action, directionalPadCache map[string]map[string][]Action) int {
	firstPadSequence := computeNumericPadSequence(code, numericPadCache)
	currentSequence := firstPadSequence
	for i := 0; i < nbRobots; i++ {
		nextSequence := computeDirectionalPadSequence(formatActions(currentSequence), directionalPadCache)
		currentSequence = nextSequence
		fmt.Printf("Sequence: %s", formatActions(currentSequence))
		fmt.Println()
	}
	return len(currentSequence) * extractNumbericPart(code)
}

func computeComplexityAdvanced(code string, nbRobots int, numericPadCache map[string]map[string][]Action) int {
	firstPadSequence := computeNumericPadSequence(code, numericPadCache)
	currentSequenceOccurences := groupFirstPadSequence(firstPadSequence)
	directionalPadCache := make(map[string][]string)
	for i := 0; i < nbRobots; i++ {
		nextSequenceOccurrences := computeDirectionalPadSequenceAdvanced(currentSequenceOccurences, directionalPadCache)
		currentSequenceOccurences = nextSequenceOccurrences
	}
	return computeSequenceLength(currentSequenceOccurences) * extractNumbericPart(code)
}

func computeSequenceLength(currentSequenceOccurences map[string]int) int {
	length := 0
	for sequence, count := range currentSequenceOccurences {
		length += len(sequence) * count
	}
	return length
}

func groupFirstPadSequence(firstPadSequence []Action) map[string]int {
	var subSequence []Action
	occurrences := make(map[string]int)
	for _, action := range firstPadSequence {
		subSequence = append(subSequence, action)
		if action == PRESS {
			if count, ok := occurrences[formatActions(subSequence)]; ok {
				occurrences[formatActions(subSequence)] = count + 1
			} else {
				occurrences[formatActions(subSequence)] = 1
			}
			subSequence = nil
		}
	}
	return occurrences
}

func extractNumbericPart(code string) int {
	re := regexp.MustCompile(`^([0-9]+)A`)
	matches := re.FindAllSubmatch([]byte(code), -1)
	numericPart := matches[0][1]
	number, _ := strconv.Atoi(string(numericPart))
	return number
}

func computeNumericPadSequence(code string, numericPadCache map[string]map[string][]Action) []Action {
	currentNumber := "A"
	var sequence []Action
	for i := 0; i < len(code); i++ {
		nextNumber := string(code[i])
		if val, ok := numericPadCache[currentNumber][nextNumber]; ok {
			sequence = append(sequence, val...)
		}
		if val, ok := numericPadCache[nextNumber][currentNumber]; ok {
			sequence = append(sequence, revertActions(val)...)
		}
		actions := findSequenceNumericPad(currentNumber, nextNumber)
		sequence = append(sequence, actions...)
		currentNumber = nextNumber

	}
	return sequence
}

func computeDirectionalPadSequence(code string, directionalPadCache map[string]map[string][]Action) []Action {
	currentNumber := "A"
	var sequence []Action
	for i := 0; i < len(code); i++ {
		nextNumber := string(code[i])
		if val, ok := directionalPadCache[currentNumber][nextNumber]; ok {
			sequence = append(sequence, val...)
			currentNumber = nextNumber
			continue
		}
		actions := findSequenceDirectionalPad(currentNumber, nextNumber)
		sequence = append(sequence, actions...)
		if _, ok := directionalPadCache[currentNumber]; !ok {
			directionalPadCache[currentNumber] = make(map[string][]Action)
		}
		directionalPadCache[currentNumber][nextNumber] = actions
		currentNumber = nextNumber

	}
	return sequence
}

func computeDirectionalPadSequenceAdvanced(currentOccurrencesCount map[string]int, subSequencesCache map[string][]string) map[string]int {
	nextOccurrencesCount := make(map[string]int)
	for sequence, count := range currentOccurrencesCount {
		if nextSubSequences, ok := subSequencesCache[sequence]; ok {
			for _, nextSubSequence := range nextSubSequences {
				if nextCount, ok := nextOccurrencesCount[nextSubSequence]; ok {
					nextOccurrencesCount[nextSubSequence] = nextCount + count
				} else {
					nextOccurrencesCount[nextSubSequence] = count
				}
			}
			continue
		}
		nextSubSequences := computeDirectionalPadSequenceUntilPress(sequence)
		subSequencesCache[sequence] = nextSubSequences
		for _, nextSubSequence := range nextSubSequences {
			if nextCount, ok := nextOccurrencesCount[nextSubSequence]; ok {
				nextOccurrencesCount[nextSubSequence] = nextCount + count
			} else {
				nextOccurrencesCount[nextSubSequence] = count
			}
		}
	}
	return nextOccurrencesCount
}

func computeDirectionalPadSequenceUntilPress(code string) []string {
	currentNumber := "A"
	var sequences []string
	for i := 0; i < len(code); i++ {
		nextNumber := string(code[i])
		actions := findSequenceDirectionalPad(currentNumber, nextNumber)
		sequences = append(sequences, formatActions(actions))
		currentNumber = nextNumber
	}
	return sequences
}

func formatActions(actions []Action) string {
	formattedActions := ""
	for _, action := range actions {
		switch action {
		case UP:
			formattedActions += "^"
		case DOWN:
			formattedActions += "v"
		case LEFT:
			formattedActions += "<"
		case RIGHT:
			formattedActions += ">"
		case PRESS:
			formattedActions += "A"
		}
	}
	return formattedActions
}

func findSequenceNumericPad(currentNumber, nextNumber string) []Action {
	keypad := buildNumericPad()
	currentPosition := findPosition(keypad, currentNumber)
	targetPosition := findPosition(keypad, nextNumber)
	updateCosts(&keypad, currentPosition, nextNumber, 0)
	return findShortestPath(keypad, currentPosition, targetPosition)
}

func findSequenceDirectionalPad(currentNumber, nextNumber string) []Action {
	keypad := buildDirectionalPad()
	currentPosition := findPosition(keypad, currentNumber)
	targetPosition := findPosition(keypad, nextNumber)
	updateCosts(&keypad, currentPosition, nextNumber, 0)
	return findShortestPath(keypad, currentPosition, targetPosition)
}

func findShortestPath(keypad KeyPad, currentPosition, targetPosition Position) []Action {
	x := targetPosition.x
	y := targetPosition.y
	var actions []Action
	currentCost := keypad.cells[targetPosition.y][targetPosition.x].cost
	for {
		if currentCost == 0 {
			actions = append(actions, PRESS)
			return groupActions(keypad.keypadType, currentPosition, actions)
		}
		if y > 0 {
			if keypad.cells[y-1][x].cost == currentCost-1 {
				actions = append(actions, DOWN)
				y -= 1
				currentCost -= 1
			}
		}
		if y < len(keypad.cells)-1 {
			if keypad.cells[y+1][x].cost == currentCost-1 {
				actions = append(actions, UP)
				y += 1
				currentCost -= 1
			}
		}
		if x > 0 {
			if keypad.cells[y][x-1].cost == currentCost-1 {
				actions = append(actions, RIGHT)
				x -= 1
				currentCost -= 1
			}
		}
		if x < len(keypad.cells[0])-1 {
			if keypad.cells[y][x+1].cost == currentCost-1 {
				actions = append(actions, LEFT)
				x += 1
				currentCost -= 1
			}
		}
	}
}

func groupActions(keypadType KeyPadType, startPosition Position, actions []Action) []Action {
	actionsGroup := make(map[Action]int)
	for _, action := range actions {
		if val, ok := actionsGroup[action]; ok {
			actionsGroup[action] = val + 1
		} else {
			actionsGroup[action] = 1
		}
	}

	if keypadType == NUMERIC {
		return groupActionsNumericKeyPad(startPosition, actionsGroup)
	} else {
		return groupActionsDirectionalKeyPad(startPosition, actionsGroup)
	}
}

func groupActionsNumericKeyPad(startPosition Position, actionsGroup map[Action]int) []Action {
	var groupedActions []Action
	goFirstLeft := startPosition.y != 3 || (startPosition.y == 3 && startPosition.x-actionsGroup[LEFT] > 0)
	goEarlyDown := startPosition.x != 0 || (startPosition.x == 0 && startPosition.y+actionsGroup[DOWN] < 3)
	if goFirstLeft {
		for i := 0; i < actionsGroup[LEFT]; i++ {
			groupedActions = append(groupedActions, LEFT)
		}
	}
	if goEarlyDown {
		for i := 0; i < actionsGroup[DOWN]; i++ {
			groupedActions = append(groupedActions, DOWN)
		}
	}
	for i := 0; i < actionsGroup[UP]; i++ {
		groupedActions = append(groupedActions, UP)
	}
	if !goFirstLeft {
		for i := 0; i < actionsGroup[LEFT]; i++ {
			groupedActions = append(groupedActions, LEFT)
		}
	}
	for i := 0; i < actionsGroup[RIGHT]; i++ {
		groupedActions = append(groupedActions, RIGHT)
	}
	if !goEarlyDown {
		for i := 0; i < actionsGroup[DOWN]; i++ {
			groupedActions = append(groupedActions, DOWN)
		}
	}
	groupedActions = append(groupedActions, PRESS)
	return groupedActions
}

func groupActionsDirectionalKeyPad(startPosition Position, actionsGroup map[Action]int) []Action {
	var groupedActions []Action
	goFirstLeft := startPosition.y != 0 || (startPosition.y == 0 && startPosition.x-actionsGroup[LEFT] > 0)
	if goFirstLeft {
		for i := 0; i < actionsGroup[LEFT]; i++ {
			groupedActions = append(groupedActions, LEFT)
		}
	}
	for i := 0; i < actionsGroup[DOWN]; i++ {
		groupedActions = append(groupedActions, DOWN)
	}
	if startPosition.x != 0 {
		for i := 0; i < actionsGroup[UP]; i++ {
			groupedActions = append(groupedActions, UP)
		}
	}
	if !goFirstLeft {
		for i := 0; i < actionsGroup[LEFT]; i++ {
			groupedActions = append(groupedActions, LEFT)
		}
	}
	for i := 0; i < actionsGroup[RIGHT]; i++ {
		groupedActions = append(groupedActions, RIGHT)
	}
	if startPosition.x == 0 {
		for i := 0; i < actionsGroup[UP]; i++ {
			groupedActions = append(groupedActions, UP)
		}
	}
	groupedActions = append(groupedActions, PRESS)
	return groupedActions
}

func updateCosts(keypad *KeyPad, currentPosition Position, targetNumber string, currentCost int) {
	currentCell := keypad.cells[currentPosition.y][currentPosition.x]
	keypad.cells[currentPosition.y][currentPosition.x].cost = currentCost
	if currentCell.value == targetNumber {
		return
	}
	nextCost := currentCost + 1
	if currentPosition.y > 0 {
		nextCell := keypad.cells[currentPosition.y-1][currentPosition.x]
		if nextCell.value != "" && nextCell.cost > nextCost {
			updateCosts(keypad, Position{currentPosition.x, currentPosition.y - 1}, targetNumber, nextCost)
		}
	}
	if currentPosition.y < len(keypad.cells)-1 {
		nextCell := keypad.cells[currentPosition.y+1][currentPosition.x]
		if nextCell.value != "" && nextCell.cost > nextCost {
			updateCosts(keypad, Position{currentPosition.x, currentPosition.y + 1}, targetNumber, nextCost)
		}
	}
	if currentPosition.x > 0 {
		nextCell := keypad.cells[currentPosition.y][currentPosition.x-1]
		if nextCell.value != "" && nextCell.cost > nextCost {
			updateCosts(keypad, Position{currentPosition.x - 1, currentPosition.y}, targetNumber, nextCost)
		}
	}
	if currentPosition.x < len(keypad.cells[0])-1 {
		nextCell := keypad.cells[currentPosition.y][currentPosition.x+1]
		if nextCell.value != "" && nextCell.cost > nextCost {
			updateCosts(keypad, Position{currentPosition.x + 1, currentPosition.y}, targetNumber, nextCost)
		}
	}

}

func findPosition(keypad KeyPad, value string) Position {
	for y, row := range keypad.cells {
		for x, cell := range row {
			if cell.value == value {
				return Position{x, y}
			}
		}
	}
	return Position{-1, -1}
}

func revertActions(actions []Action) []Action {
	var revertedActions []Action
	for i := len(actions) - 1; i >= 0; i-- {
		switch actions[i] {
		case UP:
			revertedActions = append(revertedActions, DOWN)
		case RIGHT:
			revertedActions = append(revertedActions, LEFT)
		case DOWN:
			revertedActions = append(revertedActions, UP)
		case LEFT:
			revertedActions = append(revertedActions, RIGHT)
		}
	}
	revertedActions = append(revertedActions, PRESS)
	return revertedActions
}

func parseInput(input string) []string {
	lines := strings.Split(input, "\n")
	codes := make([]string, len(lines))
	for idx, code := range lines {
		codes[idx] = code
	}
	return codes
}
