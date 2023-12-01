package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertToNumber(t *testing.T) {
	tcs := []struct {
		phrase string
		expInt int
		expOK  bool
	}{
		{
			phrase: "one3423",
			expInt: 1,
			expOK:  true,
		},
		{
			phrase: "one",
			expInt: 1,
			expOK:  true,
		},
		{
			phrase: "on",
			expInt: 0,
			expOK:  false,
		},
		{
			phrase: "two",
			expInt: 2,
			expOK:  true,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.phrase, func(t *testing.T) {
			i, ok := convertToNumber(tc.phrase)
			assert.Equal(t, tc.expInt, i)
			assert.Equal(t, tc.expOK, ok)
		})
	}
}

func TestProcessLine(t *testing.T) {
	tcs := []struct {
		phrase string
		exp    int
	}{
		{
			phrase: "two1nine",
			exp:    29,
		},
		{
			phrase: "eightwothree",
			exp:    83,
		},
		{
			phrase: "abcone2threexyz",
			exp:    13,
		},
		{
			phrase: "xtwone3four",
			exp:    24,
		},
		{
			phrase: "4nineeightseven2",
			exp:    42,
		},
		{
			phrase: "zoneight234",
			exp:    14,
		},
		{
			phrase: "7pqrstsixteen",
			exp:    76,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.phrase, func(t *testing.T) {
			assert.Equal(t, tc.exp, processLine(tc.phrase))
		})
	}
}
