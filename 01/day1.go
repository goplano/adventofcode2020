package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Expenses struct {
	Number int64
}

func main() {
	var lines map[int64]*Expenses = make(map[int64]*Expenses)

	file, err := os.Open("expenses.txt")
	if err != nil {
		log.Fatal(err)
	}
	s := bufio.NewScanner(file)
	for s.Scan() {
		n, _ := strconv.ParseInt(s.Text(), 10, 32)

		lines[n] = new(Expenses)
		lines[n].Number = n
	}
	var sum int64 = 2020
	a,b := findPairs(lines, sum)
	fmt.Println("Part One");
	fmt.Printf("%v   x  %v = %v\n", a, b, a * b)
	fmt.Println("\nPart Two");
	for c, _ := range lines {
		diff := sum - c
		a,b = findPairs(lines, diff)
		if a > 0 {
			fmt.Printf("%v   x  %v  x %v = %v\n", a, b, c,  a * b * c)
			break
		}
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
}

func findPairs(lines map[int64]*Expenses, sum int64) (int64,int64) {
	for a, _ := range lines {
		diff := sum - a
		_, ok := lines[diff]
		if ok {
			return a,diff
		}
	}
	return 0,0
}
