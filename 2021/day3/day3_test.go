package day3

import (
	"migwel/aoc/2021/util"
	"testing"
)

func TestComputePowerConsumption_Simple(t *testing.T) {
	var result = ComputePowerConsumption(`00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`)
	if result != 198 {
		t.Errorf("Got %d instead of 198", result)
	}
}

func TestComputePowerConsumption_Input(t *testing.T) {
	var result = ComputePowerConsumption(util.ReadFile("input.txt"))
	if result != 4191876 {
		t.Errorf("Got %d instead of 4191876", result)
	}
}

func TestComputeLifeSupportRating_Simple(t *testing.T) {
	var result = ComputeLifeSupportRating(`00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`)
	if result != 230 {
		t.Errorf("Got %d instead of 230", result)
	}
}

func TestComputeLifeSupportRating_Input(t *testing.T) {
	var result = ComputeLifeSupportRating(util.ReadFile("input.txt"))
	if result != 3414905 {
		t.Errorf("Got %d instead of 4191876", result)
	}
}
