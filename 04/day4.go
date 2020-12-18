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

func main() {
	var numChecked int = 0
	var numValid int = 0
	var numInvalid int = 0
	var numWithRequired int = 0
	var numWithoutReqired int = 0
	var inputFilename string
	var bits []string
	var set map[string]string
	var keys map[string]string
	var bitmap map[string]int
	keys = map[string]string{
		"byr": "Birth Year",
		"iyr": "Issue Year",
		"eyr": "Expiration Year",
		"hgt": "Height",
		"hcl": "Hair Color",
		"ecl": "Eye Color",
		"pid": "Passport ID",
		"cid": "Country ID",
	}
	bitmap = map[string]int{
		"byr": 1,
		"iyr": 2,
		"eyr": 4,
		"hgt": 8,
		"hcl": 16,
		"ecl": 32,
		"pid": 64,
		"cid": 128,
	}
	hasRequiredMask := 0
	for _, bits := range bitmap {
		hasRequiredMask = hasRequiredMask | bits
	}

	flag.StringVar(&inputFilename, "file", "passports.txt", "input file name")
	flag.Parse()

	fmt.Println("Reading ", inputFilename)
	fp := openFile(inputFilename)
	s := bufio.NewScanner(fp)
	set = make(map[string]string)
	for s.Scan() {
		line := s.Text()
		if len(line) == 0 {
			// do validity check
			if len(set) > 0 {
				numChecked++
				hasBits := inspectSetContents(keys, set, bitmap)
				hasRequired := checkForRequiredFields(hasBits, bitmap, hasRequiredMask)
				if hasRequired {
					numWithRequired++
					isValid := checkValidity(set)
					if isValid {
						numValid++
					} else {
						numInvalid++
					}
				} else {
					numWithoutReqired++
					numInvalid++
				}
			}
			// reset
			set = make(map[string]string)
		} else {
			bits = strings.Split(line, " ")
			//fmt.Printf("%q\n", bits)
			for _, str := range bits {
				parts := strings.Split(str, ":")
				//fmt.Printf("%s: %s\n", parts[0], parts[1])

				set[parts[0]] = parts[1]
			}
		}

	}
	if len(set) > 0 {
		numChecked++
		hasBits := inspectSetContents(keys, set, bitmap)
		hasRequired := checkForRequiredFields(hasBits, bitmap, hasRequiredMask)
		if hasRequired {
			numWithRequired++
			isValid := checkValidity(set)
			if isValid {
				numValid++
			} else {
				numInvalid++
			}

		} else {
			numWithoutReqired++
			numInvalid++
		}
	}
	fmt.Printf("Checked %d.  %d with all required fields.  %d without all required fields.\n", numChecked, numWithRequired, numWithoutReqired)
	fmt.Printf("Checked %d.  %d valid.  %d invalid.\n", numChecked, numValid, numInvalid)

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
	byr (Birth Year) - four digits; at least 1920 and at most 2002.
		^([0-9]{4})$
	iyr (Issue Year) - four digits; at least 2010 and at most 2020.
		^([0-9]{4})$
	eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
		^([0-9]{4})$
	hgt (Height) - a number followed by either cm or in:
		^([0-9]{2,3})(cm|in)$
	If cm, the number must be at least 150 and at most 193.
	If in, the number must be at least 59 and at most 76.
	hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
		^#[0-9a-f]{6}$
	ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
		^amb|blu|brn|gry|grn|hzl|oth$
	pid (Passport ID) - a nine-digit number, including leading zeroes.
		^[0-9]{9}$
	cid (Country ID) - ignored, missing or not.
 */
func checkValidity( set map[string]string) bool {
	if !checkPassportID(set["pid"]) {
		fmt.Printf("Invalid passport id %q\n", set["pid"])
		return false
	}
	if !checkEyeColor(set["ecl"]) {
		fmt.Printf("Invalid eye color %q\n", set["ecl"])
		return false
	}
	if !checkHairColor(set["hcl"]) {
		fmt.Printf("Invalid hair color %q\n", set["hcl"])
		return false
	}
	if !checkHeight(set["hgt"]) {
		fmt.Printf("Invalid height %q\n", set["hgt"])
		return false
	}
	if !checkYear(set["byr"], 1920, 2002) {
		fmt.Printf("Invalid birth year %q\n", set["byr"])
		return false
	}
	if !checkYear(set["iyr"], 2010, 2020) {
		fmt.Printf("Invalid issue year %q\n", set["iyr"])
		return false
	}
	if !checkYear(set["eyr"], 2010, 2030) {
		fmt.Printf("Invalid expiration year %q\n", set["eyr"])
		return false
	}
	return true
}
func checkPassportID(str string) bool {
	patStr := "^[0-9]{9}$"
	re := regexp.MustCompile(patStr)
	result := re.Find([]byte(str))
	//fmt.Printf("%q\n", result)
	return len(result) > 0
}
func checkEyeColor(str string) bool {
	patStr := "^amb|blu|brn|gry|grn|hzl|oth$"
	re := regexp.MustCompile(patStr)
	result := re.Find([]byte(str))
	//fmt.Printf("%q\n", result)
	return len(result) > 0
}
func checkHairColor(str string) bool {
	patStr := "^#[0-9a-f]{6}$"
	re := regexp.MustCompile(patStr)
	result := re.Find([]byte(str))
	//fmt.Printf("%q\n", result)
	return len(result) > 0
}
func checkHeight(str string) bool {
	patStr := "^([0-9]{2,3})(cm|in)$"
	re := regexp.MustCompile(patStr)
	result := re.FindStringSubmatch(str)
	if len(result) == 3 {
		units := result[2]
		value, err := strconv.ParseInt(string(result[1]), 10, 32)
		if err != nil {
			return false
		}
		if units == "cm" {
			return value >= 150 && value <= 193
		}
		if units == "in" {
			return value >= 59 && value <= 76
		}
	}
	return false
}

func checkYear(str string, minVal int, maxVal int) bool {
	patStr := "^([0-9]{4})$"
	re := regexp.MustCompile(patStr)
	result := re.Find([]byte(str))

	year, err := strconv.ParseInt(string(result), 10, 32)
	if err != nil {
		return false
	}
	return int(year) >= minVal && int(year) <= maxVal
}
func checkForRequiredFields(hasBits int, bitmap map[string]int, hasRequiredMask int) bool {
	return (hasBits|bitmap["cid"])&hasRequiredMask == hasRequiredMask

}

func inspectSetContents(keys map[string]string, set map[string]string, bitmap map[string]int) int {
	hasBits := 0
	for key, _ := range keys {
		//fmt.Printf("Testing %s %v\n", key, set[key])
		_, ok := set[key]
		//fmt.Printf("Map test %q\n", ok)
		if ok {
			hasBits = hasBits | bitmap[key]
			//fmt.Println(hasBits, bitmap[key])

		}
	}
	return hasBits
}
func openFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	return file
}
