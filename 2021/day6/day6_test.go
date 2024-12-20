package day6

import (
	"migwel/aoc/2021/util"
	"testing"
)

func TestCComputeNumberLanternFishes_Simple(t *testing.T) {
	var result = ComputeNumberLanternFishes(`3,4,3,1,2`, 80)
	if result != 5934 {
		t.Errorf("Got %d instead of 5934", result)
	}
}

func TestComputeNumberLanternFishes_Input(t *testing.T) {
	var result = ComputeNumberLanternFishes(util.ReadFile("input.txt"), 80)
	if result != 365131 {
		t.Errorf("Got %d instead of 365131", result)
	}
}

func TestCComputeNumberLanternFishesButLonger_Simple(t *testing.T) {
	var result = ComputeNumberLanternFishes(`3,4,3,1,2`, 256)
	if result != 26984457539 {
		t.Errorf("Got %d instead of 26984457539", result)
	}
}

func TestComputeNumberLanternFishesButLonger_Input(t *testing.T) {
	var result = ComputeNumberLanternFishes(util.ReadFile("input.txt"), 256)
	if result != 1650309278600 {
		t.Errorf("Got %d instead of 1650309278600", result)
	}
}
