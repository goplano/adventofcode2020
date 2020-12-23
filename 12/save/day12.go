package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	var inputFilename string
	var direction string = "E"
	var east int = 0
	var north int = 0

	flag.StringVar(&inputFilename, "file", "nav.txt", "input file name")
	flag.Parse()

	fmt.Printf("\neast: %d north: %d", Abs(east), Abs(north))
	fp := openFile(inputFilename)
	s := bufio.NewScanner(fp)
	for s.Scan() {

		line := s.Text()
		inst, val := parseInput(line)
		direction = newDirection(direction, inst, val)
		fmt.Printf("\t%s\t%3d\t%s\n", inst, val, direction)

		switch inst {
		case "N":
			north += val
		case "S":
			north -= val
		case "W":
			east -= val
		case "E":
			east += val
		}
		if inst == "F" {
			switch direction {
			case "E":
				east += val
			case "S":
				north -= val
			case "W":
				east -= val
			case "N":
				north += val
			}
		}
		label1 := "east"
		label2 := "north"
		if east < 0 {
			label1 = "west"
		}
		if north < 0 {
			label2 = "south"
		}
		fmt.Printf("\t%s:\t%d\t%s:\t%d", label1, Abs(east), label2, Abs(north))
	}
	fmt.Printf("\nManhattan distance; %d + %d = %d\n", Abs(east), Abs(north), Abs(east - north))

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
func newDirection(direction string, inst string, val int) string {
	if inst == "L" || inst == "R" {
		if inst == "L" {
			val = -val
		}
		num := 0
		switch direction {
		case "E":
			num = 0
		case "S":
			num = 90
		case "W":
			num = 180
		case "N":
			num = 270
		}
		num = (num + val) % 360
		if num < 0 {
			num += 360
		}
		switch num {
		case 0:
			direction = "E"
		case 90:
			direction = "S"
		case 180:
			direction = "W"
		case 270:
			direction = "N"
		}
	}
	return direction
}
func openFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	return file
}
func parseInput(line string) (string, int) {
	patStr := "^([NSEWLRF]{1})(.*)$"
	re := regexp.MustCompile(patStr)
	result := re.FindStringSubmatch(line)
	val, _ := strconv.ParseInt(result[2], 10, 32)
	return result[1], int(val)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
