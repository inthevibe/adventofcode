package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var num int
var num2 int
var total int

func main() {
	file, err := os.Open("input.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		//fmt.Println(scanner.Text())

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
}
