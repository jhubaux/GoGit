package main

import (
	"testing"
)

type testC struct {
	Input  string
	Output int
}

func TestConvert(t *testing.T) {
	conversions := []testC{{"r", ROCK},
		{"R", ROCK},
		{"s", SCISSORS},
		{"S", SCISSORS},
		{"p", PAPER},
		{"P", PAPER}}
	for _, g := range conversions {
		if Convert(g.Input) != g.Output {
			t.Error("The input", g.Input, "should be a", g.Output)
		}
	}
}

type testW struct {
	A int
	B int
	Win1    int
}

func TestWin(t *testing.T) {
	winners := []testW{{ROCK, ROCK, 0},
		{SCISSORS, SCISSORS, 0},
		{PAPER, PAPER, 0},
		{ROCK, SCISSORS, 1},
		{PAPER, SCISSORS, -1},
		{ROCK, PAPER, -1},
	}
	for _, g := range winners {
		if Wins(g.A, g.B) != g.Win1 {
			t.Error("Player1:", g.A, "Player2:", g.B, "should win",
				g.Win1)
		}
		if Wins(g.B, g.A) != g.Win1 * -1 {
			t.Error("Player1:", g.B, "Player2:", g.A, "should win",
				g.Win1)
		}
	}
}