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

type Inst struct {
	operation string
	argument  int
	seen      bool
}
type Code map[int]Inst
type Mod struct {
	orig Inst
	new  Inst
	seen bool
}
type Mods map[int]Mod

func main() {
	var inputFilename string
	var acc int = 0
	var completed bool
	var code Code

	flag.StringVar(&inputFilename, "file", "bootcode.txt", "input file name")
	flag.Parse()

	fmt.Println("Reading ", inputFilename)
	code = readFile(inputFilename)

	acc, completed = execute(code)
	if !completed {
		fmt.Printf("Infinite loop found. acc: %d\n", acc)
	} else {
		fmt.Printf("Completed. acc: %d\n", acc)
	}
	mods := extractMods(code)
	fmt.Printf("\nTRYING AGAIN\n")
	for key, mod := range mods {
		resetSeen(code)
		code[key] = mod.new
		acc, completed = execute(code)
		if !completed {
			fmt.Printf("Infinite loop found. acc: %d\n", acc)
			code[key] = mod.orig
		} else {
			fmt.Printf("Completed. acc: %d\n", acc)
			break
		}
	}


}

func extractMods(code Code) Mods {
	mods := make(Mods)
	for key, value := range code {
		if value.operation == "jmp" {
			mods[key] = Mod {
				orig: code[key],
				new: Inst{
					operation: "nop",
					argument:  0,
					seen:      false,
				},
				seen: false,
			}
		} else if value.operation == "nop" {
			mods[key] = Mod {
				orig: code[key],
				new: Inst{
					operation: "jmp",
					argument: value.argument,
					seen:      false,
				},
				seen: false,
			}
		}
	}
	return mods
}
func resetSeen(code Code)  {
	for key, value := range code {
		code[key] = Inst{
			operation: value.operation,
			argument:  value.argument,
			seen:      false,
		}
	}
}

func execute(code Code) (int, bool) {
	completed := false
	numInsts := len(code)
	acc := 0
	i := 0
	for ; i < numInsts; {
		if code[i].seen {
			fmt.Printf("Acc before instruction %d, [%s, %d] = %d\n", i, code[i].operation, code[i].argument, acc)
			break
		}
		code[i] = Inst{
			operation: code[i].operation,
			argument:  code[i].argument,
			seen:      true,
		}
		switch code[i].operation {
		case "nop":
			i++
		case "acc":
			acc += code[i].argument
			i++
		case "jmp":
			i += code[i].argument
		}
		//fmt.Printf("%d) %s %4d %5t acc: %d\n", i, code[i].operation, code[i].argument, code[i].seen, acc)

	}
	if i == numInsts {
		completed = true
	}
	return acc, completed
}

func readFile(inputFilename string) Code {
	numLines := 0
	code := make(Code)
	fp := openFile(inputFilename)
	s := bufio.NewScanner(fp)
	for s.Scan() {
		line := s.Text()
		parts := strings.Split(line, " ")

		val, err := strconv.ParseInt(string(parts[1]), 10, 32)
		if err == nil {
			code[numLines] = Inst{
				operation: string(parts[0]),
				argument:  int(val),
				seen:      false,
			}
		} else {
			fmt.Printf("%#v\n", err)
		}
		numLines++
	}
	err := s.Err()
	if err != nil {
		log.Fatal(err)
	}
	if err = fp.Close(); err != nil {
		log.Fatal(err)
	}
	return code
}
func openFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	return file
}
