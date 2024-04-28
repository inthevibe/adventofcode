package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var num int
var num2 int
var total int

type NumStruct struct {
	number   int
	position int
}

var num1pos NumStruct
var num2pos NumStruct

var zeroStruct = NumStruct{}

var numberMap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}
var round = 0

func main() {
	file, err := os.Open("input.txt")
	//defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		num = 0
		num2 = 0
		string := scanner.Text()

		for _, ch := range string {
			if ch >= '0' && ch <= '9' {
				num = int(ch-'0') * 10
				break
			}
		}

		for i := len(string) - 1; i >= 0; i-- {
			if string[i] >= '0' && string[i] <= '9' {
				num2 = int(string[i] - '0')
				break
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
		total += num + num2
		//fmt.Println("Total in this loop: ", total)
	}

	fmt.Println("Total: ", total)

	num = 0
	num2 = 0
	total = 0

	// Seek the file pointer back to the beginning of the file
	_, err = file.Seek(0, 0)
	if err != nil {
		log.Fatal(err)
	}

	scanner2 := bufio.NewScanner(file)

	for scanner2.Scan() {
		num1pos.number = 0
		num2pos.number = 0
		num1pos = zeroStruct
		num2pos = zeroStruct
		string := scanner2.Text()
		for i := 0; i <= len(string)-1; i++ {
			if string[i] >= '0' && string[i] <= '9' {
				num1pos.number = int(string[i]-'0') * 10
				num1pos.position = i
				break
			}
		}
		//fmt.Println("The num1/number is: ", num1pos.number)
		for key := range numberMap {
			if strings.Index(string, key) <= num1pos.position && strings.Index(string, key) != -1 {
				num1pos.position = strings.Index(string, key)
				num1pos.number = numberMap[key] * 10
			}
		}

		//fmt.Println("The num1/string is: ", num1pos.number)
		for z := len(string) - 1; z >= 0; z-- {
			if string[z] >= '0' && string[z] <= '9' {
				num2pos.number = int(string[z] - '0')
				num2pos.position = z
				break
			}
		}

		//fmt.Println("The num2/number is: ", num2pos.number)
		for key2 := range numberMap {
			if strings.LastIndex(string, key2) >= num2pos.position && strings.LastIndex(string, key2) != -1 {
				num2pos.position = strings.LastIndex(string, key2)
				num2pos.number = numberMap[key2]
			}
		}

		//fmt.Println("The num2/string is: ", num2pos.number)
		total += num1pos.number + num2pos.number
		//fmt.Println("The total is :", total, "Number 1: ", num1pos.number, "Number 2: ", num2pos.number, "Total", num1pos.number+num2pos.number, "in the round: ", round)
		round++
	}
	fmt.Println("Total 2: ", total)
}
