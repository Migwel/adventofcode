package day17

import (
	"math"
	"regexp"
	"strconv"
	"strings"
)

type Register int

type Instruction struct {
	opcode, operand int
}

func executeProgram(input string) string {
	aRegister, bRegister, cRegister, instructions := parseInput(input)
	output := executeInstructions(aRegister, bRegister, cRegister, instructions)
	return output
}

func copyProgram(input string) int {
	_, bRegister, cRegister, instructions := parseInput(input)
	trackingIndex := 0
	subOutput := computeSubOutput(trackingIndex, instructions)
	for i := 0; ; i++ {
		output := executeInstructions(Register(i), bRegister, cRegister, instructions)
		if output == subOutput {
			trackingIndex += 1
			if trackingIndex == 2*len(instructions) {
				return i
			}
			i = i << 3
			i -= 1
			subOutput = computeSubOutput(trackingIndex, instructions)
		}
	}
	return -1
}

func computeSubOutput(trackingIndex int, instructions []Instruction) string {
	nextValue := ""
	for i := trackingIndex; i >= 0; i-- {
		if i != trackingIndex {
			nextValue += ","
		}
		if i%2 == 0 {
			nextValue += strconv.Itoa(instructions[len(instructions)-1-i/2].operand)
		} else {
			nextValue += strconv.Itoa(instructions[len(instructions)-1-i/2].opcode)
		}
	}
	return nextValue
}

func isExpectedOutput(aRegister, bRegister, cRegister Register, instructions []Instruction, expectedOutput string) bool {
	output := ""
	for i := 0; i < len(instructions); i++ {
		instruction := instructions[i]
		instructionOutput, nextI := executeInstruction(&aRegister, &bRegister, &cRegister, instruction)
		if instructionOutput != "" {
			if output == "" {
				output = instructionOutput
			} else {
				output += "," + instructionOutput
			}
			if output == expectedOutput {
				return true
			}
			if !canBeExpectedOutput(output, expectedOutput) {
				return false
			}
		}
		if nextI != -1 {
			i = nextI - 1
		}
	}
	return false
}

func canBeExpectedOutput(output, expectedOutput string) bool {
	if len(output) > len(expectedOutput) {
		return false
	}
	for idx := 0; idx < len(output); idx++ {
		if output[idx] != expectedOutput[idx] {
			return false
		}
	}
	return true
}

func executeInstructions(aRegister, bRegister, cRegister Register, instructions []Instruction) string {
	output := ""
	for i := 0; i < len(instructions); i++ {
		instruction := instructions[i]
		instructionOutput, nextI := executeInstruction(&aRegister, &bRegister, &cRegister, instruction)
		if instructionOutput != "" {
			if output == "" {
				output = instructionOutput
			} else {
				output += "," + instructionOutput
			}
		}
		if nextI != -1 {
			i = nextI - 1
		}
	}
	return output
}

func executeInstruction(aRegister, bRegister, cRegister *Register, instruction Instruction) (string, int) {
	switch instruction.opcode {
	case 0:
		result := adv(*aRegister, computeComboOperand(*aRegister, *bRegister, *cRegister, instruction.operand))
		*aRegister = Register(result)
		return "", -1
	case 1:
		result := bxl(*bRegister, instruction.operand)
		*bRegister = Register(result)
		return "", -1
	case 2:
		result := bst(computeComboOperand(*aRegister, *bRegister, *cRegister, instruction.operand))
		*bRegister = Register(result)
		return "", -1
	case 3:
		result := jnz(*aRegister, instruction.operand)
		return "", result
	case 4:
		result := bxc(*bRegister, *cRegister)
		*bRegister = Register(result)
		return "", -1
	case 5:
		result := out(computeComboOperand(*aRegister, *bRegister, *cRegister, instruction.operand))
		return result, -1
	case 6:
		result := adv(*aRegister, computeComboOperand(*aRegister, *bRegister, *cRegister, instruction.operand))
		*bRegister = Register(result)
		return "", -1
	case 7:
		result := adv(*aRegister, computeComboOperand(*aRegister, *bRegister, *cRegister, instruction.operand))
		*cRegister = Register(result)
		return "", -1
	}
	return "", -1
}

func out(operand int) string {
	value := operand % 8
	strValue := strconv.Itoa(value)
	return strValue
}

func bxc(bRegister, cRegister Register) int {
	return int(bRegister) ^ int(cRegister)
}

func jnz(aRegister Register, operand int) int {
	if aRegister == 0 {
		return -1
	}
	return operand
}

func bst(operand int) int {
	return operand % 8
}

func adv(aRegister Register, operand int) int {
	denominator := math.Pow(2, float64(operand))
	return int(aRegister) / int(denominator)
}

func bxl(bRegister Register, operand int) int {
	return int(bRegister) ^ operand
}

func computeComboOperand(aRegister, bRegister, cRegister Register, operand int) int {
	switch operand {
	case 0, 1, 2, 3:
		return operand
	case 4:
		return int(aRegister)
	case 5:
		return int(bRegister)
	case 6:
		return int(cRegister)
	}
	return -1
}

func parseInput(input string) (Register, Register, Register, []Instruction) {
	lines := strings.Split(input, "\n")
	aRegister := parseRegister(lines[0])
	bRegister := parseRegister(lines[1])
	cRegister := parseRegister(lines[2])
	instructions := parseProgram(lines[4])
	return aRegister, bRegister, cRegister, instructions
}

func parseProgram(programString string) []Instruction {
	values := strings.Split(programString[9:], ",")
	var instructions []Instruction
	for i := 0; i < len(values); i += 2 {
		opcode, _ := strconv.Atoi(values[i])
		operand, _ := strconv.Atoi(values[i+1])
		instructions = append(instructions, Instruction{opcode, operand})
	}
	return instructions
}

func parseRegister(registerInput string) Register {
	re := regexp.MustCompile(`^Register [ABC]: ([0-9]+)$`)
	matches := re.FindAllSubmatch([]byte(registerInput), -1)
	registerValue, _ := strconv.Atoi(string(matches[0][1]))
	return Register(registerValue)
}
