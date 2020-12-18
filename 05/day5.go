package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
)

type BoardingPass struct {
	code string
	row  int
	seat int
}
type ByID []BoardingPass

func (a ByID) Len() int           { return len(a) }
func (a ByID) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByID) Less(i, j int) bool { return a[i].id() < a[j].id() }
func (m BoardingPass) id() int {
	return m.row*8 + m.seat
}
func main() {
	var inputFilename string
	var emptyId int
	numRows := 128
	numSeats := 8

	numRowCodes := int(math.Log2(float64(numRows)))
	numColCodes := int(math.Log2(float64(numSeats)))

	//var passes ByID
	lookup := make(map[int]int)

	maxID := 0

	flag.StringVar(&inputFilename, "file", "boardingpasses.txt", "input file name")
	flag.Parse()

	fmt.Println("Reading ", inputFilename)
	fp := openFile(inputFilename)
	s := bufio.NewScanner(fp)

	//str := []byte("FBFBBFFRLR")
	for s.Scan() {
		str := []byte(s.Text())
		row := findPosition(&str, 0, numRows-1, 0, numRowCodes-1)
		col := findPosition(&str, 0, numSeats-1, numRowCodes, numRowCodes+numColCodes-1)
		pass := BoardingPass{string(str), row, col}
		id := pass.id()
		//passes = append(passes,pass)
		lookup[id] = id
		fmt.Printf("Pass: %s %d %d ID: %d\n", pass.code, pass.row, pass.seat, id)
		if id > maxID {
			maxID = id
		}
	}
	fmt.Printf("MAX: %d\n", maxID)

	//sort.Sort(ByID(passes))
	//fmt.Printf("Items: (%+v)\n", lookup)
	emptyId = findGap(lookup)
	fmt.Printf("Next Available seat: %d\n", emptyId)
}

func findGap(lookup map[int]int) int {
	for key, _ := range lookup {
		gapIdx := key + 1
		nextIdx := key + 2
		_, ok := lookup[gapIdx]
		if !ok {
			_, ok := lookup[nextIdx]
			if ok {
				return gapIdx
			}
		}
	}
	return 0
}
func findPosition(direction *[]byte, low int, high int, index int, maxIndex int) int {
	//fmt.Printf("find: %s low: %d high: %d index: %d max: %d %c\n", *direction, low, high, index, maxIndex, (*direction)[index])
	if index < maxIndex {
		mid := (high-low)/2 + low + 1

		if (*direction)[index] == 'F' || (*direction)[index] == 'L' {
			return findPosition(direction, low, mid-1, index+1, maxIndex)

		} else {
			return findPosition(direction, mid, high, index+1, maxIndex)

		}
	} else
	if (*direction)[index] == 'F' || (*direction)[index] == 'L' {
		return low

	} else {
		return high

	}
}
func openFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	return file
}
