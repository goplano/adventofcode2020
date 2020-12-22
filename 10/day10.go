package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type adapter struct {
	jolts int64
	paths int
}

func (a *adapter) canConnect(jolts int64) bool {
	difference := Abs(a.jolts - jolts)
	return 1 <= difference && difference <= 3
}

type adapters []adapter

func (a adapters) Len() int           { return len(a) }
func (a adapters) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a adapters) Less(i, j int) bool { return a[i].jolts < a[j].jolts }
func (a adapters) findCompatible(jolts int64, start int) []int {
	var keys []int
	keys = make([]int, 0)
	count := len(a)
	for i := start; i < count; i++ {
		val := a[i]
		if val.jolts == 0 {
			break
		}
		if val.canConnect(jolts) {
			keys = append(keys, i)
		}
		if val.jolts > jolts+3 {
			break
		}
	}
	return keys
}

func (a adapters) getRouteCount(startJolt int64, endJolt int64, i int) int {
	if i >= 0 && a[i].paths > 0 {
		//fmt.Println("cache")
		return a[i].paths
	}
	if startJolt+3 == endJolt {
		return 1
	}
	count := 0
	keys := a.findCompatible(startJolt, i+1)

	for _, j := range keys {
		result := a.getRouteCount(a[j].jolts, endJolt, j)
		count += result
	}
	if i >= 0 {
		a[i].paths = count
	}
	return count
}

func (a adapters) printMatches(keys []int) {
	for _, i := range keys {
		fmt.Printf("[%d]: %d ", i, a[i].jolts)
	}
	fmt.Printf("\n")
}

func main() {
	var inputFilename string

	flag.StringVar(&inputFilename, "file", "device_jolts.txt", "input file name")
	flag.Parse()

	bag := make(adapters, 0)

	fp := openFile(inputFilename)
	s := bufio.NewScanner(fp)
	for s.Scan() {
		line := s.Text()
		value, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		bag = append(bag, adapter{
			jolts: value,
		})
	}
	sort.Sort(adapters(bag))

	ones := 0
	threes := 1
	switch bag[0].jolts {
	case 1:
		ones++
	case 3:
		threes++
	}
	for i := 0; i < len(bag)-1; i++ {
		switch bag[i+1].jolts - bag[i].jolts {
		case 1:
			ones++
		case 3:
			threes++
		}
	}
	fmt.Printf("1s %d x 3s %d  = %d \n", ones, threes, ones*threes)

	// traverse until goal
	endJolts := bag[len(bag)-1].jolts + 3
	routeCount := bag.getRouteCount(0, endJolts, -1)
	fmt.Printf("Total routes to %d: %d\n", endJolts, routeCount)
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
func openFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func Abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}
