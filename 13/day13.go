package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Bus struct {
	id     int64
	offset int64
}

func (b Bus) waitTime(time int64) int64 {
	trips := time / b.id
	trips++
	return b.arrivalAfter(time) % time
}
func (b Bus) arrivalAfter(time int64) int64 {
	trips := time / b.id
	trips++
	return trips * b.id
}
func (b Bus) isDeparting(time int64) bool {
	if time%b.id == 0 {
		return true
	}
	return false
}

type Buses []Bus
type ByID Buses

func (a ByID) Len() int           { return len(a) }
func (a ByID) Swap(i, j int64)      { a[i], a[j] = a[j], a[i] }
func (a ByID) Less(i, j int64) bool { return a[i].id >= a[j].id }
func (s Buses) ToStr() string {
	result := ""
	for i, b := range s {
		result += fmt.Sprintf("%3d] t - %d ⩭ 0 (mod %3d)\n",
			i, b.offset,  b.id)
	}
	return result
}
func (s Buses) Product() int64 {
	result := int64(1)
	for _, b := range s {
		result *= b.id
	}
	return result
}
func (s Buses) ChineseRemainder() int64 {
	result := int64(0)
	product := s.Product()
	for _, b := range s {
		p := product / b.id
		i:= inverse(p, b.id)
		result += p * i * (b.id-b.offset)
		fmt.Printf("%d x  %d  x %d ~~~~ %d\n", p, i, b.offset, result)
	}
	fmt.Printf("%d mod  %d  =  %d\n", result, product, result % product)

	return result % product
}
func main() {
	var inputFilename string
	var earliestDeparture int64
	var buses []string
	var shuttles Buses

	flag.StringVar(&inputFilename, "file", "notes.txt", "input file name")
	flag.Parse()

	fp := openFile(inputFilename)
	s := bufio.NewScanner(fp)
	if s.Scan() {
		val, err := strconv.ParseInt(s.Text(), 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		earliestDeparture = int64(val)
	}
	if s.Scan() {
		line := s.Text()
		buses = strings.Split(line, ",")
		for i, l := range buses {
			if l == "x" {
				continue
			}
			val, err := strconv.ParseInt(l, 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			bus := Bus{
				id:     int64(val),
				offset: int64(i) ,
			}
			shuttles = append(shuttles, bus)
		}
	}
	fmt.Printf("Earliest Departure: %d\n", earliestDeparture)
	fmt.Printf("Shuttles:\n%s\n", shuttles.ToStr())
	minID := int64(0)
	minWait := earliestDeparture

	for _, bus := range shuttles {
		wait := bus.waitTime(earliestDeparture)
		if wait < minWait {
			minWait = wait
			minID = bus.id
		}
	}
	fmt.Printf("Wait time for shuttle %d is %d after %d ANSWER: %d \n", minID, minWait, earliestDeparture, minWait*minID)
	time := shuttles.ChineseRemainder()
	fmt.Printf("Earliest timestamp is %d\n", time)
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


func inverse(a int64 ,n int64) int64 {
	t := int64(0)
	r := n
	newT := int64(1)
	newR := a

	for newR != 0 {

		q := r / newR
		t,newT = newT, t - q * newT
		r,newR = newR, r - q * newR
	}
	if r > 1 {
		log.Fatal("not invertible")
	}
	if t < 0 {
		t = t + n
	}
	return t
}

