package day1

import (
	"migwel/aoc/2021/util"
	"testing"
)

func TestComputeIncreasesNumber_Simple(t *testing.T) {
	var dist = ComputeIncreasesNumber(`199
200
208
210
200
207
240
269
260
263`)
	if dist != 7 {
		t.Errorf("Got %d instead of 7", dist)
	}
}

func TestComputeIncreasesNumber_Input(t *testing.T) {
	var dist = ComputeIncreasesNumber(util.ReadFile("input.txt"))
	if dist != 1342 {
		t.Errorf("Got %d instead of 1342", dist)
	}
}

func TestComputeIncreasesNumberSlidingWindow_Simple(t *testing.T) {
	var dist = ComputeIncreasesNumberSlidingWindow(`199
200
208
210
200
207
240
269
260
263`)
	if dist != 5 {
		t.Errorf("Got %d instead of 5", dist)
	}
}

func TestComputeIncreasesNumberSlidingWindow_Input(t *testing.T) {
	var dist = ComputeIncreasesNumberSlidingWindow(util.ReadFile("input.txt"))
	if dist != 1378 {
		t.Errorf("Got %d instead of 1378", dist)
	}
}
