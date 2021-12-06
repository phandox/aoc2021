package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func loadCourse(plan io.Reader) []Command {
	var course []Command
	scn := bufio.NewScanner(plan)
	for scn.Scan() {
		c := strings.Split(scn.Text(), " ")
		if len(c) != 2 {
			log.Fatalf("incorrect command: %q", c)
		}
		v, err := strconv.Atoi(c[1])
		if err != nil {
			log.Fatal(err)
		}
		course = append(course, Command{
			cmd: c[0],
			val: v,
		})
	}
	return course
}

func main() {
	f, err := os.Open("day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	plan := loadCourse(f)
	s := Submarine{}
	s.Navigate(plan)
	fmt.Println(s.PositionMultiplication())
}
