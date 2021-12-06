package main

import (
	"log"
)

type Position struct {
	depth      int
	horizontal int
}

type Submarine struct {
	Position
	aim int
}

func (s *Submarine) PositionMultiplication() int {
	return s.depth * s.horizontal
}

func (s *Submarine) up(i int) {
	s.aim -= i
}

func (s *Submarine) down(i int) {
	s.aim += i
}

func (s *Submarine) forward(i int) {
	s.horizontal += i
	s.depth += i * s.aim
}

func (s *Submarine) Navigate(course []Command) {
	for _, cmd := range course {
		switch cmd.cmd {
		case "up":
			s.up(cmd.val)
		case "down":
			s.down(cmd.val)
		case "forward":
			s.forward(cmd.val)
		default:
			log.Fatalf("unknown command: %q", cmd.cmd)
		}
	}
}

type Command struct {
	cmd string
	val int
}
