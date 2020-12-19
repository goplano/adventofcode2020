package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Bag struct {
	name     string
	contains map[string]int
}

func (bag Bag) add(name string, qty int) {
	_, ok := bag.contains[name]
	if ok {
		bag.contains[name] = qty + bag.contains[name]
	} else {
		bag.contains[name] = qty
	}
}

func (bag Bag) containedBagCount(bags Bags) int {
	count := 0

	for name,num := range bag.contains {
		count += num
		subBag, ok := bags[name]
		if ok {
			count += subBag.containedBagCount(bags) * num
		}
	}
	return count
}

func (bag Bag) canContain(name string, bags Bags) bool {
	_, ok := bag.contains[name]
	if ok {
		return true
	} else {
		for t := range bag.contains {
			subBag, ok := bags[t]
			if ok {
				if subBag.canContain(name, bags) {
					return true
				}
			}
		}
	}
	return false
}

type Bags map[string]Bag

func (bags Bags) find(name string) Bag {
	//code
	bag, ok := bags[name]
	if ok {
		return bag
	}
	newBag := Bag{
		name:     name,
		contains: make(map[string]int),
	}
	bags[name] = newBag
	return bags[name]
}
func main() {
	var inputFilename string
	bags := make(Bags)
	flag.StringVar(&inputFilename, "file", "rules.txt", "input file name")
	flag.Parse()

	fmt.Println("Reading ", inputFilename)
	fp := openFile(inputFilename)
	s := bufio.NewScanner(fp)
	for s.Scan() {
		line := s.Text()
		//fmt.Printf("Input: %s\n", line)
		reLine := regexp.MustCompile("^(.*) bags contain (.*).$")
		reContained := regexp.MustCompile("^([0-9]*) (.*) bag[s]{0,1}$")

		result := reLine.FindStringSubmatch(line)
		//fmt.Printf("Matches %q\n", result)
		if len(result) == 3 {
			mainBag := result[1]
			bag := bags.find(mainBag)
			containsStr := result[2]
			bags := strings.Split(containsStr, ", ")
			//fmt.Printf("Bag %q contains\n",bag)
			for _, txt := range bags {
				result := reContained.FindStringSubmatch(txt)
				if len(result) == 3 {
					qty, err := strconv.ParseInt(result[1], 10, 32)
					bagName := string(result[2])
					if err != nil {
						break
					}
					bag.add(bagName, int(qty))
					//fmt.Printf("\t %d %s\n", qty, bagName)

				}
			}

		}
	}
	count := 0
	for name, bag := range bags {
		if bag.canContain("shiny gold", bags) {
			count++
		}
		fmt.Printf("%s %#v \n", name, bag)
	}
	fmt.Printf("%d bags can contain shiny gold \n", count)
	shiny := bags.find("shiny gold")
	fmt.Printf("shiny gold holds %d bags\n",shiny.containedBagCount(bags) )
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

/*
striped orange bags contain 5 light crimson bags, 3 muted tan bags, 5 dotted coral bags, 3 plaid blue bags.
(.*) bags contain (.*)

*/
func openFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	return file
}
