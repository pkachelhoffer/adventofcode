package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProcessLine(t *testing.T) {
	tcs := []struct {
		name   string
		lines  []string
		expSum int
	}{
		{
			name: "1",
			lines: []string{
				"467..114..",
				"...*......",
				"..35..633.",
				"......#...",
				"617*......",
				".....+.58.",
				"..592.....",
				"......755.",
				"...$.*....",
				".664.598..",
			},
			expSum: 467835,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			sum := ProcessInput(tc.lines)

			assert.Equal(t, tc.expSum, sum)
		})
	}
}

func TestGetNextNumber(t *testing.T) {
	tcs := []struct {
		name                              string
		line                              string
		startIdx                          int
		expStartIdx, expEndIdx, expNumber int
		expOK                             bool
	}{
		{
			name:        "1",
			line:        "467..114..",
			startIdx:    0,
			expStartIdx: 0,
			expEndIdx:   2,
			expNumber:   467,
			expOK:       true,
		},
		{
			name:        "1",
			line:        "467..114..",
			startIdx:    3,
			expStartIdx: 5,
			expEndIdx:   7,
			expNumber:   114,
			expOK:       true,
		},
		{
			name:     "1",
			line:     "467..114..",
			startIdx: 8,
			expOK:    false,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			startIdx, endIdx, number, ok := GetNextNumber(tc.line, tc.startIdx)
			assert.Equal(t, tc.expStartIdx, startIdx)
			assert.Equal(t, tc.expEndIdx, endIdx)
			assert.Equal(t, tc.expNumber, number)
			assert.Equal(t, tc.expOK, ok)
		})
	}
}
