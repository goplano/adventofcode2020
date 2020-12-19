package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	var inputFilename string
	var yesses map[rune]int

	flag.StringVar(&inputFilename, "file", "groupanswers.txt", "input file name")
	flag.Parse()

	fmt.Println("Reading ", inputFilename)
	fp := openFile(inputFilename)
	s := bufio.NewScanner(fp)
	yesses = make(map[rune]int)
	sumYesses := 0
	sumAllYesses := 0
	numMembers := 0
	for s.Scan() {
		line := s.Text()
		if len(line) == 0 {
			if len(yesses) > 0 {
				countYesses, allYesses := processGroup(yesses, numMembers)
				sumYesses += countYesses
				sumAllYesses += allYesses
			}
			yesses = make(map[rune]int)
			numMembers = 0
		} else {
			numMembers++
			for _, runeValue := range line {
				yesses[runeValue]++
			}
		}
	}

	// do last wrap up
	if len(yesses) > 0 {
		countYesses, allYesses := processGroup(yesses, numMembers)
		sumYesses += countYesses
		sumAllYesses += allYesses

	}
	fmt.Printf("Sum of group yesses  %d All Yesses: %d\n",sumYesses, sumAllYesses)
	err := s.Err()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = fp.Close(); err != nil {
			log.Fatal(err)
		}
	}()
}
func processGroup(yesses map[rune]int, numMembers int) (int,int) {
	fmt.Printf("Processing group data %d\n", len(yesses))
	allYesses := 0
	for _, count := range yesses {
		if count == numMembers {
			allYesses++
		}
	}
	countYesses := len(yesses)
	return countYesses, allYesses
}
func openFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	return file
}
