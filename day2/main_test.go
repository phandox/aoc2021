package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecuteSubmarineCoursePlan(t *testing.T) {
	tests := []struct {
		name string
		plan []string
		want int
	}{
		{
			"example input",
			[]string{
				"forward 5",
				"down 5",
				"forward 8",
				"up 3",
				"down 8",
				"forward 2",
			},
			900,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := Submarine{}
			r := strings.NewReader(strings.Join(test.plan, "\n"))
			s.Navigate(loadCourse(r))
			got := s.PositionMultiplication()
			assert.Equal(t, test.want, got)
		})
	}
}
