package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPossible(t *testing.T) {
	tcs := []struct {
		name        string
		line        string
		limit       []CubePlay
		expPossible bool
	}{
		{
			name:        "1",
			line:        "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			limit:       []CubePlay{{CubeRed, 12}, {CubeGreen, 13}, {CubeBlue, 14}},
			expPossible: true,
		},
		{
			name:        "2",
			line:        "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			limit:       []CubePlay{{CubeRed, 12}, {CubeGreen, 13}, {CubeBlue, 14}},
			expPossible: true,
		},
		{
			name:        "3",
			line:        "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			limit:       []CubePlay{{CubeRed, 12}, {CubeGreen, 13}, {CubeBlue, 14}},
			expPossible: false,
		},
		{
			name:        "4",
			line:        "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			limit:       []CubePlay{{CubeRed, 12}, {CubeGreen, 13}, {CubeBlue, 14}},
			expPossible: false,
		},
		{
			name:        "5",
			line:        "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			limit:       []CubePlay{{CubeRed, 12}, {CubeGreen, 13}, {CubeBlue, 14}},
			expPossible: true,
		},
		{
			name:        "5",
			line:        "Game 38: 13 red, 6 blue, 1 green; 8 red, 4 green, 8 blue; 13 green, 7 red, 3 blue; 6 red, 12 green, 2 blue; 7 blue, 15 green, 5 red; 13 green, 2 blue, 13 red",
			limit:       []CubePlay{{CubeRed, 13}, {CubeGreen, 15}, {CubeBlue, 14}},
			expPossible: true,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			game := parseGame(tc.line)
			assert.Equal(t, tc.expPossible, game.IsPossible(tc.limit))
		})
	}
}

func TestParseCube(t *testing.T) {
	tcs := []struct {
		name string
		line string
		exp  Game
	}{
		{
			name: "1",
			line: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			exp: Game{
				Number: 1,
				Sets: []Set{
					{
						Cubes: []CubePlay{
							{
								Cube:  CubeBlue,
								Count: 3,
							},
							{
								Cube:  CubeRed,
								Count: 4,
							},
						},
					},
					{
						Cubes: []CubePlay{
							{
								Cube:  CubeRed,
								Count: 1,
							},
							{
								Cube:  CubeGreen,
								Count: 2,
							},
							{
								Cube:  CubeBlue,
								Count: 6,
							},
						},
					},
					{
						Cubes: []CubePlay{
							{
								Cube:  CubeGreen,
								Count: 2,
							},
						},
					},
				},
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			game := parseGame(tc.line)
			assert.Equal(t, tc.exp, game)
		})

	}
}
