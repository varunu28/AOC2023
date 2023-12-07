package day06

import "testing"

var sampleTestInput = []string{
	"Time:      7  15   30",
	"Distance:  9  40  200",
}

var actualTestInput = []string{
	"Time:        53     83     72     88",
	"Distance:   333   1635   1289   1532",
}

func TestFindProductOfNumberOfWinningWays(t *testing.T) {
	var tests = []struct {
		input    []string
		expected int
	}{
		{sampleTestInput, 288},
		{actualTestInput, 140220},
	}
	for _, test := range tests {
		if output := FindProductOfNumberOfWinningWays(test.input); output != test.expected {
			t.Errorf("[TestFindProductOfNumberOfWinningWays] Expected(%d) Received(%d)", test.expected, output)
		}
	}
}

func TestFindNumberOfWinningWaysForSingleRace(t *testing.T) {
	var tests = []struct {
		input    []string
		expected int
	}{
		{sampleTestInput, 71503},
		{actualTestInput, 39570185},
	}
	for _, test := range tests {
		if output := FindNumberOfWinningWaysForSingleRace(test.input); output != test.expected {
			t.Errorf("[TestFindNumberOfWinningWaysForSingleRace] Expected(%d) Received(%d)", test.expected, output)
		}
	}
}
