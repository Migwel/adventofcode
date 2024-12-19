package day17

import (
	"testing"
)

func TestExecuteProgram_Simple(t *testing.T) {
	var result = executeProgram(`Register A: 0
Register B: 0
Register C: 9

Program: 2,6,2,6`)
	if result != "" {
		t.Errorf("Got %s instead of empty", result)
	}
}

func TestExecuteProgram_Simple2(t *testing.T) {
	var result = executeProgram(`Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`)
	if result != "4,6,3,5,6,3,5,2,1,0" {
		t.Errorf("Got %s instead of 4,6,3,5,6,3,5,2,1,0", result)
	}
}

func TestExecuteProgram_Input(t *testing.T) {
	var result = executeProgram(`Register A: 48744869
Register B: 0
Register C: 0

Program: 2,4,1,2,7,5,1,3,4,4,5,5,0,3,3,0`)
	if result != "7,1,5,2,4,0,7,6,1" {
		t.Errorf("Got %s instead of 7,1,5,2,4,0,7,6,1", result)
	}
}

func TestCopyProgram_Simple(t *testing.T) {
	var result = copyProgram(`Register A: 2024
Register B: 0
Register C: 0

Program: 0,3,5,4,3,0`)
	if result != 117440 {
		t.Errorf("Got %d instead of 117440", result)
	}
}

func TestCopyProgram_Input(t *testing.T) {
	var result = copyProgram(`Register A: 2024
Register B: 0
Register C: 0

Program: 2,4,1,2,7,5,1,3,4,4,5,5,0,3,3,0`)
	if result != 37222273957364 {
		t.Errorf("Got %d instead of 37222273957364", result)
	}
}
