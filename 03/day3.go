package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	//var hill []int32
	var mask uint
	numTrees := 0
	numLines := 0

	downVal := flag.Int("down", 1, "an int")
	overVal := flag.Uint("over", 3, "an int")
	flag.Parse()

	var pos uint = 0
	fp := openFile("trees.txt")
	s := bufio.NewScanner(fp)
	for s.Scan() {
		mod := numLines % *downVal
		fmt.Printf("Over: %v Down %v Line: %v Mod: %v Use: %v \n", *overVal, *downVal,  numLines, mod,   mod == 0)
		line := s.Text()

		numLines++
		if mod == 0 {
			digits, length := convertToNumber(line)
			if (length - (pos % 31) - 1) <= 0 {
				mask = 1
			} else {
				shiftValue := (length - (pos % 31) - 1)
				mask = shift(1, shiftValue)
			}
			fmt.Printf("%0.31b %0.31b %s %d %d %d\n", digits, mask, line, length, pos, length-pos)
			if digits&mask != 0 {
				fmt.Println("TREE")
				numTrees++
			}
			pos += *overVal
		}
	}
	fmt.Printf("Lines: %d Trees: %d\n", numLines, numTrees)
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

func shift(num uint, times uint) uint {
	return num << times
}
func convertToNumber(line string) (uint, uint) {
	var digits uint = 0
	var length, i uint
	length = uint(len(line))
	for i = 0; i < length; i++ {
		digits = digits << 1
		if line[i] == '#' {
			digits = digits | 1
		}

	}
	return digits, length
}

func openFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

/*
00100000010011110000010010000000
..#......#..####.....#..#.......
*/
