package day5

import (
	"migwel/aoc/2021/util"
	"testing"
)

func TestComputePointsWithOverlap_Simple(t *testing.T) {
	var result = ComputePointsWithOverlap(`0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`, false)
	if result != 5 {
		t.Errorf("Got %d instead of 5", result)
	}
}

func TestComputePointsWithOverlap_Input(t *testing.T) {
	var result = ComputePointsWithOverlap(util.ReadFile("input.txt"), false)
	if result != 7380 {
		t.Errorf("Got %d instead of 7380", result)
	}
}

func TestComputePointsWithOverlapWithDiagonals_Simple(t *testing.T) {
	var result = ComputePointsWithOverlap(`0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`, true)
	if result != 12 {
		t.Errorf("Got %d instead of 12", result)
	}
}

func TestComputePointsWithOverlapWithDiagonals_Input(t *testing.T) {
	var result = ComputePointsWithOverlap(util.ReadFile("input.txt"), true)
	if result != 21373 {
		t.Errorf("Got %d instead of 21373", result)
	}
}
