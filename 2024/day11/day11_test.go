package day11

import (
	"testing"
)

func TestCountStonesAfterBlinking_Simple(t *testing.T) {
	var result = countStones(`0 1 10 99 999`, 1)
	if result != 7 {
		t.Errorf("Got %d instead of 7", result)
	}
}

func TestCountStonesAfterBlinking_Simple2(t *testing.T) {
	var result = countStones(`125 17`, 0)
	if result != 2 {
		t.Errorf("Got %d instead of 2", result)
	}
}

func TestCountStonesAfterBlinking_Simple3(t *testing.T) {
	var result = countStones(`125 17`, 1)
	if result != 3 {
		t.Errorf("Got %d instead of 3", result)
	}
}

func TestCountStonesAfterBlinking_Simple4(t *testing.T) {
	var result = countStones(`125 17`, 6)
	if result != 22 {
		t.Errorf("Got %d instead of 22", result)
	}
}

func TestCountStonesAfterBlinking_Simple5(t *testing.T) {
	var result = countStones(`125 17`, 25)
	if result != 55312 {
		t.Errorf("Got %d instead of 55312", result)
	}
}

func TestCountStonesAfterBlinking_Input(t *testing.T) {
	var result = countStones(`6 11 33023 4134 564 0 8922422 688775`, 25)
	if result != 220999 {
		t.Errorf("Got %d instead of 220999", result)
	}
}

func TestCountStonesAfterBlinking_InputButMoreBlinks(t *testing.T) {
	var result = countStones(`6 11 33023 4134 564 0 8922422 688775`, 75)
	if result != 261936432123724 {
		t.Errorf("Got %d instead of 261936432123724", result)
	}
}
