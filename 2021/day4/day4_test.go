package day4

import (
	"migwel/aoc/2021/util"
	"testing"
)

func TestComputeFinalScore_Simple(t *testing.T) {
	var result = ComputeFinalScore(`7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7`)
	if result != 4512 {
		t.Errorf("Got %d instead of 4512", result)
	}
}

func TestComputeFinalScore_Input(t *testing.T) {
	var result = ComputeFinalScore(util.ReadFile("input.txt"))
	if result != 82440 {
		t.Errorf("Got %d instead of 82440", result)
	}
}

func TestComputeFinalScoreLetSquidWin_Simple(t *testing.T) {
	var result = ComputeFinalScoreLetSquidWin(`7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7`)
	if result != 1924 {
		t.Errorf("Got %d instead of 1924", result)
	}
}

func TestComputeFinalScoreLetSquidWin_Input(t *testing.T) {
	var result = ComputeFinalScoreLetSquidWin(util.ReadFile("input.txt"))
	if result != 20774 {
		t.Errorf("Got %d instead of 20774", result)
	}
}
