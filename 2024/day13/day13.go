package day13

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type Button struct {
	xTranslation int
	yTranslation int
}

type Position struct {
	x int
	y int
}

type Machine struct {
	aButton       Button
	bButton       Button
	prizePosition Position
}

func computeFewestTokensToWin(input string) int {
	machines := parseInput(input)
	totalNbTokens := 0
	for _, machine := range machines {
		nbTokens := computeNbTokensToWin(machine)
		if nbTokens != -1 {
			totalNbTokens += nbTokens
		}
	}
	return totalNbTokens
}

func computeFewestTokensToWinNewRules(input string) int {
	machines := parseInputNewRules(input)
	totalNbTokens := 0
	for _, machine := range machines {
		nbTokens := computeNbTokensToWinNewRules(machine)
		if nbTokens != -1 {
			totalNbTokens += nbTokens
		}
	}
	return totalNbTokens
}

func computeNbTokensToWinNewRules(machine Machine) int {
	alpha, beta, err := solveEquations(machine)
	if err != nil {
		return -1
	}
	if !validAnswerNewRules(alpha, beta, machine) {
		return -1
	}
	return 3*alpha + beta
}

func computeNbTokensToWin(machine Machine) int {
	alpha, beta, err := solveEquations(machine)
	if !validAnswer(alpha, beta, machine) {
		return -1
	}
	if err != nil {
		return -1
	}
	return 3*alpha + beta
}

func solveEquations(machine Machine) (int, int, error) {
	// Applying Cramer's Rule
	denominator := machine.aButton.xTranslation*machine.bButton.yTranslation - machine.bButton.xTranslation*machine.aButton.yTranslation
	if denominator == 0 {
		return 0, 0, errors.New("Dividing by 0")
	}
	alphaNumerator := machine.prizePosition.x*machine.bButton.yTranslation - machine.bButton.xTranslation*machine.prizePosition.y
	alpha := alphaNumerator / denominator

	betaNumerator := machine.aButton.xTranslation*machine.prizePosition.y - machine.prizePosition.x*machine.aButton.yTranslation
	beta := betaNumerator / denominator
	return alpha, beta, nil
}

func validAnswerNewRules(alpha, beta int, machine Machine) bool {
	x := alpha*machine.aButton.xTranslation + beta*machine.bButton.xTranslation
	if x != machine.prizePosition.x {
		return false
	}
	y := alpha*machine.aButton.yTranslation + beta*machine.bButton.yTranslation
	if y != machine.prizePosition.y {
		return false
	}
	return true
}

func validAnswer(alpha, beta int, machine Machine) bool {
	if alpha > 100 || beta > 100 {
		return false
	}
	x := alpha*machine.aButton.xTranslation + beta*machine.bButton.xTranslation
	if x != machine.prizePosition.x {
		return false
	}
	y := alpha*machine.aButton.yTranslation + beta*machine.bButton.yTranslation
	if y != machine.prizePosition.y {
		return false
	}
	return true
}

func parseInput(input string) []Machine {
	lines := strings.Split(input, "\n")
	var machines []Machine
	for i := 0; i < len(lines); i += 4 {
		aButton := parseButton(lines[i])
		bButton := parseButton(lines[i+1])
		prizePosition := parsePrizePosition(lines[i+2])
		machine := Machine{aButton, bButton, prizePosition}
		machines = append(machines, machine)
	}
	return machines
}

func parseInputNewRules(input string) []Machine {
	lines := strings.Split(input, "\n")
	var machines []Machine
	for i := 0; i < len(lines); i += 4 {
		aButton := parseButton(lines[i])
		bButton := parseButton(lines[i+1])
		prizePosition := parsePrizePositionNewRules(lines[i+2])
		machine := Machine{aButton, bButton, prizePosition}
		machines = append(machines, machine)
	}
	return machines
}

func parsePrizePositionNewRules(input string) Position {
	re := regexp.MustCompile(`X=([0-9]+), Y=([0-9]+)`)
	matches := re.FindAllSubmatch([]byte(input), -1)
	x, _ := strconv.Atoi((string)(matches[0][1]))
	y, _ := strconv.Atoi((string)(matches[0][2]))
	return Position{x + 10000000000000, y + 10000000000000}
}

func parsePrizePosition(input string) Position {
	re := regexp.MustCompile(`X=([0-9]+), Y=([0-9]+)`)
	matches := re.FindAllSubmatch([]byte(input), -1)
	x, _ := strconv.Atoi((string)(matches[0][1]))
	y, _ := strconv.Atoi((string)(matches[0][2]))
	return Position{x, y}
}

func parseButton(input string) Button {
	re := regexp.MustCompile(`X\+([0-9]+), Y\+([0-9]+)`)
	matches := re.FindAllSubmatch([]byte(input), -1)
	xTranslation, _ := strconv.Atoi((string)(matches[0][1]))
	yTranslation, _ := strconv.Atoi((string)(matches[0][2]))
	return Button{xTranslation, yTranslation}
}
