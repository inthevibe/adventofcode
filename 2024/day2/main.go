package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	total := 0
	scanner := bufio.NewScanner(file)
OuterLoop:
	for scanner.Scan() {
		var numbers []int
		line := scanner.Text()
		parts := strings.Split(line, " ")
		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				fmt.Println("Error converting:", part)
			}
			numbers = append(numbers, num)
		}

		increasing := true
		if numbers[1]-numbers[0] < 0 {
			increasing = false
		}
		if numbers[1] == numbers[0] {
			continue OuterLoop
		}

		for i := 1; i < len(numbers); i++ {
			if increasing && numbers[i]-numbers[i-1] > 3 {
				continue OuterLoop
			}
			if !increasing && numbers[i]-numbers[i-1] < -3 {
				continue OuterLoop
			}
			if numbers[i] == numbers[i-1] {
				continue OuterLoop
			}
			if increasing && numbers[i]-numbers[i-1] < 0 {
				continue OuterLoop
			}
			if !increasing && numbers[i]-numbers[i-1] > 0 {
				continue OuterLoop
			}
		}
		total++
	}
	fmt.Println("The total is:", total)
}
