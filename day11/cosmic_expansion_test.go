package day11

import "testing"

var sampleTestInput = []string{
	"...#......",
	".......#..",
	"#.........",
	"..........",
	"......#...",
	".#........",
	".........#",
	"..........",
	".......#..",
	"#...#.....",
}

var actualTestInput = []string{
	".#...........#.........................................................................#...............#...........#........................",
	".......................#....................................................................................................................",
	"...............................#...........#.................#.............................#.............................................#..",
	"......#............................................#........................................................................................",
	"......................................................................#.....#.....................#.........................................",
	".................#.......#.......................................#................#.....#.....................#.............................",
	".........................................#.................#................................................................................",
	"................................#.............................................................#..........#..................................",
	"#.......................................................................#..............................................#.................#..",
	"..........#....................................#......................................#............................................#........",
	"................#....................................#........#.............................................................................",
	"......#........................................................................#............................................................",
	"......................#..............#..............................................................#..........#.....#......................",
	"..................................................#.......................#..................................................#..........#...",
	"............................#...............................................................................................................",
	"........#...................................................................................................................................",
	"..................#.......................#................................................................#................................",
	".............................................................#........#.....................................................................",
	"..............................#.......#......................................#.............#........................#.......................",
	".#............................................................................................................#...........................#.",
	"...................................................................................................#....................#...................",
	"....................#......................#.............#.....#..................................................................#.........",
	".................................................................................................................#..........................",
	"............................................................................................................................................",
	"....................................................................#.....................#.................................................",
	".................................#..............................................#..............#............#.............#.................",
	"...........................#............................................#...........................................#.......................",
	"......................#.........................#...........................................................................................",
	".......................................#....................................#...................................#...........................",
	".....#..........#.........................................#...............................................................................#.",
	"#..........#......................#...................................................#.......#.....#.......................................",
	".........................#..................................................................................................................",
	"....................#..............................#.............#.........................................#.......#........#...............",
	"...........................................#........................................................................................#.......",
	"...........................................................#....................#...........................................................",
	"..........#....................................................................................................#............................",
	"...............#................................#..........................#..............................................................#.",
	"..#.................................#.................................................................#.......................#.............",
	".....................#......#...............#........................#................................................#.....................",
	"..............................................................................................#....................................#........",
	"................................#.......#...................................................................................................",
	"..........#................................................#...................#.....#..............#..........#..........#.................",
	".........................#..................................................................................................................",
	"..................#...............................#...........................................................................#.............",
	".........................................................................#............................................#.....................",
	"........#.....#..............................................................................#.........#..................................#.",
	".#...............................#......#............#.............................................................................#........",
	"....................................................................................#.............#...............#........#................",
	"...........................#................................................................................................................",
	"......................#..................................#..............#................................#.............#................#...",
	"...............................................................#..............#.............................................................",
	"..................................#........................................................#........#.......................................",
	"......#.................................#.......#....................#...............................................................#......",
	".................#............#..............................................................................#.............................#",
	"............................................................................................................................................",
	"..#.....................#.............................................................................#.....................................",
	".......................................................................................#..............................#.....................",
	"...............................................................................#.............#..............................................",
	"....................#.....................#...............................................................#.................................",
	"..........#........................#............#...................#.............................................#...............#.........",
	"............................#.............................#..........................#......................................................",
	"...........................................................................................#................................................",
	".................................................................#.................................#.........................#........#.....",
	".................#........................................................#..................................#..............................",
	"......................................#......................#..............................................................................",
	".............#...........#...............................................................................#..........#.......................",
	".................................#...........#...........................................................................#..............#...",
	".......#............................................#.....#....................#..............#.....#.......................................",
	"...................#...........................................#................................................#...........................",
	"........................................................................................#...................................................",
	"...#....................#.....#............#.........................................................................................#......",
	"...............#.............................................................#..................#........#................#.................",
	".........#.....................................#............................................................................................",
	"................................................................................................................................#...........",
	".....................#.............#.........................................................................#..............................",
	".....#...............................................................................................................#..............#......#",
	"...........#.......................................#.....#.................#.................#..............................................",
	"............................................................................................................................................",
	".#.......................................#..........................................#................#......................................",
	"......................#.............#.................#.......................#...........#.............................#...................",
	"...............#.............#................................#......#.......................................................#..............",
	"............................................................................................................................................",
	"...#............................................................................................#....................#......................",
	"..........#.....................#.............................................................................#.............................",
	"...........................#...............................................................................................#...............#",
	".....................#......................................#.......#......#........#..................#...........................#........",
	".....#...............................#............#.........................................................................................",
	"............................................................................................................................................",
	".........................................................#........................................................#.........................",
	"#..............................#..........#..............................................................................#................#.",
	"....................#..............................................#.......................#.........#......................................",
	"............................................................................................................................................",
	"...........#.......................................#.............................................................................#..........",
	"...#......................#...............................................#......................#..........................................",
	"..................#.........................................................................................................................",
	"...........................................................#........#...........#.....#......#........................#.................#...",
	"........#...................................................................................................................................",
	"....................................................................................................#......#................................",
	"................................#...........#....................#.....#...................................................#........#.......",
	"........................#............#......................................#.............#.................................................",
	"...................#............................................................................................#...........................",
	"#................................................................................................................................#..........",
	"..............#.............................................#...............................................................................",
	"..............................#...............................................#.................#......................................#....",
	"....#....................#......................................#...................#...............................#.......#...............",
	"........................................#..........#........................................#................#..............................",
	"............................................................................................................................................",
	"............................................................................................................................................",
	"........................................................................................................................................#...",
	"...........................................................#..................#.................#.................#.........................",
	".......#........................................................#................................................................#..........",
	"..........................#......#......#..................................................................................#................",
	"...............................................#.....#.................#....................................................................",
	"......................................................................................................#.....................................",
	"................#....................#.....................................................................#........#.......................",
	"..............................................................................................................................#.............",
	"...........................................#.............................................................................................#..",
	"....#........................#..............................#.............#................#................................................",
	"......................#..........................................................................................#......#...................",
	"....................................................#................................#.........#............................................",
	"..........................#.......................................#.........................................................................",
	".................#........................#.....#........................................................#..................................",
	".#.........................................................#............#....................................................#..............",
	"..........#..................#.....#..................................................................................................#.....",
	".................................................................................#..........................................................",
	".....................................................#....................................#.........#...........#.....#.....................",
	".............................................#.....................#.......................................#.......................#........",
	"................................#...........................................................................................................",
	"......#........................................................#..............................#.............................................",
	"............#......................................#.....#............................#........................................#............",
	"............................................................................#...............................................................",
	".......................................#...............................#...................#..............................#...........#.....",
	".......................#..............................#..........................#..............................#...........................",
	".................................................................#...................................#......................................",
	"......................................................................................................................#...................#.",
	".................#..........................................................................................................................",
	".........#......................#.....#...........#.......................................................#.................................",
	"...........................................#.................#................#.....................................................#.......",
	"..#...............................................................#.........................................................................",
	".....................#......#......#.............................................................#............#...........#.................",
}

func TestCalculateSumOfDistances(t *testing.T) {
	var tests = []struct {
		input    []string
		expected int
	}{
		{sampleTestInput, 374},
		{actualTestInput, 9686930},
	}
	for _, test := range tests {
		if output := CalculateSumOfDistances(test.input); output != test.expected {
			t.Errorf("[TestCalculateSumOfDistances] Expected(%d) Received(%d)", test.expected, output)
		}
	}
}

func TestCalculateSumOfDistancesForOlderGalaxies(t *testing.T) {
	var tests = []struct {
		input    []string
		expected int
	}{
		{sampleTestInput, 82000210},
		{actualTestInput, 630728425490},
	}
	for _, test := range tests {
		if output := CalculateSumOfDistancesForOlderGalaxies(test.input); output != test.expected {
			t.Errorf("[TestCalculateSumOfDistancesForOlderGalaxies] Expected(%d) Received(%d)", test.expected, output)
		}
	}
}