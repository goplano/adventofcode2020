package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var inputFilename string
	var preambleLength int
	var codes []int64
	var save map[int]int64
	var badCode int64 = 0


	flag.StringVar(&inputFilename, "file", "noise.txt", "input file name")
	flag.IntVar(&preambleLength, "l",25, "preamble length")
	flag.Parse()

	fmt.Printf("file: %s preamble length %d\n", inputFilename, preambleLength)

	codes = make([]int64, preambleLength, preambleLength)
	save = make(map[int]int64)
	fp := openFile(inputFilename)
	s := bufio.NewScanner(fp)
	i := 0
	for ;  i < preambleLength && s.Scan() ; i++{
		line := s.Text()
		fmt.Printf("Input: %s\n", line)
		val,err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		codes[i] = val
		save[i] = val
	}
	for key, val := range codes {
		fmt.Printf("Key: %d Val: %d\n", key, val)
	}
	for s.Scan() {

		line := s.Text()

		val,err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		isOK := checkCode(val, codes)
		if !isOK {
			fmt.Printf("Code: %d is NOT valid\n", val)
			badCode = val
			break
		}
		save[i] = val
		fmt.Printf("Input: %d checks out\n", val)
		codes = codes[1:preambleLength]
		//for key, val := range codes {
		//	fmt.Printf("Key: %d Val: %d\n", key, val)
		//}
		codes = append(codes, val)
		//for key, val := range codes {
		//	fmt.Printf("Key: %d Val: %d\n", key, val)
		//}
		i++
	}
	if badCode != 0 {
		fmt.Printf(" bad code %d\n", badCode)
		for key, val := range save {
			fmt.Printf(" \t %d:  %d\n", key, val)
		}
		sum := int64(0)
		for i,j:=0,0; j < len(save); {
			val := save[j]
			sum += val
			fmt.Printf("val: %d i: %d j: %d sum: %d \n", val ,i,j,sum)
			if sum == badCode {
				minIdx :=0
				maxIdx :=0
				max:= int64(0)
				min:=badCode
				for idx := i;  idx <= j; idx++ {
					if save[idx] < min {
						minIdx = idx
						min = save[idx]
					}
					if save[idx] > max {
						max = save[idx]
						maxIdx = idx
					}
				}
				fmt.Printf(" %d %d = %d\n", save[minIdx], save[maxIdx], min + max)
				break
			}
			j++
			if val > badCode {
				fmt.Printf("val %d is > badcode %d\n", val, badCode)
				sum = 0
				i++
				j = i
			}

			if sum > badCode {
				fmt.Printf("sum %d is > badcode %d\n", sum, badCode)

				i++
				j = i
				sum = 0
			}



		}
	}
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

func checkCode(val int64, codes []int64) bool {
	isOK := false
	for _,a := range codes{
		if(a > val) {
			continue
		}
		for _, b := range codes {
			if(a == b) {
				continue
			}
			if a + b == val {
				isOK = true
				return isOK
			}
		}
	}
	return isOK
}
func openFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	return file
}