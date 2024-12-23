package day22

import (
	"migwel/aoc/2024/util"
	"testing"
)

func TestComputeSumOfSecrets_Simple(t *testing.T) {
	var result = ComputeSumOfSecrets(`1
10
100
2024`, 2000)
	if result != 37327623 {
		t.Errorf("Got %d instead of 37327623", result)
	}
}

func TestComputeSumOfSecrets_Input(t *testing.T) {
	var result = ComputeSumOfSecrets(util.ReadFile("input.txt"), 2000)
	if result != 202274 {
		t.Errorf("Got %d instead of 202274", result)
	}
}

func TestComputeMostBananas_Simple(t *testing.T) {
	var result = ComputeMostBananas(`1
2
3
2024`, 2000)
	if result != 23 {
		t.Errorf("Got %d instead of 23", result)
	}
}

func TestComputeMostBananas_Input(t *testing.T) {
	var result = ComputeMostBananas(util.ReadFile("input.txt"), 2000)
	if result != 202274 {
		t.Errorf("Got %d instead of 202274", result)
	}
}
