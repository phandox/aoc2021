package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlidingWindowConstruction(t *testing.T) {
	tests := []struct {
		name    string
		in      []int
		winSize int
		want    []slideWindow
	}{
		{
			"input not padded",
			[]int{
				199, 200, 208, 210, 200, 207, 240, 269, 260, 263,
			},
			3,
			[]slideWindow{
				{data: []int{199, 200, 208}},
				{data: []int{200, 208, 210}},
				{data: []int{208, 210, 200}},
				{data: []int{210, 200, 207}},
				{data: []int{200, 207, 240}},
				{data: []int{207, 240, 269}},
				{data: []int{240, 269, 260}},
				{data: []int{269, 260, 263}},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := createWindows(test.in, test.winSize)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestCountingIncrements(t *testing.T) {
	tests := []struct {
		name     string
		readings []slideWindow
		want     int
	}{
		{
			"example input",
			[]slideWindow{
				{data: []int{199, 200, 208}},
				{data: []int{200, 208, 210}},
				{data: []int{208, 210, 200}},
				{data: []int{210, 200, 207}},
				{data: []int{200, 207, 240}},
				{data: []int{207, 240, 269}},
				{data: []int{240, 269, 260}},
				{data: []int{269, 260, 263}},
			},
			5,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := CalcSeaLevelIncr(test.readings)
			assert.Equal(t, test.want, got)
		})
	}
}
