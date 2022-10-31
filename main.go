package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

var m0 = [3]string{
	" _ ",
	"| |",
	"|_|",
}

var m1 = [3]string{
	"   ",
	"  |",
	"  |",
}

var m2 = [3]string{
	" _ ",
	" _|",
	"|_ ",
}

var m3 = [3]string{
	" _ ",
	" _|",
	" _|",
}

var m4 = [3]string{
	"   ",
	"|_|",
	"  |",
}

var m5 = [3]string{
	" _ ",
	"|_ ",
	" _|",
}

var m6 = [3]string{
	" _ ",
	"|_ ",
	"|_|",
}

var m7 = [3]string{
	" _ ",
	"  |",
	"  |",
}

var m8 = [3]string{
	" _ ",
	"|_|",
	"|_|",
}

var m9 = [3]string{
	" _ ",
	"|_|",
	" _|",
}

func main() {
	if len(os.Args) != 2 {
		fmt.Print("Incorrect number of arguments.\n")
		os.Exit(1)
	}

	var e1, e2, e3 string
	fileName := os.Args[1]

	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(content), "\n")

	for i, line := range lines {
		fmt.Println(line)
		if i%3 == 0 {
			e1 = line
		} else if i%3 == 1 {
			e2 = line
		} else if i%3 == 2 {
			e3 = line
			validateAccount(e1, e2, e3)
		}
	}
}

func validateAccount(e1, e2, e3 string) {
	var c1, c2, c3, account string
	var flag int = 0

	for n := 0; n < 9; n++ {
		c1, c2, c3 = "", "", ""
		for i := 0; i < 3; i++ {
			c1 += e1[((n * 3) + i) : (n*3)+i+1]
			c2 += e2[((n * 3) + i) : (n*3)+i+1]
			c3 += e3[((n * 3) + i) : (n*3)+i+1]
		}
		account += decodeDigit(c1, c2, c3)

		length := len(account)

		if account[length:] == "?" {
			flag = 1
		} else {
			flag = 0
		}
	}
	fmt.Println(account)
	if flag == 1 {
		fmt.Println("==> " + account + " ILL")
	} else {
		fmt.Println("==> " + account + " " + checkSum(account) + "\n")
	}
}

func decodeDigit(c1, c2, c3 string) string {
	var digit = ""
	if c1 == m0[0] && c2 == m0[1] && c3 == m0[2] {
		digit = "0"
	} else if c1 == m1[0] && c2 == m1[1] && c3 == m1[2] {
		digit = "1"
	} else if c1 == m2[0] && c2 == m2[1] && c3 == m2[2] {
		digit = "2"
	} else if c1 == m3[0] && c2 == m3[1] && c3 == m3[2] {
		digit = "3"
	} else if c1 == m4[0] && c2 == m4[1] && c3 == m4[2] {
		digit = "4"
	} else if c1 == m5[0] && c2 == m5[1] && c3 == m5[2] {
		digit = "5"
	} else if c1 == m6[0] && c2 == m6[1] && c3 == m6[2] {
		digit = "6"
	} else if c1 == m7[0] && c2 == m7[1] && c3 == m7[2] {
		digit = "7"
	} else if c1 == m8[0] && c2 == m8[1] && c3 == m8[2] {
		digit = "8"
	} else if c1 == m9[0] && c2 == m9[1] && c3 == m9[2] {
		digit = "9"
	} else {
		digit = "?"
	}
	return digit
}

func checkSum(possible string) string {
	var response string
	d1, _ := strconv.ParseInt(possible[8:9], 10, 0)
	d2, _ := strconv.ParseInt(possible[7:8], 10, 0)
	d3, _ := strconv.ParseInt(possible[6:7], 10, 0)
	d4, _ := strconv.ParseInt(possible[5:6], 10, 0)
	d5, _ := strconv.ParseInt(possible[4:5], 10, 0)
	d6, _ := strconv.ParseInt(possible[3:4], 10, 0)
	d7, _ := strconv.ParseInt(possible[2:3], 10, 0)
	d8, _ := strconv.ParseInt(possible[1:2], 10, 0)
	d9, _ := strconv.ParseInt(possible[:1], 10, 0)
	if (1*d1+2*d2+3*d3+4*d4+5*d5+6*d6+7*d7+8*d8+9*d9)%11 == 0 {
		response = "OK"
	} else {
		response = "ERR"
	}
	return response
}
