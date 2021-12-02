package main

import (
	"bufio"
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

	var incCtr, curVal, lastVal int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		curVal, err = strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		if curVal > lastVal && lastVal != 0 {
			incCtr++
		}
		lastVal = curVal
	}
	log.Printf("The depth has increased %d times!", incCtr)
}
