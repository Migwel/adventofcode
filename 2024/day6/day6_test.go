package day6

import (
	"testing"
)

func TestCountDistinctPosition_Example(t *testing.T) {
	var result = countDistrinctPosition(`....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`)
	if result != 41 {
		t.Errorf("Got %d instead of 41", result)
	}
}

func TestCountDistinctPosition_Input(t *testing.T) {
	var result = countDistrinctPosition(`......#.......#...................................................................................##......#..............#........
........................#....................................#....................................................................
...#.........#...#.....##..#...#......#..........................................#................#...................#....#....#.
....................#...#...#.#............#..................................#..............#....#...............................
.....................................................................#................................#...........................
........#............#......#................#..................#................................................#...............#
...................#.........#............................................................##......................................
.........#..........#......##.............#.#...............#................................##.................#.................
...........................................................#...............................#..................##......##.....#....
...................#...........................................##...............#.....#......#............#..........#............
...............#........................#...#....#..............................................#.#......#............#...........
........#.............#...........................#..#.....................................#......................................
......#......#....#..............#...............#...#...............#.......#...................................................#
..........#............#..#.....#......................#.........#...........................#..........#.........................
#...............................................................................#.#....#.............###..................#.......
.....#.#....#....................................#............................#...............#........................#....#.....
....#......#................#....................#...................#............#................#..............................
#...................#.............................#.#...#.............#......................#.....................#..............
.................#...............#................................#..##..#..............#.......#.................................
.....................................................##.#.....................................#.............#.....................
...#.............................#................................................................................................
#.##.........................................#....#..#..............................#......#..................................#...
......................#.........##.....................................#......................#..#..........#...............#....#
.#......................................................#....#......................................#...............#....#........
............#...............................#.......................................................#.............................
.............................................................................#....................................................
..#...#.....#...............................................................................................#.....................
................................................#.#.......................#..................#............#.......................
...............#.............#.....................#........................................................#..........#.......#..
...................................#..........#......#............................................#.....#.........................
.............#..............#...........##..........................................................................#.....#.......
.............#...............#........#.................................................................#.......#.................
.#...............................................................#..............................................................#.
.#......#...........#..............................................................##.............................................
...................#.....................................#....#........................#............................#.......###...
.....#...#.....#......................#.....................................#...............................#.....................
......................#....#...#........#........#...........................#...............#...........................#....#...
..#..#................................................................................#...........................................
.....#......#.#.......................................#..........#................................................................
..........#.......................#.............#.#........##.........................................................#...........
##........##.#....................#...#...................................................#.##......#...................#.........
...........#................#.#...........#......................#..................#......#......................................
...#..............................#......#...............#.................#.#............................................#.....#.
...#..................#...............#.........#..#.............................#..........#...................#.......#.........
.....................................................................................................#.......................#....
....................................##.#.................................#......##................................#.#.............
..................#.............................................................................................#............#....
............#........#...#.....#.....#........................................#....#..............#..............................#
....................................#...........................................................................................#.
............................#.#....#.....................#........................................................................
........#....#......#.......#................................................................##............................#......
....#...................#....#.............................................................................................#......
....................#......#.........................#..................#.........#....................#..........................
#........#.....#.......#....................................#.....#...........................................#...................
.....#....................................................................................................#.#....#................
....#....................................................................#.......#........................................#.......
#...........................#.#....#......#................................................................................#......
.#..........................................................#...............#......................#.............................#
..#..#.#...................................#.......#...............#....#....................#................................#...
.........................#........................................................................................................
...................................#...........................#...........#..#...........#................#.........#............
......................................#...............#.................#..#............................#.....#..#...#............
#.......................#.........#..#...............................................................#.......................#....
............#............................#..#.....................................................................................
................#............#..................#....#........#...............................#....#.#.......................#.#..
.........................................................................#....................#..........................#........
.....................................................#...............#........................................#......#............
..........##..................##...............................................................................................#..
..................#.........................................#..#......................#.........#.....#.....#.....................
..#.........................................................................#...................................#................#
...............................#..................................................................................................
..#..........#..#..............................................................................#........................#.........
...........#.#........................................................#.....#.........................#...........................
...........................................#......#...............................................................................
.......................#................................................................................................##........
.....#...............#....................................................................................#.......#...............
.#.#.................................#........................#................#...............#......#...........................
........#........#...............#.........................^...........................#......#.................................#.
.....................#.............#..#...........................................................................................
...........#...........................................................................#.......##......#.............#............
..........#..........................................#.................#....#......##.........#.................................#.
...................#...#...................#....#...........#................................#.#..................................
.......#....#...#..................#...............................................#.....#.........#......#.#.#.......#......#....
.....#................................................#......................#...#.....................................#.#........
.##........##.......#.........................................##..................................................................
......................................#....#........#..........#..#...........................#....................#....#.........
.......................#........................#..............................#.#.............................................#..
..........................#....###......................................................................#...................#.....
...........#..................#..............#....................................................................................
......................#................#..#............................#............#.........................##...#..............
.................#.......................................................................#....#...............................#...
................................................#..............................................................................#..
..#.......#....#................#...........................#..........#.#..#.......................#.#....................#......
........#.............................#........................................#..................................................
...........................#............................................................................#.............#....#......
#...................#.......#............#..........................................................#..#....#......#..........#..#
#..........#.................................#.........................#..................................#....................##.
..........#..#...........................................................................#........................##..........#...
.............#.#.......................................#.................................#..............#.......#.................
.................#....#........#........#....#.................................#................................#.................
...........................#..........#.#..#..#........#...#....................#...#........#.......................#............
........#................................#...................................#.#...............#..................................
............................................................#.......#......#.#..........................#..#....................#.
.........#.#....................................#.....................#.....##.......#..................................#.........
........#......................................#.......##............................................................##.....#.....
.............#................................#....#..................##...........................#..#..#........................
.......................................................#...#..........#...................................#.......#........#......
.....#...........#...............................................................................................#................
.......#..................#......#.....#............#..................#.......#.............#................#...................
.....#.........................................................#...........................................#.............#.....#..
...............#..................#...............................................#...................#...........................
..#............#......#....#...#..............................#..#...................................#...#.............#..........
.........#........#.....................................##.......................#.....................#..#.............#.........
....................#....#......................#..#.##..................#................#...........#............#..............
..........#..................................#............#..#...............#.................#..........#.......................
..................#.........#........................#.........................#.........#..................................#.....
...........#.#..............#........#.................................#............#............................#................
.#.....#........................................................................................#....#........##..................
.....................#....#.................#..........#............#............................#.............#.............#.#..
................................#....#......................................................#....#............#.............#.....
..............#..#.........#....................#.......................#.......#.........#..................#....................
.............#.....#....................#.#.....................#................#...#.........................#........##........
.......#.......................#.........................................#.........................##.....#...........#...#.....#.
.....#.................#................................................................#.................................#....#..
................#.......................#...........#........#................#.......#.#.#........................#........#.....
......#..........................#......#............#.......#...#.............................#..........##.......#..........#...
................#.....#...#..............................#......#..#..............................................................
.............................#....#...............#..............#......................................#.......###..#.#..........
..........#.......#...............#............................#........#.........................................................
.................#.....................................#.......................................#................#........#.#..#...`)
	if result != 5404 {
		t.Errorf("Got %d instead of 5404", result)
	}
}

func TestFindLoopyPositions_Example(t *testing.T) {
	var result = findLoopyPositions(`....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`)
	if result != 6 {
		t.Errorf("Got %d instead of 6", result)
	}
}

func TestFindLoopyPositions_Input(t *testing.T) {
	var result = findLoopyPositions(`......#.......#...................................................................................##......#..............#........
........................#....................................#....................................................................
...#.........#...#.....##..#...#......#..........................................#................#...................#....#....#.
....................#...#...#.#............#..................................#..............#....#...............................
.....................................................................#................................#...........................
........#............#......#................#..................#................................................#...............#
...................#.........#............................................................##......................................
.........#..........#......##.............#.#...............#................................##.................#.................
...........................................................#...............................#..................##......##.....#....
...................#...........................................##...............#.....#......#............#..........#............
...............#........................#...#....#..............................................#.#......#............#...........
........#.............#...........................#..#.....................................#......................................
......#......#....#..............#...............#...#...............#.......#...................................................#
..........#............#..#.....#......................#.........#...........................#..........#.........................
#...............................................................................#.#....#.............###..................#.......
.....#.#....#....................................#............................#...............#........................#....#.....
....#......#................#....................#...................#............#................#..............................
#...................#.............................#.#...#.............#......................#.....................#..............
.................#...............#................................#..##..#..............#.......#.................................
.....................................................##.#.....................................#.............#.....................
...#.............................#................................................................................................
#.##.........................................#....#..#..............................#......#..................................#...
......................#.........##.....................................#......................#..#..........#...............#....#
.#......................................................#....#......................................#...............#....#........
............#...............................#.......................................................#.............................
.............................................................................#....................................................
..#...#.....#...............................................................................................#.....................
................................................#.#.......................#..................#............#.......................
...............#.............#.....................#........................................................#..........#.......#..
...................................#..........#......#............................................#.....#.........................
.............#..............#...........##..........................................................................#.....#.......
.............#...............#........#.................................................................#.......#.................
.#...............................................................#..............................................................#.
.#......#...........#..............................................................##.............................................
...................#.....................................#....#........................#............................#.......###...
.....#...#.....#......................#.....................................#...............................#.....................
......................#....#...#........#........#...........................#...............#...........................#....#...
..#..#................................................................................#...........................................
.....#......#.#.......................................#..........#................................................................
..........#.......................#.............#.#........##.........................................................#...........
##........##.#....................#...#...................................................#.##......#...................#.........
...........#................#.#...........#......................#..................#......#......................................
...#..............................#......#...............#.................#.#............................................#.....#.
...#..................#...............#.........#..#.............................#..........#...................#.......#.........
.....................................................................................................#.......................#....
....................................##.#.................................#......##................................#.#.............
..................#.............................................................................................#............#....
............#........#...#.....#.....#........................................#....#..............#..............................#
....................................#...........................................................................................#.
............................#.#....#.....................#........................................................................
........#....#......#.......#................................................................##............................#......
....#...................#....#.............................................................................................#......
....................#......#.........................#..................#.........#....................#..........................
#........#.....#.......#....................................#.....#...........................................#...................
.....#....................................................................................................#.#....#................
....#....................................................................#.......#........................................#.......
#...........................#.#....#......#................................................................................#......
.#..........................................................#...............#......................#.............................#
..#..#.#...................................#.......#...............#....#....................#................................#...
.........................#........................................................................................................
...................................#...........................#...........#..#...........#................#.........#............
......................................#...............#.................#..#............................#.....#..#...#............
#.......................#.........#..#...............................................................#.......................#....
............#............................#..#.....................................................................................
................#............#..................#....#........#...............................#....#.#.......................#.#..
.........................................................................#....................#..........................#........
.....................................................#...............#........................................#......#............
..........##..................##...............................................................................................#..
..................#.........................................#..#......................#.........#.....#.....#.....................
..#.........................................................................#...................................#................#
...............................#..................................................................................................
..#..........#..#..............................................................................#........................#.........
...........#.#........................................................#.....#.........................#...........................
...........................................#......#...............................................................................
.......................#................................................................................................##........
.....#...............#....................................................................................#.......#...............
.#.#.................................#........................#................#...............#......#...........................
........#........#...............#.........................^...........................#......#.................................#.
.....................#.............#..#...........................................................................................
...........#...........................................................................#.......##......#.............#............
..........#..........................................#.................#....#......##.........#.................................#.
...................#...#...................#....#...........#................................#.#..................................
.......#....#...#..................#...............................................#.....#.........#......#.#.#.......#......#....
.....#................................................#......................#...#.....................................#.#........
.##........##.......#.........................................##..................................................................
......................................#....#........#..........#..#...........................#....................#....#.........
.......................#........................#..............................#.#.............................................#..
..........................#....###......................................................................#...................#.....
...........#..................#..............#....................................................................................
......................#................#..#............................#............#.........................##...#..............
.................#.......................................................................#....#...............................#...
................................................#..............................................................................#..
..#.......#....#................#...........................#..........#.#..#.......................#.#....................#......
........#.............................#........................................#..................................................
...........................#............................................................................#.............#....#......
#...................#.......#............#..........................................................#..#....#......#..........#..#
#..........#.................................#.........................#..................................#....................##.
..........#..#...........................................................................#........................##..........#...
.............#.#.......................................#.................................#..............#.......#.................
.................#....#........#........#....#.................................#................................#.................
...........................#..........#.#..#..#........#...#....................#...#........#.......................#............
........#................................#...................................#.#...............#..................................
............................................................#.......#......#.#..........................#..#....................#.
.........#.#....................................#.....................#.....##.......#..................................#.........
........#......................................#.......##............................................................##.....#.....
.............#................................#....#..................##...........................#..#..#........................
.......................................................#...#..........#...................................#.......#........#......
.....#...........#...............................................................................................#................
.......#..................#......#.....#............#..................#.......#.............#................#...................
.....#.........................................................#...........................................#.............#.....#..
...............#..................#...............................................#...................#...........................
..#............#......#....#...#..............................#..#...................................#...#.............#..........
.........#........#.....................................##.......................#.....................#..#.............#.........
....................#....#......................#..#.##..................#................#...........#............#..............
..........#..................................#............#..#...............#.................#..........#.......................
..................#.........#........................#.........................#.........#..................................#.....
...........#.#..............#........#.................................#............#............................#................
.#.....#........................................................................................#....#........##..................
.....................#....#.................#..........#............#............................#.............#.............#.#..
................................#....#......................................................#....#............#.............#.....
..............#..#.........#....................#.......................#.......#.........#..................#....................
.............#.....#....................#.#.....................#................#...#.........................#........##........
.......#.......................#.........................................#.........................##.....#...........#...#.....#.
.....#.................#................................................................#.................................#....#..
................#.......................#...........#........#................#.......#.#.#........................#........#.....
......#..........................#......#............#.......#...#.............................#..........##.......#..........#...
................#.....#...#..............................#......#..#..............................................................
.............................#....#...............#..............#......................................#.......###..#.#..........
..........#.......#...............#............................#........#.........................................................
.................#.....................................#.......................................#................#........#.#..#...`)
	if result != 1984 {
		t.Errorf("Got %d instead of 1984", result)
	}
}
