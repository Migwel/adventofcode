package day21

import (
	"migwel/aoc/2024/util"
	"testing"
)

func TestComputeSumOfComplexities_Simple(t *testing.T) {
	var result = ComputeSumOfComplexities(`029A`)
	if result != 126384 {
		t.Errorf("Got %d instead of 126384", result)
	}
}

func TestComputeSumOfComplexities_Input(t *testing.T) {
	var result = ComputeSumOfComplexities(util.ReadFile("input.txt"))
	if result != 202274 {
		t.Errorf("Got %d instead of 202274", result)
	}
}

func TestComputeSumOfComplexitiesMoreRobots_Simple(t *testing.T) {
	var result = ComputeSumOfComplexitiesMoreRobots(`029A
980A
179A
456A
379A`)
	if result != 126384 {
		t.Errorf("Got %d instead of 126384", result)
	}
}

func TestComputeSumOfComplexitiesMoreRobots_Input(t *testing.T) {
	var result = ComputeSumOfComplexitiesMoreRobots(util.ReadFile("input.txt"))
	if result != 202274 {
		t.Errorf("Got %d instead of 202274", result)
	}
}
