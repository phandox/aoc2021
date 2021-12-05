package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	log.Println("Loading values from sonar sweep...")
	input := "day1/input.txt"
	f, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	measurements, err := loadSonarValues(f)
	if err != nil {
		log.Fatal(err)
	}
	incrCtr := CalcSeaLevelIncr(createWindows(measurements, 3))
	log.Printf("The depth has increased %d times!", incrCtr)
}

func loadSonarValues(r io.Reader) ([]int, error) {
	var vals []int
	s := bufio.NewScanner(r)
	for s.Scan() {
		v, err := strconv.Atoi(s.Text())
		if err != nil {
			return nil, err
		}
		vals = append(vals, v)
	}
	return vals, nil
}

type slideWindow struct {
	data []int
}

func (w slideWindow) Sum() int {
	var s int
	for _, v := range w.data {
		s += v
	}
	return s
}

func createWindows(input []int, winSize int) []slideWindow {
	res := make([]slideWindow, len(input)-winSize+1)
	for i := 0; i+winSize <= len(input); i++ {
		res[i] = slideWindow{data: input[i : i+winSize]}
	}
	return res
}

func CalcSeaLevelIncr(readings []slideWindow) int {
	var incrCount, lastVal int
	for _, r := range readings {
		currVal := r.Sum()
		if currVal > lastVal && lastVal != 0 {
			incrCount++
		}
		lastVal = currVal
	}
	return incrCount
}
