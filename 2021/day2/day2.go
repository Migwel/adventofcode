package day2

import (
	"strconv"
	"strings"
)

type Action int

const (
	FORWARD Action = iota
	DOWN    Action = iota
	UP      Action = iota
)

type Command struct {
	action Action
	value  int
}

type Position struct {
	x, depth int
}

type Aim struct {
	value int
}

func ComputePositionAndDepthResult(input string) int {
	commands := parseInput(input)
	position := applyCommands(commands)
	return position.x * position.depth
}

func ComputePositionAndDepthResultWithAim(input string) int {
	commands := parseInput(input)
	position := applyCommandsWithAim(commands)
	return position.x * position.depth
}

func applyCommands(commands []Command) Position {
	position := Position{0, 0}
	for _, command := range commands {
		applyCommand(command, &position)
	}
	return position
}

func applyCommand(command Command, position *Position) {
	switch command.action {
	case FORWARD:
		position.x = position.x + command.value
	case UP:
		position.depth = position.depth - command.value
	case DOWN:
		position.depth = position.depth + command.value
	}
}

func applyCommandsWithAim(commands []Command) Position {
	position := Position{0, 0}
	aim := Aim{0}
	for _, command := range commands {
		applyCommandWithAim(command, &position, &aim)
	}
	return position
}

func applyCommandWithAim(command Command, position *Position, aim *Aim) {
	switch command.action {
	case FORWARD:
		position.x = position.x + command.value
		position.depth = position.depth + command.value*aim.value
	case UP:
		aim.value = aim.value - command.value
	case DOWN:
		aim.value = aim.value + command.value
	}
}

func parseInput(input string) []Command {
	lines := strings.Split(input, "\n")
	commands := make([]Command, len(lines))
	for idx, line := range lines {
		split := strings.Split(line, " ")
		var action Action
		switch split[0] {
		case "forward":
			action = FORWARD
		case "down":
			action = DOWN
		case "up":
			action = UP
		}
		value, _ := strconv.Atoi(split[1])
		command := Command{action, value}
		commands[idx] = command
	}
	return commands
}
