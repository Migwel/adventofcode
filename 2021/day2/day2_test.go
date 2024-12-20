package day2

import (
	"migwel/aoc/2021/util"
	"testing"
)

func TestComputePositionAndDepthResult_Simple(t *testing.T) {
	var dist = ComputePositionAndDepthResult(`forward 5
down 5
forward 8
up 3
down 8
forward 2`)
	if dist != 150 {
		t.Errorf("Got %d instead of 150", dist)
	}
}

func TestComputePositionAndDepthResult_Input(t *testing.T) {
	var dist = ComputePositionAndDepthResult(util.ReadFile("input.txt"))
	if dist != 2039256 {
		t.Errorf("Got %d instead of 2039256", dist)
	}
}

func TestComputePositionAndDepthResultWithAim_Simple(t *testing.T) {
	var dist = ComputePositionAndDepthResultWithAim(`forward 5
down 5
forward 8
up 3
down 8
forward 2`)
	if dist != 900 {
		t.Errorf("Got %d instead of 900", dist)
	}
}

func TestComputePositionAndDepthResultWithAim_Input(t *testing.T) {
	var dist = ComputePositionAndDepthResultWithAim(util.ReadFile("input.txt"))
	if dist != 1856459736 {
		t.Errorf("Got %d instead of 1856459736", dist)
	}
}
