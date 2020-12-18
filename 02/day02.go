package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

/*
 min-max letter: password
 you only need to count the one letter
 how many tested
 how many are valid
 which are invaid
*/
func main() {
	goodPasswords := 0
	badPasswords := 0
	var passwd string
	var min int
	var max int
	var letter byte

	file, err := os.Open("passwords.txt")
	if(err != nil) {
		log.Fatal(err)
	}
	s := bufio.NewScanner(file)
	for s.Scan() {
		//fmt.Println(s.Text())
		fmt.Sscanf(s.Text(),"%d-%d %c: %s",&min, &max, &letter, &passwd )
//fmt.Printf("%s %d %d %c\n", passwd, min, max, letter)
		if(testPass2(min, max, letter, passwd)) {
			goodPasswords++
		} else {
			badPasswords++
		}
	}


	fmt.Printf("Good: %v bad: %v\n", goodPasswords, badPasswords)
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

}
func testPass(min int, max int , letter byte,  str string) bool {
	patStr := fmt.Sprintf("%c", letter)
	re := regexp.MustCompile(patStr)
	result := re.FindAll([]byte(str),-1)
	//fmt.Println(result);
	count := len(result)
	if count >= min && count <= max {
		return true;
	}
	return false;
}
func testPass2(pos1 int, pos2 int , letter byte,  str string) bool {

	return  (hasLetterInPos(str, pos1, letter) == true &&  hasLetterInPos(str, pos2, letter) == false) ||  (hasLetterInPos(str, pos1, letter) == false &&  hasLetterInPos(str, pos2, letter) == true)
}
func hasLetterInPos(str string, pos int, letter byte )bool {
	length := len(str)
	return  pos <= length &&   str[pos-1] == letter
}